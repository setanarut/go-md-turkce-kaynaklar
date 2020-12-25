- [Bölüm 1 - Temel Bilgiler](#bölüm-1---temel-bilgiler)
  - [Derleme](#derleme)
  - [Statik Tip Özelliğine Sahip Olma](#statik-tip-özelliğine-sahip-olma)
  - [C Benzeri Sözdizimi](#c-benzeri-sözdizimi)
  - [Toplanan Çöpler](#toplanan-çöpler)
  - [Go Kodunu Çalıştırma](#go-kodunu-çalıştırma)
    - [Main](#main)
  - [Import](#import)
  - [Değişkenler ve Tanımlamaları](#değişkenler-ve-tanımlamaları)
  - [İşlev Tanımlamaları](#i̇şlev-tanımlamaları)
  - [Devam Etmeden Önce](#devam-etmeden-önce)

# Bölüm 1 - Temel Bilgiler

Go, C benzeri bir sözdizimi ve çöp toplama özelliğine sahip derlenmiş, değişken tipi zorunlu olarak yazılmış bir dildir. Bu ne anlama geliyor?

## Derleme

Derleme, yazdığınız kaynak kodunu daha düşük düzeyli bir dile dönüştürme işlemidir -- örneğin assembly diline (Go'da olduğu gibi) veya aracı bir dile (Java ve C # gibi).

Derlenen diller ile çalışmak zevkli olmayabilir, çünkü derleme yavaş olabilir. Kodun derlenmesini beklemek için dakikalar veya saatler harcamanız gerekiyorsa hızlı bir şekilde kod değişikliği yapmak zordur. Derleme hızı Go'nun ana tasarım hedeflerinden biridir. Bu, büyük projeler üzerinde çalışan insanlar için olduğu kadar, yorumlanan diller tarafından sunulan hızlı bir geri bildirim döngüsüne alışkın olanlar için iyi bir haber.

Derlenmiş diller daha hızlı çalışma eğilimindedir ve derlenmiş dosyalar ek bağımlılıklar olmadan çalıştırılabilir (bu en azından doğrudan derlenen C, C ++ ve Go gibi diller için geçerlidir).

## Statik Tip Özelliğine Sahip Olma

Statik tipli olmak, değişkenlerin belirli bir tipte (int, string, bool, [] byte, vb.) olması gerektiği anlamına gelir. Bu, değişken bildirildiğinde tip belirtilerek veya birçok durumda derleyicinin tipi çıkarmasına izin verilerek elde edilir (kısaca örneklere bakacağız).

Statik tip ile ilgili söylenebilecek çok daha fazla şey var, ancak bunun kodlara bakarak daha iyi anlaşılacağına inanıyorum. Dinamik olarak yazılan dillere alışkınsanız, bunu hantal bulabilirsiniz. Yanlış değilsiniz, ancak özellikle statik tip ile yazmayı derleme ile eşleştirdiğinizde bir çok avantaj ortaya çıkar. Bu ikisi genellikle birbiriyle bağlantılıdır. Birine sahip olduğunuzda, normalde diğerine sahip olduğunuz doğru ama bu zor bir kural değil. Katı tip bir sistemle, bir derleyici sadece sözdizimsel hataların ötesinde problemleri tespit edebilir ve daha fazla optimizasyon sağlayabilir.

## C Benzeri Sözdizimi

Bir dilin C benzeri bir sözdizimine sahip olduğunu söylemek, C, C ++, Java, JavaScript ve C # gibi diğer C benzeri dillere alışkınsanız, Go'yu en azından yüseysel olarak tanıdık bulacaksınız demektir. Örneğin, `&&` bir boolean AND olarak kullanılır, `==` eşitliği karşılaştırmak için kullanılır, `{` ve `}` bir kapsamı başlatır ve sonlandırır ve dizi indeksleri de 0'dan başlar.

C-benzeri sözdizimi aynı zamanda  satır sonlarında noktalı virgüller ve koşullar etrafında parantezler anlamına gelir. Go, her ikisini de ortadan kaldırır, ancak yine de önceliği kontrol etmek için parantez kullanılır. Örneğin, bir `if` ifadesi şöyle olabilir:

```go
if name == "Leto" {
  print("the spice must flow")
}
```

Ve daha karmaşık durumlarda parantezler hala yararlıdır:

```go
if (name == "Goku" && power > 9000) || (name == "gohan" && power < 4000)  {
  print("super Saiyan")
}
```

Bunun ötesinde, Go C'ye sadece sözdizimi açısından değil, amaç açısından da C# veya Java'dan çok daha yakındır. Bu dilin açıklığına ve basitliğine yansır ve öğrenirken umarız size daha belirgin olmaya başlayacak.

## Toplanan Çöpler

Bazı değişkenler yaratıldıklarında tanımlanması kolay bir ömre sahiptir. Örneğin, bir işlevin (fonksiyon) yerel değişkeni, işlevden çıkıldığında kaybolur. Diğer durumlar, en azından bir derleyici için o kadar açık değildir. Örneğin, bir işlev tarafından döndürülen veya diğer değişkenler ve nesneler tarafından başvurulan bir değişkenin ömrünü belirlemek zor olabilir. Çöp toplama olmadan, değişkenin gerekli olmadığını bildiği bir noktada bu değişkenlerle ilişkili belleği boşaltmak geliştiricilere kalmıştır. Nasıl? C dilinde, tam olarak `free(str);` kullanarak.

Çöp toplayıcıları olan diller (ör. Ruby, Python, Java, JavaScript, C #, Go) bunları takip edebilir ve artık kullanılmadıklarında onları temizleyebilir. Çöp toplama yükü arttırır, ancak aynı zamanda bir dizi yıkıcı hatayı da ortadan kaldırır.

## Go Kodunu Çalıştırma

Basit bir program oluşturarak ve onu nasıl derleyeceğimizi ve çalıştıracağımızı öğrenerek yolculuğumuza başlayalım. En sevdiğiniz metin düzenleyicisini açın ve aşağıdaki kodu yazın:

```go
package main

func main() {
  println("it's over 9000!")
}
```

Dosyayı `main.go` olarak kaydedin. Şimdilik istediğiniz yere kaydedebilirsiniz; önemsiz örnekler için Go'nun ana çalışma klasöründe olmaya ihtiyacımız yok.

Ardından, bir kabuk ya da komut istemi açın ve dosyayı kaydettiğiniz dizini geçiş yapın. Benim için bu, `cd ~/code` yazmak anlamına geliyor.

Son olarak, programı aşağıdaki komutu girerek çalıştırın:

```
go run main.go
```

Eğer herşey sorunsuz çalışırsa, ekranda *it's over 9000!* yazısını görmeniz lazım.

Ama bekleyin, derleme adımı ne olacak? `go run` , kodunuzu derleyen *ve* çalıştıran kullanışlı bir komuttur. Programı oluşturmak için geçici bir dizin kullanır, çalıştırır ve sonra kendini temizler. Geçici dosyanın konumunu aşağıdaki komutu çalıştırarak görebilirsiniz:

```
go run --work main.go
```

Kodu sadece derlemek için `go build` komutunu kullanın:

```
go build main.go
```

Bu, çalıştırabileceğiniz yürütülebilir bir `main` dosyası oluşturur. Linux/OSX'te, yürütülebilir dosyaya nokta eğik çizgi ile ön ek eklemeniz gerektiğini unutmayın, bu nedenle `./main` yazmanız gerekir.

Kodu geliştirmeye devam ederkeni ya `go run` ya da `go build` kullanabilirsiniz. Kodunuzu sonlandırıp yüklemek istediğinizde çalıştırılabilir bir halini `go build` ile oluşturabilirsiniz.

### Main

Şanslıyız çünkü yeni çalıştırdığımız kod oldukça anlaşılabilir. Bir işlev oluşturduk ve yerleşik `println` işleviyle bir cümle yazdırdık. Sadece tek bir seçim olduğu için mi `go run` neyi yürüteceğini biliyor? Hayır. Go'da bir programın giriş noktası olarak, `main` paketinde `main` adlı bir işlev olmalıdır.

Daha sonraki bir bölümde paketler hakkında daha fazla konuşacağız. Şimdilik, Go'nun temellerini anlamaya odaklanırken, kodumuzu her zaman `main` pakete yazacağız.

Eğer isterseniz, kodu paket ismini değiştirip güncelleyin. Sonra `go run` komutu ile çalıştırın, mutlaka bir hata ile karşılaşacaksınız. Sonra paket ismini tekrar `main` yapın ve bu kez farklı bir işlev adı kullanın. Bu kez de başka bir hata ile karşılaşacaksınız. Aynı değişimleri yapıp aynı denemeleri `go build` ile yapın. Bu kez derlemenin başarılı olduğunu göreceksiniz, çünkü kod bu şekilde çalıştırılmadı. Bu tarz çalışma (main isminde paket ya da fonksiyon olmadan) eğer bir kütüphane yazıyorsanız doğru bir çalışma şeklidir.

## Import

Go'nun referans olmadan kullanılabilen `println` gibi bir dizi yerleşik işlevi vardır. Go'nun standart kitaplığını ve üçüncü taraf kitaplıkları kullanmadan çok ileri gidemeyiz. Go'da, `import` anahtar sözcüğü, dosyadaki kod tarafından kullanılan paketleri bildirmek için kullanılır.

Örnek programımızı biraz değiştirelim:

```go
package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) != 2 {
    os.Exit(1)
  }
  fmt.Println("It's over", os.Args[1])
}
```

Aşağıdaki gibi çalıştırabilirsiniz:

```
go run main.go 9000
```

Şimdi Go'nun iki standart paketini kullanıyoruz: `fmt` ve `os` . Ayrıca başka bir yerleşik işlev olan `len`  kullandık. `len` bir dizenin boyutunu veya metindeki harflerin sayısını veya burada gördüğümüz gibi bir dizideki öğe sayısını döndürür. Neden 2 argüman beklediğimizi merak ediyorsanız, bunun nedeni ilk argümanın (0 dizinindeki) her zaman şu anda çalışan yürütülebilir dosyanın yolu olmasıdır. (Yazdırmak için programı değiştirin ve kendiniz görün.)

Muhtemelen fonksiyon isminin önüne paket ismini `fmt.Println` şeklinde eklediğimizi fark etmişsinizdir. Bu, diğer birçok dilden farklıdır. Daha sonraki bölümlerde paketler hakkında daha fazla bilgi edineceğiz. Şimdilik, bir paketi nasıl içe aktaracağınızı ve kullanacağınızı bilmek için bu iyi bir başlangıçtır.

Go, paketleri içe aktarma konusunda katıdır. Bir paketi içe aktarırsanız ancak kullanmazsanız derlemez. Aşağıdaki kodu çalıştırmayı deneyin:

```go
package main

import (
  "fmt"
  "os"
)

func main() {
}
```

İçeri aktarılan ama kullanılmayan `fmt` ve `os` hakkında iki hata almalısınız. Bu can sıkıcı olabilir mi? Kesinlikle. Zamanla, buna alışacaksınız (yine de sinir bozucu olacak). Go bu konuda katıdır çünkü kullanılmayan paketler derlemeyi yavaşlatabilir; Kuşkusuz ki bir çoğumuzun bu başlangıç seviyesinde bu tarz problemlemimiz yok.

Dikkat edilmesi gereken başka bir şey, Go'nun standart kütüphanesinin iyi belgelendirilmiş olmasıdır. Kullandığımız `Println` işlevi hakkında daha fazla bilgi edinmek için [https://golang.org/pkg/fmt/#Println](https://golang.org/pkg/fmt/#Println) adresine gidebilirsiniz. Bölüm başlığını tıklayabilir ve kaynak kodunu görebilirsiniz. Ayrıca, Go'nun biçimlendirme özellikleri hakkında daha fazla bilgi edinmek için en üste seviyeye gidin.{/code1}

İnternet erişimi olmadan sıkışıp kalırsanız, belgeleri yerel olarak da görebilirsiniz:

```
godoc -http=:6060
```

ve tarayıcınızda `http://localhost:6060` adresini ziyaret edin

## Değişkenler ve Tanımlamaları

Değişkenler bölümünü *&nbsp;tamımlamayı ve değer atamayı x = 4 yaparak sağladığımızı* söyleyemek güzel olurdu. Maalesef Go'da işler biraz daha karmaşık. Konuşmamıza basit örneklere bakarak başlayacağız. Ardından, bir sonraki bölümde, struct oluşturmaya ve kullanmaya baktığımızda bunu genişleteceğiz. Yine de, bu konuda kendinizi gerçekten rahat hissetmeniz biraz zaman alacaktır.

*Yaaa! Bu konu ne kadar karmaşık olabilir ki?* diye düşünüyor olabilirsiniz. Örneklere bakmaya devam edelim.

Go'da değişken tanımlama ve değer atama ile başa çıkmanın en açık yolu da en ayrıntılı olanıdır:

```go
package main

import (
  "fmt"
)

func main() {
  var power int
  power = 9000
  fmt.Printf("It's over %d\n", power)
}
```

Burada `int` tipinde bir `power` değişkeni tanımlıyoruz. Varsayılan olarak, Go değişkenlere sıfır değeri atar. Tamsayılar için `0` , boolean için `false` , dizeler `""` vb. atanır. Ardından, `power` değişkenimize `9000` değeri atarız. İlk iki satırı birleştirebiliriz:

```go
var power int = 9000
```

Yine de, bu çok fazla olarak düşünülebilir. Go'nun kullanışlı kısa ve değişken tipini tahmin eden değer atama operatörü `:=` vardır:

```go
power := 9000
```

Bu kullanışlıdır ve işlevlerle de çalışır:

```go
func main() {
  power := getPower()
}

func getPower() int {
  return 9001
}
```

`:=` değişkeni tanımlamak ve değişkene bir değer atamak için kullanılır. Neden? Çünkü bir değişken iki kez tanımlanamaz (aynı kapsamda). Aşağıdaki kodu çalıştırmayı denerseniz bir hata alırsınız.

```go
func main() {
  power := 9000
  fmt.Printf("It's over %d\n", power)

  // Derleme Hatası:
  // := ile ancak yeni bir değişken kullanabilirsiniz
  power := 9001
  fmt.Printf("It's also over %d\n", power)
}
```

Derleyici *no new variables on left side of :=* mesajı ile şikayet edecektir. Bu, bir değişkeni ilk bildirdiğimizde `:=` kullanabileceğimiz, ancak sonraki atamada `=` atama operatörünü kullanmamız gerektiği anlamına gelir. Bu çok mantıklı, ancak kas hafızanızın ikisi arasında ne zaman geçiş yapacağını hatırlaması zor olabilir.

Hata mesajını dikkatli okursanız, *variables* kelimesinin çoğul olduğunu fark edeceksiniz. Çünkü Go aynı anda birden fazla değişken atamanıza izin verir ( `=` veya `:=` kullanarak):

```go
func main() {
  name, power := "Goku", 9000
  fmt.Printf("%s's power is over %d\n", name, power)
}
```

Değişkenlerden biri yeni olduğu sürece `:=` kullanılabilir. Örneğin:

```go
func main() {
  power := 1000
  fmt.Printf("default power is %d\n", power)

  name, power := "Goku", 9000
  fmt.Printf("%s's power is over %d\n", name, power)
}
```

`power` `:=` ile iki kez kullanılıyor olsa da, derleyici ikinci kez kullandığımızda şikayet etmeyecek, diğer değişkenin `name` yeni bir değişken olduğunu görecek ve `:=` kullanılmasına izin verecek. Ancak, `power` türünü değiştiremezsiniz. Bir tamsayı olarak tanımlandı ve bu nedenle yalnızca tamsayı değerler atanabilir.

Şimdilik, bilinmesi gereken son şey, içe aktarma gibi Go'nun kullanılmayan değişkenlere sahip olmanıza izin vermeyeceğidir. Örneğin,

```go
func main() {
  name, power := "Goku", 1000
  fmt.Printf("default power is %d\n", power)
}
```

`name` bildirildiği, ancak kullanılmadığı için derlenmeyecektir. Kullanılmayan paket içe aktarımları gibi bazı hatalara neden olur, ancak genel olarak kod temizliği ve okunabilirliğe yardımcı olduğunu düşünüyorum.

Tanımlama ve değer atama hakkında öğrenilecek daha çok şey var. Şimdilik, bir değişkeni tanımlayıp sıfır değerine eşitlerken `var NAME TYPE`, bir değer atayarak tanımlarken `NAME := VALUE` ve daha önce tanımlanan bir değişkene değer atarken `NAME = VALUE` kullanacağınızı unutmayın.

## İşlev Tanımlamaları

Bu, işlevlerin birden fazla değer döndürebileceğini belirtmek için iyi bir zamandır. Üç işleve bakalım: biri dönüş değeri olmayan, biri dönüş değeri olan ve diğeri iki dönüş değeri olan.

```go
func log(message string) {
}

func add(a int, b int) int {
}

func power(name string) (int, bool) {
}
```

Dönen değerlerin sonuncunu şu şekilde kullanabiliriz:

```go
value, exists := power("goku")
if exists == false {
  // handle this error case
}
```

Bazen, dönüş değerlerinden yalnızca birini önemsersiniz. Bu durumlarda, diğer değerleri `_` öğesine atarsınız:

```go
_, exists := power("goku")
if exists == false {
  // handle this error case
}
```

Bu bir tanımlamadan daha fazlasıdır. Boş tanımlayıcı olan `_` , dönüş değerinin gerçekten atanmamış olması nedeniyle özeldir. Bu `_` döndürülen ne türde olursa olsun tekrar tekrar kullanmanızı sağlar.

Son olarak, işlev bildirimleriyle karşılaşmanız muhtemel başka bir şey daha var. Parametreler aynı türü paylaşıyorsa, daha kısa bir sözdizimi kullanabiliriz:

```go
func add(a, b int) int {

}
```

Birden çok değer döndürmek sık kullandığınız bir şeydir. Ayrıca, bir değeri silmek için `_` sık sık kullanırsınız. Adlandırılmış dönüş değerleri ve biraz daha az ayrıntılı parametre bildirimi o kadar yaygın değildir. Yine de, tüm bunlara er geç başlayacaksınız, bu yüzden onları bilmek önemlidir.

## Devam Etmeden Önce

Birkaç küçük parçaya baktık ve muhtemelen bu noktada birbiri ile alakasız olduğunu düşündürebilir. Yavaşça daha büyük örnekler yapacağız ve umarım parçalar bir araya gelmeye başlayacaktır.

Dinamik bir dilden geliyorsanız, tipler ve tanımlamalar arasındaki karmaşıklık geriye doğru bir adım gibi görünebilir. Size katılmıyorum. Bazı sistemler için dinamik diller kategorik olarak daha verimlidir.

Statik tipli bir dilden geliyorsanız, muhtemelen Go ile kendinizi rahat hissedersiniz. Değerden tip belirleme ve çoklu dönüş değerleri güzel görünebilir (kesinlikle Go'ya özel olmasa da). Umarım daha fazlasını öğrendikçe, temiz ve kısa söz dizimini takdir edersiniz.

