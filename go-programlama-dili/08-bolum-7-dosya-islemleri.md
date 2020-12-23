# Bölüm 7 - Dosya İşlemleri

## Çapraz Platform Dosya Yolları

Bir işletim sisteminde dosyanın veya dizinin yolunu belirtmek için taksim veya ters-taksim işaretleri kullanırız. Fakat yazağımız program çapraz-platformsa bu durumda ne yapmamız gerekir?

Ya kendimiz bunun için bir fonksiyon oluşturacağız ya da kısa yoldan `os.PathSeperator`'ı kullanabiliriz.

Hemen örneğimizi görelim:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	s := string(os.PathSeparator)
	yol := "dosyalar" + s + "muzikler"
	fmt.Println(yol)
}
```


Her seferinde `string(os.PathSeperator)` yazmamak için `s` değişkenine atayarak kısalttık.


Windows için çıktımız:

> dosyalar\muzikler

Unix-Like için çıktımız:

> dosyalar/muzikler

## İşletim Sistemini Görme

Go programının çalıştığı işletim sistemi görmek için aşağıdaki kodları yazabilirsiniz.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {

	if r := runtime.GOOS; r == "windows" {
		fmt.Println("Windows için yönetici olarak çalıştırın.")
	} else if r == "linux" {
		fmt.Println("Linux için sudo komutu ile çalıştırın.")
	} else {
		fmt.Println("Geçersiz işletim sistemi!")
	}
}
```

GNU/Linux kullandığım için çıktım aşağıdaki gibi olacaktır.

> Linux için sudo komutu ile çalıştırın.

## Dosya Varlığı Kontrolü

Go programımızda kullancağımız bir bir dosyanın varlığını `os` paketi ile kontrol edebiliriz. Örnek programımızı görelim:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if d := "dosya.txt"; dosyaVarmı(d) {
		fmt.Println(d, "bulunuyor")
	} else {
		fmt.Println(d, "bulunmuyor!")
	}
}

