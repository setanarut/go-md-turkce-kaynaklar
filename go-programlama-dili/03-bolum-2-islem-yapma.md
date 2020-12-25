- [Bölüm 2 - İşlem Yapma](#bölüm-2---i̇şlem-yapma)
	- [Fonksiyonlar](#fonksiyonlar)
	- [Fonksiyon Çeşitleri](#fonksiyon-çeşitleri)
	- [Anonim Fonksiyonlar](#anonim-fonksiyonlar)
	- [Boş Tanımlayıcılar](#boş-tanımlayıcılar)
	- [Döngüler](#döngüler)
	- [If-Else](#if-else)
	- [Switch](#switch)
	- [Defer](#defer)

# Bölüm 2 - İşlem Yapma

## Fonksiyonlar

Fonksiyonlar içlerine parametre girilebilen ve işlemler yapabilen birimlerdir. Matematikteki fonksiyonlar ile aynı mantıkta çalışan bu birimlerden bir örneği inceleyelim.

```go
package main

import "fmt"

func topla(a int, b int) int {
	return a + b //a ve b’nin toplamını döndürür.
}

func main() {
	fmt.Println(topla(2, 5)) //2+5 sonucunu ekrana bastır
}
```

Yukarıdaki kodları ineleyecek olursak, foksiyonlarımızı oluşturmak için **func** anahtar kelimesini kullanırız. Yanına ise fonskiyonumuzun ismini yazarız. Parantez içine fonksiyonumuzun dışarıdan alacağı parametreler için değişken-tip tanımlaması yaparız. parantezin sağına ise fonksiyonun döndüreceği **return** değerinin tipini yazarız. Süslü parantezler içinde fonksiyonumuzun işlemleri bulunur. Son olarak return ile veri tipini belirlediğimiz değeri elde etmiş oluruz.

**Main** fonksiyonu içerisinde **topla\(2,5\)** fonksiyonu ile 2 ve 5 sayısının toplamını ekrana bastırmış olduk. Yani ekrana 7 sayısı verildi.

Fonksiyonlar istendiği kadar parametre alabildiği gibi, istenirse parametresiz de olabilir. Fonksiyonları veri return etmek yerine bir işlem yaptırmak içinde kullanabiliriz.

```go
package main

import "fmt"

func yazdir() {
	fmt.Println("yazı yazdırdık")
}

func main() {
	yazdir()
}
```

**yazdir** adlı fonsiyonumuzun parantezine değişken tanımlamadık ve parantezin sağına fonksiyon bloğu içerisinde **return** olmadığı için veri çıkış tipini belirtmedik. Fonksiyonumuzun içerisinde sadece ekrana yazı bastırdık.

**Fonksiyonlar Hakkında Ayrıntılı Bilgiler**

Fonksiyon parantezi içerisine değişken tanımlanırken eğer tüm değişkenlerin türleri aynı ise sadece en sağdaki değişkenin tipini belirtmeniz yeterlidir. Örnek:

```go
package main

import "fmt"

func islem(sayi int) (x, y int) { //return’un degiskenlerini tanımladık
	x = sayi / 2
	y = sayi * 2
	return //Burada sadece return yazıyor
}

func main() {
	fmt.Println(islem(10))
}
```

Yukarıda ise isimlendirilmiş **return** kullandık. return tipini yazdığımız paranteze bakacak olursa **\(x, y int\)** diyerek **return** edilecek verinin fonksiyonun blokları içerisinden çekilmesini sağladık. Böylece fonksiyon bloğununun sonundaki **return** kelimesinin yanına birşey yazmadık. Bu fonksiyonumuzun çıktısı ise **5 20** olacaktır.

## Fonksiyon Çeşitleri

Golang’ta genel olarak 3 çeşit fonksiyon yapısı bulunmaktadır. Hemen bu çeşitleri görelim.

**Variadic Fonksiyonlar**

Variadic fonksiyon tipi ile fonksiyonumuza kaç tane değer girişi olduğunu belirtmeden istediğiniz kadar değer girebilirsiniz.

Hemen örneğimize geçelim.

```go
package main

import "fmt"

func toplama(sayilar ...int) int {
    toplam := 0
    for _, n := range sayilar {
        toplam += n
    }
    return toplam
}

func main() {
    fmt.Println(toplama(3, 4, 5, 6)) //18
}
```

Yukarıdaki fonksiyonumuzu inceleyelim. Vereceğimiz sayıları toplaması için aşağıda **toplama** adında bir fonksiyon oluşturduk. Fonksiyonun parametresi içerisine, yani parantezler içerisine, **sayilar** isminde **int** tipinde bir değişken tanımladık. **…** \(üç nokta\) ile istediğimiz kadar değer alabileceğini belirttik. **toplam** değerini mantıken doğru değer vermesi için **0** yaptık. Çünkü her sayıyı toplam değikeninin üzerine ekleyecek.

**range**’in buradaki kullanım amacından bahsedeyim. **range**’i **for** döngüsü ile kullandığımızda işlem yaptığımız öğenin uzunluğuna göre işlemimizi sürdürürüz. Yani fonksiyonumuzun içine ne kadar sayı eklersek işlemimiz ona göre şekillenecektir. For ve Range işlemini daha sonraki bölümümüzde göreceğiz.

**Range** kullanımında **\_, n** şeklinde değişken tanımlamamızın sebebi, birinci değişken yani **\_**, dizinin indeksini yani sıra numarasını verir. Bizim bununla bir işimiz olmadığı için **\_** koyarak kullanmayacağımızı belirttik. İkinci değişken ise yani **n** dizinin içindeki değeri verir yani fonksiyona girdiğimiz sayıları. Sonuç olarak bu fonksiyonda **return** ile **for** işleminden sonra tüm sayıların toplamını döndürüp **main\(\)** fonksiyonu içerisinde ekrana bastırmış olduk.

**Closure \(Anonim\) Fonksiyonlar**

Closure fonksiyonlar ile değişkenlerimizi fonksiyon olarak tanımlayabiliriz. Örneğimize geçelim.

```go
package main

import "fmt"

func main() {
    toplam := func(x, y int) int {
        return x + y
    }
    fmt.Println(toplam(2, 3))
}
```

Yukarıdaki kodlarımızı inceleyecek olursak, **main** fonksiyonunun içine **toplam** adında bir değişken oluşturduk. Bu değişkenin türünün otomatik algılanması için **:=** işaretlerimizi girdik. Değişkene değer olarak anonim bir fonksiyon \(ismi olmayan fonksiyon yani\) yazdık. Bu fonksiyon **x** ve **y** adında iki tane **int** değer alıyor ve **return** kısmında bu iki değeri **int** olarak döndürüyor. Aşağıdaki **Println\(\)** fonksiyonunda ise bu değişkeni aynı bir fonksiyonmuşcasına kullandık.

**Recursive \(İç-içe\) Fonksiyonlar**

Recursive fonksiyonlar yazdığımız fonksiyonun içinde aynı fonksiyonu kullanmamız demektir. Fonksiyonumun tüm işlemler bittiğinde return olur. Örneğimize geçelim.

```go
package main

import "fmt"

func main() {
    fmt.Println(faktoriyel(4))
}

func faktoriyel(a uint) uint {
    if a == 0 {
        return 1
    }
    return a * faktoriyel(a-1)
}
```

Yukarıdaki fonksiyon ile bir sayının faktöriyelini hesaplayabiliriz. Faktöriyel hakkında kısaca bir hatırlatma yapayım. Belirlediğimiz sayıya kadar olan tüm sayıların sırasıyla çarpımınına o sayının faktöriyeli denir. Yani 4 sayısının faktöriyelini bulmak istiyorsak: 1\_2\_3\*4 işlemini yaparız. Sonuç 24’tür.

Faktöriyel fonksiyonun giriş ve çıkış tiplerini uint yapmamızın sebebi ise faktöriyel sonucunu bulmak için en geriye gidildiğinde eksi değerlere geçilmemesi içindir. Ayrıca sıfırın faktöriyeli birdir. Onun için değer sıfırsa bir return etmesini istedik. Faktöriyel fonksiyonunun en alttaki return kısmında girdiğimiz sayı ile girdiğimiz sayının bir eksiğinin faktöriyelini çarpacak. Girdiğimiz sayının bir küçüğünü bulmak içinse yeniden o sayının faktöriyelini hesaplayacak. Daha sonra aynı işlemler bu sayılar içinde yapılacak, ta ki sayı sona gelene yani en küçük uint değeri olan 0’a dayanana kadar. Daha sonra sonucu main fonksiyonu içerisinde ekrana bastırdık.

## Anonim Fonksiyonlar

Anonim fonksiyonların en büyük özelliği isimsiz olmalarıdır. \(Zaten adından da belli oluyor 🤔\) Yazıldıkları yerde direkt olarak çalışırlar. Çalışırken diğer fonksiyonlardaki gibi parametre verilemediği için fonksiyonun sonuna parametre eklenerek çalışıtırılırlar. Örneğimizi görelim:

```go
package main

import "fmt"

func main() {
	metin := "Merhaba Dünya"

	func(a string) {
		fmt.Println(a)
	}(metin)
}

```

## Boş Tanımlayıcılar

Golang kodlarımızda bazen 2 adet değer döndüren fonksiyonlar kullanırız. Bu değerlerden hangisini kullanmak istemiyorsak, değişken adı yerine **\_ \(alt tire\)** kullanırız.

Örneğimizi görelim:

```go
package main

import "fmt"

func fonksiyonumuz(girdi int) (int, int) {
	işlem1 := girdi / 2
	işlem2 := girdi / 4
	return işlem1, işlem2
}

func main() {
	ikiyeböl, dördeböl := fonksiyonumuz(16)
	fmt.Println(ikiyeböl, dördeböl)
}
```

Gördüğünüz gibi fonksiyonumuzdan dönen iki değeri de değişkenlere atadık. Eğer birini atamak istemeseydik şöyle yapardık:

```go
package main

import "fmt"

func fonksiyonumuz(girdi int) (int, int) {
	işlem1 := girdi / 2
	işlem2 := girdi / 4
	return işlem1, işlem2
}

func main() {
	ikiyeböl, _ := fonksiyonumuz(16)
	fmt.Println(ikiyeböl)
}
```

Yukarıdaki kodlarımızda fonksiyonumuzun 4’e bölme özelliğini kullanmak istemediğimizden dolayı boş tanımlama işlemi yaptık.

Boş tanımlama işlemleri çoğunlukla Golang’ta programcılar tarafından hata çıktısını kullanmak istenmediğinizde yapılıyor.

## Döngüler

Programlama ile uğraşan arkadaşlarımızın da bileceği üzere, programlama dillerinde **while, do while** ve **for** döngüleri vardır. Bu döngüler ile yapacağımız işlemin belirli koşullarda tekrarlanmasını sağlayabiliriz. Golang’ta ise diğer dillerin aksine sadece **for** döngüsü vardır. Ama bu **while** ve **do while** ile yapılanları yapamayacağımız anlamına gelmiyor. Golang’taki for döngüsü ile hepsini yapabiliriz. Yani dilin yapımcıları tek döngü komutu ile hepsini yapabilmemize olanak sağlamışlar.

Gelelim for döngüsünün kullanımına. Go’da for döngüsü parametreleri parantez içine alınmaz.

**STANDART FOR KULLANIMI**

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

**Açıklaması:**

**For** döngüsünden ayrı olarak **deger** adında **0** sayısal değerini alan bir değişken oluşturduk. **For** döngüsünde ise sadece koşul parametresini belirttlik. Yani döngü **deger** değişkeni **10** sayısından küçük olduğu zaman çalışacak. **For** kod bloğu içerisinde her döngü tekrarlandığında deger değişkeni ekrana basılacak ve deger değişkenine **+1** eklenecek.

Konsol çıktımız şu şekilde olacaktır;

> 0
>
> 1
>
> 2
>
> 3
>
> 4
>
> 5
>
> 6
>
> 7
>
> 8
>
> 9

**SADECE KOŞUL BELİRTEREK KULLANMA**

Bu **for** yazım şekli while mantığı gibi çalışır. Parametrelerde sadece koşul belirtilir.

```go
package main

import "fmt"

func main() {
	deger := 0
	for deger < 10 {
		fmt.Println(deger)
		deger++
	}
}
```

**Açıklaması:**

**For** döngüsünden ayrı olarak **deger** adında **0** sayısal değerini alan bir değişken oluşturduk. **For** döngüsünde ise sadece koşul parametresini belirttlik. Yani döngü **deger** değişkeni **10** sayısından küçük olduğu zaman çalışacak. **For** kod bloğu içerisinde her döngü tekrarlandığında deger değişkeni ekrana basılacak ve deger değişkenine **+1** eklenecek.

Konsol çıktımız şu şekilde olacaktır;

> 0
>
> 1
>
> 2
>
> 3
>
> 4
>
> 5
>
> 6
>
> 7
>
> 8
>
> 9

## If-Else

If ve Else kelimelerinin Türkçe karşılığına bakacak olursak;

**If :** Eğer, **Else :** Yoksa anlamına gelir. **If-Else** akışı koşullandırmalar için kullanılır. Diğer dillerin aksine koşul parametresi parantezler içine yazılmaz. Teorik kısmı bırakıp uygulama kısmına geçelim ki daha anlaşılır olsun

```go
if koşul {
	//Koşul sağlandığında yapılacak işlemler
} else {
	//Koşul sağlanmadığında yapılacak işlemler
}
```

Yukarıdaki kod tanımına göre örnek bir program yazalım;

```go
package main

import "fmt"

func main() {
 i := 5
 if i == 5 {
  fmt.Println("i'nin değeri 5'tir.")
 } else {
  fmt.Println("i'nin değeri 5 değildir.")
 }
}
```

Yukarıdaki kodları inceleyelim. i’nin değerini 5 verdik. if teriminin sağında i’nin 5 eşitliği koşulunu sorguladık. Eşitse ekrana i’nin değeri 5’tir. yazısını bastıracak. Değilse i’nin değeri 5 değildir. yazısı bastıracak. i’nin değeri 5 olduğu için ekrana i’nin değeri 5’tir. yazısını bastırdı. If-Else akşında else kullanmamız else’nin kod bloğunu boş bırakmamız ile aynı anlama gelir.

```go
i := 10
if i==10 {
 fmt.Println(“i’nin değeri 10’dur.”)
}
```

Yukarıda sadece **if** deyimini girdik. **Else**’yi girmedik. Burada sonuçlanan olay, **i**’nin değeri **10**’a eşitse **i**’nin değeri **10**’dur. yazısını ekrana bastırır. **Else** deyimini girmediğimiz için şartın sağlanmaması durumunda hiçbir işlem gerçekleşmez. Çıktımız **i**’nin değeri **10**’a eşit olduğu için **i**’nin değeri **10**’dur. çıkar.

**ELSE-IF KULLANIMI**

**If-Else** akışında birden fazla koşul kontrolü ekleyebiliriz. Bunu **else if** deyimi ile yapabiliriz. Kısaca bakacak olursak;

```go
i := 5
if i == 5 {
 fmt.Println("i'nin değeri 5'tir.")
} else if i==3{
 fmt.Println("i'nin değeri 3'tür.")
}else{
 fmt.Println("i'nin değeri belirsiz.")
}
```

else if deyiminin yazılışını da gördük. Açıklamaya gelirsek, else if deyimi kendinden önceki deyimin koşulunun sağlanmaması halinde bir sonraki koşulu kontrol ettirir. If-Else akışında istenildiği kadar else if deyimi eklenebilir.

  
**Koşullar İçerisinde Operatör Kullanımı**  
Koşullar içerisinden mantıksal ve ilişkisel operatörler kullanılabilir. Operatörleri görmüştük. Operatör kullanarak örnekler yapalım.

```go
package main
import "fmt"
func main() {
 i := 5
 a := 3
 b := 5
 if i != a { //Birinci Koşul
  fmt.Println("i eşit değildir a")
 }
 if i == b { //İkinci Koşul
  fmt.Println("i eşittir b")
 }
 if i == b && i > a { //Üçüncü Koşul
  fmt.Println("i eşittir b ve i büyüktür a")
 }
}
```

Çıktımız şu şekilde olacaktır;

> i eşit değildir a
>
> i eşittir b
>
> i eşittir b ve i büyüktür a

## Switch

Switch kelimesinin Türkçe’deki anlamı **anahtardır**. Switch deyimi de if-else deyimi gibi koşul üzerine çalışır. Yine teorik kısmı geçip anlaşılır olması için örnek yapalım. **case** deyimi durumu ifade eder. Koşul sağlandığı zaman işleme devam edilmez.

```go
package main
import "fmt"
func main() {
 i := 5
 switch i {
  case 5:
   fmt.Println("i eşittir 5")
  case 10:
   fmt.Println("i eşittir 10")
  case 15:
   fmt.Println("i eşittir 15")
 }
}
```

Çıktımız şu şekilde olacaktır;

> i eşittir 5

Switch’te koşulların gerçekleşmediği zaman işlem uygulamak istiyorsak bunu **default** terimi ile yaparız. Örnek;

```go
i := 5
switch i {
 case 5:
  fmt.Println("i eşittir 5")
 default:
  fmt.Println("i bilinmiyor")
}
```

**Koşulsuz Switch**  
Switch’in tanımını daha iyi anlayabilmeniz için koşulsuz switch kullanımına örnek verelim. Bu yöntemde switch deyiminin yanına koşul girmek yerine case deyiminin yanına koşul giriyoruz.

```go
package main
import "fmt"
func main() {
 i := 5
 switch {
  case i == 5: //i=5 olduğu için diğer case’ler sorgulanmaz
   fmt.Println("i eşittir 5")
  case i < 10:
   fmt.Println("i küçüktür 10")
  case i > 3:
   fmt.Println("i büyüktür 3")
 }
}
```

Çıktımız şu şekilde olacaktır;

> i eşittir 5

## Defer

Defer kelimesinin Türkçe’deki karşılığı **ertelemektir**. Bu deyimi yapacağımız işlemin başına eklersek o işlemi içerisinde bulunduğu fonksiyonun içindeki işlemlerden sonra çalıştırır. Çok karışık bir cümle kurdum ama uygulamaya geçince anlayacaksınız.

```go
package main
import "fmt"
func main() {
 defer fmt.Println("İlk Cümle")
 fmt.Println("İkinci Cümle")
}
```

Çıktımız şu şekilde olacaktır;

> İkinci Cümle
>
> İlk Cümle

Açıklamaya gelirsek ekrana **İlk Cümle** yazısını bastıran satırımızın başına **defer** terimini ekledik. **defer** eklediğimiz satır **main\(\)** fonksiyonunun içinde olduğu için **main\(\)** fonsyionundaki tüm işlemler tamamlandıktan sonra ekrana yazımızı bastırdı.  
Birden fazla defer ekleyecek olursak;

```go
package main
import "fmt"
func main() {
 defer fmt.Println("ilk Cümle")
 defer fmt.Println("İkinci Cümle")
 defer fmt.Println("Üçüncü Cümle")
 defer fmt.Println("Dördüncü Cümle")
 fmt.Println("Beşinci Cümle")
}
```

Çıktımız şu şekilde olacaktır;

> Beşinci Cümle
>
> Dördüncü Cümle
>
> Üçüncü Cümle
>
> İkinci Cümle
>
> ilk Cümle

Burdan anlıyoruz ki en baştaki defer eklenen satır en son işleme tabi tutuluyor. Hadi defer ile alakalı bir programlama alıştırması yapalım.

```go
package main
import "fmt"
func main() {
 fmt.Println("Sayıyor")
 for i := 0; i < 10; i++ {
  defer fmt.Println(i)
 }
 fmt.Println("Bitti")
}
```

Çıktımız şöyle olacaktır;

> Sayıyor
>
> Bitti
>
> 9
>
> 8
>
> 7
>
> 6
>
> 5
>
> 4
>
> 3
>
> 2
>
> 1
>
> 0

