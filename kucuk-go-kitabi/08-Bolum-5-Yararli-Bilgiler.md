- [Bölüm 5 - Yararlı Bilgiler](#bölüm-5---yararlı-bilgiler)
  - [Hata Yönetimi](#hata-yönetimi)
  - [Erteleme](#erteleme)
  - [go fmt](#go-fmt)
  - [Değer Atamalı If](#değer-atamalı-if)
  - [Boş Arayüz ve Dönüşümler](#boş-arayüz-ve-dönüşümler)
  - [Dizeler ve Bayt Dizileri](#dizeler-ve-bayt-dizileri)
  - [İşlev Tipi](#i̇şlev-tipi)
  - [Devam Etmeden Önce](#devam-etmeden-önce)

# Bölüm 5 - Yararlı Bilgiler

Bu bölümde, Go'nun başka hiçbir yere tam olarak uymayan bazı özellikleri hakkında konuşacağız.

## Hata Yönetimi

Go'nun hatalarla başa çıkmanın tercih edilen yolu istisna fırlatarak değil, dönüş değerleri ile olur. Bir dize alan ve onu bir tamsayıya dönüştürmeye çalışan `strconv.Atoi` işlevini örnek alalım:

```go
package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) != 2 {
    os.Exit(1)
  }

  n, err := strconv.Atoi(os.Args[1])
  if err != nil {
    fmt.Println("not a valid number")
  } else {
    fmt.Println(n)
  }
}
```

Kendi hata türünüzü oluşturabilirsiniz; tek gereklilik, yerleşik `error` arayüzünün sözleşmesini yerine getirmesidir:

```go
type error interface {
  Error() string
}
```

Daha yaygın olarak, `errors` paketini içe aktararak ve `New` işlevini kullanarak kendi hatalarımızı oluşturabiliriz:

```go
import (
  "errors"
)


func process(count int) error {
  if count < 1 {
    return errors.New("Invalid count")
  }
  ...
  return nil
}
```

Go'nun standart kitaplığında hata değişkenlerini kullanma konusunda yaygın bir alışkanlık vardır. Örneğin, `io` paketi şu şekilde tanımlanan bir `EOF` değişkenine sahiptir:

```go
var EOF = errors.New("EOF")
```

Bu, herkes tarafından erişilebilen (ilk harf büyük harf  olduğu için) bir paket değişkenidir (bir işlev dışında tanımlanır). Bir dosyadan veya STDIN'den okurken çeşitli işlevler bu hatayı döndürebilir. Bağlamsal olarak mantıklıysa, bu hatayı da kullanmalısınız. Geliştiriciler olarak bu singletonu kullanabiliriz:

```go
package main

import (
  "fmt"
  "io"
)

func main() {
  var input int
  _, err := fmt.Scan(&input)
  if err == io.EOF {
    fmt.Println("no more input!")
  }
}
```

Son bir not olarak, Go `panic` ve `recover` işlevlerine de sahiptir. `panic` bir istisna atmaya benzerken `recover` ise `catch` kullanımına benzer ve  nadiren kullanılırlar.

## Erteleme

Go'nun çöp toplayıcısı olmasına rağmen, bazı kaynaklar açıkça onları serbest bırakmamızı gerektirir. Örneğin, dosyaları okumayı bitirdikten sonra `Close()` işlevi ile kapatmamız gerekir. Bu tür kodlar her zaman tehlikelidir. Biz bir işlev yazıyoruz ve 10 satır kadar sonra `Close` yazmayı unutmak kolaydır. Bir diğeri sorun da, bir fonksiyonun birden fazla dönüş noktası olabilir. Go'nun buna çözümü `defer` anahtar kelimesidir:

```go
package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("a_file_to_read")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()
  // read the file
}
```

Yukarıdaki kodu çalıştırmayı denerseniz, muhtemelen bir hata alırsınız (the file doesn't exist). Bu örneği `defer` anahtarının nasıl çalıştığını göstermek için kullandık. `defer` işlev (bu durumda `main()` ) geri dönme işlemini yaptıktan sonra yürütülür. Bu şu anlama geliyor, kaynakların kullanılmaya başlatıldığı yerin yakınında defer ile kapanışı belirtmenizi sağlar ve birden fazla dönüş noktasını da kapsamış olur.

## go fmt

Go'da yazılan çoğu program aynı biçimlendirme kurallarına uyar, yani girinti için bir sekme kullanılır ve parantezler ifadeleriyle aynı satıra gider.

Biliyorum, kendi kod yazma tarzınız var ve ona bağlı kalmak istiyorsiniz. Uzun zamandır yaptığım şey bu, ama sonunda vazgeçmiş olduğuma sevindiğimi söyleyebilirim. Bunun en büyük nedeni `go fmt` komutudur. Kullanımı kolay ve etkilir (bu yüzden hiç kimse anlamsız tercihleri tartışmıyor).

Bir projenin içindeyken, biçimlendirme kuralını ona ve tüm alt projelere şu yolla uygulayabilirsiniz:

```
go fmt ./...
```

Bir şans ver. Kodunuzu girintiden daha fazlasını yapar; ayrıca alan bildirimlerini ve alfabetik olarak ithalatları hizalar.

## Değer Atamalı If

Go, değerlendirilen koşuldan önce bir değerin başlatılabildiği, biraz değiştirilmiş bir if ifadesini destekler:

```go
if x := 10; count > x {
  ...
}
```

Bu oldukça saçma bir örnek. Daha gerçekçi olarak, şöyle bir şey yapabilirsiniz:

```go
if err := process(); err != nil {
  return err
}
```

Bir ilgi çekici durum ise, bu değerlerin if dışında erişilebilir olmamalarına rağmen `else if` veya `else` bloklarında erişilebilir olmalarıdır.

## Boş Arayüz ve Dönüşümler

Çoğu nesne yönelimli dilde, genellikle `object` olarak adlandırılan yerleşik bir temel sınıf, diğer tüm sınıflar için üst sınıftır. Go, kalıtıma sahip olmadığı için, böyle bir üst sınıfa sahip değildir. Sahip olduğu ise hiç bir yöntemi olmayan boş bir arayüzdür: `interface{}`. Her tip yapı boş arayüzün yöntemlerinin tümünü uyguladığından ve arayüzler örtülü olarak uygulandığından, her tip yapı için boş arayüzün sözleşmesini yerine getirir diyebiliriz.

İstersek, aşağıdaki imzayla bir `add` işlevi yazabiliriz:

```go
func add(a interface{}, b interface{}) interface{} {
  ...
}
```

Bir arayüz değişkenini açık bir tipe dönüştürmek için şunu kullanın: `.(TYPE)` :

```go
return a.(int) + b.(int)
```

Kullanılan değişken `int` değilse, yukarıdakilerin bir hataya neden olacağını unutmayın.

Ayrıca swicth kullanarak güçlü bir tip seçimine erişebilirsiniz:

```go
switch a.(type) {
  case int:
    fmt.Printf("a is now an int and equals %d\n", a)
  case bool, string:
    // ...
  default:
    // ...
}
```

Boş arayüzü beklediğinizden daha fazla görecek ve kullanacaksınız. Kuşkusuz, temiz kod ile sonuçlanmaz. Değerleri ileri geri dönüştürmek çirkin ve tehlikelidir, ancak bazen statik bir dilde tek seçenek budur.

## Dizeler ve Bayt Dizileri

Dizeler ve bayt dizileri yakından ilişkilidir. Birini diğerine kolayca dönüştürebiliriz:

```go
stra := "the spice must flow"
byts := []byte(stra)
strb := string(byts)
```

Aslında, bu dönüştürme yöntemi çeşitli tipler için de yaygındır. Bazı işlevler açıkça bir `int32` veya bir `int64` veya bunların işaretsiz hallerini bekler. Kendinizi aşağıdaki gibi şeyler yapmak zorunda bulabilirsiniz:

```go
int64(count)
```

Yine de, bayt ve dizeler söz konusu olduğunda, muhtemelen sık sık yapacağınız bir şeydir. `[]byte(X)` veya `string(X)` kullandığınızda verilerin bir kopyasını oluşturduğunuzu unutmayın. Bu gereklidir çünkü dizeler değişmezdir.

Dizeler, `runes` denilen unicode kod noktalardan oluşur. Bir dizenin uzunluğunu alırsanız, beklediğinizi alamayabilirsiniz. Aşağıdaki kod ekrana 3 yazar:

```
fmt.Println(len("椒"))
```

`range` kullanarak bir dize içinde döngüyle gezinirseniz bayt değil rune elde edersiniz. Elbette, bir dizgiyi `[]byte` çevirdiğinizde doğru verileri alırsınız.

## İşlev Tipi

İşlevler birinci sınıf tipleridir:

```go
type Add func(a int, b int) int
```

daha sonra herhangi bir yerde kullanılabilir - alan tipi, parametre olarak, dönüş değeri olarak.

```go
package main

import (
  "fmt"
)

type Add func(a int, b int) int

func main() {
  fmt.Println(process(func(a int, b int) int{
      return a + b
  }))
}

func process(adder Add) int {
  return adder(1, 2)
}
```

Bunun gibi işlevlerin kullanılması, arayüzlerde elde ettiğimiz gibi belirli uygulamalardan bağımlılıkların ayrıştırılmasına yardımcı olabilir.

## Devam Etmeden Önce

Go ile programlamanın çeşitli yönlerine baktık. En önemlisi, hata işlemenin nasıl davrandığını ve bağlantılar ve açık dosyalar gibi kaynakların nasıl serbest bırakılacağını gördük. Birçok kişi Go'nun hata işlemeye yaklaşımından hoşlanmaz. Geriye doğru bir adım gibi hissedebilir. Bazen katılıyorum. Yine de, takip etmesi daha kolay bir kodla sonuçlandığını da düşünüyorum. `defer` , kaynak yönetimine alışılmadık ama pratik bir yaklaşımdır. Aslında, yalnızca kaynak yönetimine bağlı değildir. `defer`, bir işlevden çıkıldığında loga kaydetme gibi herhangi bir amaç için kullanabilirsiniz.

Elbette, Go'nun sunduğu tüm yeniliklere bakmadık. Ama yine de karşılaştığınız her şeyle başa çıkmak için yeterince rahat hissediyor olmalısınız.