func dosyaVarmı(isim string) bool {
	bilgi, hata := os.Stat(isim)
	if os.IsNotExist(hata) {
		return false
	}
	return !bilgi.IsDir()
}
```

**Gelelim açıklmasına:**

Dosya işlemleri yapabilmek için `os` paketini import ettik. if-else akışında geçici değişken olarak `d` değişkenine `"dosya.txt"` atayarak kontrol edilecek dosyamızın ismini belirledik.

Bu akışta `dosyaVarmı` fonksiyonunda `true` değer dönerse `dosya.txt bulunuyor` olarak çıktı almamız gerekir.

`dosyaVarmı` fonksiyonunu incelediğimizde `bilgi` ve `hata` değişkenlerine `os.Stat` ile dosyanın bilgilerini çektik. `hata` değişkeni `false` döndürürse fonksiyonun `false` döndürmesini istedik. Aynı şekilde `bilgi.IsDir()` ile dosya değil de bir dizinse `false` döndürmesini istedik.

## ioutil ile Dosya Okuma ve Yazma

**ioutil** paketi standart Golang paketleri içerisinde gelir ve dosya işlemleri yapabilmemiz için bize fonksiyonlar sağlar.

**Dosya Okuma**

Hemen örneğimize geçelim. Açıklamaları kod üzerinde ilgili alanlara yazdım.

```go
package main
import (
	"fmt"
	"io/ioutil"
)
// Hatayı kontrol etmek için fonksiyonumuz
func kontrol(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// Okunacak dosyamızı belirtiyoruz
	dosya, err := ioutil.ReadFile("dosya.txt")
	// Hata kontrolü yapıyoruz.
	kontrol(err)
	//Dosyamızın içeriğini ekrana bastırıyoruz.
	fmt.Println(string(dosya))
}
```


Okuma işlemi **byte** tipinde yapıldığı için **string\(\)** fonksiyonu ile byte tipini string tipine dönüştürüyoruz.


**Dosya Yazma**

```go
package main
import (
	"io/ioutil"
)
// Hatayı kontrol etmek için fonksiyonumuz
func kontrol(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// Yazmak istediğimiz veriyi belirtiyoruz
	veri := []byte("golangtr.org")
	// Dosya yazma işlemini başlatıyoruz.
	err := ioutil.WriteFile("dosya.txt", veri, 0644) // 0644 dosya yazdırma izni oluyor.
	// Hata kontrolü yapıyoruz.
	kontrol(err)
}
```


String tipini dosyaya yazdırmamız için önce byte tipine çevirmemiz gerekir.


Dosya yazdırma işleminde aynı isimde dosya varsa üzerine yazar.

## Bir Dizindeki Dosya ve Klasörleri Sıralama

Golang üzerinde adresini belirlediğimiz bir dizindeki dosya ve klasörleri listelemeyi göreceğiz. Örneğimize geçelim:

```go
package main
import (
	"fmt"
	"os"
)
func diziniOku(d string) {
	dizin, err := os.Open(d)
	if err != nil {
		fmt.Println("Dizin bulunamadı!")
		os.Exit(1)
	}
	defer dizin.Close()
	liste, _ := dizin.Readdirnames(0) // Açıklamada okuyun
	for _, isim := range liste {
		fmt.Println(isim)
	}
	fmt.Printf("Toplamda %d içerik bulundu.\n", len(liste))
}
func main() {
	diziniOku(".")
}
```

Yukarıdaki kodlarımızın açıklamasını görelim:  
Öncelikle **“os”** paketimizi içe aktarıyoruz. **diziniOku\(\)** fonksiyonumuzun içerisinde **dizin** adında değişken oluşturduk ve bu değişkende fonksiyonumuza **d** argümanı ile gelecek olan dizinimizi açtık. Eğer bir hata ile karşılaşırsak diye hata yakalama işlemini yaptık.  
Daha sonra **dizin** değişkenimizi **defer** ile kapattık.  
**liste** adında değişken oluşturduk. Bu değişkenimizin içerisine **dizin.Readdirnames\(0\)** diyerek tüm dosya ve klasörleri bu değişkenimizin içerisine attık. Burada sıfır kullanmamızın sebebi tüm dosya ve klasörleri okuyabilmek içindir.  
   
Hemen aşağısında **for** ve **range** ile **liste** değişkenimizdeki dosya ve klasör isimlerini isim değişkenimize bastırmak istedik. Her dosya ve klasör ayrı ayrı isim değişkenimize atandı ve ekrana bastırılmış oldu.  
   
Daha sonra **diziniOku\(\)** fonksiyonumuzun en altında **len\(liste\)** ile dosya sayımızı öğrenerek ekrana bastırdık.  
   
**main\(\)** fonksiyonumuzda ise **diziniOku\(“.”\)** diyerek nokta ile bulunduğumuz dizini okuttuk.

## XML Parsing \(Ayrıştırma\)

Bu yazımıza Golang üzerinde **XML** dosyalarını işlemeyi öğreneceğiz. Bu işlemin yapabileceğimiz hali hazırda standart Golang paketleri ile gelen **“encoding/xml”** paketi vardır. Örneğimize geçelim.  
**veri.xml** isminde aşağıdaki gibi bir belgemiz olduğunu varsayalım.

```markup
<?xml version="1.0" encoding="UTF-8"?>
<üyeler>
    <üye tip="admin">
        <isim>Ahmet</isim>
        <sosyal>
            <facebook>https://facebook.com</facebook>
            <twitter>https://twitter.com</twitter>
            <youtube>https://youtube.com</youtube>
        </sosyal>
    </üye>
    <üye tip="okuyucu">
        <isim>Mehmet</isim>
        <sosyal>
            <facebook>https://facebook.com</facebook>
            <twitter>https://twitter.com</twitter>
            <youtube>https://youtube.com</youtube>
        </sosyal>
    </üye>
