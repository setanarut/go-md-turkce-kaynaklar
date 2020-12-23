# Web Sunucu (Server)

## net/http ile Web Server Oluşturma

Golang’ta web sunucusu oluşturma çok basit bir işlemdir.  
İlk örneğimizde **localhost:5555** üzerinde çalışacak olan bir web sunucusu oluşturacağız.

```go
package main 
import (
    "fmt"
    "net/http"
)
 
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
}
 
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":5555", nil)
 
    fmt.Println("Web Sunucu")
}
```

Tarayıcınız üzerinden **localhost:5555**‘e girdiğinizde sayfada sadece **Merhaba** yazdığını göreceksiniz. Daha sonra adrese **/ahmet** yazıp girdiğiniz zaman yazının **Merhaba ahmet** olarak değiştiğini göreceksiniz.  


**Peki bu olayın açıklaması nedir?**  
**main\(\)** fonksiyonunun içerisinde 2 temel fonksiyon bulunuyor. **HandleFunc\(\)** fonksiyonu belirlediğimiz adrese girildiğinde hangi fonksiyonun çalıştırılacağınız belirliyor. **ListenAndServe\(\)** fonksiyonu ise sunucunun ayağa kalkmasını ve istediğimiz bir porttan ulaşılmasını sağlıyor.  
Eğer sunucuya dosya verme yoluyla işlem yapmasını istiyorsak aşağıdaki yönteme başvurmalıyız.  
**index.html** adında bir dosya oluşturuyoruz. İçine aşağıdakileri yazıyoruz ve kaydediyoruz.

```markup
<!DOCTYPE html>
<html lang="tr">
<head>
    <title>Sayfa Başlığı</title>
</head>
<body>
    Merhaba Dünya
</body>
</html>
```

Şimde de sunucu işlemlerini gerçekleştireceğimiz **main.go** dosyamızı oluşturalım.

```go
package main 
import (
    "fmt"
    "io/ioutil"
    "net/http"
)
 
func loadFile(fileName string) (string, error) {
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
 
func handler(w http.ResponseWriter, r *http.Request) {
    var body, _ = loadFile("index.html")
    fmt.Fprintf(w, body)
}
 
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":5555", nil)
}
```

Tarayıcıdan **localhost:5555** adresine girdiğimiz zaman oluşturmuş olduğumuz **index.html** dosyasının görüntülendiğini göreceksiniz.  
Açıklayacak olursak eğer;  
**loadFile\(\)** fonksiyonumuz **index.html** programa aktarıldığında **byte** türünde olduğu için onu okuyabileceğimiz **string** türüne çevirdi. Bu özellik programımıza **“io/ioutil”** paketi sayesinde eklendi. Geri kalan kısımdan zaten yukarıda bahsetmiştik.

## HTML Şablonlar \(Templates\)

Golang’ta HTML sayfalarına öğe yerleştirmek için şabloblar kullanılır. Bu işlemin uygulanışı, .html dosyamızın içinde Golang tarafından gelecek öğeler için işaret bırakırız. Bu işaret bu şekilde olur: `{{ kodumuz }}`  
Hemen bir örnek ile olayı anlayalım. **main.go** dosyamız şöyle olsun.

```go
package main
import (
	"fmt"
	"html/template"
	"net/http"
)
// SayfaBilgi ...
type SayfaBilgi struct {
	Başlık string
	İçerik string
}
func anasayfa(w http.ResponseWriter, r *http.Request) {
	sayfa := SayfaBilgi{Başlık: "Golang Türkiye", İçerik: "Sitemize Hoşgeldiniz"}
	şablon, _ := template.ParseFiles("sablonumuz.html")
	şablon.Execute(w, sayfa)
}
func main() {
	fmt.Println("Program Başladı")
	http.HandleFunc("/", anasayfa)
	http.ListenAndServe(":8000", nil)
}
```

İlk olarak sunucu oluşturacağımız için **“net/html”** ve şablon oluşturacağımız için de **“html/template”** paketlerini içe aktarıyoruz.  
**SayfaBilgi** adında bir **struct** metod oluşturuyoruz ve içerisine **string** değer alan **Başlık** ve **İçerik** türünü oluşturuyoruz. Bunu yapmamızın sebebi web sayfamızda sayfamızın başlığını ve içeriğini bunlar aracılığıyla şablona göndereceğiz.  
**anasayfa** adında fonksiyonumuzun içerisini inceleyelim. Bu fonksiyonumuz bir sayfa yakalayıcı fonksiyondur. **net/http** paketi için alttaki konuyu okuyabilirsiniz.

net-http-ile-web-server-olusturma.md

**sayfa** adında değişken oluşturuyoruz ve bu değişkenin **SayfaBilgi** strunct’ından olduğunu belirtip içerisine sayfa bilgilerimizi giriyoruz.  
**şablon** değişkeni oluşturduk. \(\_ alt tire yerine hata bilgilerini alan değişken koyabilirsiniz.\) **template.ParseFiles\(\)** fonksiyonu ile HTML şablonumuzu tanıttık.  
Hemen altında şablonumuzu çalıştırması için **Execute\(\)** fonksiyonundan yaralandık.  
**main\(\)** fonksiyonumda ise klasik bir server ayağa kaldırma kodları yer alıyor.  
Şimdi de **sablonumuz.html** dosyasını görelim.

