- [Bölüm 4 - Kod Organizasyonu ve Arayüzler](#bölüm-4---kod-organizasyonu-ve-arayüzler)
  - [Paketler](#paketler)
    - [Döngüsel İçe Aktarım](#döngüsel-i̇çe-aktarım)
    - [Görünürlük](#görünürlük)
    - [Paket Yönetimi](#paket-yönetimi)
    - [Bağımlılık Yönetimi](#bağımlılık-yönetimi)
  - [Arayüzler](#arayüzler)
  - [Devam Etmeden Önce](#devam-etmeden-önce)

# Bölüm 4 - Kod Organizasyonu ve Arayüzler

Şimdi kodumuzu nasıl düzenleyeceğimize bakmanın zamanı geldi.

## Paketler

Daha karmaşık kütüphaneleri ve sistemleri organize tutmak için paketler hakkında bilgi edinmeliyiz. Go'da paket adları, Go ana çalışma alanınızın dizin yapısını izler. Bir alışveriş sistemi inşa ediyor olsaydık, muhtemelen "shopping" adlı bir paket adıyla başlayıp kaynak dosyalarımızı `$GOPATH/src/shopping/` altına koyacağız.

Yine de her şeyi bu klasörün içine koymak istemeyebiliriz. Örneğin, veritabanı işlemlerini kendi klasörü içinde izole etmek isteriz. Bunu sağlamak için `$GOPATH/src/shopping/db` adıyla bir alt klasör oluştururuz. Bu alt klasördeki dosyaların paket adı sadece `db` 'dir, ancak `shopping` paketi de dahil olmak üzere başka bir paketten erişmek için `shopping/db` dosyasını içe aktarmamız gerekir.

Diğer bir deyişle, bir paketi adlandırdığınızda, `package` anahtar sözcüğü aracılığıyla, tam bir hiyerarşi (örneğin, "alışveriş" veya "db") değil, tek bir isim verirsiniz. Bir paketi içe aktarırken ise tam yolu belirtirsiniz.

Hadi deneyelim. Go ana çalışma alanınızın `src` klasörünün (Giriş bölümünde ayarladığımız) içinde, `shopping` adı verilen yeni bir klasör ve onun da içinde `db` adı verilen bir alt klasör oluşturun.

`shopping/db` içinde `db.go` adlı bir dosya oluşturun ve aşağıdaki kodu ekleyin:

```go
package db

type Item struct {
  Price float64
}

func LoadItem(id int) *Item {
  return &Item{
    Price: 9.001,
  }
}
```

Paket adının klasörün adıyla aynı olduğuna dikkat edin. Ayrıca, açıkçası, kodun içinde aslında veritabanına erişmiyoruz. Bunu sadece kodun nasıl düzenleneceğini göstermek için örnek olarak kullanıyoruz.

Şimdi, ana `shopping` klasörünün içinde `pricecheck.go` adlı bir dosya oluşturun. İçeriği aşağıdaki gibi olsun:

```go
package shopping

import (
  "shopping/db"
)

func PriceCheck(itemId int) (float64, bool) {
  item := db.LoadItem(itemId)
  if item == nil {
    return 0, false
  }
  return item.Price, true
}
```

`shopping/db` içe aktarmanın bir şekilde özel olduğunu düşünmek cazip gelebilir çünkü zaten `shopping{/ code1} paketinin klasörünün içindeyiz. Gerçekte, <code data-md-type="codespan" data-parent-segment-tag-id="5029777">$GOPATH/src/shopping/db` klasöründen içe aktarıyorsunuz, yani çalışma alanınızın `src/test` klasörünün içindeki `db` adlı bir paketiniz varsa `test/db` şeklinde çağırabilirsiniz.

Bir paket oluşturuyorsanız, gördüğümüzden daha fazlasına ihtiyacınız yok. Yürütülebilir bir dosya oluşturmak için yine de bir `main` paketine ve işlevine ihtiyacınız vardır. Bunu yapmayı tercih ettiğim yol, `main.go` adlı bir dosya ve aşağıdaki içerikle `shopping` içinde `main` adı verilen bir alt klasör oluşturmaktır:

```go
package main

import (
  "shopping"
  "fmt"
)

func main() {
  fmt.Println(shopping.PriceCheck(4343))
}
```

Artık kodunuzu `shopping` projenize girip şunu yazarak çalıştırabilirsiniz:

```
go run main/main.go
```

### Döngüsel İçe Aktarım

Daha karmaşık sistemler yazmaya başladığınızda, döngüsel içe aktarma işlemi ile karşılaşabilirsiniz. Bu, A paketi B paketini içe aktarırken B paketi A paketini (doğrudan veya dolaylı olarak başka bir paket üzerinden) aldığında olur. Bu derleyicinin izin vermeyeceği bir şeydir.

Hataya neden olmak için alışveriş yapımızı değiştirelim.

`Item` tanımını `shopping/db/db.go` dosyasından `shopping/pricecheck.go` dosyasına taşıyın. `pricecheck.go` dosyanız şimdi şöyle görünmelidir:

```go
package shopping

import (
  "shopping/db"
)

type Item struct {
  Price float64
}

func PriceCheck(itemId int) (float64, bool) {
  item := db.LoadItem(itemId)
  if item == nil {
    return 0, false
  }
  return item.Price, true
}
```

Kodu çalıştırmayı denerseniz, `db/db.go` dosyasından `Item` tanımlı değil diye hata alırsınız. Bu mantıklı. `Item` artık `db` paketinde mevcut değil; shopping paketine taşındı. `shopping/db/db.go` şu şekilde değiştirmemiz gerekir:

```go
package db

import (
  "shopping"
)

func LoadItem(id int) *shopping.Item {
  return &shopping.Item{
    Price: 9.001,
  }
}
```

Şimdi kodu çalıştırmaya çalıştığınızda, korkunç bir *import cycle not allowed*yani "döngüsel içeri aktarıma izin verilmiyor" hatası alırsınız. Bunu, paylaşılan yapıları içeren başka bir paket oluştururak çözüyoruz. Dizin yapınız şimdi şu şekilde olmalıdır:

```
$GOPATH/src
  - shopping
    pricecheck.go
    - db
      db.go
    - models
      item.go
    - main
      main.go
```

`pricecheck.go` hala `shopping/db` paketini kullanacak, `db.go` `shopping` yerine `shopping/models` kullanacak ve bu da döngüyü kıracak. Paylaşılan `Item` yapısını `shopping/models/item.go` 'a taşıdığımızdan, referans olarak `models` paketindeki `Item` yapısını kullanması için `shopping/db/db.go`'yi değiştirmemiz gerekiyor:

```go
package db

import (
  "shopping/models"
)

func LoadItem(id int) *models.Item {
  return &models.Item{
    Price: 9.001,
  }
}
```

Genellikle `models`'ten daha fazlasını ortak olarak kullanmanız gerekir, bu nedenle `utilities` ve benzeri adlı başka klasörleriniz olabilir. Bu paylaşılan paketlerle ilgili önemli kural, `shopping` paketinden veya alt paketlerden hiçbir şey almamalarıdır. Birkaç bölümde, bu tür bağımlılıkları çözmemize yardımcı olabilecek arayüzlere bakacağız.

### Görünürlük

Go, bir paketin dışında hangi tiplerin ve işlevlerin görünür olacağını tanımlamak için basit bir kural kullanır. Tip veya işlevin adı büyük harfle başlıyorsa görünürdür. Küçük bir harfle başlıyorsa değildir.

Bu aynı zamanda yapı alanları için de geçerlidir. Bir yapı alanı adı küçük harfle başlıyorsa, yalnızca aynı paket içindeki kod bunlara erişebilir.

Örneğin, eğer `items.go` dosyası aşağıdaki gibi bir işleve sahip ise:

```go
func NewItem() *Item {
  // ...
}
```

`model.NewItem()` şeklinde çağrılabilir. Ancak işlev `newItem` olarak adlandırılsaydı, işleve farklı bir paketten erişemezdik.

Devam edin ve `shopping` kodundan çeşitli işlevlerin, türlerin ve alanların adını değiştirin. Örneğin, `Item` yapısının `Price` alanını `price` diye yeniden adlandırırsanız, bir hata almanız gerekir.

### Paket Yönetimi

`run` ve `build` için kullandığımız `go` komutunun, üçüncü taraf kitaplıklarını getirmek için kullanılan bir `get` alt komutu vardır. `go get` çeşitli protokolleri destekler ancak bu örnek için Github bir kütüphane almaya çalışacağız ve bunun için `git`'in bilgisayarınızda yüklü olması gerekiyor.

Git'in kurulu olduğunu varsayarsak, bir kabuk / komut isteminden şunu girin:

```
go get github.com/mattn/go-sqlite3
```

`go get` , uzak dosyaları getirir ve bunları çalışma alanınızda uygun klasörlerde saklar. Anlamak için `$GOPATH/src` kontrol edin. Oluşturduğumuz `shopping` projesine ek olarak, artık bir `github.com` klasörü göreceksiniz. İçinde bir `go-sqlite3` klasörü içeren bir `mattn` klasörü göreceksiniz.

Çalışma alanımızda olan paketlerin nasıl içe aktarılacağından çoktan bahsettik. Yeni indirdiğimiz `go-sqlite3` paketimizi kullanmak için şu şekilde içe aktarırız:

```go
import (
  "github.com/mattn/go-sqlite3"
)
```

Bunun bir URL gibi göründüğünü biliyorum, ama aslında olan `go-sqlite3` paketini `$GOPATH/src/github.com/mattn/go-sqlite3` klasöründen alıyoruz.

### Bağımlılık Yönetimi

`go get` komutunun birkaç hilesi daha vardır. Eğer projemizde `go get` çalıştırırsak, projedeki tüm dosyaları tarar `imports` içeri aktarılan üçüncü parti kütüphaneleri bulup ve bunları indirecektir. Bir bakıma, kendi kaynak kodumuz `Gemfile` veya `package.json` dosyaları gibi kullanılır.

`go get -u` çağırırsanız, paketler güncellenir (veya `go get -u FULL_PACKAGE_NAME` yoluyla belirli bir paketi güncelleyebilirsiniz).

Bazı durumlar `go get` komutunu yetersiz bulabilirsiniz. Birincisi, bir revizyon belirtmenin bir yolu yoktur, her zaman master/head/trunk/default'u alır. Aynı kütüphanenin farklı sürümlerine ihtiyaç duyan iki projeniz varsa, bu daha da büyük bir sorundur.

Bunu çözmek için üçüncü taraf bir bağımlılık yönetimi aracı kullanabilirsiniz. Hala gelişmekteler, ancak umut vaat eden iki proje [goop](https://github.com/nitrous-io/goop) ve [godep](https://github.com/tools/godep) projeleridir . [Go-wiki](https://code.google.com/p/go-wiki/wiki/PackageManagementTools)'de daha eksiksiz bir liste bulunmaktadır.

## Arayüzler

Arayüzler, bir sözleşmeyi tanımlayan ancak bir uygulaması olmayan tiplerdir. İşte bir örnek:

```go
type Logger interface {
  Log(message string)
}
```

Bunun hangi amaca hizmet edebileceğini merak ediyor olabilirsiniz. Arayüzler, kodunuzu belirli uygulamalardan ayırmanıza yardımcı olur. Örneğin, çeşitli log şekillerimiz olabilir:

```go
type SqlLogger struct { ... }
type ConsoleLogger struct { ... }
type FileLogger struct { ... }
```

Yine de, bu somut uygulamalardan ziyade arayüze ile programlayarak, kodumuzu kolaylıkla değiştirebilir (ve test edebiliriz).

Nasıl kullanırız? Tıpkı diğer tipler gibi, bir yapının alanı olabilir:

```go
type Server struct {
  logger Logger
}
```

veya bir işlev parametresi (veya dönüş değeri):

```go
func process(logger Logger) {
  logger.Log("hello!")
}
```

C # veya Java gibi bir dilde, bir sınıf bir arayüz uyguladığında açıkça tanımlanmalıdır:

```go
public class ConsoleLogger : Logger {
  public void Logger(message string) {
    Console.WriteLine(message)
  }
}
```

Go'da bu dolaylı olarak gerçekleşir. Yapınızın bir `string` parametresi olan ve dönüş değeri olmayan bir `Log` adında fonksiyonu varsa, o zaman `Logger` olarak kullanılabilir. Bu, arayüzleri kullanmanın ayrıntı düzeyini azaltır:

```go
type ConsoleLogger struct {}
func (l ConsoleLogger) Log(message string) {
  fmt.Println(message)
}
```

KDilin kendisi küçük ve odaklanmış arayüzleri teşvik etme eğilimindedir. Standart kütüphane arayüzlerle doludur. `io` paketinde `io.Reader`,`io.Writer` ve `io.Closer` gibi yaygın kullanılan arayüzler vardır. Yalnızca `Close()` işlevini çağıracağınız bir parametreyi bekleyen bir işlev yazarsanız, kesinlikle kendi tanımladığınız bir tipte bir yapı yerine bir `io.Closer` kabul etmelisiniz.

Arayüzler içermelerde de kullanılıyor. Ve arayüzlerin kendileri de diğer arayüzlerden oluşabilir. Örneğin, `io.ReadCloser`, `io.Reader` arayüzünün yanı sıra ` io.Closer ` arayüzünden de oluşan bir arayüzdür.

Son olarak, arayüzler döngüsel içe aktarımları önlemek için yaygın olarak kullanılır. Uygulamaları olmadığı için sınırlı bağımlılıkları olacaktır.

## Devam Etmeden Önce

Sonuçta, kodunuzu Go'nun çalışma alanı etrafında nasıl yapılandırdığınız, yalnızca birkaç örnek proje yazdıktan sonra kendinizi rahat hissedeceğiniz bir şeydir. Hatırlamanız gereken en önemli şey, paket adları ile dizin yapınız arasındaki sıkı ilişkidir (sadece bir proje içinde değil, tüm çalışma alanı içinde).

Go'nun tiplerin görünürlüğünü işleme şekli basit ve etkilidir. Aynı zamanda tutarlıdır. Sabitler ve global değişkenler gibi bakmadığımız birkaç şey var, ancak emin olabilirsiniz ki, görünürlükleri aynı adlandırma kuralıyla belirlenir.

Son olarak, arayüzlerde yeniyseniz, onları tam olarak anlamanız biraz zaman alabilir. Ancak, ilk kez `io.Reader` gibi bir tip bekleyen bir işlev gördüğünüzde, yazara ihtiyaç duyduğundan fazlasını talep etmediği için teşekkür edersiniz.