</üyeler>
```

**XML Belgemizi Okuyalım**

Bu işlemimizi yaparken **“io/ioutil”** ve **“os”** paketlerimizden faydalanacağız. Hemen kodlarımızı görelim.

```go
package main
import (
	"fmt"
	"os"
)
func main() {
	// XML dosyamızı açıyoruz
	xmlDosya, err := os.Open("veri.xml")
	// Hata var mı diye kontrol ediyoruz
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("veri.xml dosyası başarıyla açıldı")
	// XML dosyamızı kapatmayı unutmuyoruz.
	defer xmlDosya.Close()
}
```

Eğer XML dosyası açılırken hata oluşmazsa çıktımız olumlu yönde olacaktır.  
Şimde XML dosyasındaki verileri struct’ımıza kaydedelim. Parsing işlemi de yapacağımızdan dolayı **“encoding/xml”** paketini de içe aktarıyoruz. Hemen kodumuz geliyor.

```go
package main
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)
type Üyeler struct {
	Alan   xml.Name `xml:"üyeler"`
	Üyeler []Üye    `xml:"üye"`
}
type Üye struct {
	Alan   xml.Name `xml:"üye"`
	Tip    string   `xml:"tip,attr"`
	İsim   string   `xml:"isim"`
	Sosyal Sosyal   `xml:"sosyal"`
}
type Sosyal struct {
	Alan     xml.Name `xml:"sosyal"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}
func main() {
	// XML dosyamızı açıyoruz
	xmlDosya, err := os.Open("veri.xml")
	// Hata var mı diye kontrol ediyoruz
	if err != nil {
		fmt.Println(err)
	}
	// XML dosyamızı kapatmayı unutmuyoruz.
	defer xmlDosya.Close()
	//XML dosyamızı okuyoruz (byte olarak geliyor)
	byteDeğer, _ := ioutil.ReadAll(xmlDosya)
	//Yerleştirme işlemi için değişken oluşturuyoruz.
	var üyeler Üyeler
	xml.Unmarshal(byteDeğer, &üyeler)
	fmt.Println(üyeler.Üyeler)
}
```

## JSON Parsing \(Ayrıştırma\)


Yazıya başlamadan önce bu konuyu yazdığı için **Latif Uluman**'a \([@latif70427517](https://twitter.com/latif70427517)\) teşekkürlerimi sunarım.


Bugünkü yazımızda _Golang_ ile _**JSON**_ parse etmeye bakacağız. Hepimizin bildiği gibi günümüzde bir _API_ \(_application programming interface_\) a veri göndermede ya da veri çekmede en sık kullanılan veri formatı _JSON_ \(_javascript object notation_\) dur. _Golang_ ile de kendi oluşturduğumuz verimizi \(_Golang struct_\) _JSON’a_ dönüştürüp bir _API’a_ request olarak gönderebilir ya da bir _API’dan_ gelen _JSON_ verisini Go programımızda kullanabiliriz.  O halde çok uzatmadan Go programımızdaki verileri nasıl _JSON’a_ dönüştürüz hemen bakalım:

**MARSHALLING \(Sıralama\)**

Evet Go programında _Go struct’ını_ JSON stringine dönüştürmek için **“encoding”** altındaki **“json”** paketini kullanıyoruz.  Kullanıma ait kod örneği aşağıdaki gibidir.

```go
package main
import (
	"encoding/json"
	"fmt"
	"log"
)
type kişi struct {
	isim    string
	soyisim string
	yaş     int
}
func main() {
	ali := kişi{
		isim:    "Ali",
		soyisim: "Veli",
		yaş:     20,
	}
	veri, err := json.Marshal(ali)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("JSON Parse Sonucu: %s", string(veri))
}
```

   
Şimdi de kodumuzu çalıştıralım ve sonucu görelim:

> JSON Parse Sonucu: {}

Çıktımıza baktığımızda bir hata olmamasına rağmen JSON string’i boş görüyoruz. Yani marshalling başarılı olmuş gözüküyor; fakat boş bir struct’ı marshal etmiş gibi gözüküyor.  
Evet durum tam da böyle. JSON marshal **sadece dışa aktarılmış \(exported\)** verileri marshal eder. Bildiğimiz gibi Golang’de export etmek için değişken ismi büyük harfle yazılmalıdır. İlk kodumuzda struct elemanlarının baş harflerini  küçük yazdığımız için hiçbiri export edilmedi. Bu yüzden aslında boş bir struct ı marshal etmeye çalışıyoruz gibi algıladı Json.Marshal\(\) fonksiyonu. Doğal olarak geriye boş bir JSON döndü. Haydi şimdi struct elemanlarının tamamını export ederek yani ilk harflerini büyük yazarak test edelim:

```go
package main
import (
	"encoding/json"
	"fmt"
	"log"
)
type kişi struct {
	İsim    string
	Soyisim string
	Yaş     int
}
func main() {
	ali := kişi{
		İsim:    "Ali",
		Soyisim: "Veli",
		Yaş:     20,
	}
	veri, err := json.Marshal(ali)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("JSON Parse Sonucu: %s", string(veri))
}
```

Ve tekrar kodumuzu derleyelim ve sonucu görelim:

> JSON Parse Sonucu: {"İsim":"Ali","Soyisim":"Veli","Yaş":20}

Evet arkadaşlar görüldüğü gibi kodumuz çalıştı. Şimdi kısaca açıklayalım programımızı:  
**7-11** satırlarda kendi **“kişi”** tipimizi oluşturduk.  **13-17** satırlarda bu tipte bir örnek oluşturduk ve **ali** değişkenine atadık.  Daha sonra ali değişkenimizi **json.Marshal\(\)** fonksiyonu kullanarak JSON’a parse ettik. Bu fonksiyondan bize 2 değer dönmektedir. Bunların bir tanesi **\[\]byte** tipinde parse edilen verimiz, diğeri ise **error** tipinde hata durumunu gösteren mesajdır. **19-22** satırlarda hatayı kontrol ettik. Ve son olarak da hatalı değilse ekrana bastık. Tabii bizim datamız \[\]byte tipindeydi, bunu daha okunur hale getirmek için string’e dönüştürdük.  
   
Evet işte bu kadar. Peki diyelim ki JSON string imizi test etmek istiyoruz ve elimizde oldukça karmaşık bir string var. Bunu tek bir satırda incelemek oldukça zahmetli olabilir. İşte bu durumda imdadımıza **json.MarshalIndent\(\)** fonksiyonu yetişiyor. Kullanımı aşağıdaki gibidir:

```go
func main() {
	ali := kişi{
		İsim:    "Ali",
		Soyisim: "Veli",
		Yaş:     20,
	}
	veri, err := json.MarshalIndent(ali, "", "    ")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("JSON Parse Sonucu:\n%s", string(veri))
}
```

Görüldüğü gibi JSON için yeni bir fonksiyon kullandık. Dikkatimizi çeken bir şey fonksiyonun ek olarak 2 paremetre içermesidir. Bunlardan ilki yani 2. parametremiz prefix olarak geçmektedir. Yani 2. parametre her satırın başına gelmektedir. 2. si ise yani 3. parametremiz indentation olarak geçmektedir. Ben onu 4 boşluk olarak ayarladım. Şimdi programımızı tekrar çalıştıralım:

> JSON Parse Sonucu:  
> {  
>     "İsim": "Ali",  
>     "Soyisim": "Veli",  
>     "Yaş": 20  
> }

Görüldüğü gibi ekrana basarken indentation ekleyerek bastı.  
Evet **“encoding/json”** paketi ile go struct’ımızı nasıl JSON’a parse edeceğimizi gördük. Artık JSON datamızı istediğimiz gibi kullanabiliriz.  
Peki tam tersi olsaydı nasıl olurdu? Yani elimizde bir JSON verisi var. Bu bir sorgunun sonucu olabilir. Bunu Go struct’ımıza nasıl çevireceğiz? Çözüm: UNMARSHALL

**UNMARSHALL**

Evet arkadaşlar unmarshal işlemi amaç olarak marshal işleminin tam tersidir. Elimizde JSON formatında bir veri vardır ve biz bunu Go struct’ına dönüştürmek istiyoruz. Bunun için **“encoding/json”** paketinde **Unmarshal** fonksiyonunu kullanırız. O halde çok uzatmadan koda bakalım:

```go
package main
import (
	"encoding/json"
	"fmt"
	"log"
)
type kişi struct {
	İsim    string
	Soyisim string
	Yaş     int
}
func main() {
	jsonVeri := []byte(`{"İsim":"Latif","Soyisim":"Uluman","Yaş":23}`)
	var goVeri kişi
	err := json.Unmarshal(jsonVeri, &goVeri)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("İsim - Soyisim: %s %s\nYaş: %d", goVeri.İsim, goVeri.Soyisim, goVeri.Yaş)
}
```

Evet görüldüğü gibi string formatındaki JSON verimizi önce **\[\]byte** formatına çevirdik sonra onu **Unmarshal** fonksiyonuna parametre olarak verdik. Sonucu da referansını verdiğimiz kişi türündeki **goVeri** değişkenine yazmak istedik. Ve **goVeri.İsim**, **goVeri.Soyisim** ve **goVeri.Yaş** ile bunlara erişmeye çalıştık. Bakalım sonuçlar nasıl:

> İsim - Soyisim: Latif Uluman  
> Yaş: 23

Görüldüğü gibi Unmarshal işlemi başarılı bir şekilde gerçekleşti.  
Peki bir API’ dan gelen JSON verimize ait özellikleri \(attribute\) tam olarak bilmeseydik nasıl bir yol izlememiz gerekirdi? Yani biz burada API’ dan isim-soyisim-yas özelliklerinin geleceğini biliyoruz; fakat bunları bilmeyebilirdik. Bu durumda unmarshal ı hangi türden bir veri tipine gerçeklememiz gerekiyor?  
Çözüm: “**map**” . Evet _**map**_ kullanabiliriz. Yani _**key-value**_ \(anahtar-değer\) ler işimizi görür. Peki türleri ne olmalıdır. “key” ler için düşündüğümüzde bu string olacağı hepimizin aklına gelecektir. Peki Value lar ne olmalıdır? Görüldüğü gibi isim türü string iken, yas integer dı. O halde hepsini karşılayabilen bir veri türü olması lazım. Aklınızda bir şeyler canlanıyor mu? Evet yardımımıza interface yetişiyor. O halde map imizin türü **map\[string\]interface{}** olabilir.Hemen bunu da bir kod örneği ile görelim:

```go
package main
import(
    "encoding/json"
    "fmt"
)
func main(){
    jsonVeri := []byte(`{"İsim":"Latif","Soyisim":"Uluman" ,"Yas":23 , "Kilo":80.25}`)
    var goVeri map[string]interface{}
    err := json.Unmarshal(jsonVeri ,&goVeri )
    if (err != nil){
        fmt.Printf("%+v" , err.Error())
        return
    }
    fmt.Printf("İsim: %+v \nSoyisim: %+v \nYas:%+v\nKilo:%+v" , goVeri["İsim"] , goVeri["Soyisim"] , goVeri["Yas"] , goVeri["Kilo"])
}
```

Programımızı çalıştırıp sonucu görelim:

> İsim: Latif  
> Soyisim: Uluman  
> Yas:23  
> Kilo:80.25

Evet görüldüğü gibi farklı türden veri tipleri olan bir json string ini go map ine dönüştürdük ve key değerleri ile de değerlere ulaştık.Evet arkadaşlar bu yazımızda nasıl bir go verisini json verisine döönüştürüp kullanacağımızı ya da tam tersi json verisini go verisine dönüştüreceğimizi gördük.

## ini Dosyası Okuma ve Düzenleme

ini dosyaları programımızın ayarlarını barındırabileceğimiz dosyalardır. Golang’de ini dosyalarını paket ekleyerek yapabiliriz. Paketimizi indirmek için aşağıdaki komutu yazıyoruz.

> go get gopkg.in/ini.v1

Paketimizi indirdikten sonra ini dosyamız üzerinde işlemler yapabiliriz.  
Aşağıdaki örneklerde kullanacağımız ini dosyası bu şekildedir. Dosyamızın ismi **ayarlar.ini** olsun.

```text
## Yorum satırımız
uygulama_modu = geliştirme
[dizinler]
veri = ./dosyalar
[sunucu]
protokol = http
port = 8000
```

**Ini Dosyası Okuma**

Dosya okuma işlemimiz dizin mantığında çalışır. Örneğimizi görelim.

```go
package main
import (
	"fmt"
	"gopkg.in/ini.v1"
)
func kontrol(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	veri, err := ini.Load("ayarlar.ini")
	kontrol(err)
	fmt.Println("Uygulama Modu:", veri.Section("").Key("uygulama_modu").String())
	fmt.Println("Veri Dizini:", veri.Section("dizinler").Key("veri").String())
	fmt.Println("Bağlantı Protokolü:", veri.Section("sunucu").Key("protokol").String())
	fmt.Println("Bağlantı Portu:", veri.Section("sunucu").Key("port").MustInt(9999))
}
```

Çıktımız şu şekilde olacaktır.

> Uygulama Modu: geliştirme  
> Veri Dizini: ./dosyalar  
> Bağlantı Protokolü: http  
> Bağlantı Portu: 8000



**Inı Dosyası Düzenleme**

Yine aynı **ayarlar.ini** dosyası üzerinde düzenlemeler yapalım. İşte örneğimiz:

```go
package main
import (
	"gopkg.in/ini.v1"
)
func kontrol(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	veri, err := ini.Load("ayarlar.ini")
	kontrol(err)
	// Değer atıyoruz.
	veri.Section("").Key("uygulama_modu").SetValue("ürün")
	// ini dosyamızı kaydetmeyi unutmuyoruz.
	veri.SaveTo("ayarlar.ini")
}
```

## Web Scrapper \(goquery\)

Bu yazıda Go dilinde nasıl basitçe web scrapper yapacağımıza bakacağız.

#### Web Scrapper Nedir?

Web Scrapper bir web sayfasındaki elementleri işleyen araçtır.

**Örnek uygulama:**

**blog.golang.org** sitesindeki blog başlıklarını listeleyen Go programının yazılması.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	blogBasliklari, err := baslikCek("https://blog.golang.org")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Başlıklar:")
	fmt.Printf(blogBasliklari)
}