```markup
<h1>{{.Başlık}}</h1>
{{.İçerik}}
```

Süslü parantezler içerisin **nokta** ile başlayan değişken yerleştirmelerini yapıyoruz. Bu değişkenler bize **SayfaBilgi** struct’ından gelmektedir.  
   
Seviyeyi biraz daha yükseltelim ve listeleme işlemi yapalım.  
**main.go** dosyamızı aşağıdaki gibi oluşturalım.

```go
package main
import (
	"fmt"
	"html/template"
	"net/http"
)
// Görev ...
type Görev struct {
	İsim       string
	Tamamlandı bool
}
// SayfaVerisi ...
type SayfaVerisi struct {
	Sayfaİsmi    string
	GörevListesi []Görev
}
func anasayfa(w http.ResponseWriter, r *http.Request) {
	sayfa := SayfaVerisi{
		Sayfaİsmi: "Görevler Listesi",
		GörevListesi: []Görev{
			{İsim: "Ekmek Al", Tamamlandı: false},
			{İsim: "Kola Al", Tamamlandı: true},
			{İsim: "Yoğurt Al", Tamamlandı: false},
		},
	}
	şablon, _ := template.ParseFiles("sablonumuz.html")
	şablon.Execute(w, sayfa)
}
func main() {
	fmt.Println("Program Başladı")
	http.HandleFunc("/", anasayfa)
	http.ListenAndServe(":8000", nil)
}
```

Bu sefer farkettiyseniz 2 tane struct metodumuz var. **Görev** struct’ımız içerisinde 2 tane değişkene ev sahipliği yapacak. **SayfaVerisi** struct’ında ise **Sayfaİsmi** ve **GörevListesi** adında elemanlar var. **GörevListesi** elemanı **Görev** struct’ı türündedir. Bu sayede içerisine dizi olarak görevler kaydedebileceğiz.  
Bir önceki örnekteki gibi **anasayfa** yakalayıcı fonksiyonumuzu oluşturuyoruz. İçerisinde **sayfa** isminde değişken oluşturuyoruz. Bu değişken içerisine sayfamızda görünmesini istediğimiz **Sayfaİsmi** ve **GörevListesi** elemanlarını giriyoruz. Hemen aşağısında ise şablonumuzu bağlama işlemlerini yapıyoruz.  
**main\(\)** fonksiyonumuz ise bir önceki örnek ile aynıdır.  
Şimdide **sablonumuz.html** dosyasını görelim.

```markup

<style>
.kirmizi{
    color:red;
}
.yesil{
    color:green;
}
</style>
<h1>{{.Sayfaİsmi}}</h1>
<ul>
    {{range .GörevListesi}}
        {{if .Tamamlandı}}
            <li class="yesil">{{.İsim}}</li>
        {{else}}
            <li class="kirmizi">{{.İsim}}</li>
        {{end}}
    {{end}}
</ul>
```

Yukarıdaki kodları incelediğinizde içerisinde **range**, **if**, **else** ve **end** gibi kelimeler göreceksiniz. **range** anahtar kelimesi Golang’taki gibi belirtilen dizinin uzunluğu kadar sıralama işlemi yapar. **{{range}}** anahtar kelimesi mutlaka **{{end}}** ile kapatılmalıdır. if-else akışının ne işe yaradığını zaten biliyorsunuzdur. Aynı şekilde **{{end}}** ile kapatılmalıdır.  
Bu işlemlerin sonunda elimize şöyle bir sonuç gemiş olacaktır.  
Tarayıcımızda http://localhost:8000 adresini açıyoruz.

### Görevler Listesi

* Ekmek Al
* Kola Al
* Yoğurt Al



**İşte kullanabileceğimiz bazı şablon kodları:**  


| Şablon Kodu | Açıklama |
| :--- | :--- |
| {{/\* yorum \*/}} | Yorum yapmak için kullanılır. |
| {{.}} | Golang’tan ana değişken için kullanılır. |
| {{.Değişkenİsmi}} | Golang’tan belirli bir değişken almak için kullanılır. |
| {{if .Tamamlandı}} {{else}} {{end}} | If-Else akışı için kullanılır. |
| {{block “içerik” .}} {{end}} | İçerik ismine bloklama tanımlar. |
| {{range .Görevler}} {{.}} {{end}} | Dizi sıralamak için kullanılır. |

## Statik Kütüphanesi ile Dosyaları Uygulamaya Gömme

Golang’ın müthiş yanlarından biri de bir uygulamayı build ettiğimizde bize tek çalıştırılabilir dosya şeklinde sunmasıdır. Fakat bu özellik sadece **.go** dosyalarını birleştirilmesinde kullanılıyor. Harici dosyalar programa gömülmüyor. Fakat **Statik** isimli kütüphane ile bu işlem mümkün kılınıyor.  
Kütüphanenin mantığından kısaca bahsedeyim. Belirlediğiniz bir dizindeki dosyaları bir kodlamaya çevirerek programın içine dosya gömmek yerine kod gömüyor. Ve bu kodu sanki dosyaymışcasına kullanabiliyoruz. Tabi ki sadece sunucu işlemlerinde işe yarar olduğunu belirtelim.  
Bu yöntemin güzel artı yönleri var.

* Programımız tek dosya halinde kalıyor.
* Programımız kapalı kaynak oluyor.

Tanıtımını yaptığımıza göre hafiften uygulamaya başlayalım.

