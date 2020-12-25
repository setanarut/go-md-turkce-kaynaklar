- [Bölüm 3 - Gruplama](#bölüm-3---gruplama)
	- [Struct](#struct)
		- [İsim Belirterek Tanımlama](#i̇sim-belirterek-tanımlama)
	- [Anonim Struct'lar](#anonim-structlar)
	- [Struct Fonksiyonlar \(Methodlar\)](#struct-fonksiyonlar-methodlar)
	- [Pointers \(İşaretçiler\)](#pointers-i̇şaretçiler)
	- [Arrays (Diziler)](#arrays-diziler)
	- [Slices (Dilimler)](#slices-dilimler)
	- [Range](#range)
	- [Map](#map)
	- [Interface (Arayüz)](#interface-arayüz)

# Bölüm 3 - Gruplama

## Struct

Go programlama dilinde sınıflar yoktur. Sınıflar yerine struct'lar \(yapılar\) vardır. Yapılar sayesinde bir nesne oluşturabilir ve bu nesneye ait özellikler oluşturabiliriz. Örnek bir struct oluşturalım.

struct örneği
```go
type kişi struct {
	isim    string
	soyİsim string
	yaş     int
}
```


`type` terimi ile yeni bir tür oluşturabiliyoruz. İsmini `kişi` olarak verdik ve türünün de `struct` olacağını söyledik. Yukarıdaki şekilde bir yapı oluşturmuş olduk. Bu yapı içerisinde `isim`, `soyİsim` ve `yaş` değişkenlerine sahip. Yukarıdaki yapı üzerinden bir nesne örneği oluşturduğumuzda örneğimiz bu değişkenlere sahip olacak.

Örnek Kullanım:
```go
package main

import "fmt"

type kişi struct {
	isim    string
	soyİsim string
	yaş     int
}

func main() {

	kişi1 := kişi{"Kaan", "Kuşcu", 23}

	fmt.Println(kişi1)

}
```


`main()` fonksiyonunun içerisini incelediğimizde, `kişi1` isminde `kişi{}` yapısında bir nesne örneği oluşturuyoruz. İçerisine oluşturucu parametreler olarak `kişi struct`'ındaki sıralamayı göz önünde bulundurarak parametrelerimi giriyoruz. Daha sonra kişi1 nesne örneğini ekrana bastırıyoruz. Çıktımız aşağıdaki gibi olacaktır:

> {Kaan Kuşcu 23}

Yukarıdaki örnekte nesneyi tanımlama sırasında değer atamasını yaptık. Nesnenin alt değişkenlerine ulaşarak da tanımlama yapabilirdik.

```go
kişi1 := kişi{"Kaan", "Kuşcu", 23}
kişi1.isim = "Ahmet"
kişi1.soyİsim = "Karaca"
kişi1.yaş = 34

fmt.Println(kişi1) //{Ahmet Karaca 34}
```

Nesne örneğini oluşturuyorken parametreleri boş bırakıp sonradan da atama yapabilirdik.

```go
kişi1 := kişi{}
kişi1.isim, kişi1.soyİsim = "M. K.", "ATATÜRK"
kişi1.yaş = 999

fmt.Println(kişi1) //{M. K. ATATÜRK 999}
```

### İsim Belirterek Tanımlama

Nesneye özel değişkenleri tanımlarken değişken ismini belirterek de tanımlama yapabiliriz.

```go
kişi1 := kişi{soyİsim: "Kuşcu", isim: "Kaan", yaş: 23}

fmt.Println(kişi1) //{Kaan Kuşcu 23}
```

Değişken ismini belirterek atama yaptığımız için sıralamaya dikkat etmemiz gerekli değildir.

## Anonim Struct'lar

Golang’ta tıpkı anonim fonksiyonlar olduğu gibi anonim struct methodlar da oluşturabiliriz. Örneğimizi görelim:

```go
package main
import "fmt"
func main() {
    kişi := struct {
        ad, soyad string
    }{"Kemal", "Atatürk"}
    fmt.Println(kişi)
}
```

Yukarıda struct’ı bir değişken içerisinde tanımladık. Bunu normal struct method olarak yazmaya kalksaydık aşağıdaki gibi yazardık.

```go
package main
import "fmt"
type insan struct {
    ad, soyad string
}
func main() {
    kişi := insan{"Kemal", "Atatürk"}
    fmt.Println(kişi)
}
```

## Struct Fonksiyonlar \(Methodlar\)

Bu bölümde bir struct'a özel nasıl fonksiyon oluşturacağımızı göreceğiz.

Örneğimizi görelim:

```go
package main

import "fmt"

type insan struct {
	isim string
	yaş  int
}

func (i insan) tanıt() {
	fmt.Printf("Merhaba, Ben %s. %d yaşındayım.", i.isim, i.yaş)
}
func main() {
	kişi := insan{"Kaan", 23}
	kişi.tanıt()
}
```

`insan` isminde bir struct tipi oluşturduk. Bu yapımızın tıpkı insanlarda olduğu gibi `isim` ve `yaş` değişkenleri var.

Hemen aşağısında bir fonksiyon oluşturduk. Bu fonksiyonumuzun özelliği ise fonksiyonun isminden önce parantez içerisinde hangi struct'ta çalışacağını belirtmemizdir. `insan` struct'ının içerindeki değişkenlere ise `i` değişkeni ile eriştik.

Daha sonra `main` fonksiyonumuzda `kişi` isminde `insan` tipinde bir nesne oluşturduk. `kişi.tanıt()` yazarak `insan` struct tipinde oluşturduğumuz nesne için olan `tanıt` fonksiyonumuzu çalıştırdık.

Çıktımızı görelim:

> Merhaba, Ben Kaan. 23 yaşındayım.

## Pointers \(İşaretçiler\)

**İşaretçiler** yani pointer’lar bir değerin bellekteki adresini tutar. Değişken atamalarında **& \(and\)** işareti değişkenin bellekteki adresini tutar. **\* \(yıldız\)** işareti ise tutulan adresteki değeri görüntüler. Tekrardan teorik kısmı kısa tutup örneğimize geçelim.

```go
package main
import "fmt"
func main() {
 i := 40
 p := &i
 fmt.Println(p) //Alacağımız benzeri çıktı: 0xc000012120
 fmt.Println(*p) //Alacağımız çıktı: 40
 *p = 35 //Dolaylı olarak i nin değerini değiştirdik
 fmt.Println(i) //Alacağımız çıktı: 35
}
```

İşaretçilerin ana görevini anlatmak gerekir ise, işaretçiler yeni bir değişken oluşturmak yerine var olan bir değişkeni işaretler ve bu değişken üzerinde işlemler yapar. Kodlar ile değişiklikler yaparak mantığını kafanızda pekiştirebilirsiniz.

## Arrays (Diziler)

Diziler içlerinde bir veya birden fazla değer tutabilen birimlerdir. Bir dizideki her değer sırasıyla numaralandırılır. Numaralandırma sıfırdan başlar. Aynı şekilde örneğe geçelim.

```go
package main
import "fmt"
func main() {
 var a [3]string
 a[0] = "Ayşe" //Birinci değer
 a[1] = "Fatma" //İkinci değer
 a[2] = "Hayriye" //Üçüncü değer
 fmt.Println(a) //Çıktımız: [Ayşe Fatma Hayriye]
 fmt.Println(a[1])//Çıktımız: Fatma
}
```

Gelelim kodlarımızın açıklamasına. **a** isminde içerisinde 3 tane **string** tipinde değer barındırabilen bir dizi oluşturduk. a dizisinin birinci değerine yani **0** indeksine **“Ayşe”** atadık. 1 ve 2 indeksine ise “Fatma” ve “Hayriye” değerlerini atadık. a dizisini ekrana bastırdığımızda köşeli parantezler içinde dizinin içeriğini gördük. a’nın 1 indeksindeki değeri bastırdığımızda ise sadece 1 indeksindeki değeri gördük. Dizinin değerlerini tek tek olarak atayabileceğimiz gibi diziyi tanımlarken de değişkenlerini atayabiliriz.

```go
package main
import "fmt"
func main() {
 a := [3]string{"Ayşe", "Fatma", "Hayriye"}
 fmt.Println(a) //Çıktımız: [Ayşe Fatma Hayriye]
}
```

## Slices (Dilimler)

Dilimler bir dizideki değerlerin istediğimiz bölümünü kullanmamıza yarar. Yani diziyi bir pasta olarak düşünürsek kestiğimiz dilimi yiyoruz sadece. Örneğimize geçelim.

```go
package main
import "fmt"
func main() {
 a := [6]int{2, 3, 5, 6, 7, 9}
 fmt.Println(a) //Çıktımız: [2 3 5 6 7 9]
 var b []int = a[2:4] //Dilimleme işlemi
 fmt.Println(b) //Çıktımız: [5 6]
}
```

İnceleme kısmına geçelim. a isminde 6 tane int tipinde değer alan bir dizi oluşturduk. Çıktımızın içeriğini görmek için ekrana bastırdık. Dilimleme işlemi olarak yorum yaptığım satırda ise a dizisinde 2 ve 4 indeksi arasındaki değerleri dizi olarak b’ye kaydettik. b dizisinin içeriğini ekrana bastırdığımızda ise dilimlenmiş alanımızı gördük. Dilimleme işleminde \[ \] içerisine dilimlemenin başlayacağı ve biteceği indeksi yazarız.

**Dilim Varsayılanları \(Sıfır Değerleri\)**

```go
package main
import "fmt"
func main() {
 a := [6]int{2, 3, 5, 6, 7, 9}
 var b []int = a[:4] //Boş bırakılan indeks 0 varsayıldı
 fmt.Println(b) //Çıktımız: [2 3 5 6]
 var c []int = a[3:] //Boş bırakıldığı için 3. index ve sonrası alındı
 fmt.Println(c) //Çıktımız: [6 7 9]
}
```

**Dilim Uzunluğu ve Kapasitesi**  
Bir dilimin **uzunluk** ve **kapasite** değeri vardır. Dilimin uzunluğunu **len\(\)** fonksiyonu ile, kapasitesini ise **cap\(\)** fonksiyonu ile hesaplarız. Örneğimize geçelim.

```go
package main
import "fmt"
func main() {
 a := [6]int{2, 3, 5, 6, 7, 9}
 b := a[2:4]
 fmt.Println("a uzunluk", len(a))
 fmt.Println("a kapasite", cap(a))
 fmt.Println("a'nın içeriği", a)
 fmt.Println("b uzunluk", len(b))
 fmt.Println("b kapasite", cap(b))
 fmt.Println("b'nin içeriği", b)
}
```

b dizisi ile a dizisini dilimlediğimiz için b dizisinin kapasitesi ve uzunluğu değişti. Uzunluk dizinin içindeki değerlerin sayısıdır. Kapasite ise dizinin maksimum alabileceği değer sayısıdır. Çıktımıza bakacak olursak;

> a uzunluk 6 
>
> a kapasite 6 
>
> a'nın içeriği \[2 3 5 6 7 9\] 
>
> b uzunluk 2 
>
> b kapasite 4 
>
> b'nin içeriği \[5 6\]

**Boş Dilimler \(Nil Slices\)**  
Boş bir dilimin varsayılan \(sıfır\) değeri **nil**’dir. Örnek olarak;

```go
package main
import "fmt"
func main() {
 var a []int
 if a == nil {
  fmt.Println("Boş")
 }
}
```

Çıktısı tahmin edeceğiniz üzere **Boş** yazısı olaraktır.

  
**Make ile Dilim Oluşturma**  
Dilimler **make** fonksiyonu ile de oluşturulabilir. Dinamik büyüklükte diziler oluşturabiliriz.

```go
a := make([]int, 5)
```

Burada make fonksiyonu ile uzunluğu 5 olan a adında bir dizi oluşturduk.

```go
a := make([]int, 0, 5)
```

Burada ise make fonksiyonu ile uzunluğu 0, kapasitesi ise 5 olan a adında bir dizi oluşturduk.

  
**Dilime Ekleme Yapma**  
Bir dilime ekleme yapmak için append fonksiyonu kullanılır. Hemen bir örnek ile kullanılışını görelim.

```go
package main
import "fmt"
func main() {
 var a []string
 fmt.Println(a) //[ ]
 a = append(a, "Ali")
 a = append(a, "Veli")
 fmt.Println(a) //[Ali Veli]
}
```

a isminde string tipinde boş bir dizi oluşturduk. Hemen ardından boş olduğunu teyit etmek için a dizisini ekrana bastırdık. Daha sonra a dizisine append fonksiyonu ile “Ali” değerini ekledik. Yine aynı yöntem ile “Veli” değerini de ekledik. Son olarak a dizisinin çıktısının ekrana bastırdığımızda değerlerin eklenmiş olduğunu gördük.

```go
fmt.Println(len(a), cap(a))
```

a dizisinin uzunluk ve kapasitesine baktığımızda aşağıdaki çıktıyı alırız.

> 2 2

## Range

**Range**, üzerinde kullanıldığı diziyi **for** döngüsü ile tekrarlayabilir. Bir dilim range edildiğinde, tekrarlama başına iki değer döndürür \(return\). Birinci değer dizinin **indeksi**, ikinci değer ise bu indeksin içindeki **değerdir**. Örneğimize geçelim.

```go
package main
import "fmt"
var isimler = []string{"Ali", "Veli", "Hasan", "Ahmet", "Mehmet"}
func main() {
 for a, b := range isimler {
  fmt.Printf("%d. indeks = %s\n", a, b)
 }
}
```

Yukarıdaki yazdığımız kodları açıklayalım. **isimler** isminde içerisinde **string** tipinde değerler olan bir **dizi** oluşturduk.  
For döngümüz ile dizinimizdeki değerleri sıralayacak bir sistem oluşturduk. Döngümüzü açıklayacak olursak, bahsettiğimiz gibi dizi üzerinde uygulanan **range** terimi iki değer döndürecek olduğundan bu değerleri kullanabilmek için **a** ve **b** adında argüman belirledik. **range** **isimler** diyerek **isimler** dizisini kullanacağımızı belirttik. Ekrana bastırma bölümümüzde ise **%** işaretleri ile sağ taraftan hangi değerleri nerede kullanacağımızı belirttik.  
Çıktımız ise şu şekilde olacaktır.

> 1. indeks = Ali
> 2. indeks = Veli
> 3. indeks = Hasan
> 4. indeks = Ahmet
> 5. indeks = Mehmet

## Map

**Map**’in Türkçe karşılığında yapacağı işlemi anlatan bir çeviri olmadığı için anlamı yerine yaptığı işi bilelim. Map ile bir değişken içerisindeki dizileri bölge olarak ayırabiliriz. Çok karmaşık bir cümle oldu. O yüzden örneğimize geçelim ki anlaşılır olsun.

```go
package main
import "fmt"
type insan struct {
 kisi1, kisi2, kisi3 string
}
func main() {
 var m map[string]insan
 m = make(map[string]insan)
 m["isim"] = insan{
  "Ali", "Veli", "Ahmet",
 }
 fmt.Println(m["isim"])
}
```

Yukarıda **insan** isminde bir **struct** metodu oluşturduk ve içerisine **string** tipinde 3 tane değişken girdik. **main\(\)** fonksiyonumuz içerisinde ise **m** adında **map** kullanarak **string** değer saklayabilen **insan** tipinde değişken oluşturduk. **m** değişkenini **make** ile **map dizisi** haline getirdik. Hemen aşağısında ise **m** değişkenine **“isim”** adında bir bölge oluşturduk ve **insan** **struct**’ında belirttiğimiz gibi 3 tane **string** değer girdik. Son olarak **m** dizisinin isim bölgesindeki değerleri ekrana bastırmasını istedik. Çıktımız şöyle olacaktır;

> {Ali Veli Ahmet}

**Birden Fazla Bölge Ekleme**  
Önceki yazımızda map ile dizilere bölgesel hale getirmeyi gördük. Şimdi de birden fazla bölgeyi nasıl yapacağımızı göreceğiz. Örneğimize geçelim.

```go
package main
import "fmt"
type insan struct {
 kisi1, kisi2, kisi3 string
}
var m = map[string]insan{
 "erkekler": insan{"Ali", "Veli", "Ahmet"},
 "kadinlar": insan{"Ayşe", "Fatma", "Hayriye"},
}
func main() {
 fmt.Println(m["erkekler"])
 fmt.Println(m["kadinlar"])
 fmt.Println(m)
}
```

Yukarıda önceki örneğimizdeki gibi **insan struct**’ı oluşturduk ve içine **3** tane **string** tipinde değer atadık. **m** adında dizi oluşturduk ve **map** ile bölgeli bir dizi olduğunu belirttik. Dizinin içerisine **“erkekler”** isminde **insan** tipinde bir bölge oluşturduk ve içine **3** tane **string** tipinde değerimizi girdik. Aynı işlemi **“kadinlar”** isimli bölge içinde yaptık. **main** fonksiyonumuz içerisinde **erkekler** ve **kadinlar** bölgemizi ekrana bastırdık. Son olarak **m** dizisindeki tüm içeriği ekrana bastırık.  
Çıktımız ise şöyle olacaktır;

> {Ali Veli Ahmet}  
> {Ayşe Fatma Hayriye}  
> map\[erkekler:{Ali Veli Ahmet} kadinlar:{Ayşe Fatma Hayriye}\]

Burada ayrıntıyı farkedelim. **m** dizisini ekrana bastırdığımızda map yeni bölgeli bir dizi olduğunu vurguluyor. Map ile bir bakıma dizi içerisine yeni bir dizi ekliyorsunuz. Tabi bunu **struct metodu** ile yapıyoruz.

  
**Bölgesel Silme İşlemi**  
**delete** fonksiyonu ile silme işlemimizi yapabiliriz. Hemen örneğimize geçelim.

```go
package main
import "fmt"
func main() {
 m := make(map[string]int) //m isminde string bölge isimli int değer taşıyan dizi
 m["sayi"] = 25 //sayi bölgesine 25 değerini yerleştirdik
 fmt.Println(m["sayi"]) //Çıktımız: 25
 delete(m, "sayi") //sayi bölgesindeki değeri sildik
 fmt.Println(m["sayi"]) //Çıktımız: 0 (sıfır)
}
```

## Interface (Arayüz)

**Interface**'in Go dili üzerindeki kullanımını basitçe açıklayalım. Interface struct nesnelerin, struct tipine göre ilişkili fonksiyonların çalışmasını sağlar. Detayına inmek için önce interface'in nasıl oluşturulduğuna bakalım.

interface oluşturma
```go
type hesap interface{
	hesapla()
}
```


Yukarıda ne yaptığımıza bakacak olursak, `type` ile `hesap` isminde bir `interface` oluşturduk. `hesapla()` ise `hesap` interface'imiz ile ilişkili olacak fonksiyonumuzun ismi olacak.

Interface'in belirli structlar üzerinde etki göstermesi gerekiyor. Bu struct'ları da oluşturalım.

```go
type toplam struct {
	sayı1 int
	sayı2 int
}

type çarpım struct{
	sayı1 int
	sayı2 int
}
```

Yukarıdaki yapılarla toplanılacak ve çarpılacak sayıları barından nesneler oluşturacağız. Sonrasında bu yapılara iliştirilen struct fonksiyonlar yazacağız.

Örnek olarak:

```go
işlem1 := toplam{5,10}

işlem2 := çarpım{5,10}
```

Şimdi de bu structlar için fonksiyonlar oluşturalım.

```go
func (t *toplam) hesapla() {
	fmt.Println(t.sayı1 + t.sayı2)
}

func (ç *çarpım) hesapla() {
	fmt.Println(ç.sayı1 + ç.sayı2)
}
```

Yukarıdaki oluşturduğumuz fonksiyonlarda dikkat edilmesi gereken nokta iki struct fonksiyonun da ismi interface içerisinde belirttiğimiz gibi `hesapla` olmasıdır.

İki fonksiyonda ismini aynı yapmamızın sebebi: oluşturduğumuz interface, nesnenin tipine göre hesapla fonksiyonunu çalıştırmasıdır. Yani işlem1 nesnesini hesap interface'i ile çaşıştırıldığında toplam struct'ı olduğunu algılayıp, toplam struct'ı ile ilişkili hesapla fonksiyonu çalışacaktır. Biraz karışık bir cümle olduğunun farkındayım. O yüzden işlem yaparak öğrenebiliriz.

İlk olarak interface'imizi parametre olarak alan bir fonksiyon oluşturalım.

```go
func hesapYap(h hesap){
	h.hesapla()
}
```

Yukarıda yaptığımız işlem çok basit. `hesap` interface tipini `h` değişkeni ile çağırdık. `h.hesapla()` ile fonksiyonumuzu çalıştırdık.

Gelelim interfacemizi nasıl kullandığımıza:

```go
package main

import "fmt"

type hesap interface {
	hesapla()
}

type toplam struct {
	sayı1 int
	sayı2 int
}

type çarpım struct {
	sayı1 int
	sayı2 int
}

func (t *toplam) hesapla() {
	fmt.Println(t.sayı1 + t.sayı2)
}

func (ç *çarpım) hesapla() {
	fmt.Println(ç.sayı1 * ç.sayı2)
}

func hesapYap(h hesap) {
	h.hesapla()
}

func main() {
	işlem1 := toplam{5, 10}

	işlem2 := çarpım{5, 10}


	//hesap interface'inden bir örnek oluşturalım
	var işlem hesap

	//işlem1'in adresini işlem interface'ine atayalım.
	işlem = &işlem1

	//interface toplam structı olduğunu algılayıp toplama işlemi yapcaktır.
	hesapYap((işlem))

	//işlem2'nin adresini işlem interface'ine atayalım.
	işlem = &işlem2

	//interface çarpım structı olduğunu algılayıp çarpma işlemi yapcaktır.
	hesapYap((işlem))
}
```

Özet geçmek gerekirse, en yukarıda interface'in tanımını yaptığım cümleyi aşağıya kopyala + yapıştır yapayım.

> Interface struct nesnelerin, struct tipine göre ilişkili fonksiyonların çalışmasını sağlar.

