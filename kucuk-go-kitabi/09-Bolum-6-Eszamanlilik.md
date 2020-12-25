- [Bölüm 6 - Eşzamanlılık](#bölüm-6---eşzamanlılık)
  - [Go Rutinleri](#go-rutinleri)
  - [Senkronizasyon](#senkronizasyon)
  - [Kanallar](#kanallar)
    - [Tamponlu Kanallar](#tamponlu-kanallar)
    - [Select](#select)
    - [Zaman Aşımı](#zaman-aşımı)
  - [Devam Etmeden Önce](#devam-etmeden-önce)

# Bölüm 6 - Eşzamanlılık

Go genellikle eşzamanlılık dostu bir dil olarak tanımlanır. Bunun nedeni, iki güçlü mekanizma için basit bir sözdizimi sağlamasıdır: go rutinleri ve kanallar.

## Go Rutinleri

Bir goroutine bir çok dilden alışkın olduğumuz thread'lere benzer, ancak işletim sistemi tarafından değil Go tarafından çalıştırılır. Bir goroutine'de çalışan kod diğer kodlarla aynı anda çalışabilir. Bir örneğe bakalım:

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("start")
  go process()
  time.Sleep(time.Millisecond * 10) // bu kötü bir pratiktir, sık kullanmayın!
  fmt.Println("done")
}

func process() {
  fmt.Println("processing")
}
```

Burada birkaç ilginç şey var, ama en önemlisi bir goroutine nasıl başlattığımızdır. Sadece `go` anahtar sözcüğünü ve ardından yürütmek istediğimiz işlevi kullanırız. Yukarıdaki gibi bir kod çalıştırmak istiyorsak, anonim bir işlev de kullanabiliriz. Ancak anonim işlevlerin yalnızca go rutinleri ile kullanılmadığını unutmayın.

```go
go func() {
  fmt.Println("processing")
}()
```

Go rutinleri oluşturmak kolaydır ve sisteme yükü azdır. Birden fazla go rutin aynı temel OS iş parçacığında çalışır. N OS iş parçacığı üzerinde çalışan M uygulama iş parçacığı (go rutin) olduğu için buna genellikle M: N iş parçacığı modeli denir. Sonuç, bir go rutinin OS iş parçacıklarından daha fazla ek yüke (birkaç KB) sahip olmasıdır. Modern donanımlarda milyonlarca go rutine sahip olmak mümkündür.

Ayrıca, eşlemenin ve programlamanın karmaşıklığı gizlenmiştir. Sadece *bu kodun aynı anda çalışması gerektiğini* söylüyoruz ve Go'nun bunu gerçekleştirmesi konusunda uğraşmasını izliyoruz.

Biz geri bizim örneğimize giderseniz, `Sleep` ile birkaç milisaniye zorunda  bekletmek olduğumuzu fark edeceksiniz . Çünkü ana süreç go rutinin çalışma şansı elde etmeden önce çalışmasını bitirebilir (ana süreç çıkmadan önce tüm alt go rutinler bitene kadar beklemez). Bunu çözmek için kodumuzu koordine etmemiz gerekiyor.

## Senkronizasyon

Go rutin oluşturmak çok basittir ve o kadar ucuzdur ki, bir çok rutin başlatabiliriz; ancak, eşzamanlı kodun koordine edilmesi gerekir. Bu soruna yardımcı olmak için Go `channels` adında yapıları sağlar. `channels` kavramına bakmadan önce, eşzamanlı programlamanın temelleri hakkında biraz bilgi sahibi olmanın önemli olduğunu düşünüyorum.

Eşzamanlı çalışan kod yazmak, değerleri nerede ve nasıl okuduğunuza ve yazdığınıza özellikle dikkat etmenizi gerektirir. Bazı yönlerden, çöp toplayıcı olmadan programlama gibi - verilerinizi yeni bir açıdan düşünmenizi ve olası tehlikelere karşı her zaman dikkatli olmanızı gerektirir. Örneğin:

```go
package main

import (
  "fmt"
  "time"
)

var counter = 0

func main() {
  for i := 0; i < 20; i++ {
    go incr()
  }
  time.Sleep(time.Millisecond * 10)
}

func incr() {
  counter++
  fmt.Println(counter)
}
```

Çıktının ne olacağını düşünüyorsunuz?

Çıktının `1, 2, ... 20` olduğunu düşünüyorsanız hem doğru düşünüyorsunuz hem de yanılıyorsunuz. Yukarıdaki kodu çalıştırırsanız, bazen bu çıktıyı alırsınız. Ancak, gerçek şu ki, davranış tanımsızdır. Neden? Potansiyel olarak `counter` adlı değişkene aynı anda yazma hakkı olan birden fazla go rutine (bu durumda iki işlev) sahibiz. Ya da, daha kötüsü, bir go rutin `counter` değerini okurken diğeri yazıyor olabilirdi.

Bu gerçekten bir tehlike mi? Evet kesinlikle. `counter++` basit bir kod satırı gibi görünebilir, ancak aslında birden fazla assembly ifadesine bölünür - ne olacağı kodu çalıştırdığınız platforma bağlıdır. Bu örneği çalıştırırsanız, sayıların genellikle garip bir sırada yazdırıldığını, sayıların çoğaltıldığını ve ya eksik olduğunu görürsünüz. Sistem çökmeleri veya rastgele bir veri parçasına erişme ve bunları artırma gibi daha kötü olasılıklar da var!

Bir değişkene güvenli bir şekilde yapabileceğiniz tek eşzamanlı şey, onu okumaktır. İstediğiniz kadar okuyucuya sahip olabilirsiniz, ancak yazıların senkronize edilmesi gerekir. Özel CPU talimatlarına dayanan bazı gerçekten atomik işlemleri kullanmak da dahil olmak üzere bunu yapmanın çeşitli yolları vardır. Bununla birlikte, en yaygın yaklaşım bir mutex kullanmaktır:

```go
package main

import (
  "fmt"
  "time"
  "sync"
)

var (
  counter = 0
  lock sync.Mutex
)

func main() {
  for i := 0; i < 20; i++ {
    go incr()
  }
  time.Sleep(time.Millisecond * 10)
}

func incr() {
  lock.Lock()
  defer lock.Unlock()
  counter++
  fmt.Println(counter)
}
```

Bir mutex kilit altındaki koda erişimi seri hale getirir. Kilidi basitçe `lock sync.Mutex`  olarak tanımlamamızın nedeni bir `sync.Mutex` varsayılan değerinin kilidinin açıl olmasıdır.

Yeterince basit görünüyor mu? Yukarıdaki örnek aldatıcıdır. Eşzamanlı programlama yaparken ortaya çıkabilecek bir dizi ciddi hata var. Her şeyden önce, hangi kodun korunması gerektiği her zaman açık değildir. Kilitleri (büyük miktarda kodu kapsayan kilitler) kullanmak cazip gelse de, bu ilk etapta eşzamanlı programlama yapmamızın nedenini zayıflatır. Genellikle akıllı kilitler istiyoruz; başka bir şekilde, aniden tek şeritli bir yola dönüşen on şeritli bir otoyolla karşılaşırız.

Diğer sorun kilitlenmelerle ilgilidir. Tek bir kilitle, bu bir sorun değildir, ancak aynı kodun etrafında iki veya daha fazla kilit kullanıyorsanız, goroutineA'nın lockA'yı tuttuğu, ancak lockB'ye erişmesi gerektiğinde, goroutineB'nin lockB'yi tuttuğu, ancak erişime ihtiyacı olduğu gibi karmaşık durumlara sahip olmak tehlikeli derecede kolaydır.

Aslında biz kilidi serbest bırakmayı unutursak, tek kilit ile kilitlenme *mümkündür*. Bu, çok kilitli bir kilitlenme kadar tehlikeli değildir (çünkü bunların tespit edilmesi *gerçekten* zor), ancak ne olduğunu görebilmeniz için aşağıdaki kodu çalıştırmayı deneyin:

```go
package main

import (
  "time"
  "sync"
)

var (
  lock sync.Mutex
)

func main() {
  go func() { lock.Lock() }()
  time.Sleep(time.Millisecond * 10)
  lock.Lock()
}
```

Eşzamanlı programlamada şimdiye kadar gördüğümüzden daha fazlası var. Bir kere, okuma-yazma mutex adı verilen başka bir ortak mutex de var. Bu iki kilitleme işlevi ortaya koyar: biri okuma için kilitlemek ve diğeri yazma için kilitlemek. Bu ayrım, yazmanın ayrıcalıklı olmasını sağlarken aynı anda birden fazla okuyucuya izin verir. `sync.RWMutex` böyle bir kilittir. Bir `sync.Mutex` `Lock` ve `Unlock` yöntemlerine ek olarak, `RLock` ve `RUnlock` yöntemlerini de gösterir; burada `R` , *read* anlamına gelir. Okuma-yazma muteksleri yaygın olarak kullanılırken, geliştiricilere ek bir yük getirir: şimdi sadece verilere erişirken değil, başka durumlarda da dikkat etmeliyiz.

Ayrıca, eşzamanlı programlamanın bir kısmı, mümkün olan en dar kod parçasına erişimi serileştirmekle ilgili değildir; aynı zamanda çoklu go rutinleri koordine etmekle ilgilidir. Örneğin, 10 milisaniye beklemek özellikle zarif bir çözüm değildir. Bir go rutin 10 milisaniyeden fazla sürerse ne olur? Ya daha az zaman alırsa ve sadece CPU zamanını israf edersek? Ayrıca, sadece go rutinlerin bitmesini beklemek yerine, birine "*hey, işlemek için yeni verilerim var!*" söylemek istersek nasıl olur?

Bunlar `kanallar` olmadan yapılabilir olan tüm şeylerdir. Kesinlikle daha basit durumlar için, `sync.Mutex` ve `sync.RWMutex` gibi ilkelleri kullanmanız **gerektiğine** inanıyorum, ancak bir sonraki bölümde göreceğimiz gibi, `kanallar` eşzamanlı programlamayı daha temiz ve daha az hataya yatkın hale getirmeyi amaçlıyor.

## Kanallar

Eşzamanlı programlama ile ilgili zorluk, veri paylaşımından kaynaklanmaktadır. Go rutinleriniz veri paylaşmıyorsa, bunları senkronize etme konusunda endişelenmenize gerek yoktur. Ancak bu, tüm sistemler için bir seçenek değildir. Aslında, pek çok sistem tam tersi amaç göz önünde bulundurularak oluşturulmuştur: birden fazla talep arasında veri paylaşmak. Bir bellek içi önbellek veya bir veritabanı bunun iyi örnekleridir. Bu giderek yaygınlaşan bir gerçeklik haline geliyor.

Kanallar, paylaşılan verileri büyük resimden çıkararak eşzamanlı programlama yapmayı kolaylaştırır. Kanal, veri aktarmak için kullanılan go rutinler arasındaki bir iletişim hattıdır. Başka bir deyişle, verileri olan bir go rutin, bir kanal aracılığıyla başka bir go rutine bu verileri gönderebilir. Sonuç, herhangi bir anda, yalnızca bir go rutinin verilere erişimi olmasıdır.

Bir kanalın, diğer her şey gibi, bir tipi vardır. Bu kanalımızdan geçireceğimiz veri tipidir. Örneğin, bir tamsayıyı iletmek için kullanılabilecek bir kanal oluşturmak için şunları yaparız:

```go
c := make(chan int)
```

Bu kanalın türü `chan int` . Bu nedenle, bu kanalı bir işleve geçirmek için imzamız şöyle olmalıdır:

```go
func worker(c chan int) { ... }
```

Kanallar iki işlevi destekler: alma ve gönderme. Bir kanala şunu yaparak veri göndeririz:

```
CHANNEL <- DATA
```

ve şunu yaparak veri alırız

```
VAR := <-CHANNEL
```

Ok işareti, verinin aktığı yönü gösterir. Gönderirken, veriler kanala akar. Alma sırasında, veriler kanaldan dışarı akar.

İlk örneğimize bakmadan önce bilmemiz gereken son şey, bir kanala alma ve bir kanaldan göndermenin kilitleme özelliğinin olmasıdır. Yani, bir kanaldan veri alırken, veri bulunana kadar go rutinin yürütülmesi devam etmez. Benzer şekilde, bir kanala gönderdiğimizde, veri alınana kadar yürütme devam etmez.

Gelen verileri ayrı go rutinlerde işlemek istediğimiz bir sistem düşünün. Bu çok sık karşılaştığımız bir istektir. Gelen verileri kabul eden go routine içinfr yoğun veri işlememizi yapsaydık, istemcilerin zaman aşımına uğrama riskiyle karşı karşıya kalırdık. İlk önce işleyici kodumuzu yazacağız. Bu basit bir işlev olabilir, ancak daha önce bu şekilde kullanılan go rutinleri görmediğimiz için bir yapının yöntemi olarak yapacağım:

```go
type Worker struct {
  id int
}

func (w Worker) process(c chan int) {
  for {
    data := <-c
    fmt.Printf("worker %d got %d\n", w.id, data)
  }
}
```

İşleyicimiz basit. Veriler hazır olana kadar bekler ve sonra "işler". Elbette, bunu bir döngüde yapar, sonsuza kadar daha fazla verinin gönderilmesini bekler.

Bunu kullanmak için yapacağımız ilk şey bir kaç işleyici başlatmaktır:

```go
c := make(chan int)
for i := 0; i < 5; i++ {
  worker := &Worker{id: i}
  go worker.process(c)
}
```

Ve sonra şu şekilde onlara biraz iş verebiliriz:

```go
for {
  c <- rand.Int()
  time.Sleep(time.Millisecond * 50)
}
```

Çalıştırmak için kodun şöyle bir araya getirelim:

```go
package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  c := make(chan int)
  for i := 0; i < 5; i++ {
    worker := &Worker{id: i}
    go worker.process(c)
  }

  for {
    c <- rand.Int()
    time.Sleep(time.Millisecond * 50)
  }
}

type Worker struct {
  id int
}

func (w *Worker) process(c chan int) {
  for {
    data := <-c
    fmt.Printf("worker %d got %d\n", w.id, data)
  }
}
```

Hangi işleyicinin hangi verileri alacağını bilmiyoruz. Bildiğimiz, Go'nun garanti ettiği tek şey, bir kanala gönderdiğimiz verilerin yalnızca tek bir alıcı tarafından alınacağıdır.

Paylaşılan tek durumun, aynı anda güvenli bir şekilde veri alıp gönderebileceğimiz kanal olduğuna dikkat edin. Kanallar, ihtiyacımız olan tüm senkronizasyon kodunu sağlar ve ayrıca herhangi bir zamanda yalnızca bir go rutinin belirli bir veri parçasına erişmesini sağlar.

### Tamponlu Kanallar

Yukarıdaki kod göz önüne alındığında, işleyebileceğimizden daha fazla veri geliyorsa ne olur? Veri aldıktan sonra işleyici işlevi uyku moduna geçirerek bunu simüle edebilirsiniz:

```go
for {
  data := <-c
  fmt.Printf("worker %d got %d\n", w.id, data)
  time.Sleep(time.Millisecond * 500)
}
```

Olan şey, ana kodumuzda, kullanıcının gelen verilerini kabul eden kod (rastgele bir sayı üreteci ile simüle ettik) kanala gönderdiği için bloke oluyor çünkü alabilecek müsait bir alıcı yok.

Verilerin işlendiğine dair garantilere ihtiyaç duyduğunuz durumlarda, muhtemelen istemciyi engellemeyi seçmek istersiniz. Tersi durumlarda, bu garantileri gevşetmeyi seçebilirsiniz. Bunu yapmak için birkaç popüler strateji vardır. Birincisi, verileri tamponlamaktır. Hazır bir işleyici yoksa, verileri geçici olarak bir tür kuyrukta saklamak isteyebiliriz. Kanallarda bu tamponlama özelliği yerleşik olarak bulunmaktadır. Kanalımızı `make` ile oluşturduğumuzda, kanalımıza bir uzunluk verebiliriz:

```go
c := make(chan int, 100)
```

Bu değişikliği yapabilirsiniz, ancak işlemenin hala dalgalı olduğunu fark edeceksiniz. Tamponlu kanallar daha fazla kapasite eklemez; sadece bekleyen bir iş kuyruğu ve ani bir artışla başa çıkmanın iyi bir yolunu sunarlar. Örneğimizde, sürekli olarak işleyicilerimizin işleyebileceğinden daha fazla veri gönderiyoruz.

Bununla birlikte, tamponlu kanalın ne olduğunu, aslında kanalın `len` değerine bakarak anlayabiliriz:

```go
for {
  c <- rand.Int()
  fmt.Println(len(c))
  time.Sleep(time.Millisecond * 50)
}
```

Dolduruluncaya kadar büyüyüp büyüdüğünü görebilirsiniz, bu noktada kanalımıza gönderme tekrar engellenmeye başlacaktır.

### Select

Tamponla  bile, iletileri bırakmaya başlamamız gereken bir nokta vardır. Bir işleyicinin serbest bırakacağı umuduyla sonsuz miktarda bellek kullanamayız. Bunun için biz Go'nun `select` kavramını kullanırız.

Sözdizimsel olarak, `select` swicth kullanımına çok benzer. Bununla kanalın veri gönderimine müsait olmadığı zaman için kod sağlayabiliriz. İlk olarak, `select` nasıl çalıştığını açıkça görebilmemiz için kanalımızın tampon özelliğini kaldıralım:

```go
c := make(chan int)
```

Sonra, `for` döngümüzü değiştiriyoruz:

```go
for {
  select {
  case c <- rand.Int():
    // buraya kod yazılabilir
  default:
    // burası boş olarak da bırakılabilir, kanalın bırakıldığına dair birşey söylenmek istenmediği zamanlarda
    fmt.Println("dropped")
  }
  time.Sleep(time.Millisecond * 50)
}
```

Saniyede 20 mesaj gönderiyoruz, ancak işleyicilerimiz saniyede yalnızca 10 mesaj işleyebilir; böylece mesajların yarısı boşa gider.

Bu bizim `select` ile neleri başarabileceğini sadece başlangıçtır . Select'in temel amacı, birden fazla kanalı yönetmektir. Birden fazla kanal verildiğinde, `select` birincisi kullanılabilir hale gelene kadar engellenir. Hiçbir kanal yoksa, eğer varsa `default` seçeneği yürütülür. Birden fazla kanal kullanılabilir olduğunda biri rastgele seçilir.

Oldukça gelişmiş bir özellik olduğu için bu davranışı gösteren basit bir örnek bulmak zor. Bir sonraki bölüm bunu açıklamaya yardımcı olabilir.

### Zaman Aşımı

İletileri tampona almanın yanı sıra basitçe yoketmeye de baktık. Bir başka popüler seçenek de zaman aşımıdır. Bir süre beklemeye hazırız, ama sonsuza kadar değil. Bu aynı zamanda Go'da başarılması kolay bir şeydir. Kuşkusuz, sözdizimini takip etmek biraz zor olabilir, ancak bu dışarıda bırakılmayacak kadar düzgün ve kullanışlı bir özelliktir.

Maksimum süre belirlemek için, `time.After` işlevini kullanabiliriz. Ona bakalım, sonra büyünün ötesine bakmaya çalışalım. Bunu kullanmak için veri gönderen kodumuz:

```go
for {
  select {
  case c <- rand.Int():
  case <-time.After(time.Millisecond * 100):
    fmt.Println("timed out")
  }
  time.Sleep(time.Millisecond * 50)
}
```

`time.After` bir kanal döndürür, böylece `select` yapabiliriz. Kanala, belirtilen süre dolduktan sonra yazılır. Bu kadar. Bundan daha büyülü bir şey yok. Merak ediyorsanız, `after` işlevinin kodu şöyle görünebilir:

```go
func after(d time.Duration) chan bool {
  c := make(chan bool)
  go func() {
    time.Sleep(d)
    c <- true
  }()
  return c
}
```

Yazdığımız `select` koduna geri dönersek, deneyebileceğimiz bir kaç şey vardır. İlk olarak, `default` durumu geri eklerseniz ne olur? Tahmin edebilir misin? Deneyin. Neler olup bittiğinden emin değilseniz, kullanılabilir kanal yoksa `default` seçeneğinin hemen tetiklendiğini unutmayın.

Ayrıca, `time.After` `chan time.Time` tipinde bir kanal döner. Yukarıdaki örnekte, kanala gönderilen değeri atıyoruz. Eğer isterseniz, alabilirsiniz de:

```go
case t := <-time.After(time.Millisecond * 100):
  fmt.Println("timed out at", t)
```

`select` uygulamamıza çok dikkat edin. Her zaman  `c`'ye gönderiyoruz ama fakat `time.After` tipinde kanaldan da veri alabiliyoruz. `select` kanallardan alma, gönderme veya herhangi bir kanal kombinasyonundan bağımsız olarak aynı şekilde çalışır:

- İlk kullanılabilir kanal seçilir.
- Birden fazla kanal varsa, rastgele bir kanal seçilir.
- Hiçbir kanal yoksa, varsayılan durum yürütülür.
- Varsayılan durum yoksa select kilitlenir.

Son olarak, bir `for` içinde `select` kulllanmak çok yaygındır. Örneğin:

```go
for {
  select {
  case data := <-c:
    fmt.Printf("worker %d got %d\n", w.id, data)
  case <-time.After(time.Millisecond * 10):
    fmt.Println("Break time")
    time.Sleep(time.Second)
  }
}
```

## Devam Etmeden Önce

Eşzamanlı programlama dünyasında yeniyseniz, hepsi bir anda oldukça zor görünebilir. Kategorik olarak çok daha fazla dikkat ve özen gerektirir. Go bunu her aşamada kolaylaştırmayı amaçlıyor.

Go rutinler, eşzamanlı kodu çalıştırmak için gerekenleri etkili bir şekilde soyutlar. Kanallar, veri paylaşımını ortadan kaldırarak veri paylaşıldığında meydana gelebilecek bazı ciddi hataların giderilmesine yardımcı olur. Bu sadece hataları ortadan kaldırmaz, aynı zamanda eşzamanlı programlamaya yaklaşımını değiştirir. Kodun sorun çıkarıcı alanlarından ziyade mesaj geçişi ile ilgili eşzamanlılığı düşünmeye başlarsınız.

Bunu söyledikten sonra, `sync` ve `sync/atomic` paketlerde bulunan çeşitli senkronizasyon ilkellerinden hala geniş çapta faydalanıyorum. Her ikisiyle de rahat olmanın önemli olduğunu düşünüyorum. Öncelikle kanallara odaklanmanızı öneririm, ancak kısa ömürlü bir kilit gerektiren basit bir örnek gördüğünüzde, bir mutex veya okuma-yazma mutex'i kullanmayı düşünün.