// URL adresinden blog başlıklarını çekecek fonksiyon
func baslikCek(url string) (string, error) {

	// HTML'i çek
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// goquery dökümanına çevir
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// liste oluştur
	basliklar := ""
	doc.Find(".title a").Each(func(i int, s *goquery.Selection) {
		basliklar += "- " + s.Text() + "\n"
	})
	return basliklar, nil
}

```

Açıklaması:  
goquery kütüphanesini bilgisayarımıza indiriyoruz.

> go get github.com/PuerkitoBio/goquery

`baslikCek` fonksiyonuna URL adresini girdik. Zaten bu fonksiyonu da bi oluşturduk. Hata kontrolü yaptıktan sonra başlıkları yazdırdık.

`baslikCek` fonksiyonuna baktığımızda;  
İlk önce url adresini, yani içindeki elementleri, çektik. goquery dökümanına çevirdik. Burada dikkat edilmesi gereken nokta, `resp` değişkeni bizim çektiğimiz url adresidir. Daha sonra liste olarak oluşturduk. Liste oluşturma işleminde `.title` sınıfına ait ve `a` etiketinde olan elementleri sıralamasını istedik. Element seçim işlemi jQuery selector mantığında çalışır.

Çıktımız:


Başlıklar:

* Announcing the 2019 Go Developer Survey
* Go.dev: a new hub for Go developers
* Go Turns 10
* Go Modules: v2 and Beyond
* Working with Errors in Go 1.13


## chromedp \(Web Driver\)

**Chromedp** paketi, harici bağımlılıklar \(Selenium veya PhantomJS gibi\) olmadan Go'da **Chrome DevTools Protokolü**nü destekleyen tarayıcıları çalıştırmanın daha hızlı ve daha basit bir yoludur. Harici bağımlılık yoktur derken, tabi ki sisteminizde Google Chrome'un yüklü olması gerekiyor. Chromedp'ye headless modu gerektiği için minimum Chrome sürümünüz 59 olması gerekiyor.

##### Paketi yüklemek için:

> go get -u github.com/chromedp/chromedp

### Örnek.1

```go
package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	//chrome örneği oluşturalım
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)...,
	)
	//headless: false ayarlayarak pencerenin görünmesini istedik
	
	//chrome nesnesini defer ile kapatmayı unutmuyoruz
	defer cancel()

	//yeni durum oluşturuyoruz
	ctx, cancel = chromedp.NewContext(ctx)

	//aynı şekilde defer ile penceremizide kapatıyoruz
	defer cancel()

	//Twitter isminin kaydedileceği değişkeni oluşturalım
	var twitterName string

	//chromedp.Run() içerisinde tarayıcıda yapılacak işlemleri yazıyoruz.
	err := chromedp.Run(ctx, //önce durumu (hangi pencere) olacağını belirtiyoruz

		//tarayıcının gitmsini istediğimiz adresi yazalım
		chromedp.Navigate(`https://kaanksc.com/posts/webview-statik-uygulama-ornegi3/`),
		
		//css seçici ile belirttiğimiz elementin yüklenmesini bekleyelim
		chromedp.WaitVisible(`.single__contents > p:nth-child(16) > a:nth-child(1)`, chromedp.ByQuery),
		
		//Tıklanılacak nesneyi yine css seçici ile belirtelim
		chromedp.Click(`.single__contents > p:nth-child(16) > a:nth-child(1)`, chromedp.ByQuery),
		//Bu işlemden sonra twitter'a gidecek

		//Twitter profilinde adın gösterildiği yeri css seçici ile beklemesini istedik
		chromedp.WaitVisible(`div.r-1b6yd1w:nth-child(1) > span:nth-child(1)`, chromedp.ByQuery),

		//belirttiğimiz css seçicisi ile elementin içindeki yazıyı twitterName değişkenine atayalım
		chromedp.Text(`div.r-1b6yd1w:nth-child(1) > span:nth-child(1)`, &twitterName),

		//burdan sonra tarayıcı penceresi kapanacak
	)

	//hata kontrolü yapaım
	if err != nil {
		log.Fatal(err)
	}

	//son olarak twitterName içindeki değişkeni ekrana bastıralım
	log.Printf("Twitter İsim:%s\n", twitterName)
}