> go get github.com/rakyll/statik

Konsola yukarıdakini yazarak kütüphanemizi indiriyoruz. Öncelikle dosya ve klasör yapımızı aşağıdaki gibi ayarlıyoruz.

![Proje klas&#xF6;r&#xFC;m&#xFC;z&#xFC;n dizin d&#xFC;zeni](./goruntuler/statik.png)

Kodlamaya dönüştürülmesini istediğimiz klasör ile işlem yapıyoruz. Yani **public** klasörü ile. Aşağıdaki komutu **Proje klasörümüz** içerisindeyken yazıyoruz.

> statik -src=/public/klasörünün/adresi -f

Bu işlemle birlikte **public** klasörümüzün yanına statik isimli bir klasör oluşturduk ve içine **statik.go** isimli dosya oluşturmuş olduk. Bu dosyanın içerisinde bizim **public** klasörümüzün kodlanmış hali mevcuttur.  
Ve sırada **main.go** dosyamızı oluşturmakta. Aşağıdaki kodlarıda **main.go** dosyamıza yazıyoruz.

```go
package main
import (
    "net/http"
    "github.com/rakyll/statik/fs"
    _ "./statik" //Oluşturulmuş statik.go dosyasının konumu
)
func main() {
    statikFS, _ := fs.New()
    http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
    http.ListenAndServe(":5555", nil)
}
```

Gerekli kütüphanelerimizi ekledikten sonra **main\(\)** fonksiyomuzun içeriğini inceleyelim.

**statikFS** ve **\_** adında değişkenlerimizi tanımladık. Bu değişkenlerimizi fonksiyonel değişkendir. **\_** koymamızın sebebi **error** çıktısını kullanmak istemediğimizdendir. Eğer lazım olursa kullanabilirsiniz. **fs.New\(\)** diyerek **statikFS** değişkenimizi bir dosya sistemi olarak tanıttık. Daha sonra sunucu oluşturak anadizine ulaşılmak istendiğinde oluşturduğumuz dosya sistemine bağlanmasını sağladık. Artık dosya sistemimize **localhost:5555** üzerinden ulaşılabilir oldu.

## Gin Web Kütüphanesi

Gin Go'da yazılmış bir web kütüphanesidir. Performans ve üretkenlik odaklıdır. Sizlere basitçe web sunucu ve api oluşturmanız için kolaylık sağlar.

Kurulum için:

> go get -u github.com/gin-gonic/gin

Daha sonra yine aynı yöntemle projemize dahil edebiliriz.

```go
import "github.com/gin-gonic/gin"
```

Basit bir web sunucu oluşturma örneği:

```go
package main

import (
	// kütüphanemizi içeri aktaralım
	"github.com/gin-gonic/gin"
)

func main() {
	//gin'in varsayılan ayarlarında bir yönlendirici oluşturalım.
	router := gin.Default()

	//anasayfayı inde fonksiyonumuz yakalayacak
	router.GET("/", index)

	//daha sonra sunucuyu başlatıyoruz
	router.Run()
}

//anasayfayı yakalayacak olan fonksiyonumuz
func index(c *gin.Context) {
	//c ile gin nesnemize bağlam oluşturduk.
	//c'yi kullanarak artık gin özelliklerine erişebiliriz.

	//sayfaya düz yazı gönderdik
	c.String(200, "Merhaba Dünya")
	//Buradaki 200 sunucudan bir cevap geldiğini anlamına gelir
}
```

Programımızı çalıştırdığımızda aşağıdaki gibi konsol çıktısı alacağız.

> \[GIN-debug\] \[WARNING\] Creating an Engine instance with the Logger and Recovery middleware already attached.
>
> \[GIN-debug\] \[WARNING\] Running in "debug" mode. Switch to "release" mode in production.
>
> * using env:   export GIN\_MODE=release
> * using code:  gin.SetMode\(gin.ReleaseMode\)
>
> \[GIN-debug\] GET / --&gt; main.index \(3 handlers\)  
> \[GIN-debug\] Listening and serving HTTP on :8080

Bu çıktıyı incelediğimizde, Gin'in debug \(hata ayıklama\) modunda çalıştığını söylüyor ve hemen aşağısında sunucumuz ürün haline gelince Gin'i Release Moduna nasıl alacağımızı gösteriyor. Son olarak ise web sunucumuzun `8080` portunda çalıştığınız gösteriyor.

Yukarıdaki örnekte web sunucumuz varsayılan olarak `8080` protunda çalışacaktır. Bunun sebebi `router.Run()`'a parametre olarak port numarası vermememizdir.

Örneğe göre [http://localhost:8080](http://localhost:8080) adresine gittiğimizde komut satırında yeni detaylar belirecek. Tıpkı aşağıdaki gibi:

![Komut sat&#x131;r&#x131; bilgisi](./goruntuler/go-gin.png)

Bu bilgileri inceleyelim. İlk kayıt anasayfaya bağlanılmaya çalışıldığında alınmış. Bu kayıtta bağlantının zamanını, durum kodunu, bağlantı süresi, bağlantı yöntemini ve hangi adrese bağlantı denendiğini yazıyor. Hemen altındaki ise sitenin ikonuna istek yapmış fakat site ikonumuz bulunmadığı için **404 durum kodu**nu almış. Bu kısımdan da bağlantı isteklerini görebildiğimizi öğrenmiş olduk.

### Çıktı Tipleri

#### JSON Çıktı Verme

```go
func index(c *gin.Context) {
	//JSON Fonksiyonunu kullanıyoruz.
	c.JSON(200, gin.H{
		"ad":    "kaan",
		"soyad": "kuşcu",
	})
}
```

Sonucumuz aşağıdaki gibi olacaktır.

![JSON &#xE7;&#x131;kt&#x131;s&#x131;](./goruntuler/gin-json.png)

#### XML Çıktı Verme

```go
//xml için örnek bir yapı oluşturalım
type xmlYapı struct {
	Ad    string `xml:"ad"`
	Soyad string `xml:"soyad"`
}

//anasayfayı yakalayacak olan fonksiyonumuz
func index(c *gin.Context) {
	//xml için örnek bir nesne oluşturduk
	xmlOrnek := xmlYapı{"kaan", "kuşcu"}

	//xml başlığını gönderelim
	c.Writer.WriteString(xml.Header) //<?xml version="1.0" encoding="UTF-8"?>

	//xml nesnesini XML fonksiyonu ile yolladık
	c.XML(200, xmlOrnek)
}
```

Bu kodlar sonucunda sayfamızı açtığımızda "kaankuşcu" sonucu göreceğiz. XML tipinde görmek için sayfanıza sağ tıklayıp "_Sayfa Kaynağını Gör_"e tıklayarak kontrol edebilirsiniz.

![XML sonucu](./goruntuler/image%20%282%29.png)



#### Template Kullanımı

Template hakkında bilginiz yoksa önce aşağıdaki dökümanı okumanız önerilir.

html-sablonlar-templates.md

Gin'de template \(şablon\) işlemleri bayağı kolaylaştırılmış. Ufak bir örnek uygulama yazalım. Öncelikle projemizin ana dizinine `templates` isimli bir klasör oluşturalım ve içerisine `index.html` dosyası oluşturalım. `index.html` dosyamızın içeriği ise aşağıdaki gibi olsun.

```markup
<html>
	<h1>
		{{ .başlık }}
	</h1>
</html>
```

Burada `{{ .başlık }}` yerine Go'dan değer göndereceğiz.

`main.go` dosyamız ise aşağıdaki gibi olsun.

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Burada templates klasörünün içindeki tüm şablonları
	//yüklemesini isteyelim.
	router.LoadHTMLGlob("templates/*")

	router.GET("/", index)

	router.Run()
}

func index(c *gin.Context) {

	//HTML şablonunu almak için
	//HTML fonksiyonunu kullanıyoruz.
	c.HTML(200, "index.html", gin.H{
		//Şablondaki başlık yerine Anasayfa yazısını yollayalım.
		"başlık": "Anasayfa",
	})
}
```

Web sunucumuza bağlandığımızda ise **Anasayfa** yazdığını görebiliriz.

### Statik Dosyaları Yayınlama

Web sunucumuzda kullanacağımız Css, JS vb. statik dosyalarımız olabilir. Bunun için `Static` fonksiyonunu kullanabiliriz.

Statik dosyalarımızı projemizin ana dizindeki `statik` klasöründe barındırdığımızı varsayalım.

```go
func main() {
	router := gin.Default()

	//(yönlendirme, klasör-ismi)
	router.Static("/static", "./statik")
	
	router.GET("/", index)

	router.Run(":9000")
}
```

`statik` klasörümüzün içerisinde `index.js` adında bir dosya olduğunu varsayarsak `http://localhost:9000/static/index.js` adresinden ulaşabiliriz.

### Bağlantı Metodları

Örnek bağlantı oluştururken GET metoduna değindik. Metodları test ediyorken **Postman**'i kullanabilirsiniz. Ben bu konuda **curl** komut satırı aracını kullanacağım. Detaylarına bakacak olursak:

#### GET Metodu

`GET` metodu web sunucumuza normal bağlantı yapılırken kullanılır. Hazır bir kaynağı yüklemek için kullanılır.

```go
router.GET("/", index)
```

`index` fonksiyonu ile `GET` metodlu anasayfayı yakalayabilirsiniz.

#### POST Metodu

POST metodu genellikle form gönderimlerinde kullanılır. Yeni bir kaynak oluşturmak için kullanılır. \(Yeni kayıt oluşturma, yeni gönderi oluşturma vb...\)

Örnek kullanımını görelim.

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Burada templates klasörünün içindeki tüm şablonları
	//yüklemesini isteyelim.
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndex)
	router.POST("/", postIndex)

	router.Run(":9000")
}

func getIndex(c *gin.Context) {
	c.String(200, "GET metodu ile bağlanıldı.")
}

func postIndex(c *gin.Context) {
	c.String(200, "POST metodu ile bağlanıldı.")
}
```

Yukarıdaki örnekte anasayfa `GET` ile bağlanıldığında `getIndex`, `POST` ile bağlanıldığında `postIndex` fonksiyonu çalışacak. Tarayıcımızdan girdiğimizde "_GET metodu ile bağlanıldı._" yazısını görürüz. `POST` metodu ile bağlanmak için komut satırına şu komutları yazalım.

> curl -X POST http://localhost:9000

Çıktısı "_POST metodu ile bağlanıldı._" olacaktır.

POST metodu üzerinden değer almayı görelim.

```go
//json verisi için yapımız
type kişi struct {
	Ad    string `json:"ad"`
	Soyad string `json:"soyad"`
}

func postIndex(c *gin.Context) {
	//posttan gelen json'ın kaydedileceği değişken
	var postkişi kişi

	//postan gelen json'ı postkişi'ye atayalım
	c.BindJSON(&postkişi)


	c.String(200, "JSON Veri:")
	//json'ı tekrar post ile gösterelim
	c.JSON(200, postkişi)
}
```

Komut satırına aşağıdaki komutu yazarak çıktısını görebilirsiniz.

> curl -X POST -H "Content-Type: application/json" -d '{"ad":"kaan","soyad":"kuşcu"}' [http://localhost:9000](http://localhost:9000)

#### Diğer Metodlar

Diğer metodlardan kısaca bahsedelim:

* **PATCH metodu:** Bir kaynak üzerindeki belirli bir alanı değiştirmek için kullanılır.
* **DELETE metodu:** Sunucudaki bir kaynağı silmeye yarar.
* **PUT metodu:** Bir kaynağın yerine başka bir kaynağı koymaya yarar. \(Komple değiştirme\)
* **HEAD metodu:** Sunucuya tıpkı GET metodu gibi fakat sadece başlığı olan bir istek gönderir.
* **OPTIONS metodu:** Sunucunun desteklediği metodları kontrol etmek için kullanılır.

### Adreslendirme

#### Parametre ile Adreslendirme

Örneğin:

![Gin parametre &#xF6;rne&#x11F;i](./goruntuler/gin-param.png)

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/blog/:yazı", blog)
	//Buradaki :yazı bizim parametremiz
	//Bu parametre ile hangi blog yazısını
	//göstereceğimizi belirleyeceğiz.

	router.Run(":9000")
}
func blog(c *gin.Context) {
	//yazı parametresininde geleni sayfaİsmi değişkenine atayalım.
	sayfaİsmi := c.Param("yazı")
	c.String(200, "Şuanda "+sayfaİsmi+" blogunu okuyorsunuz.")
}
```

Yukarıdaki örneğe göre [http://localhost:9000/blog/gin%20ile%20sunucu%20geli%C5%9Ftirme](http://localhost:9000/blog/gin%20ile%20sunucu%20geli%C5%9Ftirme) adresine gittiğimizde "_Şuanda gin ile sunucu geliştirme blogunu okuyorsunuz._" yazısı ile karşılacağız.

Tabi ki birden fazla parametre ekleyebilirsiniz.

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/blog/:yazar/:yazı", blog)

	router.Run(":9000")
}
func blog(c *gin.Context) {
	yazar := c.Param("yazar")
	yazı := c.Param("yazı")
	c.String(200, "Yazar: "+yazar+" Yazı: "+yazı)
}
```

#### Querystring \(Sorgu dizesi\) ile Adreslendirme

Örneğin:

![Gin sorgu sizesi &#xF6;rne&#x11F;i](./goruntuler/gin-querystring.png)

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/arama", arama)

	router.Run(":9000")
}
func arama(c *gin.Context) {
	tür := c.Query("tur")
	sıralama := c.Query("siralama")
	c.String(200, tür+" türünden filmler "+sıralama+" olarak sıralanıyor.")
}
```

Yukarıdaki örneğe göre [http://localhost:9000/arama?tur=bilim-kurgu&siralama=imdb](http://localhost:9000/arama?tur=bilim-kurgu&siralama=imdb) adresine girdiğimizde "_bilim-kurgu türünden filmler imdb olarak sıralanıyor._" yazılı bir sonuç elde edeceğiz.



## gRPC

### gRPC Nedir?

gRPC, Google tarafından geliştirilen, açık kaynaklı uzaktan prosedür çağırma \(rpc\) kütüphanesi \(sistemi\)'dir. Aktarım için `HTTP/2` kullanır. Arayüz tanımlama dili olarak `protokol tamponları (buffers)` kullanır. Kimlik doğrulama, Çift yönlü akış, engelleme ve engelleme olmayan bağlantılar, iptal ve zaman aşımı işlemleri için kullanılır.

JSON ve HTTP API'larının aksine, daha katı bir spesifikasyona sahiptirler. Tek bir spesifikasyona sahip olduğu için, gRPC tartışmalarını ortadan kaldırır ve geliştiriciye zaman kazandırır. Çünkü gRPC, platformlar ve uygulamalar arasında tutarlıdır.

Özet olarak, uzaktan prosedür çağırarak client \(müşteri\) ve server \(sunucu\) programları arasındaki iletişimi sağlayan bir kütüphanedir.

### Protokol Tamponları \(Buffers\) Nedir?

İki uygulama arasındaki iletişimin nasıl olacağını belirleyen sözleşmedir.

### Örnek Uygulama Yapalım

Bu örneğimizde sunucuya mesaj gönderip, karşılığında mesaj alan bir uygulama yazacağız. Bir nevi chat uygulaması olacak.

Öncelikle, iletişim protokülünü Go koduna dönüştürebilmemiz \(generating\) için `protoc`'yi bilgisayarımıza kuralım.

İndirmek için [buradaki](https://github.com/protocolbuffers/protobuf/releases/tag/v3.13.0) adrese gidin. Buradaki versiyon v3.13.0 versiyonu. Daha güncel versiyonlar için [buradaki](https://github.com/protocolbuffers/protobuf/releases/) sayfayı kontrol etmeyi unutmayın.

![](./goruntuler/protoc-download.png)

GNU/Linux İS kullanan arkadaşlarımız uygulama depolarından protobuf adıyla yükleyebilirler. Böylece PATH olarak eklemelerine gerek kalmaz.

### Sisteme Yol Olarak Ekleme

Eğer GitHub sayfasından indirdiyseniz, sisteminize dosya yolunu tanıtmanız gerekir. Bunun için örnek olarak;

#### Windows üzerinde

Örnek olarak `protoc.exe` dosyamız `C:\\protoc\bin` klasörü içerisinde olsun.

Başlar menüsünde _"Ortam değişkenleri"_ yazarak aratalım. Açılan pencerede _Gelişmiş_ sekmesinde alt tarafta _Ortam Değişkenleri_ butonuna basalım.

_PATH_ seçeneğini _Düzenle_ diyerek `protoc.exe`'nin konumunu ekleyelim.

#### GNU/Linux, MacOS vs. Bash Kullanan Sistemler Üzerinde

`protoc` çalıştırılabilir dosyamızı örnek olarak `Home (Ev)` dizinine atmış olalım.

**Örnek:** `~/protoc/bin` olsun.

Ev dizinimizdeki `.bashrc` dosyamızın en altına şunları ekleyelim.

```text
export PATH="~/protoc/bin:$PATH"
```

`.bashrc` dosyasını kaydettikten sonra komut satırına `source ~/.bashrc` yazarak dosyada yaptığımız değişimi onaylayalım.

##### Protoc'yi kontrol edelim

Komut satırına aşağıdakileri yazarak `protoc`'nin versiyonuna bakalım.

```text
protoc --version
```

`Protoc`'ye Golang desteğini eklemek için aşağıdaki paketleri kuralım.

```text
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go-grpc
```

Aynı şekilde Windows üzerinde `C:\\Users\isim\go\bin` klasörünü ekli değilse ortam değişkenlerine ekleyelim. Unix-like sistemler için `~/go/bin`'i path'e ekleyelim.

Yukarıda belirttiğimiz örnek `chat` uygulamasını yazmak için aşağıdaki gibi bir proje yapımız olacak.

```text
.
├── chat
│   └── chat.go
├── chat.proto
├── client.go
└── server.go
```

Client ve Server arasındaki iletişim protokolünü belirleyelim.

`chat.proto` dosyamız şöyle olsun.

```groovy
//Söz dizimi için proto3'ü kullanacağız
syntax = "proto3";

//Go paketi için ayarlarımız
package chat;
option go_package=".;chat";

//iletişimde kullanılacak mesajın yapısı
message Message {
    string body = 1;
    //buradaki 1 sıra numarasıdır. Değer ataması değildir.
}

//rpc servisimiz
service ChatService {
    rpc SayHello(Message) returns (Message) {}
}
```

Yukarıdaki `chat.proto` dosyasını Go koduna dönüştürelim. Bunun için proje dizinindeyken aşağıdaki komutu yazalım.

> protoc --go\_out=plugins=grpc:chat chat.proto

Bu işlem sonucunda chat klasörümüzüm içerisinde `chat.pb.go` dosyamız oluşacak. Bu dosyanın en başında yorum olarak düzenleme yapmamızın uygun olamayacağı yazıyor. O yüzden bu dosyayla çok uğraşmamakta fayda var.

### Server \(Sunucu\) Oluşturalım

`chat` klasörümüzdeki `chat.go` dosyasını oluşturalım.

```go
package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	//Burası şuanlık boş olacak
}

//server nesnemize iliştirilecek fonksiyonumuz
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	//Client'ten gelen mesajı ekrana bastıralım
	log.Printf("Clientten mesaj alındı: %s", message.Body)

	//mesaj gelince değer döndürelim (mesaj ve hata)
	return &Message{Body: "Server'dan merhaba!"}, nil //hata nil olsun
}
```

Go modules'tan faydalanarak chat paketini diğer paketlerde kullanabiliriz. Bunun için proje dizinindeyken komut satırına:

> go mod init github.com/ksckaan1/grpcOrnek

yazalım. Buradaki yaptığımız şey, GitHub'a paket yüklemeden sanal olarak yazdığımız `chat` paketini kullanabileceğiz. Bu komuttan sonra proje dizinimizde `go.mod` dosyası oluşacak. Böylece chat klasörünü paket olarak diğer yerlerde de kullanabileceğiz.

`server.go` dosyamızı oluşturalım.

```go
package main

import (
	"log"
	"net"

	"github.com/ksckaan1/grpcExample/chat"
	"google.golang.org/grpc"
)

func main() {

	//tcp dinleyelim
	lis, err := net.Listen("tcp", ":9080")

	//hata kontrolü yapalım
	if err != nil {
		log.Fatalf("9080 portu dinlenirken hata oluştu: %v \n", err)
	}

	//chat server nesnesi oluşturalım
	s := chat.Server{}

	//grpc server'ı oluşturalım.
	grpcServer := grpc.NewServer()

	//tcp sunucu ile grpc'yi bağlayalım.
	chat.RegisterChatServiceServer(grpcServer, &s)

	//hata kontrolü
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpcServer oluşturulurken hata : %v", err)
	}
}
```

`go run server.go` komutunu yazarak test edelim. Eğer çıktı vermiyorsa ve program kapanmıyorsa server dinleniyor demektir. Yani şuana kadar başarılıyız.

### Client \(Müşteri\) Oluşturalım

Şimdi de Server'a mesaj yollayabilmek için Client'i oluşturalım. `client.go` dosyamız aşağıdaki gibi olsun.

```go
package main

import (
	"log"

	//chat klasörünü kullanacağımız için
	//go mod init üzerinde belirlediğimiz deponun
	//sonuna /chat ekliyoruz.
	"github.com/ksckaan1/grpcExample/chat"
	"golang.org/x/net/context"

	//grpc kütüphanemiz
	"google.golang.org/grpc"
)

func main() {

	//grpc client bağlantı nesnesi
	var conn *grpc.ClientConn

	//grpc ye 9080 portunu dinlemesini söyleyelim
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("9080 portu dinlenirken hata oluştu: %v", err)
	}

	//bağlantığı kapatmayı unutmayalım
	defer conn.Close()

	//Client nesnesi oluşturalım
	c := chat.NewChatServiceClient(conn)

	//server'a gönderilecek mesajı belirleyelim
	message := chat.Message{
		Body: "Client'ten merhabalar!",
	}
	
	//Server'dan dönen cevabı alalım
	response, err := c.SayHello(context.Background(), &message)

	//hata kontrolü yapalım.
	if err != nil {
		log.Fatalf("SayHello fonksiyonu çağırılıken hata oluştu: %v", err)
	}

	//server'dan gelen mesajı ekrana bastıralım
	log.Printf("Server'dan gelen cevap: %s", response.Body)
}
```

Yukarıda yazdığımız kodların çalışma mantığını açıklayacak olursak:

1. Server `9080` portunu belirlediğimiz protokol sözleşmesine \(_chat.proto_\) göre dinliyor.
2. `Client.go` dosyamız çalıştırılınca server'a mesaj yolluyor.
3. Client'ten gelen mesajı server ekrana bastırıyor ve bunun karşılığında server client'e cevap veriyor.
4. Server'dan gelen cevabı da client ekrana bastırıyor.
5. Son olarak client programımız sonlanıyor \(kapanıyor\).

Gif olarak sonucu göstermek gerekirse:

![gRPC &#xD6;rne&#x11F;i](./goruntuler/golang-grpc-example.gif)

## Heroku'da Go Uygulaması Yayınlama

### Öncelikle Bilmeyenler İçin Heroku Nedir?

Kısaca Heroku, JavaScript, Go, Ruby, Python, Scala, PHP, Java, Clojure ile geliştirdiğimiz sunucu uygulamalarını ücretsiz barındırabileceğimiz bir platformdur.

Aşağıdaki bağlantıdaki blog yazısını okumanızı tavsiye ederim.

{% embed url="https://ceaksan.com/tr/heroku-nedir

### Projemizi Planlayalım

Bu örneğimizde bir web sunucu oluşturacağız. Öncelikle `Go modules` kullanacağımız için projemizi kullanıcının go dizinine oluştumalıyız.

Komut satırını açalım ve aşağıdaki komutu yazarak bahsettiğimiz dizine geçelim.

Windows'ta:

> cd C:\\Users\%username%\go\src

GNU/Linux ve MacOS'te:

> cd ~/go/src

Bu konuma proje dizinimizi oluşturalım

> mkdir heroku-app

`heroku-app` klasörü projemizin ana dizini olacak. Aşağıdaki komut ile proje ana dizinimize girelim.

> cd heroku-app

Daha sonra bu dizini `code .` komutu ile VSCode üzerinde açalım.

`main.go` dosyamızı oluşturalım ve aşağıdaki gibi olsun.

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", anaSayfa)
	http.ListenAndServe(":"+port, nil)
}

func anaSayfa(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	fmt.Fprintf(w, "Merhaba Dünya!\nKullanılan Port: "+port)
}
```

Yukarıda normal web sunucu oluşturma kodlarından biraz farklı işlemler var. Bunları açıklamak gerekir ise:

`port` değişkenimiz sistemden `string` tipinde `PORT` ortam değişkenini alıyor. Yani `port` değişkeni sunucumuzun hangi portta çalışacağını belirliyor. Uygulamamızı Heroku'ya yükledikten sonra sistemimiz Heroku olduğu için port ortam değişkenini Heroku'dan almış olacağız. Sunucunun çalışacağı portu Heroku belirlediği için portu kendimiz kodlar içinde belirleyemiyoruz.

`http.ListenAndServe()` fonksiyonuna da parametre olarak `":"+port` vererek devam ediyoruz.

Sunucumuzun ana dizinini yakalacak olan anaSayfa fonksiyonumuza bakalım.

Yine burada sistemden portu istedik. Hemen aşağısında "Merhaba Dünya!" ve kullanılan portun çıktısını vermesini sağladık. Kodlarımız artık hazır.

Bu projemizde dışarıdan bir pakete ihtiyacımız olmadı. Hepsi Go'nun hazır paketlerinden. Eğer dışarıdan paketler olsa ne yapacaktık? Hadi hemen görelim.

Komut satırına go modules için aşağıdaki komutu yazalım.

```bash
go mod init
```

> Eğer projenizi go/src klasörü içinde oluşturmazsanız bu komut hata verecektir.

![go modules örnek](./goruntuler/go-mod-init.png)

Böylece go.mod dosyamızı oluşturduk. Dışarıdan paket bağımlılıklarını yüklemek için aşağıdaki komutu yazalım.

```bash
go mod vendor
```

Bu komutu yazdığınızda paket bağımlılığınız yoksa aşağıdaki gibi bir çıktı alacaksınız.

![vendor &#xF6;rne&#x11F;i](./goruntuler/no-vendor.png)

Eğer paket bağımlılığınız varsa projenizin ana dizininde vendor adında bir klasör oluşacak ve bu klasörün içinde dış paketlerin kaynak kodları bulunanacak.

### Versiyon sistemli hale getirelim

Heroku platformu versiyon kontrol sistemi ile çalıştığı için, öncelikle git projemizi oluşturalım. Projemizin ana dizinindeyken komut satırına:

> git init

Daha sonra oluşturduğumuz projeyi staging'e almak için:

> git add .

yazalım. Commit oluşturmak için ise:

> git commit -m "Heroku uygulamamı oluşturdum."

### Heroku'da Yayınlama

Öncelikle Heroku'nun komut satırı uygulamasını bilgisayarımıza kuralım.

Windows, MacOS ve Ubuntu için [bu adresten](https://devcenter.heroku.com/articles/heroku-cli#download-and-install) kurabilirsiniz.

Arch Linux ve türevleri için kolaylık açısından AUR üzerinden `heroku-cli-bin` aratarak kurabilirsiniz.

Uygulamayı kurduktan sonra Heroku Hesabımıza bağlayalım.

Komut satırına aşağıdakileri yazalım.

> heroku login

Şöyle bir çıktı verecek:

![Heroku cli Login](./goruntuler/heroku-login.png)

q tuşuna basınca giriş yapmayı iptal eder. O yüzden giriş yapmak için herhangi bir tuşa başabilirsiniz. \(Lütfen bilgisayarınızın güç tuşuna basmayın 🙂\)

Daha sonra varsayılan tarayıcınız üzerinden giriş yapma sayfası açılacak. Heroku hesabınıza girdikdek sonra tarayıcınızda girişin başarılı olduğunu söylecek.

![Heroku taray&#x131;c&#x131; giri&#x15F;i](./goruntuler/heroku-login-warning.png)

Komut satırında da aşağıdaki gibi bir çıktı göreceksiniz. Kendi bilgilerim olduğu için birazını sansürledim.

![Heroku cli ba&#x15F;ar&#x131;l&#x131; giri&#x15F; ](./goruntuler/heroku-login-success.png)

Böylece başarıyla giriş yapmış olduk.

Heroku projemizi oluşturalım.

> heroku create

Şöyle bir çıktı alacağız.

![heroku uygulama olu&#x15F;turma](./goruntuler/heroku-create.png)

Yazdığımız kodları Heroku uygulamamıza yükleyelim.

> git push heroku master

Bu komut sonrasında aşağıdakine benzer bir sonuç almanız gerekir.

```go
[kaanksc@KAANHP heroku-app]$ git push heroku master
Enumerating objects: 4, done.
Counting objects: 100% (4/4), done.
Delta compression using up to 4 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 477 bytes | 477.00 KiB/s, done.
Total 4 (delta 0), reused 0 (delta 0), pack-reused 0
remote: Compressing source files... done.
remote: Building source:
remote: 
remote: -----> Go app detected
remote: -----> Fetching jq... done
remote: -----> Fetching stdlib.sh.v8... done
remote: -----> 
remote:        Detected go modules via go.mod
remote: -----> 
remote:        Detected Module Name: heroku-app
remote: -----> 
remote:  !!    The go.mod file for this project does not specify a Go version
remote:  !!    
remote:  !!    Defaulting to go1.12.17
remote:  !!    
remote:  !!    For more details see: https://devcenter.heroku.com/articles/go-apps-with-modules#build-configuration
remote:  !!    
remote: -----> New Go Version, clearing old cache
remote: -----> Installing go1.12.17
remote: -----> Fetching go1.12.17.linux-amd64.tar.gz... done
remote: -----> Determining packages to install
remote:        
remote:        Detected the following main packages to install:
remote:                 heroku-app
remote:        
remote: -----> Running: go install -v -tags heroku heroku-app 
remote: heroku-app
remote:        
remote:        Installed the following binaries:
remote:                 ./bin/heroku-app
remote:        
remote:        Created a Procfile with the following entries:
remote:                 web: bin/heroku-app
remote:        
remote:        If these entries look incomplete or incorrect please create a Procfile with the required entries.
remote:        See https://devcenter.heroku.com/articles/procfile for more details about Procfiles
remote:        
remote: -----> Discovering process types
remote:        Procfile declares types -> web
remote: 
remote: -----> Compressing...
remote:        Done: 3.6M
remote: -----> Launching...
remote:        Released v3
remote:        https://obscure-ocean-33068.herokuapp.com/ deployed to Heroku
remote: 
remote: Verifying deploy... done.
To https://git.heroku.com/obscure-ocean-33068.git
 * [new branch]      master -> master
```

Yukarıdaki çıktıya göre aşağıdaki işaretlediğim yerde uygulamamızın adresi olacak.

![Heroku Push sonu&#xE7;](./goruntuler/heroku-push.png)

Bu adres tabiki de sizde farlı olacak. Buradan girip uygulamanızı kontrol edebilirsiniz. Benim sonucum ise şu şekilde:

![Site sonucu](./goruntuler/heroku-site.png)

