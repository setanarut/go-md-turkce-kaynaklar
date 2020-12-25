- [Bölüm 5](#bölüm-5)
  - [init\(\) Fonksiyonu \(Ön Yükleme\)](#init-fonksiyonu-ön-yükleme)
  - [Import \(Kütüphane Ekleme\) Yöntemleri](#import-kütüphane-ekleme-yöntemleri)
  - [Dışa Aktarma \(Exporting\)](#dışa-aktarma-exporting)
  - [Print Fonksiyonu Birkaç İnceleme](#print-fonksiyonu-birkaç-i̇nceleme)
  - [Format ve Kaçış Karakterleri](#format-ve-kaçış-karakterleri)
  - [Çok Satırlı String Oluşturma](#çok-satırlı-string-oluşturma)
  - [Sprintf](#sprintf)
  - [Golang'te Kullanıcıdan Giriş Alma](#golangte-kullanıcıdan-giriş-alma)
  - [Testing \(Test Etme\)](#testing-test-etme)
  - [Panic & Recover](#panic--recover)

# Bölüm 5

## init\(\) Fonksiyonu \(Ön Yükleme\)

Golang’te bir uygulama çalışırken genelde çalışan ilk fonksiyon **main\(\)** fonksiyonu oluyor. Bazen programın açılışında ayarlamamız gereken ön durumlar oluşuyor. İşte **init\(\)** fonksiyonu bize bu imkanı sunuyor. Ufak bir örnekle yazdıklarıma anlam katalım.

```go
package main
import "fmt"
func init() {
	fmt.Println("init fonksiyonu yüklendi")
}
func main() {
	fmt.Println("main Fonksiyonu yüklendi")
}
```

Çıktımız aşağıdaki gibi olacaktır.

> init fonksiyonu yüklendi  
> main Fonksiyonu yüklendi

Golang’taki **init\(\)** fonksiyonunun kullanımı, farklı dillerdeki aynı işlevi gören fonksiyonlara oranla daha kolaydır. Örnek olarak **init\(\)** fonksiyonunda veritabanı bağlantımızı, kayıt defteri işlemlerimizi veya sadece bir kez yapmamız gereken işleri yapabiliriz. Buna imkan sağlayan mantığı aşağıdaki örnekte görelim. Bu örnekte global tanımlanmış değişkenin değerini **init\(\)** fonksiyonunda değiştirdiğimizde **main\(\)** gibi farklı fonksiyonlarda kullanabildiğimizi göreceğiz.

```go
package main
import "fmt"
var değişken string
func init() {
	değişken = "Merhaba Dünya"
}
func main() {
	fmt.Println(değişken)
}
```

Çıktımız ise şöyle olacaktır.

> Merhaba Dünya

İşte **init\(\)** fonksiyonunun böyle bir güzelliği var. Benzer bir işlevi ancak pointers \(işaretçiler\) ile yapabiliriz.

```go
package main
import "fmt"
var değişken string = "Naber"
func değiştir(değişken *string) {
	*değişken = "Merhaba Dünya"
}
func main() {
	değiştir(&değişken)
	fmt.Println(değişken)
}
```

O da gördüğünüz gibi uzun bir işlem oluyor.

## Import \(Kütüphane Ekleme\) Yöntemleri

Bu yazıda sizlere Golang’ta paket import etmenin tüm yöntemlerini göstereceğim.  
**1. Yöntem**

```go
import "fmt"
```

**fmt** paketini import ettik.  
**2. Yöntem**

```go
import (
    "fmt"
    "net/http"
)
```

Birden fazla paket import ettik  
**3. Yöntem**

```go
import f "fmt"
```

**fmt** paketini import edip **f** olarak kullanacağımızı belirttik. Örnek olarak **fmt.Println\(\)** yazmak yerine **f.Println\(\)** yazacağız.  
**4. Yöntem**

```go
import . "fmt"
```

Dikkat ederseniz, **import** kelimesinden sonra **nokta** koyduk. Bu işlem sayesinde **fmt.Println\(\)** yazmak yerine sadece **Println\(\)** yazarak aynı işi yapmış oluruz.  
**5. Yöntem**

```go
import _ "fmt"
```

Bazen Golang yazarken kütüphaneyi ekleyip kullanmadığımız zamanlar olur. Böyle durumlarda program çalıştırılırken veya derlenirken kullandığınız editör veya ide bu bölümü silebilir. import ederken **\_ \(alt tire\)** koyarak bunun üstesinden gelebiliriz.

## Dışa Aktarma \(Exporting\)

Golang’ta dışa altarma çok basit bir olaydır. Diğer programlama dillerinde public anahtar kelimesi olarak gördüğümüz bu olayın Golang’ta nasıl yapıldığına bakalım. Golang’ta bunun için bir anahtar kelime yoktur. Dışa aktarılmasını istediğimiz öğeyi oluştururken baş harfini büyük yazarız. Örnek olarak:

```go
func Topla(x, y int) int {
    return x + y
}
```

Gördüğünüz gibi Topla\(\) fonksiyonunun baş harfini büyük yazdır. Peki dışa aktarma hangi durumlarda yapılır.

* Bir paket oluşturup başka bir paket içerisinden dışa aktarılan öğeyi kullanmak istiyorsak,
* Projemiz birden fazla .go dosyası içeriyorsa ve bir sayfadaki öğeyi başka sayfada da kullanmak istiyorsak,

dışa aktarma yöntemi işimizi görecektir.  
Fonksiyonları dışa aktarabildiğimiz gibi değişkenleri ve sabitleride dışa aktarbiliriz. Örnek olarak:

```go
var Degisken = string("değişken değerimiz")
const Sabit = string("sabit değerimiz")
```

Dışa aktarma olayı Golang’ta bu kadar basittir.

## Print Fonksiyonu Birkaç İnceleme

Print fonksiyonu Go dilinde komut satırı üzerinde yazdırma işlemi yapmak için kullanılır. Print fonksiyonunun en çok kullanılan 3 çeşidine bakalım.

**Print\(\) Fonksiyonu**  
Bu fonksiyonun içine parametreler girerek ekrana yazdırma işlemi yapabiliriz.

```go
fmt.Print(“Merhaba Dünya!”)
```

Çıktımız şu şekilde olacaktır:

> Merhaba Dünya

**Println\(\) Fonksiyonu**  
Bu fonksiyon ile içine parametre girerek ekrana yazdırma işlemi yapabiliriz. Yazdırma işlemi yaptıktan sonra bir alt satıra geçer.

```go
fmt.Println(“satır1”)
fmt.Println(“satır2”)
```

Çıktımız şu şekilde olacaktır;

> satır1  
> satır2

**Printf\(\) Fonksiyonu**  
Gelelim işimizi görecek olan Printf\(\) fonksiyonuna. Bu fonksiyon sayesinde metinsel bölümlerin arasına değişken yerleştirebiliriz.

```go
dil:=”Go”
yıl:=2007
fmt.Printf(“%s dili %d yılından beri geliştiriliyor.”,dil,yıl)
```

Çıktımız şu şekilde olacaktır;

> Go dili 2007 yılından beri geliştiriliyor.

## Format ve Kaçış Karakterleri

**Format Karakterleri ve Kullanım Alanları**  
Format karakterleri metinsel bir ifadede \(string\), dizgiyi formatlandırmak için kullanılır. Yani bir metinde değişken yerleri biçimlendirmeye yarar.

| Format Karakteri | Açıklama |
| :--- | :--- |
| %T | Değişkenin tipini verir |
| %t | Boolean değeri verir |
| %d | Int \(tamsayı\) değeri verir |
| %b | Sayının binary \(ikili\) karşılığını verir |
| %c | Karakter değerini verir |
| %x | Sayının hexadecimal \(onaltılı\) karşılığını verir |
| %f | Float \(ondalıklı\) değeri verir |
| %s | String \(dizgi-metin\) değeri verir |
| %v | Değeri otomatik belirler |

Hemen bir örnek yapalım.

```go
package main

import "fmt"

func main() {
    isim := "Kaan"
    yaş := 23
    kilo := 71.3
    evli := false

    fmt.Printf("İsim: %s, Yaş: %d, Kilo: %f, Evli: %t", isim, yaş, kilo, evli)
}
```

Yukarıdaki kodlara göre şöyle bir çıktı alacaksınız:

> İsim: Kaan, Yaş: 23, Kilo: 71.300000, Evli: false

Kilo olarak girdiğimiz değer uzun olarak görüntülendi. Bunu değiştirmek için aşağıdaki yöntem uygulanır.

```go
fmt.Printf("İsim: %s, Yaş: %d, Kilo: %.1f, Evli: %t", isim, yaş, kilo, evli)
```

Yukarıdaki kodda farkedeceğiniz üzere `kilo` değişkeni için olan format karakterini `%.1f` olarak değiştirdik. Bu küsüratlı sayılarda noktadan sonra 1 karakter gelebileceğini gösteriyor. Çıktımız: `71.3` olarak değişecektir.

> İsim: Kaan, Yaş: 23, Kilo: 71.3, Evli: false

warning
Format karakterleri **Printf** ve **Scanf** gibi fonksiyonlarda kullanılabilir. Bu fonksiyonların ortak özellikleri adında **f** harfi olmasıdır.


**Kaçış Karakterleri ve Kullanım Alanları**  
Kaçış karakterleri de format karakterleri gibi metinlere etki eder. Kaçış karakterlerini kod yazma zamanında yapamadığımız işlemler için kullanırız.

| Kaçış Karakteri | Açıklama |
| :--- | :--- |
| \a | Komut satırında zil sesi çıkartır |
| \b | Silme tuşu görevini görür |
| \f | Merdiven metin yazar |
| \n | Yeni satıra geçer |
| \r | Return eder |
| \t | Tab tuşu gibi boşluk bırakır \(4 boşluk\) |
| \v | Dikey boşluk bırakır |
| \ | Ters-taksim yazar |
| \' | Tek tırnak yazar |
| \" | Çift tırnak yazar |

Gelelim örneğimize:

```go
fmt.Print("Bir\nİki\tÜç\\Dört")
```

Çıktımız şöyle olacaktır:

> Bir  
> İki Üç\Dört

## Çok Satırlı String Oluşturma

Çok satırlı string oluşturmak için \(\`\) işaretini kullanırız. Türkçe klavyeden **alt gr** ve **virgül** tuşuna basarak bu işareti koyabilirsiniz. İşte örnek kodumuz;

```go
package main
import "fmt"
func main() {
    yazi := `Bu bir
    çok satırlı
    yazı örneğidir.
`
    fmt.Printf("%s", yazi)
}
```

## Sprintf

Sprintf fonksiyonu fmt paketine dahil bir fonksiyondur. Bu fonksiyon değişkenlere formatlı atama yapmamıza yardımcı olur. Örneğimizi görelim:

```go
package main

import (
	"fmt"
)

func main() {
	isim := "Kaan"

	isimTip := fmt.Sprintf("isim değişkeni %T tipindedir.", isim)

	fmt.Println(isimTip)
}
```

## Golang'te Kullanıcıdan Giriş Alma

Golang’te diğer programlama dillerinde de olduğu gibi kullanıcıdan değer girişi alınabilir. Böylece programımızı interaktif hale getirmiş oluruz.

  
**Scan\(\) Fonksiyonu**  
Bu fonksiyon boşluğa kadar olan kelimeyi kaydeder. Yeni satır boşluk olarak sayılır. Kullanımını görelim.

```go
var yazi string
fmt.Scan(&yazi) //yazi değişkenine değer girilmesini istedik.
fmt.Println(“\n”+yazi)
```

Yukarıda yazdığımız kodları inceleyecek olursak, belleğe yazi isimli string türünde bir değişken kaydettik. Kullanıcının girişte bulunabilmesi için **Scan\(\)** fonksiyonunu kullandık. Bu fonksiyonun içerisine **&yazi** yazdık. Bu sayede kullanıcının girdiği değer **yazi** değişkeninin içerisine kaydedilebilecek. Daha sonra **yazi** değişkenini ekrana bastırdık ve bizim yazdığımız değer görüntülendi. Scan fonksiyonunda dikkat edilmesi gereken nokta kullanıcı istediği kadar kelime girse bile programın ilk kelimeyi değer olarak alacağıdır. **Scan\(\)** fonksiyonu boş giriş kabul etmez.

**Scanf\(\) Fonksiyonu**  
**Scanf\(\)** fonksiyonu **Printf\(\)** fonksiyonu gibi format içerir. Bu fonksiyon ile kullanıcının girişini bölüp birkaç değişkene kaydedebiliriz. Hemen kullanımını görelim.

```go
var kelime1, kelime2 string
fmt.Scanf(“%s %s”,&kelime1,&kelime2)
fmt.Println(kelime1)
fmt.Println(kelime2)
```

Yukarıda yazdığımız kodları inceleyecek olursak, **kelime1** ve **kelime2** adında **string** türünde değişkenler belirledik. **Scanf\(\)** fonksiyonu ile **Printf\(\)**’den benzer olarak, değişkenlerin yerleştirileceği yerleri değil de, bu sefer değişkenlerin alınacağı yerleri belirtiyoruz. **%s %s** arasındaki **boşluk** sayesinden kullanıcı boşluk bırakınca girdiyi **2** değere bölebileceğiz. Hemen yanında ise içine atanacak değişkenlerimizi belirtiyoruz. Böylelikle kullanıcı giriş bölümünden Go Dili yazdığında **Go**’yu **kelime1**’in içine **Dili** de **kelime2** içine yerleştirecek. **Scanf\(\)**, boş giriş kabul eder.

**Reader ile Satır Olarak Değer Alma**  
Aşağıdaki yöntem ile bir satır yazıyı giriş olarak alabilirsiniz.

```go
giris := bufio.NewReader( os.Stdin)
yazi, _ := giris.ReadString('\n')
```


**Scan** komutu ile kelime alamadığınızda **Reader** ile deneyebilirsiniz.


## Testing \(Test Etme\)

Hücrelerin vücumudaki yapı birimi olduğu gibi, aynı şekilde her bileşen de yazılımın birer parçasıdır. Yazılımın sağlıklı bir şekilde çalışabilmesi için, her bileşenin güvenilir bir şekilde çalışması gerekir.  
Aynı şekilde vücudumuzun sağlığı hücrelerin güvenilirliği ve verimliliğine bağlı olduğu gibi, yazılımın düzgün çalışması bileşenlerin güvenilirliği ve verimliliğine bağlıdır.  
Biraz biyoloji dersi gibi oldu ama sonuçta aynı mantığı yürütebiliriz.  


**Peki bileşenler nedir?**  
Yazılımın çalışması için yazılmış her bir kod parçasına denir. Bu bileşenlerin yazılımımızın sağlıklı bir şekilde çalıştırdığından emin olmamız gerekir.  
Peki bu bileşenlerin sağlamlık kontrolünü nasıl gerçekleştiririz? Tabiki **test** ederek.  
Bir test aşamsının Golang’ta nasıl göründüğünü görelim.

```go
import "testing"
func TestFunc(t *testing.T){
    t.error() //testin başarısız olduğunu bildirir.
}
```

Yukarıdaki işlem Golang’ta yapılan bir birim testin temel yapısıdır. Yerleşik **testing** paketi, Golang’ın standart paketleri içerisinde gelir. Birim testi, **\*testing.T** türündeki elemanı kabul eden ve bu elemanı göre hata yayınlayan bir bir işlemdir.  
Bu fonksiyonların adı büyük harfle başlamalı ve birleşik olan adın devamı da bütük harfle başlamalıdır. Yani **camel-case** olmalıdır.  
`TestFunc olmalıdır ve Testfunc olmamalıdır.`  
Uygulama örneğimize geçelim.  
Bir proje klasörü oluşturalım ve **main.go** dosyamız şöyle olsun.

```go
package main
import "fmt"
func Merhaba(isim string) (çıktı string) {
	çıktı = "Merhaba " + isim
	return
}
func main() {
	selamla := Merhaba("Kaan")
	fmt.Println(selamla)
}
```



**main.go** dosyamızda fonksiyona adını girdiğimiz kişiyi selamlıyor. Buraya kadar gayet basit bir program. Fonksiyonlarımızı test edeceğimiz için baş harflerini büyük yazmayı unutmuyoruz. Böylelikle fonksiyonlarımızı dışarı aktarabiliriz. Test fonksiyonumuzun çalışma mantığını görmek için **main\_test.go** dosyamıza bakalım.

```go
package main
import "testing"
func TestMerhaba(t *testing.T) {
	if Merhaba("Kaan") != "Merhaba Kaan" {
		t.Error("Merhaba Fonksiyonunda bir sıkıntı var!")
	}
}
```

Yukarıda ise **main.go** sayfamızdaki **Merhaba** fonksiyonunu test etmek için **TestMerhaba** adında fonksiyon oluşturduk. **t \*testing.T** ibaresi ile bu işlemin test etmeye yönelik bir işlem olduğunu belirttik.  
Fonksiyonun içerisine baktığımızda, **Merhaba\(“Kaan”\)** işleminin sonucu **“Merhaba Kaan”** olmadığı zaman test hatası vermesini istedik. Ve gözükecek hatayı belirttik.  
Test işlemi yapmak için aşağıdaki komutları komut satırına yazıyoruz.

> go test

Yukarıdaki yazdığımız kodlara göre şöyle bir çıktımızın olması gerekir.

> PASS ok  
> \_/home/ksc10/Desktop/deneme 0.002s

Eğer **TestMerhaba** fonksiyonunda test koşuluna **“Merhaba Kaan”** yerine **“Merhaba Ahmet”** yazsaydık, aşağıdaki gibi bir **go test** çıktımız olurdu.

> --- FAIL: TestMerhaba \(0.00s\)  
>        main\_test.go:7: Merhaba Fonksiyonunda bir sıkıntı var!  
> FAIL  
> exit status 1  
> FAIL \_/home/ksc10/Desktop/deneme 0.002s

**Go Test Komutları**

| Komut | Açıklama |
| :--- | :--- |
| go test | İçerisinde bulunduğu projenin tüm test fonksiyonlarını test eder. |
| go test -v | Her test için ayrı bilgi verir. |
| go test -timeout 30s | 30 saniye zaman aşımı ile test eder. |
| go test -run TestMerhaba | Sadece belirli bir fonksiyonu test eder. |

**Örnek kullanımı:**  
**main\_test.go** dosyamızdaki **TestMerhaba** fonksiyonumuzu **10 saniye** zaman aşımı ile test edecek komut

> go test -timeout 30s -run TestMerhaba

Bu yazımızda Golang’de test işleminin nasıl yapıldığını gördük. Mantığını daha iyi kavramak için bir proje üzerinde gerekli olduğu yerde kullanmamız gerekir.

## Panic & Recover

**Panic** ve **Recover**, Golang’de hata ayıklama için kullanılan anahtar kelimelerdir. Size bunu daha iyi ve akılda kalıcı anlatmak için teorik anlatım yerine uygulamalı öğretim yapmak istiyorum. Böylece daha akılda kalıcı olur.  
Aşağıda **panic** durumu oluşturan bir örnek göreceğiz:

```go
package main
func main() {
    sayilar := make([]int, 5)
    sayilar[6] = 10
}
```

Yukarıda **make** fonksiyonu ile **sayilar** adında uzunluğu **5** birimden oluşan bir **int** dizi oluşturduk. Bu bildiğimiz sayısal 5 tane değişken tutan bir dizi aslında. Ama altında **sayilar** dizisinin **6**. indeksine **10** değerini atamak istedik. Fakat **sayilar** dizesinin 6. indeksi mantıken bulunmamakta. Bu haldeyken programımız **panic** hatası verecektir ve çıktımız aşağıdaki gibi olacaktır.

> panic: runtime error: index out of range  
> goroutine 1 \[running\]:  
> main.main\(\)  
> /home/ksc10/Desktop/deneme/main.go:5 +0x11  
> exit status 2

İstersek biz de kritik bir bilginin nil girilmesi gibi durumlarda programı durdurabiliriz. Bunun için **panic\(\)** fonksiyonunu kullanacağız. Hemen bir örnek yapalım.

```go
package main

import (
    "fmt"
)

func TamIsim(Ad *string, Soyad *string) {
    if Ad == nil {
        panic("Ad nil olamaz")
    }
    if Soyad == nil {
        panic("Soyad nil olamaz")
    }
    fmt.Printf("%s %s\n", *Ad, *Soyad)
    fmt.Println("TamIsim fonksiyonu bitti")
}

func main() {
    Ad := "Yusuf"
    TamIsim(&Ad, nil)
    fmt.Println("Ana fonksiyon da bitti")
}
```

Çıktımız burada:

> panic: Soyad nil olamaz  
> goroutine 1 \[running\]:  
> main.TamIsim\(0xc00007df30, 0x0\)  
> /Users/Y/Desktop/main.go:12 +0x19a  
> main.main\(\)  
> /Users/Y/Desktop/main.go:20 +0x65  
> exit status 2

Burada **Soyad** değişkeni tanımsız olduğu için programımız durdu. Aynı şekilde **recover\(\)** fonksiyonu ile **panic\(\)** fonksiyonundan gelen veriyi alabilir, ana fonksiyonumuzun kapanmasına da engel olabiliriz. Bunun için de bir örnek yapalım.

```go
package main

import (
    "fmt"
)

func TamIsim(Ad *string, Soyad *string) {
    if Ad == nil {
        panic("Ad nil olamaz")
    }
    if Soyad == nil {
        panic("Soyad nil olamaz")
    }
    fmt.Printf("%s %s\n", *Ad, *Soyad)
    fmt.Println("TamIsim fonksiyonu bitti")
}

func main() {
    Ad := "Yusuf"
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panik Yok : ", r)
        }
    }()
    TamIsim(&Ad, nil)
    fmt.Println("Ana fonksiyon da bitti")
}
```

Çıktımız burada :

> Panik Yok : Soyad nil olamaz

