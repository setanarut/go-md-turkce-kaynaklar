- [Bölüm 1 - Ön Eğitim](#bölüm-1---ön-eğitim)
  - [Golang Hakkında](#golang-hakkında)
  - [Sıkça Sorulan Sorular](#sıkça-sorulan-sorular)
    - [Kökenler](#kökenler)
      - [Projenin amacı nedir?](#projenin-amacı-nedir)
      - [Projenin tarihçesi nedir?](#projenin-tarihçesi-nedir)
      - [Gopher maskotu nereden geliyor?](#gopher-maskotu-nereden-geliyor)
      - [Dilin adı Go mu Golang mi?](#dilin-adı-go-mu-golang-mi)
      - [Neden yeni bir dil yarattınız?](#neden-yeni-bir-dil-yarattınız)
      - [Go'nun ataları nelerdir? \(veya esinlendikleri\)](#gonun-ataları-nelerdir-veya-esinlendikleri)
      - [Tasarımda yol gösterici ilkeler nelerdir?](#tasarımda-yol-gösterici-ilkeler-nelerdir)
    - [Kullanım](#kullanım)
      - [Google, Go'yu dahili olarak kullanıyor mu?](#google-goyu-dahili-olarak-kullanıyor-mu)
      - [Başka hangi şirketler Go kullanıyor?](#başka-hangi-şirketler-go-kullanıyor)
      - [Go programları C / C ++ programlarıyla bağlantılı mı?](#go-programları-c--c--programlarıyla-bağlantılı-mı)
      - [Hangi IDE'lerin Go desteği var?](#hangi-idelerin-go-desteği-var)
      - [Go, Google'ın protokol tamponlarını \(buffers\) destekliyor mu?](#go-googleın-protokol-tamponlarını-buffers-destekliyor-mu)
      - [Go ana sayfasını başka bir dile çevirebilir miyim?](#go-ana-sayfasını-başka-bir-dile-çevirebilir-miyim)
    - [Tasarım](#tasarım)
      - [Go çalışma zamanına \(runtime\) sahip mi?](#go-çalışma-zamanına-runtime-sahip-mi)
      - [Unicode tanımlayıcılarından n'aber?](#unicode-tanımlayıcılarından-naber)
      - [Go neden X özelliğine sahip değil?](#go-neden-x-özelliğine-sahip-değil)
      - [Go'da neden jenerik tipler yok?](#goda-neden-jenerik-tipler-yok)
      - [Neden Go'da Exceptions yok?](#neden-goda-exceptions-yok)
      - [Go'da neden assertions \(iddialar\) yok?](#goda-neden-assertions-iddialar-yok)
      - [Neden CSP fikirleri üzerine eşzamanlılık inşa etmelisiniz?](#neden-csp-fikirleri-üzerine-eşzamanlılık-inşa-etmelisiniz)
      - [Neden Thread yerine Goroutine?](#neden-thread-yerine-goroutine)
      - [Map işlemleri neden atomik olarak tanımlanmıyor?](#map-işlemleri-neden-atomik-olarak-tanımlanmıyor)
      - [Dil değişikliğimi kabul edecek misiniz?](#dil-değişikliğimi-kabul-edecek-misiniz)
    - [Türler](#türler)
      - [Go nesne yönelimli bir dil midir?](#go-nesne-yönelimli-bir-dil-midir)
      - [Yöntemlerin dinamik gönderimini nasıl edinebilirim?](#yöntemlerin-dinamik-gönderimini-nasıl-edinebilirim)
      - [Neden tür mirası \(kalıtım\) yok?](#neden-tür-mirası-kalıtım-yok)
      - [len\(\) neden bir metod değilde fonksiyondur?](#len-neden-bir-metod-değilde-fonksiyondur)
      - [Go neden method ve operatör overloading desteklemiyor?](#go-neden-method-ve-operatör-overloading-desteklemiyor)
      - [Go neden "uygular" \(implements\) tanımlamalarına sahip değil?](#go-neden-uygular-implements-tanımlamalarına-sahip-değil)
      - [Türümün bir arayüzü karşıladığını nasıl garanti edebilirim?](#türümün-bir-arayüzü-karşıladığını-nasıl-garanti-edebilirim)
      - [T türü neden Equal arayüzünü karşılamıyor?](#t-türü-neden-equal-arayüzünü-karşılamıyor)
      - [Bir \[\]T'yi bir \[\]interface'e dönüştürebilir miyim?](#bir-tyi-bir-interfacee-dönüştürebilir-miyim)
      - [T1 ve T2 aynı temel türe sahipse \[\] T1'i \[\] T2'ye dönüştürebilir miyim?](#t1-ve-t2-aynı-temel-türe-sahipse--t1i--t2ye-dönüştürebilir-miyim)
      - [Nil hata değerim neden nil'e eşit değil?](#nil-hata-değerim-neden-nile-eşit-değil)
      - [C'de olduğu gibi neden untagged unions yok?](#cde-olduğu-gibi-neden-untagged-unions-yok)
      - [Go'da neden varyant türleri yok?](#goda-neden-varyant-türleri-yok)
      - [Go neden ortak değişken sonuç türlerine sahip değil?](#go-neden-ortak-değişken-sonuç-türlerine-sahip-değil)
    - [Değerler](#değerler)
      - [Go neden örtük sayısal dönüşümler sağlamaz?](#go-neden-örtük-sayısal-dönüşümler-sağlamaz)
      - [Sabitler Go'da nasıl çalışır?](#sabitler-goda-nasıl-çalışır)
      - [Map neden built-in \(yerleşik\)?](#map-neden-built-in-yerleşik)
      - [Map neden dilimlere \(slices\) anahtar \(key\) olarak izin vermiyor?](#map-neden-dilimlere-slices-anahtar-key-olarak-izin-vermiyor)
      - [Diziler değer iken neden map'ler, dilimler ve kanallar referanslarıdır?](#diziler-değer-iken-neden-mapler-dilimler-ve-kanallar-referanslarıdır)
    - [Kod Yazma](#kod-yazma)
      - [Kütüphaneler nasıl belgelenir?](#kütüphaneler-nasıl-belgelenir)
      - [Go programlama stili kılavuzu var mı?](#go-programlama-stili-kılavuzu-var-mı)
      - [Yamaları Go kütüphanelerine nasıl gönderirim?](#yamaları-go-kütüphanelerine-nasıl-gönderirim)
      - ["go get" bir depoyu klonlarken neden HTTPS kullanıyor?](#go-get-bir-depoyu-klonlarken-neden-https-kullanıyor)
      - ["Go get" kullanarak paket sürümlerini nasıl yönetmeliyim?](#go-get-kullanarak-paket-sürümlerini-nasıl-yönetmeliyim)
    - [İşaretçiler ve Tahsis](#i̇şaretçiler-ve-tahsis)
      - [Fonksiyon parametreleri ne zaman değere göre aktarılır?](#fonksiyon-parametreleri-ne-zaman-değere-göre-aktarılır)
      - [Bir arayüze işaretçi ne zaman kullanmalıyım?](#bir-arayüze-işaretçi-ne-zaman-kullanmalıyım)
      - [Metodları \(struct fonksiyonlar\) değerler veya işaretçiler üzerinde tanımlamalı mıyım?](#metodları-struct-fonksiyonlar-değerler-veya-işaretçiler-üzerinde-tanımlamalı-mıyım)
      - [Make ve New arasındaki fark nedir?](#make-ve-new-arasındaki-fark-nedir)
      - [Bir int türün 64bitlik makinedeki büyüklüğü nedir?](#bir-int-türün-64bitlik-makinedeki-büyüklüğü-nedir)
      - [Bir değişkenin yığın üzerinde mi yoksa yığın üzerinde tahsis edildiğini nasıl anlarım?](#bir-değişkenin-yığın-üzerinde-mi-yoksa-yığın-üzerinde-tahsis-edildiğini-nasıl-anlarım)
      - [Go işlemim neden bu kadar çok sanal bellek kullanıyor?](#go-işlemim-neden-bu-kadar-çok-sanal-bellek-kullanıyor)
    - [Eşzamanlılık](#eşzamanlılık)
      - [Hangi işlemler atomiktir? Mutexler ne olacak?](#hangi-işlemler-atomiktir-mutexler-ne-olacak)
      - [Programım neden daha fazla CPU ile daha hızlı çalışmıyor?](#programım-neden-daha-fazla-cpu-ile-daha-hızlı-çalışmıyor)
      - [CPU sayısını nasıl kontrol edebilirim?](#cpu-sayısını-nasıl-kontrol-edebilirim)
      - [Gorutin olarak çalışan kapanışlarda ne olur?](#gorutin-olarak-çalışan-kapanışlarda-ne-olur)
    - [Fonksiyonlar ve Metodlar \(Structlar için fonksiyonlar\)](#fonksiyonlar-ve-metodlar-structlar-için-fonksiyonlar)
      - [T ve \* T'nin neden farklı metod atamaları var?](#t-ve--tnin-neden-farklı-metod-atamaları-var)
    - [Kontrol Akışı](#kontrol-akışı)
      - [Go neden ?: operatörüne sahip değil?](#go-neden--operatörüne-sahip-değil)
    - [Paketler ve Testing](#paketler-ve-testing)
      - [Nasıl çok dosyalı paket oluştururum?](#nasıl-çok-dosyalı-paket-oluştururum)
      - [Nasıl birim testi yazarım?](#nasıl-birim-testi-yazarım)
      - [Test için en sevdiğim yardımcı fonksiyon nerede?](#test-için-en-sevdiğim-yardımcı-fonksiyon-nerede)
      - [X neden standart kütüphanede yok?](#x-neden-standart-kütüphanede-yok)
    - [İmplemantasyonlar](#i̇mplemantasyonlar)
      - [Derleyicileri oluşturmak için hangi derleyici teknolojisi kullanılır?](#derleyicileri-oluşturmak-için-hangi-derleyici-teknolojisi-kullanılır)
      - [Çalışma zamanı desteği nasıl uygulanır?](#çalışma-zamanı-desteği-nasıl-uygulanır)
      - [Basit programım neden büyük dosya boyutlu?](#basit-programım-neden-büyük-dosya-boyutlu)
      - [Kullanılmayan değişken/içe aktarım ile ilgili uyarıları durdurabilir miyim?](#kullanılmayan-değişkeniçe-aktarım-ile-ilgili-uyarıları-durdurabilir-miyim)
      - [Virüs tarama yazılımım neden Go dağıtımıma veya derlenmiş ikili dosyama virüs bulaştığını düşünüyor?](#virüs-tarama-yazılımım-neden-go-dağıtımıma-veya-derlenmiş-ikili-dosyama-virüs-bulaştığını-düşünüyor)
    - [Performans](#performans)
      - [Go, X karşılaştırmasında neden kötü performans gösteriyor?](#go-x-karşılaştırmasında-neden-kötü-performans-gösteriyor)
    - [C'den Değişiklikler](#cden-değişiklikler)
      - [Sözdizimi C'den neden bu kadar farklı?](#sözdizimi-cden-neden-bu-kadar-farklı)
      - [Tanımlamalar neden ters yönlü yapılır?](#tanımlamalar-neden-ters-yönlü-yapılır)
      - [Neden işaretçi aritmetiği yok?](#neden-işaretçi-aritmetiği-yok)
      - [Neden ++ ve -- tanımlama değilde ifadedir? Ve neden ön ek değilde son ektir?](#neden--ve----tanımlama-değilde-ifadedir-ve-neden-ön-ek-değilde-son-ektir)
      - [Neden parantez var ama noktalı virgül yok? Ve neden sonraki satırda açılış ayracı koyamazsınız?](#neden-parantez-var-ama-noktalı-virgül-yok-ve-neden-sonraki-satırda-açılış-ayracı-koyamazsınız)
      - [Neden çöp toplama yapılıyor? Çok masraflı olmayacak mı?](#neden-çöp-toplama-yapılıyor-çok-masraflı-olmayacak-mı)
  - [Go Derleyicisi Kurulumu](#go-derleyicisi-kurulumu)
  - [VSCode Go Eklentisi Yükleme](#vscode-go-eklentisi-yükleme)
  - [Merhaba Dünya](#merhaba-dünya)
  - [VSCode Varsayılan Hata Ayıklayıcıyı Seçme](#vscode-varsayılan-hata-ayıklayıcıyı-seçme)
  - [Farklı Platformlara Build \(İnşa\) Etme](#farklı-platformlara-build-i̇nşa-etme)
  - [Klasör Build Etme](#klasör-build-etme)
  - [Paketler](#paketler)
  - [Yorum Satırı](#yorum-satırı)
  - [Veri Tipleri](#veri-tipleri)
  - [Aritmetik Operatörler](#aritmetik-operatörler)
  - [İlişkisel Operatörler](#i̇lişkisel-operatörler)
  - [Mantıksal Operatörler](#mantıksal-operatörler)
  - [Atama Operatörleri](#atama-operatörleri)
  - [Değişkenler ve Atanması](#değişkenler-ve-atanması)
  - [Sabitler](#sabitler)
  - [Kod Gruplama İşlemi](#kod-gruplama-i̇şlemi)
  - [Tür Dönüşümü](#tür-dönüşümü)

# Bölüm 1 - Ön Eğitim

## Golang Hakkında

Golang \(diğer adıyla Go\), **Google**’ın **2007** yılından beri geliştirdiği açık kaynaklı programlama dilidir. Daha çok alt-sistem programlama için tasarlanmış olup, derlenebilir ve **statik** tipli bir dildir. İlk versiyonu **Kasım 2009**‘da çıkmıştır. Derleyicisi olan **“gc” \(Go Compiler\)** açık kaynak olarak birçok işletim sistemi için geliştirilmiştir.

Golang, Google mühendislerinden olan **Robert Griesemer, Rob Pike ve Ken Thompson** tarafından ilk olarak deney amaçlı ortaya çıkmıştır. Diğer dillerdeki eleştirilen sorunlara çözüm getirmek ve iyi yönlerini korumak için çıkarılmıştır.

Bu adamların daha önceden bulunmuş olduğu projelere bakacak olursak Google’ın gerçekten bu kişileri cımbızla seçtiğini anlayabiliriz. İşte Golang’ın yaratıcılarının bulunduğu projeler:

| ![Ken Thompson](./goruntuler/ken.png) | ![Rob Pike](./goruntuler/rob.jpg) | ![Robert Griesemer](./goruntuler/robert.png) |
| :---: | :---: | :---: |
| **Ken Thompson:** *B, C, UNIX ve UTF-8* | **Rob Pike:** *UNIX ve UTF-8* | **Robert Griesemer:** *Hotspot ve JVM (Java Sanal Makinesi)* |

**Dilin Özellikleri**

* Statik yazılmıştır, fakat dinamik rahatlığındadır.
* Büyük sistemlere ölçeklenebilir.
* Üretken ve okunabilir olması, çok fazla zorunlu anahtar kelime ve tekrarlamaların kullanılmaması
* Tümleşik Geliştirme Ortamına \(IDE\) ihtiyaç duyulmaması fakat desteklemesi
* Ağ ve çoklu işlemleri desteklemesi
* Değişken tanımında tür belirtimi isteğe bağlıdır.
* Hızlı derlenme süresi
* Uzak paket yöneticisi \(go get\)

**Örnek Merhaba Dünya Uygulaması**

```go
   package main

   import “fmt”

   func main(){
       fmt.Println(“Merhaba Dünya!”)
   }
```

## Sıkça Sorulan Sorular

> [https://golang.org/doc/faq](https://golang.org/doc/faq) adresindeki içerikten tercüme edilmiştir.

### Kökenler

#### Projenin amacı nedir?

Go'nun başlangıcında, sadece on yıl önce, programlama dünyası bugünden farklıydı. Yazılım üretimi genellikle C ++ veya Java ile yapılıyordu.GitHub mevcut değildi. Çoğu bilgisayar henüz çok işlemcili değildi. Visual Studio ve Eclipse dışında, internette ücretsiz olarak birkaç IDE veya diğer üst düzey araçlar mevcuttu.

Bu arada, sunucu yazılımı geliştirmede birlikte çalıştığımız dilleri kullanmak için karşılaştığımız gereksiz karmaşıklıktan dolayı hayal kırıklığına uğradık. Bilgisayarlar, C, C ++ ve Java gibi diller ilk geliştirildiğinden beri çok daha hızlı hale gelmişti, ancak programlama eylemi tam olarak o kadar ilerlememişti. Ayrıca, çok işlemcinin evrensel hale geldiği açıktı, ancak çoğu dil onları verimli ve güvenli bir şekilde programlamak için çok az yardım sundu.

Bir adım geri gitmeye ve teknoloji geliştikçe önümüzdeki yıllarda yazılım mühendisliğine hangi önemli sorunların hâkim olacağını ve yeni bir dilin bunları nasıl çözebileceğini düşünmeye karar verdik. Örneğin, çok çekirdekli CPU'ların yükselişi, bir dilin bir tür eşzamanlılık veya paralellik için birinci sınıf destek sağlaması gerektiğini savundu. Buna mütakiben kaynak yönetimini büyük bir eşzamanlı programda izlenebilir hale getirmek için, çöp toplama veya en azından bir tür güvenli otomatik bellek yönetimi gerekliydi.

Bu düşünceler, Go'nun önce bir dizi fikir ve aranan veri, sonra da bir dil olarak ortaya çıktığı bir dizi tartışmaya yol açtı. Kapsamlı bir hedef, Go'nun çalışan programcıya takım oluşturmayı etkinleştirerek, kod biçimlendirme gibi sıradan görevleri otomatikleştirerek ve büyük kod tabanlarında çalışmanın önündeki engelleri kaldırarak daha fazla yardım etmesiydi.

Go'nun hedeflerinin ve bunlara nasıl ulaşıldığının veya en azından bunlara nasıl yaklaşıldığının çok daha kapsamlı bir açıklaması [Go at Google: Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article) adlı makalede mevcuttur.

#### Projenin tarihçesi nedir?

Robert Griesemer, Rob Pike ve Ken Thompson 21 Eylül 2007'de beyaz tahtada yeni bir dilin hedeflerini çizmeye başladılar. Birkaç gün içinde hedefler bir şeyler yapma planına ve bunun ne olacağına dair adil bir fikre yerleşti. Tasarım, ilgisiz çalışmaya paralel olarak yarı zamanlı devam etti. Ocak 2008'de Ken, fikirleri araştırmak için bir derleyici üzerinde çalışmaya başladı; çıktı olarak C kodunu üretti. Yıl ortasına gelindiğinde, dil tam zamanlı bir proje haline geldi ve bir üretim derleyicisini deneyecek kadar yerleşti. Mayıs 2008'de Ian Taylor, taslak şartnameyi kullanarak Go için bir GCC ön ucunu bağımsız olarak kullanmaya başladı. Russ Cox 2008'in sonlarında katıldı ve dilin ve kitaplıkların prototipten gerçeğe taşınmasına yardımcı oldu.

Go, 10 Kasım 2009'da halka açık bir açık kaynak projesi oldu. Topluluktan sayısız kişi fikirlere, tartışmalara ve kodlara katkıda bulundu.

Şu anda dünya çapında milyonlarca Go programcısı \(gopher\) var ve her gün daha fazlası var olacak. Go'nun başarısı beklentilerimizi çok aştı.

#### Gopher maskotu nereden geliyor?

Maskot ve logo, Plan 9 tavşanı [Glenda](https://9p.io/plan9/glenda.html)'yı da tasarlayan [Renée French](https://reneefrench.blogspot.com/) tarafından tasarlandı. Gopher hakkında bir [blog yazısı](https://blog.golang.org/gopher), birkaç yıl önce bir [WFMU](https://wfmu.org/) tişört tasarımı için kullandığı birinden nasıl türetildiğini açıklıyor. Logo ve maskot, [Creative Commons Attribution 3.0 lisansı](https://creativecommons.org/licenses/by/3.0/) kapsamındadır. Gopher, özelliklerini ve bunların doğru şekilde nasıl temsil edileceğini gösteren bir [model sayfası](https://golang.org/doc/gopher/modelsheet.jpg)na sahiptir. Model sayfası ilk olarak 2016 yılında Renée'nin Gophercon'da yaptığı bir [konuşma](https://www.youtube.com/watch?v=4rw_B4yY69k)da gösterildi. Kendine özgü özellikleri var; o Go gopher, herhangi bir yaşlı sincap değil.

#### Dilin adı Go mu Golang mi?

Dilin adı Go. "Golang" takma adının ortaya çıkmasının nedeni, web sitesinin bizim için mevcut olmayan go.org değil, [golang.org](https://golang.org/) olmasıdır. Yine de çoğu golang adını kullanır ve etiket olarak kullanışlıdır. Örneğin, dil için Twitter etiketi "\#golang" dir. Dilin adı ne olursa olsun sadece Go'dur.

**Bir yan not:** [Resmi logo](https://blog.golang.org/go-brand)nun iki büyük harfi olmasına rağmen, dil adı GO değil, Go yazılmıştır.

#### Neden yeni bir dil yarattınız?

Go, Google'da yaptığımız iş için mevcut diller ve ortamlarla ilgili hayal kırıklığından doğdu. Programlama çok zor hale geldi ve dil seçimi kısmen suçluydu. Etkili derlemeyi, verimli yürütmeyi veya programlama kolaylığını seçmek gerekiyordu; üçü de aynı ana dilde mevcut değildi. C ++ veya daha az ölçüde Java yerine Python ve JavaScript gibi dinamik olarak yazılmış dillere geçerek güvenlik ve verimlilik üzerinde kolaylık seçebilmeliydi programcılar.

Endişelerimizde yalnız değildik. Yıllar sonra, programlama dilleri için oldukça sessiz bir manzaraya sahip olan Go, programlama dili geliştirmeyi tekrar aktif, neredeyse yaygın bir alan haline getiren birkaç yeni dilin \(Rust, Elixir, Swift ve daha fazlası\) ilklerinden biriydi.

Go, yorumlanmış, dinamik olarak yazılmış bir dilin programlama kolaylığını statik olarak yazılmış, derlenmiş bir dilin verimliliği ve güvenliğiyle birleştirmeye çalışarak bu sorunları ele aldı. Bu dil aynı zamanda ağa bağlı ve çok çekirdekli bilgi işlem desteğiyle modern olmayı hedefliyordu. Son olarak, Go ile çalışmanın hızlı olması amaçlanmıştır: Tek bir bilgisayarda büyük bir yürütülebilir dosya oluşturmak en fazla birkaç saniye sürecektir. Bu hedeflere ulaşmak için bir dizi dilbilimsel konuyu ele almak gerekiyor: ifade edici ama hafif bir yazım sistemi; eşzamanlılık ve çöp toplama; katı bağımlılık belirtimi; ve bunun gibi. Bunlar kütüphaneler veya araçlar tarafından iyi bir şekilde ele alınamaz; yeni bir dil aranıyordu.

[Go at Google](https://talks.golang.org/2012/splash.article) makalesi, Go dilinin tasarımının arkasındaki arka planı ve motivasyonu tartışmanın yanı sıra bu SSS'de sunulan yanıtların çoğu hakkında daha fazla ayrıntı sağlar.

#### Go'nun ataları nelerdir? \(veya esinlendikleri\)

Go, çoğunlukla Pascal / Modula / Oberon ailesinden \(bildirimler, paketler\) önemli girdiler ve ayrıca Newsqueak ve Limbo \(eşzamanlılık\) gibi Tony Hoare'nin CSP'sinden esinlenen dillerden bazı fikirlerle, C ailesinde \(temel sözdizimi\) bulunur. Ancak, her yerde yeni bir dildir. Her açıdan dil, programcıların ne yaptığını ve programlamayı nasıl yapacağını, en azından yaptığımız programlama türünü, daha etkili, yani daha eğlenceli olması düşünülerek tasarlandı.

#### Tasarımda yol gösterici ilkeler nelerdir?

Go tasarlandığında, Java ve C ++, en azından Google'da, sunucuları yazmak için en yaygın kullanılan dillerdi. Bu dillerin çok fazla defter tutma \(muhasebecilik\) ve tekrarlama gerektirdiğini hissettik. Bazı programcılar, verimlilik ve yazım güvenliği pahasına Python gibi daha dinamik, akıcı dillere geçerek tepki gösterdi. Verimliliğe, güvenliğe ve akıcılığa tek bir dilde sahip olmanın mümkün olması gerektiğini hissettik.

Go, kelimenin her iki anlamında da yazma miktarını azaltmaya çalışır. Tasarımı boyunca dağınıklığı ve karmaşıklığı azaltmaya çalıştık. İleri tanımlamalar ve başlık \(header\) dosyaları yoktur; her şey tam olarak bir kez ilan edilir. Başlatma anlamlı, otomatik ve kullanımı kolaydır. Sözdizimi temiz ve anahtar kelimelerde hafiftir. Yazımda gereksizlik \(foo.Foo \* myFoo = new \(foo.Foo\)\): = tanımla-ve-türet yapısı kullanılarak basit tür türetmesi ile azaltılır. Ve belki de en radikal olanı, tür hiyerarşisi yoktur: türler sadece ilişkilerdir. Bu basitleştirmeler, Go'nun etkileyici olmasına rağmen karmaşıklıktan ödün vermeden anlaşılabilir olmasını sağlar.

Bir diğer önemli ilke de kavramları ortogonal tutmaktır. Yöntemler her tür için uygulanabilir; yapılar \(struct\) verileri temsil ederken arayüzler \(interface\) soyutlamayı temsil eder. Ortogonallik, nesneler birleştiğinde ne olduğunu anlamayı kolaylaştırır.

###  Kullanım

#### Google, Go'yu dahili olarak kullanıyor mu?

Evet. Go, Google içindeki üretimde yaygın olarak kullanılmaktadır. Bunun kolay bir örneği [golang.org](https://golang.org/)'un arkasındaki sunucudur. [Google App Engine](https://developers.google.com/appengine/)'de bir üretim yapılandırmasında çalışan godoc belge sunucusudur.

Daha önemli bir örnek, Google'ın Chrome ikili \(binary\) dosyalarını ve apt-get \(Debian paket yöneticisi oluyor kendisi\) paketleri gibi diğer büyük yüklenebilir dosyaları sunan indirme sunucusu `dl.google.com`'dur.

Go, Google'da kullanılan tek dil değildir, ancak site güvenilirliği mühendisliği \(SRE\) ve büyük ölçekli veri işleme dahil olmak üzere birçok alan için anahtar dildir.

#### Başka hangi şirketler Go kullanıyor?

Go kullanımı dünya çapında artıyor, ancak hiçbir şekilde yalnızca bulut bilişim alanında değil. Go'da yazılan birkaç büyük bulut altyapı projesi Docker ve Kubernetes'tir, ancak çok daha fazlası vardır.

Yine de sadece bulut değil. Go Wiki, Go kullanan birçok şirketin bazılarını listeleyen ve düzenli olarak güncellenen bir [sayfa](https://github.com/golang/go/wiki/GoUsers) içerir. Wiki'de ayrıca, dili kullanan şirketler ve projeler hakkındaki [başarı hikayeleri](https://github.com/golang/go/wiki/SuccessStories)ne bağlantılar içeren bir sayfa vardır.

#### Go programları C / C ++ programlarıyla bağlantılı mı?

C ve Go'yu aynı adres alanında birlikte kullanmak mümkündür, ancak bu doğal bir uyum değildir ve özel arayüz yazılımı gerektirebilir. Ayrıca, C'yi Go koduyla ilişkilendirmek, Go'nun sağladığı bellek güvenliği ve yığın yönetimi özelliklerinden mahrum kalmaktır. Bazen bir sorunu çözmek için C kitaplıklarını kullanmak kesinlikle gereklidir, ancak bunu yapmak her zaman saf Go kodunda bulunmayan bir risk unsuru getirir, bu yüzden dikkatli olun.

C'yi Go ile kullanmanız gerekiyorsa, nasıl devam edeceğiniz Go derleyici uygulamasına bağlıdır. Go ekibi tarafından desteklenen üç Go derleyici uygulaması vardır. Bunlar, GCC arka ucunu kullanan varsayılan derleyici gccgo ve LLVM altyapısını kullanan biraz daha az olgun bir gollvm'dir.

Gc, C'den farklı bir arama kuralı ve bağlayıcı kullanır ve bu nedenle doğrudan C programlarından veya tam tersi şekilde çağrılamaz. [Cgo](https://golang.org/cmd/cgo/) programı, C kitaplıklarının Go kodundan güvenli bir şekilde çağrılmasına izin veren bir "yabancı işlev arabirimi" mekanizması sağlar. SWIG  C ++ kitaplıklarına erişme yeteneğini genişletir.

Gccgo ve gollvm ile cgo ve SWIG'i de kullanabilirsiniz. Geleneksel bir API kullandıklarından, büyük bir dikkatle, bu derleyicilerden gelen kodu doğrudan GCC / LLVM-derlenmiş C veya C ++ programlarına bağlamak da mümkündür. Ancak, bunu güvenli bir şekilde yapmak, ilgili tüm diller için çağrı kurallarının anlaşılmasını ve ayrıca Go'dan C veya C ++ çağrılırken yığın sınırlarının dikkate alınmasını gerektirir.

#### Hangi IDE'lerin Go desteği var?

Go projesi özel bir IDE içermez, ancak dil ve kitaplıklar, kaynak kodunu analiz etmeyi kolaylaştıracak şekilde tasarlanmıştır. Sonuç olarak, en iyi bilinen editörler ve IDE'ler, doğrudan veya bir eklenti aracılığıyla Go destekler. İyi Go desteğine sahip tanınmış IDE'lerin ve editörlerin listesi Emacs, Vim, VSCode, Atom, Eclipse, Sublime, IntelliJ \(Goland adlı özel bir varyant aracılığıyla\) ve daha fazlasını içerir. En sevdiğiniz ortam, Go'da programlama için üretken bir ortam olabilir.

#### Go, Google'ın protokol tamponlarını \(buffers\) destekliyor mu?

Ayrı bir açık kaynak projesi, gerekli derleyici eklentisini ve kitaplığı sağlar. [github.com/golang/protobuf/](https://github.com/golang/protobuf) adresinde mevcuttur.

#### Go ana sayfasını başka bir dile çevirebilir miyim?

Kesinlikle. Geliştiricileri kendi dillerinde Go Language siteleri oluşturmaya teşvik ediyoruz. Ancak, sitenize Google logosunu veya markasını eklemeyi seçerseniz \([golang.org](https://golang.org/)'da görünmez\), [www.google.com/permissions/guidelines.html](https://www.google.com/permissions/guidelines.html) adresindeki yönergelere uymanız gerekecektir.

###  Tasarım

#### Go çalışma zamanına \(runtime\) sahip mi?

Go, her Go programının bir parçası olan, çalışma zamanı adı verilen kapsamlı bir kitaplığa sahiptir. Çalışma zamanı kitaplığı çöp toplama, eşzamanlılık, yığın yönetimi ve Go dilinin diğer kritik özelliklerini uygular. Dil için daha merkezi olmasına rağmen, Go'nun çalışma zamanı C kitaplığı olan `libc`'ye benzer.

Bununla birlikte, Go'nun çalışma zamanının Java çalışma zamanı tarafından sağlananlar gibi bir sanal makine içermediğini anlamak önemlidir. Go programları, yerel makine koduna \(veya bazı varyant uygulamaları için JavaScript veya WebAssembly\) önceden derlenir. Bu nedenle, terim genellikle bir programın çalıştığı sanal ortamı tanımlamak için kullanılsa da, Go'da "çalışma zamanı" kelimesi yalnızca kritik dil hizmetleri sağlayan kitaplığa verilen addır.

#### Unicode tanımlayıcılarından n'aber?

Go'yu tasarlarken aşırı ASCII merkezli olmadığından emin olmak istedik. Bu da tanımlayıcıların alanını 7 bitlik ASCII'nin sınırlarından genişletmek anlamına geliyordu. Go kuralı — tanımlayıcı karakterler Unicode tarafından tanımlandığı gibi harf veya rakam olmalıdır — anlaşılması ve uygulanması kolaydır, ancak kısıtlamaları vardır. Örneğin, karakterleri birleştirmek tasarım tarafından hariç tutulur ve bu, Devanagari gibi bazı dilleri hariç tutar.

Bu kuralın talihsiz bir sonucu daha var. Dışa aktarılan bir tanımlayıcının bir büyük harfle başlaması gerektiğinden, bazı dillerde karakterlerden oluşturulan tanımlayıcılar tanım gereği dışa aktarılamaz. Şimdilik tek çözüm, açıkça tatmin edici olmayan X 日本語 gibi bir şey kullanmaktır.

Dilin en eski sürümünden bu yana, diğer yerel dilleri kullanan programcıları barındırmak için tanımlayıcı alanının en iyi şekilde nasıl genişletilebileceği konusunda önemli ölçüde düşünülmüştür. Tam olarak ne yapılacağı aktif bir tartışma konusu olmaya devam ediyor ve dilin gelecekteki bir versiyonu tanımlayıcı tanımında daha liberal olabilir. Örneğin, Unicode organizasyonunun tanımlayıcılar için önerilerinden bazı fikirleri benimseyebilir. Ne olursa olsun, Go'nun en sevdiğimiz özelliklerinden biri olan harf durumunun tanımlayıcıların görünürlüğünü belirleme şeklini korurken \(veya belki genişletirken\) uyumlu bir şekilde yapılmalıdır.

 Şimdilik, daha sonra programları bozmadan genişletilebilecek basit bir kuralımız var. Belirsiz tanımlayıcıları kabul eden bir kuraldan kesinlikle kaynaklanabilecek hataları önleyen bir kural.

#### Go neden X özelliğine sahip değil?

Her dil yeni özellikler içerir ve birinin en sevdiği özelliği atlar. Go, programlamanın mutluluğu, derleme hızı, kavramların ortogonalitesi ve eşzamanlılık ve çöp toplama gibi özellikleri destekleme ihtiyacı göz önünde bulundurularak tasarlandı. En sevdiğiniz özellik, uymadığından, derleme hızını veya tasarımın netliğini etkilediği veya temel sistem modelini çok zorlaştıracağı için eksik olabilir.

Go'nun X özelliğinin eksik olması sizi rahatsız ediyorsa, lütfen bizi affedin ve Go'nun sahip olduğu özellikleri araştırın. X'in eksikliğini ilginç yollarla telafi ettiklerini görebilirsiniz.

#### Go'da neden jenerik tipler yok?

Jenerikler bir noktada eklenebilir. Bazı programcıların yaptığını bilsek de, onlar için bir aciliyet hissetmiyoruz.

Go, zaman içinde bakımı kolay olacak sunucu programları yazmak için bir dil olarak tasarlandı. \(Daha fazla arka plan için [bu makale](https://talks.golang.org/2012/splash.article)ye bakın.\) Tasarım, ölçeklenebilirlik, okunabilirlik ve eşzamanlılık gibi şeylere odaklandı. Polimorfik programlama o zamanlar dilin hedefleri için gerekli görünmüyordu ve bu yüzden basitlik için dışarıda bırakıldı.

Dil artık daha olgundur ve bir tür genel programlamayı düşünmek için alan vardır. Ancak, bazı uyarılar var.

Jenerikler kullanışlıdır, ancak tip sistem ve çalışma süresinde karmaşıklık açısından bir maliyete sahiptirler. Karmaşıklığa orantılı değer veren bir tasarım henüz bulamadık, ancak üzerinde düşünmeye devam ediyoruz. Bu arada, Go'nun yerleşik haritaları \(map\) ve dilimlerinin \(slice\) yanı sıra konteynerler oluşturmak için boş arayüzü kullanma yeteneği \(açık kutudan çıkarma ile\), çoğu durumda kod yazmanın mümkün olduğu anlamına gelir. Bu, daha az sorunsuz olsa da jeneriklerin sağlayacağı şeyi yapar.

Konu açık kalır. Go için iyi bir jenerik çözüm tasarlamaya yönelik önceki birkaç başarısız girişime bir göz atmak için [bu öneriye](https://golang.org/issue/15292) bakın.

#### Neden Go'da Exceptions yok?

Try-catch-final deyiminde olduğu gibi bir kontrol yapısına istisnaların birleştirilmesinin kıvrımlı kodla sonuçlandığına inanıyoruz. Ayrıca, programcıları, bir dosyayı açamama gibi çok fazla sıradan hatayı istisnai olarak etiketlemeye teşvik etme eğilimindedir.

Go farklı bir yaklaşım sergiliyor. Düz hata işleme için, Go'nun çoklu değer dönüşleri, dönüş değerini aşırı yüklemeden bir hatayı bildirmeyi kolaylaştırır. [Go'nun diğer özellikleriyle birlikte kanonik bir hata türü](https://golang.org/doc/articles/error_handling.html), hata işlemeyi keyifli hale getirir ancak diğer dillerdekinden oldukça farklıdır.

Go ayrıca, gerçekten istisnai koşullardan sinyal almak ve kurtarmak için birkaç yerleşik işleve sahiptir. Kurtarma mekanizması, yalnızca bir hatadan sonra yıkılan bir işlevin durumunun bir parçası olarak yürütülür; bu, felaketi ele almak için yeterlidir, ancak ekstra kontrol yapıları gerektirmez ve iyi kullanıldığında temiz hata işleme koduyla sonuçlanabilir. Ayrıntılar için [Defer, Panic ve Recover makalesi](https://golang.org/doc/articles/defer_panic_recover.html)ne bakın. Bu dökümanda ise şu konulara bakabilirsiniz.

 page-ref page="../boeluem-2/defer.md

../boeluem-5/panic-and-recover.md

Ayrıca, [Hatalar blog gönderisi](https://blog.golang.org/errors-are-values)nin bir Go'da hataları temiz bir şekilde işleme yaklaşımı, hataların sadece değerler olduğundan, Go'nun tam gücünün hata işlemede kullanılabileceğini göstererir.

#### Go'da neden assertions \(iddialar\) yok?

Go, assertions sağlamaz. İnkar edilemez derecede kullanışlıdırlar, ancak bizim deneyimlerimiz, programcıların bunları doğru hata işleme ve raporlama hakkında düşünmekten kaçınmak için koltuk değneği olarak kullanmasıdır. Doğru hata işleme, sunucuların önemli olmayan bir hatadan sonra çökmek yerine çalışmaya devam etmesi anlamına gelir. Doğru hata raporlama, hataların doğrudan ve yerinde olduğu anlamına gelir ve programcıyı büyük bir çökme izini yorumlamaktan kurtarır. Hataları gören programcı koda aşina olmadığında kesin hatalar özellikle önemlidir.

Bunun bir çekişme noktası olduğunu anlıyoruz. Go dilinde ve kütüphanelerde modern uygulamalardan farklı birçok şey vardır, çünkü bazen farklı bir yaklaşımın denemeye değer olduğunu düşünüyoruz.

#### Neden CSP fikirleri üzerine eşzamanlılık inşa etmelisiniz?

Eşzamanlılık ve çok iş parçacıklı programlama, zaman içinde zorluklarla ilgili bir ün geliştirmiştir. Bunun kısmen [pthreads](https://en.wikipedia.org/wiki/POSIX_Threads) gibi karmaşık tasarımlardan ve kısmen de mutexler, koşul değişkenleri ve bellek engelleri gibi düşük seviyeli ayrıntılara aşırı vurgu yapılmasından kaynaklandığına inanıyoruz. Daha yüksek seviyeli arayüzler, kapakların altında hala mutexler ve benzeri şeyler olsa bile çok daha basit bir kod sağlar.

Eşzamanlılık için üst düzey dil desteği sağlamanın en başarılı modellerinden biri, Hoare'nin İletişim Sıralı Süreçlerinden veya CSP'den gelir. Occam ve Erlang, CSP'den kaynaklanan iki iyi bilinen dildir. Go'nun eşzamanlılık ilkelleri, ana katkısı birinci sınıf nesneler olarak güçlü kanal kavramı olan aile ağacının farklı bir bölümünden türemiştir. Daha önceki birkaç dil ile ilgili deneyimler, CSP modelinin bir prosedürel dil çerçevesine çok iyi uyduğunu göstermiştir.

#### Neden Thread yerine Goroutine?

Goroutinler, eşzamanlılığın kullanımını kolaylaştırmanın bir parçasıdır. Bir süredir ortalıkta olan fikir, işlevleri - koroutinleri - bağımsız olarak bir dizi iş parçacığına çoğaltmaktır. Bir coroutine bloke edildiğinde, örneğin bir bloke edici sistem çağrısını çağırarak bloke edildiğinde, çalışma zamanı aynı işletim sistemi iş parçacığındaki diğer koroutinleri otomatik olarak farklı, çalıştırılabilir bir iş parçacığına taşır, böylece bunlar engellenmez. Programcı bunların hiçbirini görmez, önemli olan budur.

Gorutinler dediğimiz sonuç çok ucuz olabilir: sadece birkaç kilobayt olan yığın hafızasının ötesinde çok az ek yüke sahiptirler. Yığınları küçültmek için Go'nun çalışma zamanı yeniden boyutlandırılabilir, sınırlı yığınlar kullanır. Yeni basılmış bir gorutine birkaç kilobayt verilir ve bu neredeyse her zaman yeterlidir. Değilse, çalışma zamanı yığını otomatik olarak depolamak için belleği büyütür \(ve küçültür\), böylece birçok gorutinin mütevazı bir bellek miktarında yaşamasına izin verir. CPU ek yükü, işlev çağrısı başına yaklaşık üç ucuz talimatın ortalamasını alır. Bu aynı adres alanında yüz binlerce gorutin oluşturmak için pratik. Gorutinler sadece iş parçacıkları olsaydı, sistem kaynakları çok daha az sayıda tükenirdi.

#### Map işlemleri neden atomik olarak tanımlanmıyor?

Uzun tartışmalardan sonra, haritaların tipik kullanımının birden fazla gorutinden güvenli erişim gerektirmediğine karar verildi ve bu durumlarda, map muhtemelen zaten senkronize edilmiş daha büyük bir veri yapısının veya hesaplamanın bir parçasıydı. Bu nedenle, tüm map işlemlerinin bir mutex yakalamasını gerektirmek çoğu programı yavaşlatır ve birkaçına güvenlik ekler. Kontrolsüz map erişiminin programı çökertebileceği anlamına geldiğinden, bu kolay bir karar değildi.

Dil, atomik map güncellemelerini engellemez. Güvenilmeyen bir programı barındırırken olduğu gibi gerektiğinde, uygulama map erişimini birbirine bağlayabilir.

Map erişimi, yalnızca güncellemeler yapılırken güvenli değildir. Tüm gorutinler yalnızca okuduğu - range için döngü \(for\) kullanarak yineleme dahil map'teki öğeleri aradığı ve öğelere atayarak veya silme işlemleri yaparak map'i değiştirmediği sürece, map'e eşzamanlı olarak erişmeleri güvenlidir.

Map kullanımının düzeltilmesine yardımcı olarak, bazı Dil uygulamaları, bir map eşzamanlı yürütmeyle güvenli olmayan bir şekilde değiştirildiğinde çalışma zamanında otomatik olarak rapor veren özel bir denetim içerir.

#### Dil değişikliğimi kabul edecek misiniz?

İnsanlar genellikle dilde iyileştirmeler yapılmasını önerir - [posta listesi](https://groups.google.com/group/golang-nuts) bu tür tartışmaların zengin bir geçmişini içerir - ancak bu değişikliklerin çok azı kabul edilmiştir. Go açık kaynaklı bir proje olmasına rağmen, dil ve kitaplıklar, en azından kaynak kodu düzeyinde mevcut programları bozan değişiklikleri önleyen bir uyumluluk vaadi ile korunmaktadır \(programların güncel kalması için ara sıra yeniden derlenmesi gerekebilir\). Öneriniz Go 1 spesifikasyonunu ihlal ederse, değeri ne olursa olsun fikri dikkate alamayız bile.

Go'nun gelecekteki büyük bir sürümü Go 1 ile uyumsuz olabilir, ancak bu konuyla ilgili tartışmalar daha yeni başladı ve kesin olan bir şey var: süreçte ortaya çıkan bu tür uyumsuzluklar çok az olacak. Dahası, uyumluluk vaadi, eski programların bu durum ortaya çıktığında adapte olması için ileriye dönük otomatik bir yol sağlamaya bizi teşvik ediyor. Teklifiniz Go 1 spesifikasyonuyla uyumlu olsa bile, Go'nun tasarım hedeflerinin ruhuna uygun olmayabilir. [Google'da Go: Yazılım Mühendisliği Hizmetinde Dil Tasarımı](https://talks.golang.org/2012/splash.article), Go'nun kökenlerini ve tasarımının arkasındaki motivasyonu açıklıyor.

###  Türler

#### Go nesne yönelimli bir dil midir?

Evet ve hayır. Go'nun türleri ve yöntemleri olmasına ve nesneye yönelik bir programlama stiline izin vermesine rağmen, tür hiyerarşisi yoktur. Go'daki "interface" kavramı, kullanımının kolay ve bazı açılardan daha genel olduğuna inandığımız farklı bir yaklaşım sağlar. Alt sınıflara benzer - ancak aynı olmayan - bir şey sağlamak için diğer türlere türleri gömmenin yolları da vardır. Dahası, Go'daki yöntemler C ++ veya Java'dakinden daha geneldir: düz, "unboxed" tamsayılar gibi yerleşik türler dahil her tür veri için tanımlanabilirler. Yapılar \(sınıflar\) ile sınırlı değildirler.

Ayrıca, bir tür hiyerarşisinin olmaması, Go'daki "nesnelerin" C ++ veya Java gibi dillerden çok daha hafif hissetmesini sağlar.

#### Yöntemlerin dinamik gönderimini nasıl edinebilirim?

Yöntemleri dinamik olarak göndermenin tek yolu bir interface kullanmaktır. Bir yapı veya başka herhangi bir somut türdeki yöntemler her zaman statik olarak çözümlenir.

#### Neden tür mirası \(kalıtım\) yok?

Nesne yönelimli programlama, en azından en iyi bilinen dillerde, türler arasındaki ilişkilerin, genellikle otomatik olarak türetilebilen ilişkilerin çok fazla tartışılmasını içerir. Go farklı bir yaklaşım benimser.

Go'da programcının önceden iki türün ilişkili olduğunu bildirmesini istemek yerine, Go'da bir tür, yöntemlerinin bir alt kümesini belirten herhangi bir interface'i otomatik olarak tatmin eder. Muhasebe tutmayı azaltmanın yanı sıra, bu yaklaşımın gerçek avantajları vardır. Türler, geleneksel çoklu kalıtımın karmaşıklıkları olmaksızın birçok arabirimi aynı anda karşılayabilir. Arayüzler çok hafif olabilir — bir hatta sıfır metotlu bir arayüz faydalı bir kavramı ifade edebilir. Arayüzler, yeni bir fikir ortaya çıkarsa veya test için orijinal tiplere açıklama yapmadan eklenebilir. Türler ve arabirimler arasında açık ilişkiler olmadığından, yönetilecek veya tartışılacak tür hiyerarşisi yoktur. Bu fikirleri benzer bir şey inşa etmek için kullanmak mümkündür. Örneğin, fmt.Fprintf'in sadece bir dosyaya değil herhangi bir çıktıya biçimlendirilmiş yazdırmayı nasıl sağladığını veya bufio paketinin dosya G / Ç'sinden nasıl tamamen ayrı olabileceğini veya görüntü paketlerinin sıkıştırılmış görüntü dosyalarını nasıl oluşturduğunu görün.

Tüm bu fikirler, tek bir yöntemi \(Yazma\) temsil eden tek bir arayüzden \(io.Writer\) kaynaklanır. Ve bu sadece yüzeyi çizer. Go'nun arayüzlerinin, programların nasıl yapılandırıldığı üzerinde derin bir etkisi vardır. Alışmak biraz zaman alır, ancak bu örtük tip bağımlılığı Go ile ilgili en verimli şeylerden biridir.

#### len\(\) neden bir metod değilde fonksiyondur?

Bu konuyu tartıştık, ancak pratikte işlevler iyi olduğu için len ve arkadaşlarını uygulamaya koymaya karar verdik ve temel türlerin interface'i \(Go türü anlamında\) hakkındaki soruları karmaşıklaştırmadı.

#### Go neden method ve operatör overloading desteklemiyor?

Tür eşleştirmesi de yapması gerekmiyorsa method gönderimi basitleştirilmiştir. Diğer dillerle edindiğimiz deneyimler bize, aynı isimde ancak farklı imzalara sahip çeşitli yöntemlere sahip olmanın bazen yararlı olduğunu, ancak pratikte kafa karıştırıcı ve kırılgan olabileceğini söyledi. Yalnızca isme göre eşleştirme ve türlerde tutarlılık gerektirme, Go'nun tür sisteminde büyük bir basitleştirici karardı. Operatörün aşırı yüklenmesi ile ilgili olarak, mutlak bir gereklilikten daha fazla bir kolaylık gibi görünüyor. Yine de, onsuz işler daha basit.

#### Go neden "uygular" \(implements\) tanımlamalarına sahip değil?

Bir Go türü, o arayüzün metodlarını uygulayarak bir arayüzü tatmin eder, başka bir şey değil. Bu özellik, arayüzlerin mevcut kodu değiştirmeye gerek kalmadan tanımlanmasına ve kullanılmasına izin verir. Kaygıların ayrılmasını teşvik eden ve kodun yeniden kullanımını geliştiren ve kod geliştikçe ortaya çıkan kalıplar üzerine inşa etmeyi kolaylaştıran bir tür [yapısal yazım ](https://en.wikipedia.org/wiki/Structural_type_system)sağlar. Arayüzlerin anlamsallığı, Go'nun çevik ve hafif yapısının ana nedenlerinden biridir. Daha fazla ayrıntı için [tür mirası hakkındaki soru](https://golang.org/doc/faq#inheritance)ya bakın.

#### Türümün bir arayüzü karşıladığını nasıl garanti edebilirim?

Derleyiciden, uygun şekilde, T için sıfır değerini veya T'ye işaretçi kullanarak bir atama yaparak T türünün arayüzü I uygulayıp uygulamadığını kontrol etmesini isteyebilirsiniz:

```go
type T struct{}
var _ I = T{}       // T'nin I'ya tanımlandığını doğrulayın.
var _ I = (*T)(nil) // *T'nin I'ya tanımlandığını doğrulayın.
```

T \(veya \* T, buna göre\) I'ya uygulamazsa, hata derleme zamanında yakalanacaktır. Bir arayüzün kullanıcılarının onu uyguladıklarını açıkça beyan etmelerini isterseniz, arayüzün yöntem kümesine açıklayıcı bir ada sahip bir yöntem ekleyebilirsiniz. Örneğin:

```go
type Fooer interface {
    Foo()
    ImplementsFooer()
}
```

Bir tür daha sonra bir Fooer olmak için ImplementsFooer yöntemini uygulamalı, gerçeği açıkça belgelendirmeli ve [go doc](https://golang.org/cmd/go/#hdr-Show_documentation_for_package_or_symbol)'un çıktısında duyurmalıdır.

```go
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

Çoğu kod, arayüz fikrinin faydasını sınırladıkları için bu tür kısıtlamaları kullanmaz. Yine de bazen benzer arayüzler arasındaki belirsizlikleri çözmek için gereklidirler.

#### T türü neden Equal arayüzünü karşılamıyor?

Bu basit arayüzü, kendisini başka bir değerle karşılaştırabilen bir nesneyi temsil edecek şekilde düşünün:

```go
type Equaler interface {
    Equal(Equaler) bool
}
```

ve bu tür, T:

```go
type T int
func (t T) Equal(u T) bool { return t == u } // Equaler'ı karşılamıyor
```

Bazı polimorfik tip sistemlerde benzer durumdan farklı olarak, T, Equaler'ı uygulamaz. T.Equal'ın bağımsız değişken türü T'dir, tam anlamıyla gerekli olan Equaler türü değildir.

Go'da tip sistemi Equal argümanını desteklemez; Bu, Equaler'ı uygulayan T2 türünde gösterildiği gibi programcının sorumluluğundadır:

```go
type T2 int
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }  //Equalerı karşılar
```

Bu bile diğer tür sistemler gibi değildir, çünkü Go'da Equaler'ı karşılayan herhangi bir tür argüman olarak T2.Equal'a aktarılabilir ve çalışma zamanında argümanın T2 türünde olup olmadığını kontrol etmeliyiz. Bazı diller bu garantiyi derleme zamanında verir. İlgili bir örnek diğer tarafa gider:

```go
type Opener interface {
   Open() Reader
}

func (t T3) Open() *os.File
```

Go'da T3, başka bir dilde olsa da Opener'ı karşılamıyor. Go'nun tip sisteminin bu gibi durumlarda programcı için daha az şey yaptığı doğru olsa da, alt tiplemenin olmaması arayüz memnuniyeti ile ilgili kuralları belirtmeyi çok kolaylaştırır: fonksiyonun adları ve imzaları tam olarak arayüzünkiler mi?

Go kuralının verimli bir şekilde uygulanması da kolaydır. Bu avantajların, otomatik tip promosyon eksikliğini telafi ettiğini düşünüyoruz. Bir gün, bir tür polimorfik yazım benimsemesi durumunda, bu örneklerin fikrini ifade etmenin ve ayrıca bunların statik olarak kontrol edilmesini sağlamanın bir yolu olacağını umuyoruz.

#### Bir \[\]T'yi bir \[\]interface'e dönüştürebilir miyim?

Direkt olarak değil. İki tür bellekte aynı temsile sahip olmadığından dil belirtimine göre buna izin verilmez. Öğeleri tek tek hedef dilime kopyalamak gerekir. Bu örnek, bir int dilimini bir interface{} dilimine dönüştürür:

```go
t := []int{1, 2, 3, 4}
s := make([]interface{}, len(t))
for i, v := range t {
    s[i] = v
}
```

#### T1 ve T2 aynı temel türe sahipse \[\] T1'i \[\] T2'ye dönüştürebilir miyim?

Bu kod örneğinin bu son satırı derlenmiyor.

```go
type T1 int
type T2 int
var t1 T1
var x = T2(t1) // Tamam
var st1 []T1
var sx = ([]T2)(st1) // tamam değil
```

Go'da türler, her adlandırılmış türün \(muhtemelen boş\) bir yöntem kümesine sahip olması nedeniyle metodlara yakından bağlıdır. Genel kural, dönüştürülen türün adını değiştirebilmeniz \(ve dolayısıyla muhtemelen yöntem kümesini değiştirebilmeniz\), ancak bileşik türdeki öğelerin adını \(ve yöntem kümesini\) değiştirememenizdir. Go, tür dönüşümleri konusunda açık olmanızı gerektirir.

#### Nil hata değerim neden nil'e eşit değil?

Kapakların altında, arayüzler iki öğe olarak uygulanır; bir T türü ve bir V değeri V, int, struct veya işaretçi gibi somut bir değerdir, hiçbir zaman arabirim değildir ve T türüne sahiptir. Örneğin, int değeri 3 bir arabirimde, ortaya çıkan arabirim değeri şematik olarak \(T = int, V = 3\) olur. V değeri, aynı zamanda arayüzün dinamik değeri olarak da bilinir, çünkü belirli bir arayüz değişkeni, programın yürütülmesi sırasında farklı V değerlerini \(ve karşılık gelen T tiplerini\) tutabilir. Bir arabirim değeri, yalnızca V ve T'nin her ikisi de ayarlanmamışsa sıfırdır \(T = nil, V ayarlanmamış\). Özellikle, bir sıfır arabirim her zaman bir sıfır türünü tutacaktır. Bir arabirim değerinin içinde  _int türünde bir sıfır gösterici saklarsak, iç tür, işaretçinin değerine bakılmaksızın_  int olacaktır: \(T = \* int, V = nil\). Bu nedenle, bu tür bir arayüz değeri, içindeki işaretçi değeri V sıfır olduğunda bile sıfır olmayacaktır. Bu durum kafa karıştırıcı olabilir ve bir sıfır değeri gibi bir arayüz değerinin içinde saklandığında ortaya çıkar. hata dönüşü:

```go
func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = ErrBad
	}
	return p // herzaman nil olmayan hata döndürür
}
```

Her şey yolunda giderse, işlev bir nil p döndürür, dolayısıyla dönüş değeri bir hata arabirim değeridir \(T = \* MyError, V = nil\). Bu, arayanın döndürdüğü hatayı nil ile karşılaştırması durumunda, kötü bir şey olmasa bile her zaman bir hata varmış gibi görüneceği anlamına gelir. Çağırana uygun bir sıfır hatası döndürmek için, işlevin açık bir nil döndürmesi gerekir:

```go
func returnsError() error {
	if bad() {
		return ErrBad
	}
	return nil
}
```

Hataları döndüren fonksiyonların, hatanın doğru şekilde oluşturulmasını garantilemeye yardımcı olmak için  _`MyError` gibi somut bir tür yerine \(yukarıda yaptığımız gibi\) imzalarında her zaman hata türünü kullanmaları iyi bir fikirdir. Örnek olarak,_ [_os.Open_](https://golang.org/pkg/os/#Open)_, sıfır değilse bile, her zaman somut_  [\*os.PathError](https://golang.org/pkg/os/#PathError) türünde olsa bile bir hata döndürür. Burada açıklananlara benzer durumlar, arayüzler her kullanıldığında ortaya çıkabilir. Arayüzde herhangi bir somut değer saklandıysa arayüzün `nil` olmayacağını unutmayın. Daha fazla bilgi için [Yansıma Yasaları](https://golang.org/doc/articles/laws_of_reflection.html)'na bakın.

#### C'de olduğu gibi neden untagged unions yok?

Untagged unions, Go'nun bellek güvenliği garantilerini ihlal eder.

#### Go'da neden varyant türleri yok?

Cebirsel türler olarak da bilinen değişken türleri, bir değerin diğer türlerden birini, ancak yalnızca bu türleri alabileceğini belirtmenin bir yolunu sağlar. Sistem programlamasında yaygın bir örnek, bir hatanın, örneğin bir ağ hatası, bir güvenlik hatası veya bir uygulama hatası olduğunu belirtir ve arayanın, hatanın türünü inceleyerek sorunun kaynağını ayırt etmesine izin verir.

Başka bir örnek, her bir düğümün farklı bir türde olabileceği bir sözdizimi ağacıdır: bildirim, ifade, atama vb. Go'ya varyant türleri eklemeyi düşündük, ancak tartışmadan sonra, arayüzlerle kafa karıştırıcı şekillerde çakıştıkları için bunları dışarıda bırakmaya karar verdik.

Bir varyant türünün öğelerinin kendileri arayüz olsaydı ne olurdu? Ayrıca, hangi varyant türlerinin adreslendiğinden bazıları zaten dil tarafından kapsanmaktadır. Hata örneğini, hatayı tutmak için bir arayüz değeri ve durumları ayırt etmek için bir tür anahtarı kullanarak ifade etmek kolaydır. Sözdizimi ağacı örneği, o kadar zarif olmasa da yapılabilir.

#### Go neden ortak değişken sonuç türlerine sahip değil?

Kovaryant \(ortak değişken\) sonuç türleri, benzer bir arayüzün

```go
type Copyable interface {
	Copy() interface{}
}
```

yöntemle karşılanmalı

```go
func (v Value) Copy() Value
```

Çünkü Değer, boş arayüzü uygular. Go metodunda türlerin tam olarak eşleşmesi gerekir, bu nedenle Value Copyable'ı uygulamaz. Go, bir türün ne yaptığı kavramını - yöntemlerini - türün uygulamasından ayırır. İki yöntem farklı türler döndürürse, aynı şeyi yapmazlar. Kovaryant sonuç türleri isteyen programcılar genellikle arayüzler aracılığıyla bir tür hiyerarşisini ifade etmeye çalışırlar. Go'da arayüz ve uygulama arasında temiz bir ayrım olması daha doğaldır.

###  Değerler

#### Go neden örtük sayısal dönüşümler sağlamaz?

C'deki sayısal türler arasında otomatik dönüşümün rahatlığı, neden olduğu kafa karışıklığından daha ağır basmaktadır. Bir ifade ne zaman işaretsizdir? Değer ne kadar büyük? Taşıyor mu? Sonuç, üzerinde yürütüldüğü makineden bağımsız olarak taşınabilir mi? Ayrıca derleyiciyi karmaşıklaştırır; "Olağan aritmetik dönüşümlerin" uygulanması kolay değildir ve mimariler arasında tutarsızdır. 

Taşınabilirlik nedenlerinden dolayı, koddaki bazı açık dönüşümler pahasına işleri net ve anlaşılır hale getirmeye karar verdik. Go'daki sabitlerin tanımı - işaretsiz ve boyut ek açıklamaları içermeyen keyfi kesinlik değerleri - yine de sorunları önemli ölçüde iyileştirir. Bununla ilgili bir ayrıntı, C'den farklı olarak int ve int64'ün, int 64 bitlik bir tür olsa bile farklı türler olmasıdır. İnt türü geneldir; Bir tamsayının kaç bit tuttuğunu önemsiyorsanız, Go sizi açık sözlü olmaya teşvik eder.

#### Sabitler Go'da nasıl çalışır?

Go, farklı sayısal türlerdeki değişkenler arasında dönüşüm konusunda katı olmasına rağmen, dildeki sabitler çok daha esnektir. 23, 3.14159 ve math.Pi gibi değişmez sabitler, keyfi bir hassasiyetle ve taşma veya yetersizlik olmadan bir tür ideal sayı alanı kaplar. Örneğin, math.Pi'nin değeri kaynak kodda 63 basamak olarak belirtilir ve değeri içeren sabit ifadeler, bir float64'ün tutabileceğinin ötesinde hassasiyeti korur. Yalnızca sabit veya sabit ifade bir değişkene \(programdaki bir bellek konumuna\) atandığında, normal kayan nokta özelliklerine ve hassasiyetine sahip bir "bilgisayar" numarası olur.

Ayrıca, bunlar tiplenmiş değerler değil, sadece sayılar oldukları için, Go'daki sabitler değişkenlere göre daha özgürce kullanılabilir, böylece katı dönüştürme kuralları etrafındaki bazı gariplikleri yumuşatır.Aşağıdaki  gibi ifadeler yazılabilir:

```go
sqrt2 := math.Sqrt(2)
```

derleyiciden şikayet etmeden ideal sayı 2, math.Sqrt çağrısı için güvenli ve doğru bir şekilde float64'e dönüştürülebilir.

[Sabitler](https://blog.golang.org/constants) başlıklı bir blog yazısı bu konuyu daha ayrıntılı olarak araştırıyor. Ve bu dökümandan sabitleri inceleyebilirsiniz.

sabitler.md

#### Map neden built-in \(yerleşik\)?

Stringler ile aynı nedeni şudur: o kadar güçlü ve önemli bir veri yapısıdır ki sözdizimsel destek ile mükemmel bir uygulama sağlamak, programlamayı daha keyifli hale getirir. Go'nun map uygulamasının, kullanımların büyük çoğunluğuna hizmet edecek kadar güçlü olduğuna inanıyoruz. Belirli bir uygulama özel bir uygulamadan yararlanabiliyorsa, bir tane yazmak mümkündür, ancak sözdizimsel olarak o kadar kullanışlı olmayacaktır; bu makul bir değiş tokuş gibi görünüyor.

#### Map neden dilimlere \(slices\) anahtar \(key\) olarak izin vermiyor?

Harita arama, dilimlerin uygulanmadığı bir eşitlik operatörü gerektirir. Eşitliği uygulamazlar çünkü eşitlik bu tür türlerde iyi tanımlanmamıştır; sığ ve derin karşılaştırma, işaretçi ile değer karşılaştırması, özyinelemeli türlerle nasıl başa çıkılacağı gibi birçok husus vardır. Bu konuyu tekrar gözden geçirebiliriz - ve dilimler için eşitliği uygulamak mevcut programları geçersiz kılmaz - ancak dilimlerin eşitliğinin ne anlama geldiği konusunda net bir fikir olmadan, şimdilik bunu dışarıda bırakmak daha kolaydı.

Go 1'de, önceki sürümlerden farklı olarak, yapılar ve diziler için eşitlik tanımlanmıştır, bu nedenle bu tür türler eşleme anahtarları olarak kullanılabilir. Yine de dilimler hala bir eşitlik tanımına sahip değil.

#### Diziler değer iken neden map'ler, dilimler ve kanallar referanslarıdır?

Bu konuyla ilgili çok fazla tarih var. Eskiden, map'ler ve kanallar sözdizimsel olarak işaretçilerdi ve işaretçi olmayan bir örneği bildirmek veya kullanmak imkansızdı. Ayrıca dizilerin nasıl çalışması gerektiğiyle uğraştık. Sonunda, işaretçilerin ve değerlerin katı bir şekilde ayrılmasının, dili kullanmayı zorlaştırdığına karar verdik. Bu türlerin ilişkili, paylaşılan veri yapılarına referans görevi görecek şekilde değiştirilmesi bu sorunları çözdü. Bu değişiklik, dile üzücü bir karmaşıklık kattı ancak kullanılabilirlik üzerinde büyük bir etkiye sahipti: Go, tanıtıldığında daha üretken ve rahat bir dil haline geldi.

###  Kod Yazma

#### Kütüphaneler nasıl belgelenir?

Go'da yazılmış, kaynak koddan paket belgelerini çıkaran ve bunu bildirimlere, dosyalara vb. Bağlantılar içeren bir web sayfası olarak hizmet veren bir program var. [golang.org/pkg/](https://golang.org/pkg/) adresinde bir örnek çalışıyor. Aslında, `godoc` sitenin tamamını [golang.org/](https://golang.org/) adresinde uygulamaktadır.

Bir godoc örneği, görüntülediği programlarda sembollerin zengin, etkileşimli statik analizlerini sağlamak üzere yapılandırılabilir; detaylar [burada](https://golang.org/lib/godoc/analysis/help.html) listelenmiştir.

Komut satırından belgelere erişim için, [go](https://golang.org/pkg/cmd/go/) aracı aynı bilgilere bir metinsel arayüzü sağlayan bir [doc](https://golang.org/pkg/cmd/go/#hdr-Show_documentation_for_package_or_symbol) alt komutunu vardır.

#### Go programlama stili kılavuzu var mı?

Kesinlikle tanınabilir bir "Go stili" olmasına rağmen açık bir stil kılavuzu yoktur.

Go, adlandırma, düzen ve dosya organizasyonu ile ilgili kararlara rehberlik edecek kurallar oluşturmuştur. [Effective Go](https://golang.org/doc/effective_go.html) belgesi bu konularla ilgili bazı tavsiyeler içermektedir. Daha doğrusu, gofmt programı, amacı düzen kurallarını uygulamak olan güzel bir yazıcıdır; yorumlamaya izin veren olağan yapılması ve yapılmaması gerekenler özetinin yerini alır. Depodaki tüm Go kodu ve açık kaynak dünyasının büyük çoğunluğu gofmt aracılığıyla çalıştırılmıştır.

[Go Code Review Comments](https://golang.org/s/comments) başlıklı belge, programcılar tarafından genellikle gözden kaçırılan Go deyiminin ayrıntılarıyla ilgili çok kısa denemelerden oluşan bir derlemedir. Go projeleri için kod incelemeleri yapan kişiler için kullanışlı bir referanstır.

#### Yamaları Go kütüphanelerine nasıl gönderirim?

Kütüphane kaynakları, arşivin `src` dizinindedir. Önemli bir değişiklik yapmak istiyorsanız, lütfen başlamadan önce posta listesinde tartışın.

Nasıl ilerleyeceğiniz hakkında daha fazla bilgi için [Go projesine katkıda bulunmak](https://golang.org/doc/contribute.html) belgesine bakın.

#### "go get" bir depoyu klonlarken neden HTTPS kullanıyor?

Şirketler genellikle giden trafiğe yalnızca standart TCP bağlantı noktaları 80 \(HTTP\) ve 443 \(HTTPS\) üzerinden izin vererek, TCP bağlantı noktası 9418 \(git\) ve TCP bağlantı noktası 22 \(SSH\) dahil olmak üzere diğer bağlantı noktalarında giden trafiği engeller. HTTP yerine HTTPS kullanırken, git varsayılan olarak sertifika doğrulamasını zorlayarak ortadaki adam, gizli dinleme ve kurcalama saldırılarına karşı koruma sağlar. Bu nedenle `go get`komutu, güvenlik için HTTPS kullanır. Git, HTTPS üzerinden kimlik doğrulaması yapacak veya HTTPS yerine SSH kullanacak şekilde yapılandırılabilir. HTTPS üzerinden kimlik doğrulaması yapmak için `$ HOME / .netrc` dosyasına git'in başvurduğu bir satır ekleyebilirsiniz:

```text
machine github.com login USERNAME password APIKEY
```

GitHub hesapları için parola [kişisel bir erişim belirteci](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) olabilir.

Git, belirli bir önekle eşleşen URL'ler için HTTPS yerine SSH kullanacak şekilde de yapılandırılabilir. Örneğin, tüm GitHub erişiminde SSH kullanmak için şu satırları `~/.gitconfig` dosyanıza ekleyin:

```text
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

#### "Go get" kullanarak paket sürümlerini nasıl yönetmeliyim?

Projenin başlangıcından bu yana, Go'nun açık bir paket sürümleri kavramı yoktu, ancak bu değişiyor. Versiyon oluşturma, özellikle büyük kod tabanlarında önemli bir karmaşıklık kaynağıdır ve tüm Go kullanıcılarına sağlanmaya uygun olacak kadar geniş bir çeşitlilikteki durumlarda ölçekte iyi çalışan bir yaklaşım geliştirmek biraz zaman almıştır. Go 1.11 sürümü, Go modülleri biçiminde go komutuna paket sürüm oluşturma için yeni, deneysel destek ekler. Daha fazla bilgi için [Go 1.11 sürüm notlarına](https://golang.org/doc/go1.11#modules) ve [go komutu belgeleri](https://golang.org/cmd/go#hdr-Modules__module_versions__and_more)ne bakın.

Gerçek paket yönetimi teknolojisinden bağımsız olarak, "go get" ve daha büyük Go araç zinciri, farklı içe aktarma yollarına sahip paketlerin izolasyonunu sağlar. Örneğin, standart kitaplığın `html/template` ve `text/template`, her ikisi de "package template" olsa bile bir arada bulunur.

Bu gözlem, paket yazarları ve paket kullanıcıları için bazı tavsiyelere yol açar. Kamusal kullanıma yönelik paketler, geriye dönük uyumluluğu korumaya çalışmalıdır. gelişmek. [Go 1 uyumluluk yönergeleri](https://golang.org/doc/go1compat.html) burada iyi bir referanstır: dışa aktarılan adları kaldırmayın, etiketli bileşik değişmez değerleri teşvik edin, vb. Farklı işlevler gerekiyorsa, eskisini değiştirmek yerine yeni bir ad ekleyin. Tam bir ara verilmesi gerekiyorsa, yeni bir içe aktarma yolu ile yeni bir paket oluşturun. Harici olarak sağlanan bir paket kullanıyorsanız ve beklenmedik şekillerde değişebileceğinden endişeleniyorsanız, ancak henüz Go modüllerini kullanmıyorsanız, en basit çözüm onu ​​yerel deponuza kopyalamaktır. Bu, Google'ın dahili olarak kullandığı yaklaşımdır ve "vendoring" adı verilen bir teknik aracılığıyla go komutuyla desteklenir. Bu, bağımlılığın bir kopyasını, onu yerel bir kopya olarak tanımlayan yeni bir içe aktarma yolu altında depolamayı içerir. Ayrıntılar için [tasarım belgesi](https://golang.org/s/go15vendor)ne bakın.

###  İşaretçiler ve Tahsis

#### Fonksiyon parametreleri ne zaman değere göre aktarılır?

C ailesindeki tüm dillerde olduğu gibi, Go'daki her şey değer bazında aktarılır. Diğer bir deyişle, bir işlev, değeri parametreye atayan bir atama ifadesi varmış gibi, her zaman iletilmekte olan şeyin bir kopyasını alır. Örneğin, bir fonksiyona bir int değerinin iletilmesi, int'in bir kopyasını oluşturur ve bir işaretçi değerinin iletilmesi, işaretçinin bir kopyasını oluşturur, ancak gösterdiği verilerin bir kopyasını oluşturmaz. \(Bunun yöntem alıcılarını nasıl etkilediğiyle ilgili bir tartışma için sonraki bölüme bakın.\) Map ve dilim \(slice\) değerleri, işaretçiler gibi davranır: bunlar, temel alınan map'e veya dilim verilerine işaretçiler içeren tanımlayıcılardır.

Bir map'in veya dilim değerinin kopyalanması, işaret ettiği verileri kopyalamaz. Bir arayüz değerinin kopyalanması, arayüz değerinde depolanan şeyin bir kopyasını oluşturur. Arayüz değeri bir yapı \(struct\) içeriyorsa, arayüz değerinin kopyalanması yapının bir kopyasını oluşturur. Arabirim değeri bir işaretçi tutuyorsa, arayüz değerinin kopyalanması işaretçinin bir kopyasını oluşturur, ancak yine işaret ettiği veriyi oluşturmaz. Bu tartışmanın, işlemlerin anlambilim. Gerçek uygulamalar, optimizasyonlar anlambilimini değiştirmediği sürece kopyalamayı önlemek için optimizasyonları uygulayabilir.

#### Bir arayüze işaretçi ne zaman kullanmalıyım?

Neredeyse hiç. Arayüz değerlerine yönelik işaretçiler yalnızca, gecikmeli değerlendirme için bir arayüz değerinin türünü gizlemeyi içeren nadir, zor durumlarda ortaya çıkar.

Bir arayüz bekleyen bir işleve bir arayüz değerine bir işaretçi iletmek yaygın bir hatadır. Derleyici bu hatadan şikayet eder, ancak durum yine de kafa karıştırıcı olabilir, çünkü bazen bir [arayüzü karşılamak için bir işaretçi](https://golang.org/doc/faq#different_method_sets) gerekir. Buradaki fikir, somut bir tipe bir işaretçi bir arayüzü karşılayabilmesine rağmen, bir istisna dışında, bir arayüze yönelik bir işaretçinin bir arayüze asla karşılamayacağıdır.

Değişken tanımlamayı düşünün,

```go
var w io.Writer
```

fmt.Fprintf yazdırma fonksiyonu, ilk argüman olarak io.Writer'ı karşılayan bir değer alır - bu, kurallı Write metodunu uygulayan bir şeydir. Böylece yazabiliriz

Ancak w'nin adresini geçersek, program derlenmeyecektir.

```go
fmt.Fprintf(&w, "hello, world\n") //Derleme-zamanı hatası.
```

Bunun tek istisnası, herhangi bir değerin, hatta bir arayüze işaretçi bile, boş arayüze türündeki bir değişkene \(interface{}\) atanabilmesidir. Öyle olsa bile, değerin bir arayüze işaretçi olması neredeyse kesinlikle bir hatadır; sonuç kafa karıştırıcı olabilir.

#### Metodları \(struct fonksiyonlar\) değerler veya işaretçiler üzerinde tanımlamalı mıyım?

```go
func (s *MyStruct) pointerMethod() { } // işaretçi metod
func (s MyStruct)  valueMethod()   { } // değer metod
```

İşaretçilere alışkın olmayan programcılar için bu iki örnek arasındaki ayrım kafa karıştırıcı olabilir, ancak durum aslında çok basittir. Bir tür üzerinde bir metodu tanımlarken, alıcı \(yukarıdaki örneklerde `s`\) tam olarak metodun bir argümanıymış gibi davranır. Alıcının bir değer olarak mı yoksa bir işaretçi olarak mı tanımlanacağı aynı sorudur, o halde, bir fonksiyon bağımsız değişkeninin bir değer mi yoksa bir işaretçi mi olması gerektiği sorusudur.

Dikkat edilmesi gereken birkaç nokta var. İlk ve en önemlisi, metodun alıcıyı değiştirmesi gerekiyor mu? Eğer öyleyse, alıcının bir işaretçi olması gerekir. \(Dilimler ve map'ler referans görevi görür, bu nedenle hikayeleri biraz daha inceliklidir, ancak örneğin bir yöntemdeki bir dilimin uzunluğunu değiştirmek için alıcının yine de bir işaretçi olması gerekir.\)

Yukarıdaki örneklerde, `pointerMethod` alanlarını değiştirirse `s`, arayan bu değişiklikleri görecektir, ancak `valueMethod`, arayanın argümanının bir kopyasıyla çağrılır \(bu, bir değeri iletmenin tanımıdır\), bu nedenle yaptığı değişiklikler görünmez olacaktır arayana.

Bu arada, Java metodlarında alıcılar her zaman işaretçilerdir, ancak işaretçi doğaları bir şekilde gizlenmiştir \(ve dile değer alıcıları eklemek için bir öneri vardır\). Sıradışı olan, Go'daki değer alıcılarıdır. İkincisi, verimliliğin dikkate alınmasıdır. Alıcı büyükse, örneğin büyük bir yapıya sahipse, bir işaretçi alıcı kullanmak çok daha ucuz olacaktır.

Sırada tutarlılık var. Türün bazı metodlarının işaretçi alıcıları olması gerekiyorsa, geri kalanı da olmalıdır, bu nedenle metod kümesi türün nasıl kullanıldığına bakılmaksızın tutarlıdır. Ayrıntılar için [metod kümeleri](https://golang.org/doc/faq#different_method_sets) bölümüne bakın.

Temel türler, dilimler ve küçük yapılar gibi türler için, bir değer alıcısı çok ucuzdur, bu nedenle metodun anlambilimi bir işaretçi gerektirmedikçe, bir değer alıcısı verimli ve nettir.

#### Make ve New arasındaki fark nedir?

**Kısaca:** new, bellek ayırır. Make ise dilim, map ve kanal türlerini başlatır.

Daha fazla ayrıntı için [Effective Go'nun ilgili bölümü](https://golang.org/doc/effective_go.html#allocation_new)ne bakın.

#### Bir int türün 64bitlik makinedeki büyüklüğü nedir?

Int ve uint boyutları uygulamaya özgüdür ancak belirli bir platformda birbirleriyle aynıdır. Taşınabilirlik için, belirli bir değer boyutuna dayanan kod, int64 gibi açıkça boyutlandırılmış bir tür kullanmalıdır. 32-bit makinelerde derleyiciler varsayılan olarak 32-bit tamsayılar kullanırken, 64-bit makinelerde tamsayılar 64 bit'e sahiptir. \(Tarihsel olarak bu her zaman doğru değildi.\)

Öte yandan, kayan nokta skalerleri ve karmaşık türler her zaman boyutlandırılır \(kayan nokta veya karmaşık temel türler yoktur\), çünkü programcılar kayan noktalı \(burada float'dan bahsediyor\) sayıları kullanırken hassasiyetin farkında olmalıdır. Bir \(türlenmemiş - untyped\) kayan nokta sabiti için kullanılan varsayılan tür float64'tür. Böylece `foo: = 3.0, float64` türünde bir değişken `foo` bildirir. Bir \(türlenmemiş\) sabit tarafından başlatılan `float32` değişkeni için değişken türü, değişken bildiriminde açıkça belirtilmelidir:

```go
var foo float32 = 3.0
```

Alternatif olarak, sabite `foo: = float32 (3.0)` 'da olduğu gibi dönüşümlü bir tür verilmelidir.

#### Bir değişkenin yığın üzerinde mi yoksa yığın üzerinde tahsis edildiğini nasıl anlarım?

Doğruluk açısından, bilmenize gerek yok. Go'daki her değişken, kendisine referans olduğu sürece var olur. Uygulama tarafından seçilen depolama konumu, dilin anlambilimiyle ilgisizdir.

Depolama konumu, verimli programlar yazma üzerinde bir etkiye sahiptir. Mümkün olduğunda, Go derleyicileri, o işlevin yığın çerçevesindeki bir işleve yerel olan değişkenleri tahsis eder. Ancak, derleyici işlev döndükten sonra değişkene başvurulmadığını kanıtlayamazsa, derleyicinin işaretçi hatalarının sarkmasını önlemek için değişkeni çöpte toplanan yığın üzerinde tahsis etmesi gerekir.

Ayrıca, yerel bir değişken çok büyükse, onu yığın yerine yığın üzerinde depolamak daha mantıklı olabilir. Mevcut derleyicilerde, bir değişkenin adresi alınmışsa, bu değişken öbek üzerinde tahsisat için bir adaydır. Bununla birlikte, temel bir kaçış analizi, bu tür değişkenlerin işlevin geri dönüşünü geçemez ve tanımsız yığındır.

#### Go işlemim neden bu kadar çok sanal bellek kullanıyor?

Go bellek ayırıcısı, ayırmalar için bir alan olarak büyük bir sanal bellek bölgesini ayırır. Bu sanal bellek, belirli Go işlemi için yereldir; rezervasyon, diğer bellek işlemlerini mahrum etmez.

Bir Go işlemine ayrılan gerçek bellek miktarını bulmak için Unix `top` komutunu kullanın ve RES \(Linux\) veya RSIZE \(macOS\) sütunlarına bakın.

###  Eşzamanlılık

#### Hangi işlemler atomiktir? Mutexler ne olacak?

Go'daki işlemlerin atomikliğine ilişkin bir açıklama [Go Bellek Modeli](https://golang.org/ref/mem) belgesinde bulunabilir.

Düşük seviyeli senkronizasyon ve atomik ilkeller, [sync](https://golang.org/pkg/sync) ve [sync/atomic](https://golang.org/pkg/sync/atomic) paketlerinde mevcuttur.

Bu paketler, referans sayılarını artırmak veya küçük ölçekli karşılıklı dışlamayı garanti etmek gibi basit görevler için iyidir. Eşzamanlı sunucular arasında koordinasyon gibi daha yüksek seviyeli işlemler için, daha yüksek seviyeli teknikler daha güzel programlara yol açabilir ve Go, bu yaklaşımı kendi düzenleri ve kanalları aracılığıyla destekler. Örneğin, programınızı belirli bir veri parçasından her seferinde yalnızca bir gorutin sorumlu olacak şekilde yapılandırabilirsiniz. Bu yaklaşım, orijinal [Go atasözü](https://www.youtube.com/watch?v=PAAkCSZUG1c) ile özetlenmiştir.

Hafızayı paylaşarak iletişim kurmayın. Bunun yerine, iletişim kurarak hafızayı paylaşın. Bu kavramın ayrıntılı bir tartışması için, [İletişimi Kurarak Belleği Paylaşma](https://golang.org/doc/codewalk/sharemem/) [kod yürüyüşüne ve ilgili](https://blog.golang.org/2010/07/share-memory-by-communicating.html) makalesine bakın. Büyük eşzamanlı programlar muhtemelen bu iki araç setinden ödünç alır.

#### Programım neden daha fazla CPU ile daha hızlı çalışmıyor?

Bir programın daha fazla CPU ile daha hızlı çalışıp çalışmadığı, çözmekte olduğu soruna bağlıdır. Go dili, gorutinler ve kanallar gibi eşzamanlılık ilkelleri sağlar, ancak eşzamanlılık yalnızca temeldeki sorun doğası gereği paralel olduğunda paralellik sağlar.

Doğası gereği sıralı olan sorunlar, daha fazla CPU eklenerek hızlandırılamazken, paralel olarak yürütülebilen parçalara ayrılabilen sorunlar bazen dramatik bir şekilde hızlandırılabilir. Bazen daha fazla CPU eklemek bir programı yavaşlatabilir. Pratik anlamda, eşitleme veya iletişim için yararlı hesaplamalar yapmaktan daha fazla zaman harcayan programlar, birden çok işletim sistemi iş parçacığı kullanırken performans düşüşü yaşayabilir. Bunun nedeni, iş parçacıkları arasında veri geçişinin, önemli bir maliyeti olan bağlamları değiştirmeyi gerektirmesidir ve bu maliyet daha fazla CPU ile artabilir. Örneğin, Go spesifikasyonundaki ana elek örneğinin pek çok gorutini başlatmasına rağmen önemli bir paralelliği yoktur; sayısını artırmak thread'ler \(CPU'lar\) hızlandırmaktan çok yavaşlatır.

Bu konu hakkında daha fazla ayrıntı için [Concurrency is not Parallelism](https://blog.golang.org/2013/01/concurrency-is-not-parallelism.html) başlıklı konuşmaya bakın.

#### CPU sayısını nasıl kontrol edebilirim?

Yürütülen programlı dizinler için aynı anda kullanılabilen CPU sayısı, varsayılan değeri mevcut CPU çekirdeği sayısı olan `GOMAXPROCS` kabuk ortam değişkeni tarafından kontrol edilir. Paralel yürütme potansiyeli olan programlar bu nedenle bunu varsayılan olarak çoklu CPU'lu bir makinede gerçekleştirmelidir. Kullanılacak paralel CPU sayısını değiştirmek için ortam değişkenini ayarlayın veya çalışma zamanı desteğini farklı sayıda iş parçacığı kullanacak şekilde yapılandırmak için çalışma zamanı paketinin benzer adlandırılmış işlevini kullanın. 1 olarak ayarlamak, bağımsız gorutinleri sırayla yürütmeye zorlayarak gerçek paralellik olasılığını ortadan kaldırır. Çalışma zamanı, birden çok bekleyen G / Ç isteğine hizmet vermek için `GOMAXPROCS` değerinden daha fazla iş parçacığı ayırabilir. 

`GOMAXPROCS` yalnızca aynı anda kaç tane gorutinin çalıştırabileceğini etkiler; sistem çağrılarında keyfi olarak daha fazlası engellenebilir. Go'nun gorutin programlayıcısı, zamanla gelişmesine rağmen olması gerektiği kadar iyi değildir. Gelecekte daha iyi olabilir OS iş parçacığı kullanımını optimize eder. Şimdilik, performans sorunları varsa, `GOMAXPROCS`'u uygulama bazında ayarlamak yardımcı olabilir.

#### Gorutin olarak çalışan kapanışlarda ne olur?

Eşzamanlı kapanışlar kullanıldığında bazı karışıklıklar ortaya çıkabilir. Aşağıdaki programı düşünün:

```go
func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // çıkmadan önce tüm goroutinlerin tamamlanmasını bekle
    for _ = range values {
        <-done
    }
}
```

Bir kimse yanlışlıkla çıktı olarak `a, b, c`'yi görmeyi bekleyebilir. Bunun yerine muhtemelen göreceğiniz şey `c, c, c`'dir. Bunun nedeni, döngünün her yinelemesinin `v` değişkeninin aynı örneğini kullanmasıdır, bu nedenle her kapanış bu tek değişkeni paylaşır. Kapatma çalıştığında, `fmt.Println` yürütüldüğünde `v` değerini yazdırır, ancak gorutin başlatıldıktan sonra `v` değiştirilmiş olabilir. Bunu ve diğer sorunları ortaya çıkmadan önce tespit etmeye yardımcı olmak için [go vet](https://golang.org/cmd/go/#hdr-Run_go_tool_vet_on_packages) komutunu kullanın.

`v`'nin geçerli değerini, başlatıldığında her kapanışa bağlamak için, her yinelemede yeni bir değişken oluşturmak için iç döngü değiştirilmelidir. Bir yol, değişkeni kapanışa bir argüman olarak iletmektir:

```go
    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
```

Bu örnekte, `v`'nin değeri anonim işleve argüman olarak aktarılır. Bu değere daha sonra fonksiyonun içinde `u` değişkeni olarak erişilebilir. Daha da kolay olanı, tuhaf görünebilecek ancak Go'da iyi çalışan bir tanımlama stili kullanarak yeni bir değişken oluşturmaktır:

```go
    for _, v := range values {
        v := v // create a new 'v'.
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }
```

Her yineleme için yeni bir değişken tanımlamayan dilin bu davranışı, geçmişe bakıldığında bir hata olabilir. Daha sonraki bir sürümde ele alınabilir, ancak uyumluluk için Go sürüm 1'de değiştirilemez.

###  Fonksiyonlar ve Metodlar \(Structlar için fonksiyonlar\)

#### T ve \* T'nin neden farklı metod atamaları var?

[Go spesifikasyonunun](https://golang.org/ref/spec#Types) dediği gibi, T türünün metod kümesi, alıcı türü T olan tüm metodlardan oluşurken, karşılık gelen işaretçi türü  _T, alıcı_  T veya \*T olan tüm yöntemlerden oluşur. Bu,  _\*T yöntem kümesidir. T'yi içerir, ancak tersini içermez. Bu ayrım, bir arabirim değeri bir_  T işaretçisi içeriyorsa, bir yöntem çağrısı işaretçinin başvurusunu kaldırarak bir değer elde edebildiği için ortaya çıkar, ancak bir arabirim değeri bir T değeri içeriyorsa, bir işaretçi elde etmek için bir yöntem çağrısı için güvenli bir yol yoktur. \(Bunu yapmak, bir yöntemin arayüz içindeki değerin içeriğini değiştirmesine izin verir ve bu, dil spesifikasyonu tarafından izin verilmemektedir.\) Derleyicinin yönteme geçmek için bir değerin adresini alabildiği durumlarda bile, yöntem değeri değiştirirse, değişiklikler çağırıcıda kaybolur. Örnek olarak, bytes.Buffer'ın Write yöntemi bir işaretçi yerine bir değer alıcısı kullanıyorsa, bu kod:

```go
var buf bytes.Buffer
io.Copy(buf, os.Stdin)
```

standart girdiyi `buf`'ın kendisine değil buf'ın bir `Copy` elemanına kopyalar. Bu neredeyse hiçbir zaman istenen davranış değildir.

###  Kontrol Akışı

#### Go neden ?: operatörüne sahip değil?

Go'da üçlü test işlemi yoktur. Aynı sonucu elde etmek için aşağıdakileri kullanabilirsiniz:

```go
if expr {
    n = trueVal
} else {
    n = falseVal
}
```

Go'da olmayan `?:`, dilin tasarımcılarının işlemin aşılmaz karmaşık ifadeler yaratmak için çok sık kullanıldığını görmeleridir. If-else formu, daha uzun olmasına rağmen, tartışmasız bir şekilde daha nettir. Bir dilin yalnızca bir koşullu kontrol akışı yapısına ihtiyacı vardır.

###  Paketler ve Testing

#### Nasıl çok dosyalı paket oluştururum?

Paketin tüm kaynak dosyalarını tek başlarına bir dizine koyun. Kaynak dosyalar isteğe bağlı olarak farklı dosyalardan öğelere başvurabilir; ileri tanımlamalara veya bir başlık dosyasına gerek yoktur. Birden çok dosyaya bölünmek dışında, paket tek dosyalık bir paket gibi derlenecek ve test edilecektir.

#### Nasıl birim testi yazarım?

Paket kaynaklarınızla aynı dizinde `_test.go` ile biten yeni bir dosya oluşturun. Bu dosyanın içine, `"testing"` i içe aktarın ve formun işlevlerini yazın

```go
func TestFoo(t *testing.T) {
    ...
}
```

Bu dizindeyken `go test` komutunu çalıştırın. Bu komut dosyası Test fonksiyonlarını bulur, bir test ikili dosyası oluşturur ve çalıştırır.

Daha fazla ayrıntı için [How to Write Go Code](https://golang.org/doc/code.html) belgesine, [testing](https://golang.org/pkg/testing/) paketine ve [go test](https://golang.org/cmd/go/#hdr-Test_packages) alt komutuna bakın.

#### Test için en sevdiğim yardımcı fonksiyon nerede?

Go'nun standart test paketi, birim testleri yazmayı kolaylaştırır, ancak onaylama işlevleri gibi diğer dillerin test çerçevelerinde sağlanan özelliklerden yoksundur. Bu belgenin daha önceki bir bölümü Go'nun neden iddialara sahip olmadığını ve aynı argümanlar testlerde assert kullanımı için de geçerli olduğunu açıkladı. Doğru hata işleme, biri başarısız olduktan sonra diğer testlerin çalıştırılmasına izin vermek anlamına gelir, böylece hatayı ayıklayan kişi neyin yanlış olduğunu tam olarak görebilir. İsPrime'ın 2, 3, 5 ve 7 \(veya 2, 4, 8 ve 16 için\) için yanlış cevap verdiğini rapor etmek, isPrime'ın 2 için yanlış cevap verdiğini ve bu nedenle hayır cevabını vermekten daha yararlıdır. daha fazla test yapıldı. Test hatasını tetikleyen programcı, başarısız olan koda aşina olmayabilir.

İyi bir hata mesajı yazmak için harcanan zaman artık daha sonra test sona erdiğinde karşılığını veriyor. Bununla ilgili bir nokta, test çerçevelerinin, koşullu ifadeler, kontroller ve baskı mekanizmaları ile kendi mini dillerine dönüşme eğiliminde olmasıdır. ancak Go tüm bu yeteneklere zaten sahiptir; neden onları yeniden yarattın? Go'da testler yazmayı tercih ederiz; öğrenilmesi gereken daha az dildir ve bu yaklaşım, testleri basit ve anlaşılması kolay tutar.

İyi hatalar yazmak için gereken ekstra kod miktarı tekrarlayıcı ve ezici görünüyorsa, test, bir veri yapısında tanımlanmış bir girdi ve çıktı listesi üzerinde yinelenen tabloya dayalıysa daha iyi çalışabilir \(Go, veri yapısı değişmezleri için mükemmel bir desteğe sahiptir\). İyi bir test ve iyi hata mesajları yazma işi daha sonra birçok test senaryosuna göre amortismana tabi tutulacaktır. Standart Go kitaplığı, fmt paketi için [biçimlendirme testleri gibi açıklayıcı örnekler](https://golang.org/src/fmt/fmt_test.go)le doludur.

#### X neden standart kütüphanede yok?

Standart kütüphanenin amacı, çalışma zamanını desteklemek, işletim sistemine bağlanmak ve biçimlendirilmiş G / Ç ve ağ iletişimi gibi birçok Go programının gerektirdiği temel işlevleri sağlamaktır. Ayrıca, kriptografi ve HTTP, JSON ve XML gibi standartlar için destek dahil olmak üzere web programlama için önemli öğeler içerir.

Neyin dahil edileceğini tanımlayan net bir kriter yok çünkü uzun zamandır bu tek Go kütüphanesi idi. Bununla birlikte, bugün neyin ekleneceğini tanımlayan kriterler var. Standart kitaplığa yeni eklemeler nadirdir ve dahil etme çıtası yüksektir. Standart kitaplığa dahil edilen kod, büyük bir sürekli bakım maliyeti taşır \(genellikle orijinal yazar dışındaki kişiler tarafından karşılanır\), [Go 1 uyumluluk taahhüdü](https://golang.org/doc/go1compat.html)ne \(API'deki herhangi bir kusur için engelleme düzeltmeleri\) tabidir ve Go sürümüne tabidir.

Program, hata düzeltmelerinin kullanıcılara hızlı bir şekilde sunulmasını önler. Yeni kodların çoğu standart kitaplığın dışında yaşamalı ve [go tool](https://golang.org/cmd/go/) ile erişilebilir olmalıdır. git komuta al. Bu tür kodların kendi bakımcıları, yayın döngüsü ve uyumluluk garantileri olabilir. Kullanıcılar [godoc.org](https://godoc.org/) adresinde paketleri bulabilir ve belgelerini okuyabilir. Standart kütüphanede `log/syslog` gibi gerçekten ait olmayan parçalar olsa da, Go 1 uyumluluk vaadi nedeniyle kitaplıktaki her şeyi korumaya devam ediyoruz. Ancak çoğu yeni kodu başka bir yerde yaşamaya teşvik ediyoruz.

###  İmplemantasyonlar

#### Derleyicileri oluşturmak için hangi derleyici teknolojisi kullanılır?

Go için birkaç üretim derleyicisi ve çeşitli platformlar için geliştirilmekte olan birkaç başka derleyici vardır. Varsayılan derleyici `gc`, go komutunun desteğinin bir parçası olarak Go dağıtımına dahildir. Gc, başlangıçta önyüklemenin zorlukları nedeniyle C dilinde yazılmıştır - bir Go ortamı kurmak için bir Go derleyicisine ihtiyacınız vardır. Ancak işler gelişti ve Go 1.5 sürümünden bu yana derleyici bir Go programı oldu. Derleyici, bu [tasarım belgesi](https://golang.org/s/go13compiler)nde ve konuşmada açıklandığı gibi otomatik çeviri araçları kullanılarak C'den Go'ya dönüştürüldü.

Bu nedenle, derleyici artık "kendi kendine barındırılıyor", bu da önyükleme sorunuyla yüzleşmemiz gerektiği anlamına geliyor. Çözüm, çalışan bir C kurulumunda olduğu gibi, çalışan bir Go kurulumuna sahip olmaktır. Kaynağından yeni bir Go ortamının nasıl ortaya çıkarılacağının hikayesi [burada](https://golang.org/s/go15bootstrap) ve [burada](https://golang.org/doc/install/source) anlatılmaktadır.

Gc, Go'da özyinelemeli bir iniş ayrıştırıcıyla yazılır ve yine Go'da yazılan ancak Plan 9 yükleyiciye dayanan özel bir yükleyici kullanır, ELF / Mach-O / PE ikili dosyaları oluşturmak için. Projenin başlangıcında gc için LLVM kullanmayı düşündük, ancak performans hedeflerimizi karşılamak için çok büyük ve yavaş olduğuna karar verdik. Geriye dönüp bakıldığında daha önemli olan, LLVM ile başlamak, Go'nun gerektirdiği ancak standart C kurulumunun bir parçası olmayan bazı ABI ve yığın yönetimi gibi ilgili değişiklikleri uygulamaya koymayı zorlaştırırdı. Ancak şimdi yeni bir [LLVM uygulaması](https://go.googlesource.com/gollvm/) bir araya gelmeye başlıyor.

`Gccgo` derleyicisi, standart GCC arka ucuna bağlı özyinelemeli bir iniş ayrıştırıcısı ile C ++ ile yazılmış bir ön uçtur. Go'nun bir Go derleyicisini uygulamak için iyi bir dil olduğu ortaya çıktı, ancak asıl amacı bu değildi. Başından beri kendi kendine barındırılmaması, Go'nun tasarımının ağa bağlı sunucular olan orijinal kullanım durumuna odaklanmasına izin verdi. Go'nun kendisini erken derlemesine karar vermiş olsaydık, derleyici yapımı için daha çok hedeflenmiş bir dil bulabilirdik, bu değerli bir hedef ama sahip olduğumuz değil başlangıçta.

Gc bunları kullanmasa da \(henüz\), Go paketinde yerel bir sözcük ve ayrıştırıcı mevcuttur ve ayrıca yerel bir [tür denetleyici](https://golang.org/pkg/go/types) de vardır.

#### Çalışma zamanı desteği nasıl uygulanır?

Yine, önyükleme sorunları nedeniyle, çalışma zamanı kodu başlangıçta çoğunlukla C ile yazılmıştır \(küçük bir montajlayıcı ile\), ancak o zamandan beri Go'ya çevrilmiştir \(bazı assembler bitleri hariç\). Gccgo'nun çalışma zamanı desteği `glibc` kullanır. Gccgo derleyicisi, altın bağlayıcıda yapılan son değişikliklerle desteklenen, segmentli yığınlar adı verilen bir teknik kullanarak gorutinleri uygular. `Gollvm` benzer şekilde ilgili `LLVM` altyapısı üzerine inşa edilmiştir.

#### Basit programım neden büyük dosya boyutlu?

Gc araç zincirindeki bağlayıcı, varsayılan olarak statik bağlantılı ikili dosyalar oluşturur. Bu nedenle tüm Go ikili dosyaları, dinamik tür kontrollerini, yansımayı ve hatta panik zamanı yığın izlerini desteklemek için gerekli çalışma zamanı tür bilgileriyle birlikte Go çalışma zamanını içerir.

Linux'ta `gcc` kullanılarak statik olarak derlenen ve bağlanan basit bir C "merhaba, dünya" programı, bir `printf` uygulaması da dahil olmak üzere yaklaşık `750 kB`'dir. `fmt.Printf` kullanan eşdeğer bir Go programı birkaç megabayt ağırlığındadır, ancak bu daha güçlü çalışma zamanı desteği ve tür ve hata ayıklama bilgileri içerir. Gc ile derlenen bir Go programı `-ldflags = -w` bayrağına bağlanarak `DWARF` oluşturmayı devre dışı bırakarak ikili dosyadan hata ayıklama bilgilerini kaldırabilir, ancak başka bir işlev kaybı olmaz. Bu, ikili boyutu önemli ölçüde azaltabilir.

#### Kullanılmayan değişken/içe aktarım ile ilgili uyarıları durdurabilir miyim?

Kullanılmayan bir değişkenin varlığı bir hatayı gösterebilirken, kullanılmayan içe aktarmalar derlemeyi yavaşlatır, bir program zamanla kod ve programcıları biriktirdikçe önemli hale gelebilir. Bu nedenlerle Go, kullanılmayan değişkenler veya içe aktarmalar içeren programları derlemeyi reddeder, uzun vadeli oluşturma hızı ve program netliği için kısa vadeli kolaylık ticaretini yapar.

Yine de, kod geliştirirken, bu durumları geçici olarak oluşturmak yaygındır ve program derlenmeden önce bunları düzenlemenin gerekmesi can sıkıcı olabilir. Bazıları, bu kontrolleri kapatmak veya en azından uyarılara indirgemek için bir derleyici seçeneği talep etti. Böyle bir seçenek eklenmemiştir, çünkü derleyici seçenekleri dilin anlamını etkilememelidir ve Go derleyicisi uyarıları değil, yalnızca derlemeyi engelleyen hataları bildirdiği için. Hiçbir uyarı almamanın iki nedeni vardır. İlk olarak, şikayet etmeye değerse, kodu düzeltmeye değer. \(Ve düzeltmeye değmiyorsa, buna değmez İkinci olarak, derleyicinin uyarılar oluşturması, uygulamayı, derlemeyi gürültülü hale getirebilecek zayıf durumlar hakkında uyarmaya ve düzeltilmesi gereken gerçek hataları maskelemeye teşvik eder. Yine de durumu ele almak kolaydır. Geliştirme sırasında kullanılmayan şeylerin devam etmesine izin vermek için boş tanımlayıcıyı kullanın.

```go
import "unused"

// This declaration marks the import as used by referencing an
// item from the package.
var _ = unused.Item  // TODO: Delete before committing!

func main() {
    debugData := debug.Profile()
    _ = debugData // Used only during debugging.
    ....
}
```

Günümüzde çoğu Go programcısı, [Goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) adlı bir araç kullanıyor, bu da Go kaynak dosyasını doğru içe aktarmaları elde etmek için otomatik olarak yeniden yazıyor ve uygulamada kullanılmayan içe aktarma sorununu ortadan kaldırıyor. Bu program, bir Go kaynak dosyası yazıldığında otomatik olarak çalışması için çoğu düzenleyiciye kolayca bağlanır.

#### Virüs tarama yazılımım neden Go dağıtımıma veya derlenmiş ikili dosyama virüs bulaştığını düşünüyor?

Bu, özellikle Windows makinelerde yaygın bir durumdur ve neredeyse her zaman yanlış bir pozitiftir. Ticari virüs tarama programları, diğer dillerden derlenenler kadar sık ​​görmedikleri Go ikili dosyalarının yapısı nedeniyle genellikle karıştırılır. Go dağıtımını yeni yüklediyseniz ve sistem virüs bulaştığını bildirirse, bu kesinlikle bir hatadır.

Gerçekten kapsamlı olmak için, sağlama toplamını [indirilenler sayfasındakil](https://golang.org/dl/)erle karşılaştırarak indirmeyi doğrulayabilirsiniz. Her durumda, raporun hatalı olduğuna inanıyorsanız, lütfen virüs tarayıcınızın tedarikçisine bir hata bildirin. Belki zamanla virüs tarayıcıları Go programlarını anlamayı öğrenebilir.

### Performans

#### Go, X karşılaştırmasında neden kötü performans gösteriyor?

Go'nun tasarım hedeflerinden biri, karşılaştırılabilir programlar için C'nin performansına yaklaşmaktır, ancak bazı kıyaslamalarda, [golang.org/x/exp/shootout](https://go.googlesource.com/exp/+/master/shootout/)'taki birkaçı da dahil olmak üzere, oldukça zayıf bir şekilde işliyor. En yavaş olanlar, karşılaştırılabilir performansın sürümlerinin Go'da bulunmadığı kütüphanlere bağlıdır. Örneğin, [pidigits.go](https://go.googlesource.com/exp/+/master/shootout/pidigits.go), çok duyarlıklı bir matematik paketine bağlıdır ve Go'nun aksine C sürümleri [GMP](https://gmplib.org/)'yi \(optimize edilmiş derleyicide yazılmıştır\) kullanır. Normal ifadelere \(örneğin [regex-dna.go](https://go.googlesource.com/exp/+/master/shootout/regex-dna.go)\) dayanan karşılaştırmalar, esasen Go'nun yerel [regexp paketi](https://golang.org/pkg/regexp)ni PCRE gibi olgun, yüksek düzeyde optimize edilmiş düzenli ifade kitaplıklarıyla karşılaştırıyor.

Kıyaslama oyunları kapsamlı ayarlamalarla kazanılır ve kıyaslamaların çoğunun Go sürümleri dikkat gerektirir. Karşılaştırılabilir C ve Go programlarını ölçerseniz \([reverse-complement.go](https://go.googlesource.com/exp/+/master/shootout/reverse-complement.go) bir örnektir\), iki dilin ham performansta bu paketin gösterdiğinden çok daha yakın olduğunu göreceksiniz.

Yine de, iyileştirme için yer var. Derleyiciler iyidir ancak daha iyi olabilir, birçok kütüphane büyük performans çalışmasına ihtiyaç duyar ve çöp toplayıcı henüz yeterince hızlı değildir. \(Öyle olsa bile, gereksiz çöp üretmemeye özen göstermenin çok büyük bir etkisi olabilir.\)

Her durumda, Go genellikle çok rekabetçi olabilir. Dil ve araçlar geliştikçe birçok programın performansında önemli gelişmeler olmuştur. Bilgilendirici bir örnek için [Go programlarını profilleme ](https://blog.golang.org/2011/06/profiling-go-programs.html)hakkındaki blog gönderisine bakın.

###  C'den Değişiklikler

#### Sözdizimi C'den neden bu kadar farklı?

Tanımlama sözdizimi dışında, farklılıklar büyük değildir ve iki arzudan kaynaklanır. İlk olarak, sözdizimi çok fazla zorunlu anahtar sözcük, tekrar veya gizem olmadan hafif hissetmelidir. İkinci olarak, dil analiz edilmesi kolay olacak şekilde tasarlanmıştır ve bir sembol tablosu olmadan ayrıştırılabilir. Bu, hata ayıklayıcılar, bağımlılık çözümleyicileri, otomatikleştirilmiş belge çıkarıcıları, IDE eklentileri vb. Gibi araçlar oluşturmayı çok daha kolaylaştırır. C ve onun soyundan gelenler bu bakımdan herkesin bildiği gibi zordur.

#### Tanımlamalar neden ters yönlü yapılır?

C'ye alışkınsanız, yalnızca geriye doğrudurlar. C'de, fikir bir değişkenin türünü ifade eden bir ifade gibi bildirilmesidir, bu güzel bir fikirdir, ancak tür ve ifade gramerleri çok iyi karışmaz ve sonuçlar kafa karıştırıcı olabilir; fonksiyon işaretlerini düşünün. Go, çoğunlukla ifade ve tür sözdizimini ayırır ve bu da işleri basitleştirir \(işaretçiler için \* öneki kullanmak, kuralı kanıtlayan bir istisnadır\). C'de tanımlamalar

```c
int* a, b;
```

Go'da ise

```go
var a, b *int
```

her ikisinin de işaretçi olduğunu beyan eder. Bu daha net ve daha düzenli. Ayrıca: = kısa tanımlama formu, tam değişken bildiriminin şu sırayı sunması gerektiğini savunur `: = değer`

```go
var a uint64 = 1
```

aşağıdaki ile aynı etkidedir.

```go
a := uint64(1)
```

Ayrıştırma, yalnızca ifade dilbilgisi olmayan türler için de farklı bir dilbilgisine sahip olarak basitleştirilmiştir; `func` ve `chan` gibi anahtar kelimeler her şeyi netleştirir. 

Daha fazla ayrıntı için [Go'nun Tanımlama Sözdizimi](https://golang.org/doc/articles/gos_declaration_syntax.html) hakkındaki makaleye bakın.

#### Neden işaretçi aritmetiği yok?

Emniyet. İşaretçi aritmetiği olmadan, hatalı bir şekilde başarılı olan geçersiz bir adresi asla türetemeyecek bir dil oluşturmak mümkündür. Derleyici ve donanım teknolojisi, dizi indekslerini kullanan bir döngünün işaretçi aritmetiği kullanan bir döngü kadar verimli olabileceği noktaya kadar gelişmiştir. Ayrıca, işaretçi aritmetiğinin olmaması, çöp toplayıcının uygulanmasını basitleştirebilir.

#### Neden ++ ve -- tanımlama değilde ifadedir? Ve neden ön ek değilde son ektir?

İşaretçi aritmetiği olmadan, ön ve sonek artırma operatörlerinin uygunluk değeri düşer. Bunları ifade hiyerarşisinden tamamen kaldırarak, ifade sözdizimi basitleştirilir ve ++ ve - \(f \(i ++\) ve p \[i\] = q \[++ i\] 'yi düşünün\) değerlendirme sırasına ilişkin karışık sorunlar da ortadan kaldırılır. . Sadeleştirme önemlidir. sonek ve öneke gelince, her ikisi de iyi çalışır, ancak son ek sürümü daha gelenekseldir; Önek ısrarı, adı ironik bir şekilde bir sonek artışını içeren bir dil kitaplığı olan STL ile ortaya çıktı.

#### Neden parantez var ama noktalı virgül yok? Ve neden sonraki satırda açılış ayracı koyamazsınız?

Go, C ailesindeki herhangi bir dille çalışan programcıların aşina olduğu bir sözdizimi olan ifade gruplaması için ayraç parantezleri kullanır. Ancak noktalı virgüller insanlar için değil, ayrıştırıcılar içindir ve biz onları olabildiğince ortadan kaldırmak istedik. Go, bu amaca ulaşmak için BCPL'den bir numara ödünç alır: İfadeleri ayıran noktalı virgüller biçimsel dilbilgisi içindedir, ancak bir ifadenin sonu olabilecek herhangi bir satırın sonuna sözcük yazar tarafından önden bakmadan otomatik olarak enjekte edilir.

Bu, pratikte çok iyi çalışır, ancak bir küme ayracı stilini zorlama etkisine sahiptir. Örneğin, bir işlevin açılış ayracı tek başına bir satırda görünemez. Bazıları, korsenin bir sonraki satırda yaşamasına izin vermek için lexer'in ileriye bakması gerektiğini savundu. Biz anlaşamadık. Go kodunun [gofmt](https://golang.org/cmd/gofmt/) tarafından otomatik olarak biçimlendirilmesi amaçlandığından, bazı stillerin seçilmesi gerekir. Bu stil, C veya Java'da kullandığınızdan farklı olabilir, ancak Go farklı bir dildir ve gofmt'nin stili diğerleri kadar iyidir.

Daha önemli - çok daha da önemlisi - tüm Go programları için tek, programatik olarak zorunlu bir formatın avantajları, belirli bir stilin algılanan dezavantajlarından büyük ölçüde ağır basar. Go'nun stilinin, Go'nun etkileşimli bir uygulamasının standart sözdizimini özel kurallar olmaksızın her seferinde bir satır olarak kullanabileceği anlamına geldiğini unutmayın.

#### Neden çöp toplama yapılıyor? Çok masraflı olmayacak mı?

Sistem programlarındaki en büyük muhasebe kaynaklarından biri, tahsis edilen nesnelerin yaşam sürelerini yönetmektir. Manuel olarak yapıldığı C gibi dillerde, önemli miktarda programcı zamanı tüketebilir ve genellikle zararlı hataların sebebidir. Yardımcı olmak için mekanizmalar sağlayan C ++ veya Rust gibi dillerde bile, bu mekanizmalar yazılımın tasarımı üzerinde önemli bir etkiye sahip olabilir ve genellikle kendi başına programlama ek yükü ekleyebilir. Bu tür programcı genel giderlerini ortadan kaldırmanın kritik olduğunu düşündük ve son birkaç yıldaki çöp toplama teknolojisindeki ilerlemeler, ağa bağlı sistemler için uygun bir yaklaşım olabileceğine dair yeterince ucuz ve yeterince düşük gecikme süresiyle uygulanabileceğine dair bize güven verdi.

Eşzamanlı programlamanın zorluğunun çoğunun kökleri nesne yaşam süresi problemine dayanır: nesneler iş parçacıkları arasında geçerken, güvenli bir şekilde serbest kalmalarını garanti etmek zahmetli hale gelir. Otomatik çöp toplama, eşzamanlı kodu çok daha kolay hale getirir yazmak. Elbette, eşzamanlı bir ortamda çöp toplamayı uygulamak başlı başına bir zorluktur, ancak bunu her programda değil bir kez karşılamak herkese yardımcı olur. Son olarak, eşzamanlılık bir yana, çöp toplama arabirimleri daha basit hale getirir çünkü bunlar arasında belleğin nasıl yönetildiğini belirtmelerine gerek yoktur. Bu, Rust gibi dillerde kaynakları yönetme sorununa yeni fikirler getiren son çalışmaların yanlış yönlendirildiği anlamına gelmez; bu çalışmayı teşvik ediyoruz ve nasıl geliştiğini görmekten heyecan duyuyoruz.

Ancak Go, yalnızca çöp toplama ve çöp toplama yoluyla nesne yaşam sürelerini ele alarak daha geleneksel bir yaklaşım benimsiyor. Mevcut uygulama, bir işaret ve süpür toplayıcıdır. Makine çok işlemcili ise, toplayıcı ana programla paralel olarak ayrı bir CPU çekirdeği üzerinde çalışır. Son yıllarda toplayıcı üzerinde yapılan büyük çalışma, büyük yığınlar için bile duraklama sürelerini sık sık milisaniyenin altındaki aralığa indirdi ve bu da çöplere yönelik başlıca itirazlardan birini ortadan kaldırıyor ağa bağlı sunucularda toplama. Algoritmayı iyileştirmeye, ek yükü ve gecikmeyi daha da azaltmaya ve yeni yaklaşımları keşfetmeye yönelik çalışmalar devam ediyor. Go ekibinden Rick Hudson'ın 2018 [ISMM açılış konuşması](https://blog.golang.org/ismmkeynote) şimdiye kadarki ilerlemeyi anlatıyor ve bazı gelecekteki yaklaşımlar öneriyor.

Performans konusuna gelince, Go'nun programcıya bellek düzeni ve tahsisi üzerinde, çöp toplama dillerinde tipik olandan çok daha fazla, hatırı sayılır bir kontrol sağladığını unutmayın. Dikkatli bir programcı, dili iyi kullanarak çöp toplama ek yükünü önemli ölçüde azaltabilir; Go'nun profil oluşturma araçlarının bir gösterimi de dahil olmak üzere, çalışılmış bir örnek için [Go programlarının profilini çıkarma](https://blog.golang.org/2011/06/profiling-go-programs.html) hakkındaki makaleye bakın.

## Go Derleyicisi Kurulumu

Öncelikle bilgisayarımız üzerinde nasıl Golang geliştireceğimize bakalım. Geliştirmek derken Golang programı oluşturacağımızı kastediyorum. Öncelikle Golang’ın resmi sitesinden Golang programını \(derleyici\) indiriyoruz.

Buradan İndirebilirsiniz

[https://golang.org/dl/](https://golang.org/dl/)

Golang’ın basit bir kurulumu var o yüzden kurulumu atlıyorum.

Linux İS kullananlara tavsiyem, kullandığınız dağıtımın uygulama deposundan Golang’ı indirin. Sizin için daha kolay olur.

## VSCode Go Eklentisi Yükleme

Golang’ı indirdiğimize göre bize Golang kodlarımızı yazacağımız bir Tümleşik Geliştirme Ortamı \(IDE\) lazım. IDE’ler kodlarımızı yazarken kodların doğruluğunu kontrol eder ve kod yazarken önerilerde bulunur. Bu da kod yazarken işimizi kolaylaştırır.

Benim tavsiyem çoğu kodlama dilini yazarken kullandığım ve Golang yazanların da popüler olarak kullandığı **Visual Studio Code** programı.

Buradan İndirebilirsiniz

[https://code.visualstudio.com/Download](https://code.visualstudio.com/Download)

Linux İS kullananlara yine kullandıkları dağıtımın uygulama deposundan indirmelerini tavsiye ediyorum.

Visual Studio Code’dan ilerki zamanlarda **vscode** olarak bahsedeceğim.

Go eklentisinin düzgün bir şekilde kurulabilmesi için bilgisayarımızda **git** komut-satırı uygulaması bulunması gerekir. Çünkü eklentinin yüklenmesinden sonra Go eklentisi VSCode için 15 civarı aracı otomatik indirecek. Git’in yüklü olup olmadığını öğrenmek için komut satırına aşağıdakileri yazın:

Eğer versiyon numarasını gördüyseniz yüklü demektir. Eğer yüklü değilse veya git’i güncellemek istiyorsanız ki, mutlaka öneririm, aşağıda nasıl yükleneceğini görebilirsiniz.

`git --version`

Windows MacOS

Windows için: [Buradan İndirebilirsiniz](https://git-scm.com/download/win)

MacOS: [Buradan İndirebilirsiniz.](https://git-scm.com/download/mac)

**GNU/Linux İşletim Sistemleri**

**Debian/Ubuntu**

`sudo apt-get install git`

**Fedora**

`dnf install git`

**Arch Linux**

`sudo pacman -S git`

**Gentoo**

`emerge --ask --verbose dev-vcs/git`

**openSUSE**

`zypper install git`

**Mageia**

`urpmi git`

Git kurulumunu da yaptığımıza göre VSCode için Go eklentisini kurabiliriz.

![VS Code Go Eklentisi](./goruntuler/deepinscreenshot_select-area_20191006213837.png)

Vscode’un sol tarafından **Extension** \(Eklentiler\) sekmesine geçiyoruz. Arama kutusuna **go** yazıyoruz. Resimde de seçili olan Go eklentisini yeşil **Install \(Yükle\)** butonuna basarak yüklüyoruz. Yeniden başlatma isterse başlatmayı unutmayın.

Daha sonra **main.go** adında bir dosya oluşturup VSCode ile açalım. main.go dosyasının içerisine rastgele birşeyler yazdığımızda VSCode sağ alt tarafta bize uyarı verecektir.

![VS Code Go Uyar&#x131;s&#x131;](./goruntuler/golang-plugins-all.png)

**Install All** diyerek Go eklentisi araçlarının kurulumunu başlatalım.

Kurulum bize 15 civarı araç kuracak. Başarı ile kurulan araçların yanında **SUCCEED** ibaresi yer alır.

Tüm araçlarımız başarıyla kurulunca artık VSCode üzerinden Go geliştirmeye hazır olacaksınız.

## Merhaba Dünya

Programlama dünyasında gelenektir, bir programlama dili öğrenilirken ilk önce ekrana **“Merhaba Dünya”** çıktısı veren bir program yazılır. Biz de geleneği bozmadan Golang üzerinde Merhaba Dünya uygulaması yazalım. İlk önce kodları görelim. Daha sonra açıklamasını görelim.

```go
package main

import “fmt”

func main(){
    fmt.Println(“Merhaba Dünya!”)
}
```

Şimdi yukarıdaki kodlarda neler yaptığımıza gelelim.

**Package**, kod sayfalarımız arasında iletişimde bulunabilmemizi sağlar. Bu sayede içerisinde **package** değeri aynı olan kod sayfaları birbirleriyle iletişim halinde olma yeteneği kazanır. Yukarıdaki örnekte package uygulama olan sayfalar birbiriyle iletişim kurabilir.

**import “fmt”** ile Golang dilinde kullanılan ana işlemler için olan kütüphanemizi içeri aktardık.

**func main\(\)** ise programımızın çalışacağı ana bölümün fonksiyonudur. Yanındaki süslü parantezler { } içine yazdığımız kodlar ile programımızda çeşitli işlemler yapabileceğiz. Derlenmiş bir uygulama ilk olarak **main** fonksiyonuna bakar ve buradaki kodları çalıştırır.

Yukarıda **import** ettiğimiz **fmt** kütüphanesi içinden **Println** fonksiyonu ile ekranımıza **“Merhaba Dünya”** yazısını bastırdık. Gelelim programımızın derlenmesine. Daha önceden programlama dilleriyle geçmişi olmayan arkadaşlarımız için derlenme şöyle anlatılabilir. Yazdığımız Golang dili insanların kolaylıkla programlama yapabilmesi için oluşturulmuş bir dildir.

Ama makine \(bilgisayar\) bu yazdıklarımızı anlamaz. O yüzden her derlenen dilde olduğu gibi Golang’ın da yazdıklarımızı makinenin anlayacağı makine diline çeviren derleyicisi vardır.

Makinemiz üzerinde çalıştırılabilir bir dosya üretmek için kodlarımızın derlenmesi gereklidir. Vscode üzerinden kodlarımızın derlenip çalışması için **F5** tuşuna basıyoruz. Böylece programımızı test edebiliriz. Eğer vscode üzerinden derliyorsanız yazdığımız kodların hemen altında **DEBUG CONSOLE** bölümünde kodumuzun sonuç çıktısını görebiliriz.

Çıktımızı inceleyecek olursak, **API server listening at :127.0.0.1:44830** ibaresinde gerçekleşen olay, Golang kodlarımızı çalıştırdığımızda oluşturulan **44830** portlu **127.0.0.1** yerel sunucusu \(localhost\) üzerinden kodlarımızın sürüş testini gerçekleştirdik. Hemen aşağısına da çıktımızı verdi.

Eğer vscode üzerinden değil de, konsol üzerinden yapmak isterseniz, oluşturmuş olduğumuz main.go dosyasının bulunduğu dizin \(klasör\) içerisinde konsol uygulamamızı açıyoruz. Windows’ta cmd veya Powershell’dir. Unix sistemlerde terminal diye geçer. İki tarafa da yazacağımız komutlar aynıdır. O yüzden hangisinden yazdığınız farketmez.

Kodlarımızı sadece denemek istiyorsak yazacağımız komut::

`go run main.go //main.go yerine .go dosyamızın ismi gelecek`

Eğer çalıştırılabilir bir dosya oluşturmak istiyorsak \(Windows’ta .exe\)

`go build main.go //main.go yerine .go dosyamızın ismi gelecek`

Böylece bu işlemleri yaptığımız dizin içerisine çalıştırılabilir bir dosya oluşturmuş olduk.

Windows üzerinden konsola **main** yazarak uygulamayı açabilir veya klasörde **main.exe** dosyasına çift tıklayarak uygulamayı çalıştırabilirsiniz.

Linux üzerinden ise terminale derlenmiş main çıktınızın bulunduğu dizinde **./main** yazarak çalıştırabilirsiniz.

Böylece ilk uygulamamızı yazmış olduk. Tabi şu ana kadar görmemiş olduğumuz kodlar gördük. Onların açıklamaları da ileriki bölümlerde olacak. Şimdilik gelenek diye ilk bölümde **Merhaba Dünya** uygulaması yazdık.

## VSCode Varsayılan Hata Ayıklayıcıyı Seçme

VSCode Go programlama yapıyorken, **F5** tuşuna bastığınızda üst tarafta hata ayıklayıcıyı seçmenizi ister. Her **F5** tuşuna bastığınızda hata ayıklayıcı seçim ekranı çıksın istemiyorsanız, yani herzaman bir hata ayıklayıcıyı kullansız istiyorsanız, yapacaklarınız çok basit.

VSCode üzerinde ekranın sol tarafından Run sekmesine geçelim.

![Ad&#x131;m.1](./goruntuler/image%20%281%29.png)

Biz varsayılan olarak **Go Hata Ayıklayıcısı**'nı seçeceğiz.

**Go**'yu varsayılan hale getirmek içinse, ekranın solundan `create a launch.json file` bağlantısına tıklayalım.

![Ad&#x131;m.2](./goruntuler/defaultdebugger2.png)

Üst tarafta açılan ekrandan **Go**'yu seçelim.

![Ad&#x131;m.3](./goruntuler/defaultdebugger1.png)

Seçtikten sonra VSCode bizim için `launch.json` adında bir dosya oluşturacak. Bu dosya **F5** tuşuna bastığımızda gerçekleşecek olayları barındıryor. Dikkat edeceğimiz nokta type bölümünde go yazıyor olması.

![Ad&#x131;m.4](./goruntuler/defaultdebugger3.png)

Daha sonra `launch.json` dosyamızı kaydedip kapatabiliriz.

Bir sonra hata ayıklama işleminde **Go** otomatik çalışacaktır.

## Farklı Platformlara Build \(İnşa\) Etme

Golang projemizi build \(inşa\) ederken, yani çalıştırılabilir bir dosya üretirken **go build dosya.go** şeklinde build ederiz. Bu işlem varsayılan olarak kullanmakta olduğumuz işletim sistemi için build işlemi yapar. Yani Windows kullanıyorsak Windows’ta çalışmak üzere Linux İS kullanıyorsak Linux İS’te çalışmak üzere dosya oluşturur. Aynı şekilde sistemimizin mimarisi 32bit ise 32bit için, 64bit ise 64bit için çalıştırılabilir dosya üretir. Örnek olarak sistemimiz Windows 64bit ise oluşturduğumuz çalıştırılabilir dosya \(exe dosyası\) sadece Windows 64bitlerde çalışır.

Eğer farklı işletim sistemi için bir çalıştırılabilir dosya üretmek istiyorsak aşağıdaki komutu kullanmamız gerekir.

`env GOOS=hedef-sistem GOARCH=hedef-mimari go build hedef-dosya`

Örnek olarak, **main.go** dosyamızı **Windows 32bit** sistemler için build etmek istersek aşağıdaki komutları girmemiz gerekir.

`env GOOS=windows GOARCH=386 go build main.go`

Bu işlem ile main.exe adında bir dosya elde ederiz. Bu dosya Windows 32bit sistemlerde çalışabilir. Biliyorsunuz ki 32bit uygulamalar 64bit sistemlerde çalışır ;fakat 64bit uygulamalar 32bit sistemlerde çalışmaz. Onun için bir uygulama build ediyorken bunu aklınızdan çıkarmayın. Golang’ın dosya build edebildiği işletim sistemi ve mimari sayısı oldukça fazladır.

Golang’ın build edebildiği tüm işletim sistemi ve mimarilerine bakmak gerekir ise;

| GOOS | GOARCH |
| :--- | :--- |
| android | arm |
| darwin | 386 |
| darwin | amd64 |
| darwin | arm |
| darwin | arm64 |
| dragonfly | amd64 |
| freebsd | 386 |
| freebsd | amd64 |
| freebsd | arm |
| linux | 386 |
| linux | amd64 |
| linux | arm |
| linux | arm64 |
| linux | ppc64 |
| linux | ppc64le |
| linux | mips |
| linux | mipsle |
| linux | mips64 |
| linux | mips64le |
| netbsd | 386 |
| netbsd | arm64 |
| netbsd | arm |
| openbsd | 386 |
| openbsd | arm64 |
| openbsd | arm |
| plan9 | 386 |
| plan9 | amd64 |
| solaris | amd64 |
| windows | 386 |
| windows | amd64 |

\* Tablo harf sıralamasına göre yapılmıştır.

Birkaç örnek daha yapmak gerekirse;

**Windows 64bit Build**

`env GOOS=windows GOARCH=amd64 go build main.go`

**Linux 32bit Build**

`env GOOS=linux GOARCH=386 go build main.go`

**MacOS 64bit Build**

`env GOOS=darwin GOARCH=amd64 go build main.go`

## Klasör Build Etme

Oluşturduğumuz **.go** dosyaları birden fazla parçadan oluşuyorsa aşağıdaki komut ile klasör içerisindeyken build işlemini yapabiliriz.

`go build . //sonda nokta işareti var`

Eğer klasörün dışından build işlemi yapacaksak nokta yerine klasörün yolunu girmemiz gerekir. Aşağıda gösterilmiştir.

Build işleminde sadece **.go** dosyaları derlenir. Tüm dosyalar işlenip tek başına çalıştırılabilir dosyaya dönüştürülür. Yani yanındaki Html, Css, Js vs. türünde dosyalar paket içine alınmaz. Tümüyle dosyları paketlemek için ek kütüphaneler kullanabilirsiniz. Önerim **Statik** isimli kütüphanedir. İlerideki bölümlerde zaten bu kütüphaneyi kullanıyor olacağız.

`go build /klasör/yolunu/buraya/girin`

## Paketler

Her Golang programı paketlerden oluşur ve kendisi de bir pakettir.

```go
package uygulama

import "fmt"

func main() {
 fmt.Println("Merhaba Dünya") // Çıktımız
}
```

**package** terimi ile programımızın paket adını belirleriz. Hemen aşağısında da **import “fmt”** ile **fmt** paketini çektiğimizi görebilirsiniz. Yani çektiğimiz fmt paketi de bir programdır. Bilin bakalım fmt paketinin ilk satırında ne yazıyor? Tabiki de **package fmt**. Yani package ile programımızın ismini tanımlıyoruz. Bu ismi kullanarak diğer paketler ile iletişimde bulunabiliriz.

**import** terimi ise yazıldığı pakete başka bir paketten bir yetenek aktarmaya yarar. Yetenekten kastım, import edilen paketin içinde fonksiyonlar mı var? Struct’lar mı var? vs. onları içeri aktarır.

```go
import (
 “fmt”
 ”math”
)
```

Yukarıda birden fazla paket import etmeyi görüyoruz. **math** paketi bize ileri matematiksel işlemler yapabilmemiz için gerekli fonksiyonları sağlar.

## Yorum Satırı

Diğer dillerde de olduğu gibi Golang’ta yorum satırı özelliği mevcuttur. Yorum satırı derleyici tarafından işlenmez. Yani görmezden gelinir. Bu bölüme kendiniz için açıklama vs. bilgiler yazabilirsiniz. Golang’ta yorum satırı oluşturmak için 2 yöntem mevcuttur.

**// Çift Taksim Yöntemi**

Bu yöntem ile derlenmesini istemediğimiz yazının başına çift taksim ekleyerek görmezden gelinmesini sağlıyoruz.

```go
//Buraya herhangi birşey yazabilirsiniz
```

**/\* \*/ Taksim-Yıldız Yöntemi**

Bu yöntem ile birden fazla satırın derlemede görmezden gelinmesini sağlayabiliriz.

```go
/* Buraya
herhangi
birşey
yazabilirsiniz */
```

## Veri Tipleri

**Integer Türler**

Öncelikle tüm integer türleri bir görelim;

int, int8, int16, int32, int64

uint, uint8, uint16, uint32, uint64, uintptr

Bu veri tipleri içerisinde sayısal değerleri depolayabiliriz. Fakat şunu da unutmamalıyız. Her sayısal veri tipinin depolayabildiği maksimum bit vardır. Örnek olarak uint8 veri tipinin maksimum uzunluğu 8 bit’tir. Bitler 0 ve 1 sayılarından oluşur. 8 bit demek 8 haneli 1 ve 0 sayısı demektir. Int8 maksimum alabileceği sayı derken 11111111 \(8 tane 1\), yani onluk sistemde 255 sayısına denk gelir. int 8 ise pozitif olarak +127, negatif olarak -128 maksimum değerinin alabilir. \(127+128=255\). int16 +32767 ve -32768 maksimum değerlerini alır. Int32 +2147483647 ve -2147483648 maksimum değerlerini alır. Int64 +9223372036854775807 ve -9223372036854775808 maksimum değerini alır.

U harfi ile başlayan sayı veritiplerinde ise sayının değeri pozitif veya negatif işarette değildir. Sadece bir sayısal değerdir. U’nun anlamı unassigned yani işaretsizdir. Uint8 0-255 arası, uint16 0-65535, uint32 0-42967295 arası, uint64 0-18446744073709551615 arası değerler alabilir. Uintptr ise yazdığınız sayıya göre alanı belirlenir.

Integer sayısal veri tipleri içierisinden bahsedebileceğimiz son tipler ise int ve uint. Int ve uint veri tipleri kullanmış olduğumuz işletim sistemi 32bit ise 32bit değer alırlar, 64bit ise 64bit değer alırlar. Sayısal bir değer atanacağı zaman en çok kullanılan veri tipleridir. Genellikle int daha çok kullanılır. Eğer çok meşakkatli bir program yazmayacaksanız int kullanmanız önerilir.

**Byte Veri Tipi:** uint8 ile aynıdır.

**Rune:** int32 ile aynıdır. Unicode karakter kodlarını ifade eder.

**Float Türler**

Float türleri integer türlerden farklı olarak küsüratlı sayıları tutar. Örnek: 3.14

**Lütfen Dikkat!**

Küsüratlı sayılar İngiliz-Amerikan sayı sistemine göre nokta koyarak ifade edilir. Türk sistemindeki gibi virgül \(3,14\) ile ifade edilmez.

**float32:** 32bitlik değer alabilir.

**float64:** 64 değer alabilir.

**Complex Türler**

Complex türleri içerisinde gerçel küsüratlı \(float\) ve sanal sayılar barındırabilir. Türkçe’de karmaşık sayılar diye adlandırılır.

**complex64:** Gerçel float32 ve sanal sayı değeri barındırır.

**complex128:** Gerçel float64 ve sanal sayı değeri barındırır.

Sayısal türler bu şekildedir.

**BOOLEAN VERİ TİPİ**

Boolean yani mantıksal veri tipi bir durumun var olması halinde olumlu \(true\) değer, var olmaması halinde olumsuz \(false\) değer alan veri tipidir.

**STRING VERİ TİPİ**

String yani dizgi veri tipi içerisinde metinsel ifadeler barındırır. Örnek olarak “Golang çok güzel ama ingilicce”. String veri tipi değeri çift tırnak \( “Değer” \) içine yazılır. Diğer dillerdeki gibi tek tırnak \( ‘Değer’ \) insiyatifi yoktur. Tek tırnakla kullanım başka bir amaç içindir. İlerde onu da göstereceğim.

**Özet olarak Veri Tipleri**

Veri tipleri atanacak değerlerimizi RAM üzerinde depolamak için kullandığımız araçlardır. Tam sayı değerler için Integer veri tiplerini, ondalık sayılar için Float veri tiplerini, mantıksal değerler için Boolean veri tipini, metinsel değerler için String veri tipini kullanırız. Karmaşık sayı değerleri için ise Complex veri tipini kullanırız.

”Türkiye” = String Tipi

1881 = Integer Tipi

10,5 = Float Tipi

True = Boolean Tipi

2+3i = Complex Tipi

**Veri Tiplerinin Varsayılan Değerleri**

Veri tipleri içerisine değer atanmadan oluşturulduğu zaman varsayılan bir değer alır.

Sayısal Tipler için 0,

Boolean Tipi için false,

String Tipi için “” \(Boş dizgi\) değeri alır.

## Aritmetik Operatörler

Aritmetik operatörler programlamada matematiksel işlemler yapabilmemize olanak sağlar.

| Operatör | Açıklama | Örnek |
| :--- | :--- | :--- |
| + | Toplar | 2+5 |
| - | Çıkarır | 10-3 |
| \* | Çarpar | 3\*4 |
| / | Böler | 10/2 |
| % | Bölümden kalanı verir \(Mod\) | 10%3 |
| ++ | 1 arttırır | 1++ |
| -- | 1 eksiltir | 1-- |

## İlişkisel Operatörler

İlişkisel operatörler programlamada iki veriyi birbiriyle karşılaştırabilmemize olanak sağlar. Karşılaştırma doğrulanıyorsa **true** değer, doğrulanmıyorsa **false** değer alır.

<table>
  <thead>
    <tr>
      <th style="text-align:left">Operat&#xF6;r</th>
      <th style="text-align:left">A&#xE7;&#x131;klama</th>
      <th style="text-align:left">&#xD6;rnek</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">==</td>
      <td style="text-align:left">&#x130;ki verinin e&#x15F;itli&#x11F;i</td>
      <td style="text-align:left">2==2</td>
    </tr>
    <tr>
      <td style="text-align:left">!=</td>
      <td style="text-align:left">&#x130;ki verinin e&#x15F;itsizli&#x11F;i</td>
      <td style="text-align:left">2!=3</td>
    </tr>
    <tr>
      <td style="text-align:left">&gt;</td>
      <td style="text-align:left">
        <ol>
          <li>verinin 2. veriden b&#xFC;y&#xFC;kl&#xFC;&#x11F;&#xFC;</li>
        </ol>
      </td>
      <td style="text-align:left">5&gt;3</td>
    </tr>
    <tr>
      <td style="text-align:left">&lt;</td>
      <td style="text-align:left">
        <ol>
          <li>verinin 2. veriden k&#xFC;&#xE7;&#xFC;kl&#xFC;&#x11F;&#xFC;</li>
        </ol>
      </td>
      <td style="text-align:left">4&lt;6</td>
    </tr>
    <tr>
      <td style="text-align:left">&gt;=</td>
      <td style="text-align:left">
        <ol>
          <li>verinin 2. veriden b&#xFC;y&#xFC;k veya e&#x15F;itli&#x11F;i</li>
        </ol>
      </td>
      <td style="text-align:left">4&gt;=6</td>
    </tr>
    <tr>
      <td style="text-align:left">&lt;=</td>
      <td style="text-align:left">
        <ol>
          <li>verinin 2. veriden k&#xFC;&#xE7;&#xFC;k veya e&#x15F;itli&#x11F;i</li>
        </ol>
      </td>
      <td style="text-align:left">3&lt;=8</td>
    </tr>
  </tbody>
</table>

## Mantıksal Operatörler

Mantıksal operatörler birden fazla mantıksal veriyi kontrol eder ve kullandığımız operatöre göre mantıksal değer döndürür.

| Operatör | Açıklama | Örnek |
| :--- | :--- | :--- |
| && | VE operatörüdür. 2 koşulda doğruysa true değer verir | 2==2 && 2==3 \(false\) |
| \|\| | VEYA operatörüdür. 2 koşuldan ikisi veya biri doğru ise true değer verir | 2==2 \|\| 2==3 \(true\) |
| ! | DEĞİL operatörüdür. Koşul sonucunun tersini verir | !\(2==2\) \(false\) |

## Atama Operatörleri

Atama operatörleri değişkenlere ve sabitlere değer atamak için kullanılır. Aşağıdaki tabloda c’nin değeri 10’dur. \(c=10\)

| Operatör | Açıklama | Örnek |
| :--- | :--- | :--- |
| = | Atama Operatörüdür | c=2 |
| += | Kendiyle toplar | c+=2 |
| -= | Kendinden çıkarır | c-=2 |
| \*= | Kendiyle çarpar | c\*=2 |
| /= | Kendine böler | c/=2 |
| %= | Kendine bölümünden kalanı atar | c%=3 |

## Değişkenler ve Atanması

**Değişkenler** içerisinde değer barındırarak RAM’e kaydettiğimiz bilgilerdir. Değişkenler programımızın işleyişinde önemli bir role sahiptir. Değişkenleri şu şekillerde atayabiliriz. Değişkenler **var** ifadesi ile atanır. Tabi ki zorunlu değildir.

```go
var isim string = “Ali”
var yas int = 20
var ogrenci bool = true
```

Yukarıdaki yazdıklarımızı inceleyecek olursak;

**var** ile değişken atadığımızı belirtiyoruz. **isim** diye bir değişken adı atadık ve içine **“Ali**” değerinde **string** tipinde bir değer yerleştirdik. String tipi değerler çift tırnak içine yazılır.

Aynı şekilde **yas** adında değişken oluşturduk. **yas** değişkeni içerisine **int** tipinde **20** değerini yerleştirdik Son olarak **ogrenci** adında bir değişken oluşturduk ve **true** değerinde **boolean** tipinde bir atama yaptık.

Golang’ta değişken adı oluştururken Türkçe karakterler kullanabiliriz. Örnek olarak **ogrenci** yerine öğrenci yazabilirdik. Ama başka bir programlama diline geçtiğinizde Türkçe harf desteklememesi halinde alışkanlıklarınızı değiştirmeniz gerekecek. O yüzden Türkçe karakter kullanmamanızı tavsiye ederim. Tabi ki zorunlu değil.

Programlama dillerinde, matematiğin aksine **= \(eşittir\)** işareti eşitlik için değil, atamalar için kullanılır.

Değişkenlerin atanması için farklı yöntemler de var. Diğer yöntemlere değinmek gerekirse;

Değişken atamasında illaki değişkenin veri tipini belirtmemiz gerekmez. Yazdığımız değere göre Golang otomatik olarak veri tipini algılar.

```go
var isim = “Ali”
var yas = 20
var ogrenci = true
```

**isim** değişkeninin değerini çift tırnak arasına yazdığımız için **string** veri tipinde olduğunu algılayacaktır.

**yas** değişkeninin değerini sayı olarak girdiğimiz için **int** tipinde olduğunu algılar. Eğer 20 değil de 2.12312 gibi küsüratlı bir değer girseydik veri tipini **float** olarak algılardı.

**ogrenci** değişkeninin değerini mantıksal olarak girdiğimiz için **boolean** veri tipinde olduğunu algılayacaktır.

En basit şekilde değişken ataması yapmak istersek;

```go
isim:=”Ali”
yas:=20
ogrenci:=true
```

Başına **var** eklemeden de değişken atamak mümkündür. Bu şekilde yapmak için **:=** işaretlerini kullanırız. Aynı şekilde bu yöntemde de verinin tipi otomatik algılanır.

Eğer değişken tanımlar iken değer kısmını boş bırakırsak yani; **var yas int** şeklinde yazarsak, önceki konuda da bahsettiğimiz gibi varsayılan olarak **0** değerini alır.

## Sabitler

Sabitler de değişkenler gibi değer alır. Fakat adından da anlaşılabileceği üzere verilen değer daha sonradan değiştirilemez.

Sabitler tanımlanırken başına const eklenir. Örnek olarak;

```go
const isim string = “Ali”
const isim=”Veli”
```

danger
**const** ile **:=** beraber kullanılamaz.

Yanlış kullanım: **const isim := “Ali”**

Doğru kullanım: **const isim = “Ali”**


Örnek olarak bir sabitin değerini atandıktan sonra değiştirmeye çalışalım. Aramızda ne olacağını merak eden çılgınlar olabilir.

Bu şekilde yazıp kodlarımızı derlediğimizde hata almamız kaçınamaz. Derlediğimizde **cannot assign to isim** hatasını verecektir. Yani diyor ki **isim’e atanamıyor**.

```go
const isim string = “Ali”
isim = “Ali”
const yas = 20
```

## Kod Gruplama İşlemi

Kod gruplama işlemi çok basit bir işlemdir. Bu işlem sayesinde aynı objeler bloklara göre farklı çalışabilir. Kodları gruplama için süslü parantez kullanırız. Örneğimizi görelim.

```go
package main

import "fmt"

func main() {
    değişken := "bir"
    {
        değişken := "iki"
        fmt.Println(değişken)
    }
    fmt.Println(değişken)
}
```

Çıktımızı gördükten sonra kodları açıklayayım.

> iki
>
> bir

Yukarıda **değişken** isminde değişken oluşturduk. Hemen aşağısına süslü parantez oluşturduk. İçine yine değişken adında bir değişken tanımladık. Bu iki değişken aynı kod bloğunda bulunmadığı için birbirleri ile alakası olmayacaktır. Aslında ikisi de aynı değişkendir. Sadece içindeki bloğa göre farklı bir değeri vardır. Bunu anlamanın en basit yolu pointer ile bellek adresine bakmaktır. Şimdi de örneğimizin o versiyonunu görelim.

```go
package main

import "fmt"

func main() {
    değişken := "bir"
    {
        değişken := "iki"
        fmt.Println(değişken)
        fmt.Println(&değişken)
    }
    fmt.Println(değişken)
    fmt.Println(&değişken)
}
```

**& \(and\)** işareti ile değişkenin bellekteki adresini öğrenebiliriz.

Çıktımız şöyle olacaktır;

> iki
>
> 0xc00008c1d0
>
> bir
>
> 0xc00008c1c0

Gördüğünüz gibi bellek adresi 2 sonuçta da farklı gözüküyor. Bu ikisinin de ayrı değişkenler olduğuna işaret ediyor.

## Tür Dönüşümü

Tür dönüşümü şu şekilde gerçekleştirilir.

**`tür(değer)`**

Örnek olarak bakmak gerekir ise;

```go
i := 42
f := float64(i)
u := uint(f)
```

Yukarıdaki yapılan işlemleri açıklayacak olursak eğer, **i** adında bir `int` değişken tanımladık. **f** adındaki değişkende **i** değişkenini `float64` türüne dönüştürdük. **u** adındaki değişkende ise **f** değişkenini `uint` türüne çevirdik.

Tüm türler arasında bu şekilde dönüşüm gerçekleştiremezsiniz. Bir sayıyı string tipine dönüştürmek istediğimizde ne olacağına bakalım.

```go
deneme := string(8378)
fmt.Println(deneme)
```

**deneme** adındaki değişkenimizin içinde **8378** sayısını **string** türüne dönüştürdük ve hemen aşağısına **deneme**’nin aldığı değeri ekrana bastırması için kodumuzu yazdık.

Aldığımız konsol çıktısı şu şekilde olacaktır.

> ₺

Yani Türk Lirası simgesi çıkacaktır. Sayılar string türüne dönüştürüldüğünde karakter olarak değer alır.