```

Yukarıdaki örnekte yeni chrome penceresi oluşturma, tıklama, elementin yüklenmesini bekleme, element içindeki yazıyı alma ve adrese gitme gibi işlemlerin nasıl yapıldığını gördük.

### Örnek.2

Go Playground linkinden Go kodlarını çeken bir uygulama yazalım.

```go
package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	//chrome örneği oluşturalım
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true), //Bu sefer headless çalışmasını istedik
			//yani chrome pecneresi açılmayacak
		)...,
	)

	//chrome nesnesini defer ile kapatmayı unutmuyoruz
	defer cancel()

	//yeni durum oluşturuyoruz
	ctx, cancel = chromedp.NewContext(ctx)

	//aynı şekilde defer ile penceremizide kapatıyoruz
	defer cancel()

	//go Kodlarının kaydedileceği değişkeni oluşturalım
	var goKodu string

	//chromedp.Run() içerisinde tarayıcıda yapılacak işlemleri yazıyoruz.
	err := chromedp.Run(ctx, //önce durumu (hangi pencere) olacağını belirtiyoruz

		//tarayıcının gitmsini istediğimiz adresi yazalım
		chromedp.Navigate(`https://play.golang.org/p/a_SoTzENmV7`),

		//tarayıcının yüklenemesini bekleyeceği elementi css seçici ile yazıyoruz
		chromedp.WaitVisible(`#code`, chromedp.ByQuery),

		//textContent ile yazı alanı içeriğini çekebiliriz.
		chromedp.TextContent(`#code`, &goKodu, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
	}

	//son oalrak go kodlarını ekrana bastıralım
	log.Printf("Go Kodu:\n%s", goKodu)
}

```

Yukarıdaki örnekte headless modda çalışmayı ve yazı kutusu \(input veya textarea\) içindeki yazıları almayı öğrendik.

Daha fazla bilgi için [https://github.com/chromedp/chromedp](https://github.com/chromedp/chromedp),

daha fazla örnek için [https://github.com/chromedp/examples](https://github.com/chromedp/examples) adresine bakabilirsiniz.

