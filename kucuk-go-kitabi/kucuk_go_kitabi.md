- [Kitap Hakkında](#kitap-hakkında)
  - [Lisans](#lisans)
  - [En Son Sürüm](#en-son-sürüm)
- [Giriş](#giriş)
  - [Yazardan Bir Not](#yazardan-bir-not)
- [Başlarken](#başlarken)
  - [OSX / Linux](#osx--linux)
  - [Windows](#windows)
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
- [Bölüm 2 - Veri Yapıları](#bölüm-2---veri-yapıları)
  - [Tanımlamalar ve Başlangıç Değerleri](#tanımlamalar-ve-başlangıç-değerleri)
  - [Veri Yapılarındaki İşlevler](#veri-yapılarındaki-i̇şlevler)
  - [Oluşturucu Yöntemler](#oluşturucu-yöntemler)
  - [New](#new)
  - [Yapıların Alanları](#yapıların-alanları)
  - [İçerme](#i̇çerme)
    - [Yeni Görev Yükleme](#yeni-görev-yükleme)
  - [İşaretçiler mi, Değerler mi?](#i̇şaretçiler-mi-değerler-mi)
  - [Devam Etmeden Önce](#devam-etmeden-önce-1)
- [Bölüm 3 - Map, Array ve Slice](#bölüm-3---map-array-ve-slice)
  - [Diziler](#diziler)
  - [Dilimler](#dilimler)
  - [Eşlemeler (Map)](#eşlemeler-map)
  - [İşaretçiler ve Değerler](#i̇şaretçiler-ve-değerler)
  - [Devam Etmeden Önce](#devam-etmeden-önce-2)
- [Bölüm 4 - Kod Organizasyonu ve Arayüzler](#bölüm-4---kod-organizasyonu-ve-arayüzler)
  - [Paketler](#paketler)
    - [Döngüsel İçe Aktarım](#döngüsel-i̇çe-aktarım)
    - [Görünürlük](#görünürlük)
    - [Paket Yönetimi](#paket-yönetimi)
    - [Bağımlılık Yönetimi](#bağımlılık-yönetimi)
  - [Arayüzler](#arayüzler)
  - [Devam Etmeden Önce](#devam-etmeden-önce-3)
- [Bölüm 5 - Yararlı Bilgiler](#bölüm-5---yararlı-bilgiler)
  - [Hata Yönetimi](#hata-yönetimi)
  - [Erteleme](#erteleme)
  - [go fmt](#go-fmt)
  - [Değer Atamalı If](#değer-atamalı-if)
  - [Boş Arayüz ve Dönüşümler](#boş-arayüz-ve-dönüşümler)
  - [Dizeler ve Bayt Dizileri](#dizeler-ve-bayt-dizileri)
  - [İşlev Tipi](#i̇şlev-tipi)
  - [Devam Etmeden Önce](#devam-etmeden-önce-4)
- [Bölüm 6 - Eşzamanlılık](#bölüm-6---eşzamanlılık)
  - [Go Rutinleri](#go-rutinleri)
  - [Senkronizasyon](#senkronizasyon)
  - [Kanallar](#kanallar)
    - [Tamponlu Kanallar](#tamponlu-kanallar)
    - [Select](#select)
    - [Zaman Aşımı](#zaman-aşımı)
  - [Devam Etmeden Önce](#devam-etmeden-önce-5)
- [Sonuç](#sonuç)

# Kitap Hakkında

## Lisans

The Little Go Book, Attribution-NonCommercial-ShareAlike 4.0 International lisansı ile lisanslanmıştır. Bu kitap için ödeme yapmamalısınız.

Kitabı kopyalamak, dağıtmak, değiştirmek veya yayınlamakta özgürsünüz. Ancak, kitabı her zaman bana (Karl Seguin) atfetmenizi istiyorum ve ticari amaçlarla kullanmayın.

Lisansın tam metnini şu adreste görebilirsiniz:

[https://creativecommons.org/licenses/by-nc-sa/4.0/](https://creativecommons.org/licenses/by-nc-sa/4.0/)

## En Son Sürüm

Bu kitabın en son kaynağını şu adreste bulabilirsiniz: [https://github.com/karlseguin/the-little-go-book](https://github.com/karlseguin/the-little-go-book).

# Giriş

Yeni dil öğrenmek söz konusu olduğunda her zaman bir aşk-nefret ilişkisi yaşadım. Bir yandan, diller yaptığımız şeyler için o kadar temeldir ki, küçük değişiklikler bile ölçülebilir bir etkiye sahip olur. Bir şeye tıklama sırasında oluşan *aha* anı program yazma şeklinize kalıcı bir etkisi olabilir ve diğer dillerden olan beklentilerinizi tanımlayabilir. Daha detaylı bakarsak, dil tasarımı oldukça değişkendir. Yeni anahtar kelimeler, tür sistemi, kodlama stili, yeni kütüphaneler, topluluklar ve paradigmalar öğrenmek, uyum sağlanması zor görünen bir iştir. Öğrenmemiz gereken diğer şeylerle karşılaştırıldığında, yeni dil öğrenmek genellikle zamanımızın zayıf bir yatırımı gibi hissettiriyor.

Bununla birlikte, ilerlemek *zorundayız*. Birbirini takip eden adımları tekrar tekrar atmaya istekli olmak *zorundayız* çünkü diller yaptığımız işin temelini oluşturuyor. Değişiklikler genellikle aşama aşama olmakla birlikte, geniş bir kapsama sahip olma eğilimindedir ve üretkenliği, okunabilirliği, performansı, test edilebilirliği, bağımlılık yönetimini, hata yönetimi, dokümantasyon, profil oluşturma, topluluklar, standart kütüphaneler ve bir çok şeyi etkilerler. *Bin bıcak darbesiyle ölümü* söylemenin güzel bir yolu var mı?

Bu bizi önemli bir soru ile baş başa bırakıyor: **neden Go?** Benim için iki temel neden var. Birincisi, nispeten basit bir standart kütüphaneye sahip nispeten sade basit bir dildir. Birçok yönden, Go'nun gelişemeye uygun doğası, son birkaç on yılda dillere eklendiğini gördüğümüz birçok karmaşıklığı basitleştirmektir. Diğer bir neden ise, birçok geliştirici için mevcut cephaneliği tamamlayacı niteliğidir.

Go bir sistem dili olarak inşa edildi (örn. işletim sistemleri, aygıt sürücüleri) ve bu nedenle C ve C ++ geliştiricilerine yönelikti. Go ekibine göre ve benim için de kesinlikle doğru olan, uygulama geliştiricileri (sistem geliştiricileri değil) birincil Go kullanıcıları haline geldi. Neden? Sistem geliştiricileri için yetkili bir şekilde konuşamam, ancak web siteleri, web servisler, masaüstü uygulamaları ve benzerlerini inşa eden bizler için kısmen düşük seviyeli sistem uygulamaları ile üst düzey uygulamalar arasına konumlandırılabilecek bir sistem ihtiyacından ortaya çıkmıştır.

Belki bir mesajlaşma, önbellekleme, hesaplama-ağır veri analizi, komut satırı arayüzü, günlük tutma veya izleme. Hangi etiketi vereceğimi bilmiyorum, ancak kariyerim boyunca, sistemler karmaşıklık içinde büyümeye devam ettikçe ve eşzamanlılık on binler seviyesine ölçüldüğü için, özel altyapı tipi sistemlere artan bir ihtiyaç olduğu açıktır. Bu tür sistemleri Ruby veya Python veya başka bir şeyle inşa *edebilirsiniz* (ve birçok kişi yapar), ancak bu tür sistemler daha katı bir tip sistemden ve daha yüksek performanstan daha çok yararlanabilir. Benzer şekilde, web sitelerini (ve birçok kişinin yaptığı gibi) oluşturmak için Go *kullanabilirsiniz*, ama yine de geniş bir marj içinde, bu tür sistemler için Node veya Ruby tercih ederim.

Go'nun mükemmel olduğu başka alanlar da var. Örneğin, derlenmiş bir Go programını çalıştırırken herhangi bir bağımlılığı yoktur. Kullanıcılarınızda Ruby veya JVM kurulu olup olmadığını ve kurulu ise hangi sürüm olduğunu endişe etmenize gerek yoktur. Bu nedenle Go, komut satırı arayüz programları ve dağıtmanız gereken diğer yardımcı program türleri (örneğin bir log toplayıcı) için giderek daha popüler hale gelmektedir.

Açıkça söylemek gerekirse, Go'yu öğrenmek zamanınızı verimli bir şekilde kullanmaktır. Go'yu öğrenmek ve hatta ustalaşmak için uzun saatler harcamak zorunda kalmazsınız ve çabalarınızdan pratik bir şey elde edersiniz.

## Yazardan Bir Not

Bu kitabı birkaç nedenden dolayı yazmakta tereddüt ettim. Birincisi Go'nun kendi belgelerinin, özellikle de [Effective Go'nun](https://golang.org/doc/effective_go.html) çok sağlam olması.

Diğeri ise bir dil hakkında bir kitap yazmamdaki rahatsızlığım. The Little MongoDB Book kitabını yazdığımda, çoğu okuyucunun ilişkisel veritabanı ve modellemenin temellerini anladığını varsaymak güvenliydi. The Little Redis Book ile, bir anahtar değer deposuna aşinalık kazanabilir ve oradan başlayarak öğrenebilirsiniz.

Önde gelen paragraflar ve bölümler hakkında düşündüğüm gibi, aynı varsayımları yapamayacağımı biliyorum. Bazıları için kavramın yeni olacağını, diğerlerinin *Go'nun arayüzlerinden* çok daha fazlasına ihtiyaç duymayacağını bilerek arayüzler hakkında ne kadar zaman harcıyorsunuz? Nihayetinde, bazı parçaların çok sığ ya da çok ayrıntılı olup olmadığını bana bildireceğinizi bilmek beni rahatlatıyor. Bunu kitap için harcanan emeğin ücereti olarak düşünün.

# Başlarken

Go ile biraz oynamak istiyorsanız, hiçbir şey yüklemeden çevrimiçi kod çalıştırmanıza izin veren [Go Playground'a](https://play.golang.org/) göz atmalısınız. Bu, [Go'nun tartışma forumunda](https://groups.google.com/forum/#!forum/golang-nuts) ve StackOverflow gibi yerlerde yardım ararken Go kodunu paylaşmanın en yaygın yoludur.

Go'yu yüklemek basittir. Kaynaktan yükleyebilirsiniz, ancak önceden derlenmiş dosyalardan birini kullanmanızı öneririm. [İndirme sayfasına gittiğinizde](https://golang.org/dl/), çeşitli platformlar için kurulum dosyaları görürsünüz. Bunlardan kaçınalım ve Go'yu nasıl kuracağımızı öğrenelim. Göreceğiniz üzere zor değil.

Basit örnekler dışında, Go kodunuz bir ana çalışma klasörü içindeyken çalışmak üzere tasarlanmıştır. Çalışma klasörü `bin` , `pkg` ve `src` alt klasörlerinden oluşan bir klasördür. Gitmek için kendi tarzını takip etmeye zorlayabilirsin - ama yapma.

Normalde projelerimi `~/code ` içine koyarım. Örneğin, `~/code/blog` blogumu içeriyor. Go için çalışma alanım `~/code/go` ve Go destekli blogum `~/code/go/src/blog `'da olacaktır.

Kısacası, projelerinizi koymak istediğiniz her yerde `src` alt klasörü içeren bir `go` klasörü oluşturun.

## OSX / Linux

Platformunuz için `tar.gz` dosyasını indirin. OSX için büyük olasılıkla `go#.#.#.darwin-amd64-osx10.8.tar.gz` dosyasını seçersiniz, burada `#.#.#` Go'nun en son sürümüdür.

Dosyayı `/usr/local` klasörü altına `tar -C /usr/local -xzf go#.#.#.darwin-amd64-osx10.8.tar.gz` komutu ile açın.

İki ortam değişkeni ayarlayın:

1. `GOPATH`, ana çalışma klasörünüzü gösterir, benim için bu, `$HOME/code/go` .
2. Go'nun binary dosyasını sistem `PATH` listesine eklemeliyiz.

Bunları bir kabuktan aşağıdaki gibi ayarlayabilirsiniz:

```
echo 'export GOPATH=$HOME/code/go' >> $HOME/.profile
echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.profile
```

Bu değişkenleri her kabuk oturumunda aktif olmasını isteyebilirsiniz. Kabuğunuzu kapatıp yeniden açabilir veya `source $HOME/.profile` komutunu çalıştırabilirsiniz.

Hangi sürümü kullandığınızı `go version` komutunu çalıştırarak görebilirsiniz, muhtemelen `go version go1.3.3 darwin/amd64` gibi görünen bir çıktı alırsınız.

## Windows

En son sürüm zip dosyasını indirin. Bir x64 sistemindeyseniz `go#.#.#.windows-amd64.zip` dosyasını indirmeniz gerekcektir, burada `#.#.#` Go'nun en son sürümüdür.

Seçtiğiniz bir yerde açın. `c:\Go` iyi bir seçimdir.

İki ortam değişkeni ayarlayın:

1. `GOPATH` ana çalışma klasörünüzü gösterir. Bu `c:\users\goku\work\go` gibi bir şey olabilir.
2. `PATH` ortam değişkeninize `c:\Go\bin` ekleyin.

Ortam değişkenleri, `System` kontrol panelinin `Advanced` sekmesindeki `Environment Variables` düğmesiyle ayarlanabilir. Bazı Windows sürümleri bu kontrol panelini `System` kontrol panelindeki `Advanced System Settings` seçeneğiyle sağlar.

Bir komut istemi açın ve `go version` . Umarım `go version go1.3.3 windows/amd64` gibi görünen bir çıktı alırsınız.

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

# Bölüm 2 - Veri Yapıları

Go, C ++, Java, Ruby ve C # gibi nesne yönelimli (OO) bir dil değildir. Nesneleri veya kalıtımları yoktur ve bu nedenle OO ile ilişkili polimorfizm ve yeni görev yükleme gibi pek çok kavramı yoktur.

Go'nun sahip olduğu, yöntemlerle ilişkilendirilebilen veri yapılardır. Go ayrıca basit ama etkili bir kompozisyon biçimini de destekler. Genel olarak, daha basit bir kodla sonuçlanır, ancak OO'nun sunduğu şeylerden bazılarını kaçıracağınız durumlar olacaktır. ( *Kalıtıma karşı kompozisyonun* eski bir savaş olduğunu ve Go'nun bu konuda sağlam bir duruş sergileyen ilk dil olduğunu belirtmek gerekir.)

Go alıştığınız gibi OO yapmasa da, bir struct yapısının tanımı ile bir sınıfın tanımı arasında birçok benzerlik olduğunu fark edeceksiniz. Basit bir örnek aşağıdaki `Saiyan` yapısıdır:

```go
type Saiyan struct {
  Name string
  Power int
}
```

Yakında bu yapıya nasıl yöntem ekleyebileceğimizi göreceğiz, tıpkı bir sınıfın parçası olarak eklediğimiz yöntemler gibi. Bunu yapmadan önce, tanımlamalara daha detaylı dalmalıyız.

## Tanımlamalar ve Başlangıç Değerleri

Değişkenlere ve tanımlamalara ilk baktığımızda, yalnızca tamsayılar ve dizgiler gibi yerleşik tiplere baktık. Şimdi yapılar hakkında konuştuğumuza göre, bu konuşmayı işaretçiler (pointer) içerecek şekilde genişletmeliyiz.

Veri yapımızın bir değerini yaratmanın en basit yolu:

```go
goku := Saiyan{
  Name: "Goku",
  Power: 9000,
}
```

*Not:* Yukarıdaki yapıda sonda yer alan `,` gereklidir. Bu olmadan, derleyici bir hata verecektir. Özellikle bunun aksini uygulayan bir dil veya biçim kullandıysanız, gereken tutarlılığı takdir edersiniz.

Alanların tümüne hatta herhangi birine değer ayamanız gerekmiyor. Bunların her ikisi de geçerlidir:

```go
goku := Saiyan{}

// ya da

goku := Saiyan{Name: "Goku"}
goku.Power = 9000
```

Atanmamış değişkenlerin sıfır değeri olduğu gibi veri yapısındaki atanmamış alanların da sıfır değeri olur.

Ayrıca, alan adını atlayabilir ve alan bildirimlerinin sırasına güvenebilirsiniz (sadelik sağlamak için bunu sadece birkaç alana sahip yapılar için yapmalısınız):

```go
goku := Saiyan{"Goku", 9000}
```

Yukarıdaki örneklerin tümü, `goku` adıyla bir değişken tanımlamak ve buna bir değer atamaktır.

Çoğu zaman, doğrudan değerimizle ilişkilendirilmiş bir değişken değil de, değerimize bir işaretçi olan bir değişken kullanmak isteriz. İşaretçi bir bellek adresidir; gerçek değeri nerede bulacağınız söyler. Bu bir dolaylama düzeyidir. Gerçek hayattan örnek vermek gerekirse, bir evde olmak ile evin yönünü bilmek arasındaki farktır.

Neden gerçek değer yerine değere bir işaretçi istiyoruz? Go'nun, argümanları bir işleve aktarma biçiminden dolayı: kopya olarak aktardığı için. Bunu bilerek düşünün, aşağıdaki kod ekrana ne yazdırır?

```go
func main() {
  goku := Saiyan{"Goku", 9000}
  Super(goku)
  fmt.Println(goku.Power)
}

func Super(s Saiyan) {
  s.Power += 10000
}
```

Cevap 19000 değil 9000. Neden? Çünkü `Super` orijinal `goku` değişkeninin bir kopyasını alır ve değişiklikleri ona yapar, yapılan değişiklikler `Super` işlevini çağıran bağlama yansımaz. Bu işlemi muhtemelen beklediğiniz gibi yapabilmek için, işleve bir değişken işaretçisi göndermemiz gerekir:

```go
func main() {
  goku := &Saiyan{"Goku", 9000}
  Super(goku)
  fmt.Println(goku.Power)
}

func Super(s *Saiyan) {
  s.Power += 10000
}
```

İki değişiklik yaptık. Birincisi, `&` operatörünün değerimizin adresini almak için kullanılmasıdır (buna operatörün *adresi* denir). Sonra, `Super` işlevinin beklediği parametre tipini değiştirdik. `Saiyan` tipinde bir değer bekliyordu ama şimdi `*Saiyan` tipinde bir adres bekliyor, burada `*X` *X tipindeki değere işaretçi* anlamına geliyor. Açıkçası `Saiyan` ve `*Saiyan` tipleri arasında bazı ilişkiler vardır, ancak bunlar iki farklı tiptir.

`goku's` değişkeninin hala `Super` işlevi içinde kopyalandığını unutmayın, ama bu kez kopyalanan `goku'nun` değeri değil adresidir. Bu kopya orijinalle aynı adrestir, bu da dolaylı aktarmanın bizi sağladığı bir şeydir. Bir restoranın yol tarifini kopyalamak olarak düşünün. Sahip olduğunuz bir kopya, ama yine de orijinal ile aynı restorana işaret ediyor.

İşaret ettiği yeri değiştirmeye çalışarak bunun bir kopya olduğunu kanıtlayabiliriz (aslında yapmak isteyeceğiniz bir şey değil):

```go
func main() {
  goku := &Saiyan{"Goku", 9000}
  Super(goku)
  fmt.Println(goku.Power)
}

func Super(s *Saiyan) {
  s = &Saiyan{"Gohan", 1000}
}
```

Yukarıdaki, bir kez daha ekrana 9000 yazdırır. Ruby, Python, Java ve C# dahil olmak üzere bir çok dil böyle davranır. Go ve bir dereceye kadar C#, sadece gerçeği daha görünür yapmıştır.

İşaretçi kopyalamanın karmaşık bir yapıyı kopyalamaktan daha ucuz olacağı da açıktır. 64 bit makinede, bir işaretçi 64 bit büyüklüğündedir. Birçok alanı olan karmaşık bir yapımız varsa, kopya oluşturmak daha pahalı olabilir. İşaretçilerin gerçek değeri, değerleri paylaşmanıza izin vermesidir. `Super` işlevinin `goku` değişkeninin bir kopyasını değiştirmesini mi veya paylaşılan `goku` değerini değiştirmesini mi istiyoruz?

Bütün bunlarla her zaman bir işaretçi isteyeceğinizi söyleyemeyiz. Bu bölümün sonunda, yapılarla neler yapabileceğimizden biraz daha fazlasını gördükten sonra, işaretçi-değer sorusunu yeniden inceleyeceğiz.

## Veri Yapılarındaki İşlevler

Bir işlevi (yöntemi) bir yapı ile ilişkilendirebiliriz:

```go
type Saiyan struct {
  Name string
  Power int
}

func (s *Saiyan) Super() {
  s.Power += 10000
}
```

Yukarıdaki kodda, `*Saiyan` tipinin `Super` yönteminin **alıcısı** olduğunu söylüyoruz. `Super` yöntemini aşağıdaki gibi çağırıyoruz:

```go
goku := &Saiyan{"Goku", 9001}
goku.Super()
fmt.Println(goku.Power) // ekrana 19001 yazacak
```

## Oluşturucu Yöntemler

Veri yapılarının oluşturucu yöntemleri yoktur. Bunun yerine, istenen türde bir örneği (fabrika gibi) döndüren bir işlev oluşturursunuz:

```go
func NewSaiyan(name string, power int) *Saiyan {
  return &Saiyan{
    Name: name,
    Power: power,
  }
}
```

Bu model birçok geliştiriciyi yanlış şekilde yönlendiriyor. Bir yandan, oldukça hafif bir söz dizimsel değişiklik; diğer yandan, biraz daha az bölümlendirilmiş hissettiriyor.

Oluşturucu yöntemimiz bir işaretçi döndürmek zorunda değildir; aşağıdaki satırlar kesinlikle geçerlidir:

```go
func NewSaiyan(name string, power int) Saiyan {
  return Saiyan{
    Name: name,
    Power: power,
  }
}
```

## New

Oluşturucuların olmamasına rağmen, Go bir tip için gereken belleği ayırmak için kullanılan yerleşik bir `new` işlevine sahiptir. `new(X)` in sonucu `&X{}` aynıdır:

```go
goku := new(Saiyan)
// eşittir
goku := &Saiyan{}
```

Hangisini kullanacağınız size kalmış, ancak çoğu insanın başlangıç için alanları olduğunda ikincisini de tercih ettiğini göreceksiniz, çünkü daha okunabilir kabul edilir:

```go
goku := new(Saiyan)
goku.name = "goku"
goku.power = 9001

//vs

goku := &Saiyan {
  name: "goku",
  power: 9000,
}
```

Hangi yaklaşımı seçerseniz seçin, yukarıdaki oluşturma yöntemi modelini izlerseniz, kodunuzun geri kalanını tanımlama ayrıntılarını bilmekten ve endişelenmekten koruyabilirsiniz.

## Yapıların Alanları

Yukarıda gördüğümüz örnekte `Saiyan` sırasıyla `string` ve `int` tiplerinde `Name` ve `Power` adlarında iki alanı (veri alanı) vardır. Alanlar, diziler, map'ler, arabirimler ve işlevler gibi henüz bahsetmediğimiz diğer yapılar ve tipler dahil olmak üzere herhangi bir tip olabilir.

Örneğin, `Saiyan` tanımımızı şöyle genişletebiliriz:

```go
type Saiyan struct {
  Name string
  Power int
  Father *Saiyan
}
```

aşağıdaki gibi değer ataması yapabiliriz:

```go
gohan := &Saiyan{
  Name: "Gohan",
  Power: 1000,
  Father: &Saiyan {
    Name: "Goku",
    Power: 9001,
    Father: nil,
  },
}
```

## İçerme

Go, bir yapıyı diğerine dahil etme eylemi olan içermeyi destekler. Bazı dillerde buna trait veya mixin denir. Açık bir içerme mekanizmasına sahip olmayan diller çoğu zaman bunu daha uzun yoldan yapabilir. Java'da, *kalıtımla* yapıları genişletme imkanı vardır, ancak bunun bir seçenek olmadığı bir senaryoda, şöyle bir mixin yazılacaktır:

```java
public class Person {
  private String name;

  public String getName() {
    return this.name;
  }
}

public class Saiyan {
  // Saiyan is said to have a person
  private Person person;

  // we forward the call to person
  public String getName() {
    return this.person.getName();
  }
  ...
}
```

Bu oldukça sıkıcı olabilir. Her `Person` sınıfına ait her yöntem `Saiyan` sınıfı için çoğaltmak gerekmektedir. Go bu sıkıcılığı önler:

```go
type Person struct {
  Name string
}

func (p *Person) Introduce() {
  fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
  *Person
  Power int
}

// and to use it:
goku := &Saiyan{
  Person: &Person{"Goku"},
  Power: 9001,
}
goku.Introduce()
```

`Saiyan` yapısının `*Person` tipinde bir alanı vardır. Açık bir alan adı vermediğimiz için, içeri aktarılan tipin alanlarına ve işlevlerine dolaylı olarak erişebiliriz. Ancak bunun yanında Go derleyicisi tamamen geçerli olan bir alan adı da *verdi*:

```go
goku := &Saiyan{
  Person: &Person{"Goku"},
}
fmt.Println(goku.Name)
fmt.Println(goku.Person.Name)
```

Yukarıdakilerin her ikisi de "Goku" yazdıracaktır.

İçeri aktarma kalıtımdan daha mı iyidir? Birçok kişi bunun kod paylaşmanın daha sağlam bir yolu olduğunu düşünüyor. Kalıtım kullanırken, sınıfınız üst sınıfınıza sıkı sıkıya bağlıdır ve sonunda davranış yerine hiyerarşiye odaklanırsınız.

### Yeni Görev Yükleme

Yeni görev yükleme veri yapılarına özgü olmasa da, burada ele almaya değer. Basitçe, Go yeni görev yüklemeyi desteklemez. Bu nedenle, `Load` , `LoadById` , `LoadByName` ve benzeri birçok işlevi görürsünüz (ve yazarsınız).

Ancak, içeri aktarma gerçekten sadece bir derleyici hilesi olduğundan, içeri aktarılmış bir tipin yönteminin üzerine "yazabiliriz". Örneğin, `Saiyan` yapısının kendi `Introduce` yöntemi olabilir:

```go
func (s *Saiyan) Introduce() {
  fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}
```

İçeri aktarılan tipin yöntemine de her zaman erişilebilir, örneğin `s.Person.Introduce()`.

## İşaretçiler mi, Değerler mi?

Go ile program yazarken, kendinize *bunu değer olarak mı kullanayım yoksa işaretçi olarak mı?* diye sormanız normaldir. Bu sorunun cevabını bulmamıza yardımcı olacak iki güzel haberim var. İlki, aşağıdakilerden hangisinden söz ettiğimize bakılmaksızın cevap aynıdır:

- Yerel değişken ataması
- Bir yapıdaki veri alanı
- Bir işlevden değer döndürme
- Bir fonksiyonun parametreleri
- Bir yöntemin alıcısı

İkinci olarak, emin değilseniz, bir işaretçi kullanın.

Daha önce gördüğümüz gibi, değişkenleri değer olarak iletmek verileri değiştirilemez hale getirmek için harika bir yoldur (bir işlevin yaptığı değişiklikler işlevi çağıran koda yansıtılmaz). Bazen, bu isteyeceğiniz davranıştır, ancak çoğunlukla değildir.

Verileri değiştirmek istemeseniz bile, büyük yapıların bir kopyasını oluşturma maliyetini göz önünde bulundurun. Tersine, küçük yapılarınız olabilir, örneğin:

```go
type Point struct {
  X int
  Y int
}
```

Bu gibi durumlarda, yapının kopyalanma maliyeti muhtemelen herhangi bir dolaylama olmaksızın doğrudan `X` ve `Y` erişebilmekle dengelenir.

Yine, bunların hepsi oldukça ince vakalardır. Binlerce veya muhtemelen on binlerce noktayı yinelemediğiniz sürece bir fark görmezsiniz.

## Devam Etmeden Önce

Pratik bir bakış açısından, bu bölüm yapıları, bir yapının bir örneğinin bir fonksiyonun alıcısı haline getirilmesini ve mevcut Go'nun tip sistemi bilgimize işaretçileri de eklemiştir. Aşağıdaki bölümler, yapılar hakkında bildiklerimizi ve keşfettiğimiz özellikleri temel alacak.

# Bölüm 3 - Map, Array ve Slice

Şimdiye kadar bir dizi basit tip ve yapı gördük. Şimdi dizilere, dilimlere (slice) ve map'lere bakma zamanı.

## Diziler

Python, Ruby, Perl, JavaScript veya PHP'den (ve daha fazlasından) geliyorsanız, muhtemelen *dinamik dizilerle* programlama yapmaya alışıksınızdır. Bunlar, veri eklendikçe kendilerini yeniden boyutlandıran dizilerdir. Go'da diğer birçok dil gibi diziler de sabittir. Bir dizinin tanımlanması için boyutu belirtmemizi gerektirir ve boyut belirtildikten sonra dizi büyüyemez:

```go
var scores [10]int
scores[0] = 339
```

Yukarıdaki dizi indeks `scores[0]` ile `scores[9]` arasında en fazla 10 skor saklayabilir. Dizideki indeks aralığı dışında bir değere erişme girişimleri derleyici veya çalışma zamanı hatasına neden olur.

Dizileri tanımlarken de değer verebiliriz:

```go
scores := [4]int{9001, 9333, 212, 33}
```

Dizinin uzunluğunu elde etmek için `len` kullanabiliriz. `range` dizi üzerinde döngü yapmak için kullanılabilir:

```go
for index, value := range scores {

}
```

Diziler verimli ancak katıdır. Sıklıkla ele alacağımız değerlerin sayısını bilmeyiz. Bu durumlar için dilimler kullanışlıdır.

## Dilimler

Go'da nadiren, eğer gerçekten ihtiyaç varsa, doğrudan dizileri kullanırsınız. Bunun yerine dilimleri kullanırsınız. Dilim, bir dizinin bir bölümünü kapsayan ve temsil eden hafif bir yapıdır. Bir dilim oluşturmanın birkaç yolu vardır ve hangisini ne zaman kullanacağımıza daha sonra detaylı bakacağız. Birinci yol, bir diziyi nasıl oluşturduğumuza ilişkin küçük bir varyasyon:

```go
scores := []int{1,4,293,4,9}
```

Dizi tanımlanmasından farklı olarak, dilim tanımlanmasında köşeli parantez içinde bir uzunluk bildirilmez. Bu iki tanımlamanın nasıl farklı çalıştığını anlamak için, bir dilim oluşturmak için farklı bir yol olan `make` kullanımına bakalım:

```go
scores := make([]int, 10)
```

`new` yerine `make` kullanıyoruz çünkü bir dilim oluşturmak için sadece belleği ayırmaktan daha fazlası var ( `new` de olan budur). Özellikle, altta yatan dizi için bellek ayırmalı ve dilimi başlatmalıyız. Yukarıda, uzunluğu 10 ve kapasitesi 10 olan bir dilim oluştururuz. Uzunluk, dilimin boyutudur, kapasite, temel dizinin boyutudur. `make` kullanırken ikisini ayrı ayrı belirleyebiliriz:

```go
scores := make([]int, 0, 10)
```

Bu 0 uzunluğunda sahip ama 10 kapasiteli bir dilim oluşturur. (Eğer dikkat ettiyseniz, `make` ve `len` işlevlerine Go yeni bir görev *yüklemiş* durumda. Bu bazılarını çok kızdırıyor ama Go geliştiricilerin kullanımına açık olmayan özellikleri bazen kendi kullanan bir dildir.)

Uzunluk ve kapasite arasındaki etkileşimi daha iyi anlamak için bazı örneklere bakalım:

```go
func main() {
  scores := make([]int, 0, 10)
  scores[7] = 9033
  fmt.Println(scores)
}
```

İlk örneğimiz hata verir. Neden? Dilimimizin uzunluğu 0 olduğu için. Evet, temel alınan dizinin 10 öğesi vardır, ancak bu öğelere erişmek için dilimimizi açıkça genişletmemiz gerekir. Bir dilimi genişletmenin bir yolu `append` yöntemdir:

```go
func main() {
  scores := make([]int, 0, 10)
  scores = append(scores, 5)
  fmt.Println(scores) // ekrana [5] yazar
}
```

Ancak bu, orijinal kodumuzun amacını değiştirir. 0 uzunluğunda bir dilimi genişletere ilk öğeye değeri atar. Herhangi bir nedenle, hata veren kodumuzda 7 indeksli öğeye değer atamak istiyorduk. Bunu yapmak için dilimi yeniden ayarlayabiliriz:

```go
func main() {
  scores := make([]int, 0, 10)
  scores = scores[0:8]
  scores[7] = 9033
  fmt.Println(scores)
}
```

Bir dilimi en fazla ne kadar boyutlandırabiliriz? Bu durumda 10 olan dizi kapasitesine kadar. Bunun *dizilerin sabit uzunluklu sorununu gerçekten çözmediğini* düşünüyor olabilirsiniz. `append ` oldukça özel işlev. Altta yatan dizi doluysa, yeni daha büyük bir dizi oluşturur ve değerleri kopyalar (bu tam olarak PHP, Python, Ruby, JavaScript, vb dillerde dinamik dizilerin çalışma şeklidir). Bu nedenle, `append` kullanılan yukarıdaki örnekte, `append` tarafından döndürülen değeri `scores` değişkenimize yeniden atamak zorunda kaldık, çünkü `append` orijinalde daha fazla yer yoksa yeni bir değer yaratmış olabilir.

Size Go'nun 2x algoritmasıyla diziler oluşturduğunu söylersem, aşağıdaki kodda ne olacağını tahmin edebilir misiniz?

```go
func main() {
  scores := make([]int, 0, 5)
  c := cap(scores)
  fmt.Println(c)

  for i := 0; i < 25; i++ {
    scores = append(scores, i)

    // if our capacity has changed,
    // Go had to grow our array to accommodate the new data
    if cap(scores) != c {
      c = cap(scores)
      fmt.Println(c)
    }
  }
}
```

`scores` başlangıç kapasitesi 5'tir. 25 tane değer tutmak için 10, 20 ve son olarak 40 kapasiteyle 3 kez genişletilmesi gerekecektir.

Son bir örnek olarak şunları göz önünde bulundurun:

```go
func main() {
  scores := make([]int, 5)
  scores = append(scores, 9332)
  fmt.Println(scores)
}
```

Burada çıktı `[0, 0, 0, 0, 0, 9332]` şeklinde olacak. Belki bunun `[9332, 0, 0, 0, 0]` olacağını düşündünüz. Bu insana mantıklı gelebilir. Bir derleyici için, zaten 5 değer içeren bir dilime bir değer eklemesini söylüyorsunuz.

Nihayetinde, bir dilime başlangıç değeri atamak için dört yaygın yol vardır:

```go
names := []string{"leto", "jessica", "paul"}
checks := make([]bool, 10)
var names []string
scores := make([]int, 0, 20)
```

Hangisini ne zaman kullanırız? İlki çok fazla açıklamaya ihtiyaç duymamalı. Bunu, dizide istediğiniz değerleri önceden bildiğinizde kullanırsınız.

İkincisi, bir dilimin belirli indeksine yazarken kullanışlıdır. Örneğin:

```go
func extractPowers(saiyans []*Saiyan) []int {
  powers := make([]int, len(saiyans))
  for index, saiyan := range saiyans {
    powers[index] = saiyan.Power
  }
  return powers
}
```

Üçüncü yol nil bir dilim oluşturur ve eleman sayısı bilinmediğinde, `append` ile birlikte kullanılır.

Son yol bir başlangıç kapasitesi belirlememizi sağlar; kaç öğeye ihtiyacımız olacağına dair genel bir fikrimiz varsa yararlı olur.

Boyutu bilseniz bile, `append` kullanılabilir. Bu büyük ölçüde bir tercih meselesi:

```go
func extractPowers(saiyans []*Saiyan) []int {
  powers := make([]int, 0, len(saiyans))
  for _, saiyan := range saiyans {
    powers = append(powers, saiyan.Power)
  }
  return powers
}
```

Dizilere erişim için dilimler güçlü bir kavramdır. Birçok dil bir dizi dilimleme kavramına sahiptir. Hem JavaScript hem de Ruby dizilerinin bir `slice` yöntemi vardır. `[START..END]` kullanarak Ruby'de veya `[START:END]` Python'da bir dilim alabilirsiniz. Bununla birlikte, bu dillerde, bir dilim aslında orijinalinin değerleri kopyalanan yeni bir dizidir. Ruby alırsak, aşağıdakilerin çıktısı nedir?

```ruby
scores = [1,2,3,4,5]
slice = scores[2..4]
slice[0] = 999
puts scores
```

Cevap `[1, 2, 3, 4, 5]` . Çünkü `slice` , değerlerin kopyalarını içeren tamamen yeni bir dizidir. Şimdi, Go eşdeğerini düşünün:

```go
scores := []int{1,2,3,4,5}
slice := scores[2:4]
slice[0] = 999
fmt.Println(scores)
```

Çıktı `[1, 2, 999, 4, 5]` olur.

Bu, kodlama şeklinizi değiştirir. Örneğin, bir dizi fonksiyonu bir pozisyon parametresi alsın. JavaScript'te, bir dizedeki ilk alanı bulmak istiyorsak (evet, dilimler dizelerde de çalışır!) İlk beş karakterden sonra şunu yazarız:

```javascript
haystack = "the spice must flow";
console.log(haystack.indexOf(" ", 5));
```

Go'da dilimleri kullanırız:

```go
strings.Index(haystack[5:], " ")
```

Yukarıdaki örnekten görebiliyoruz ki, `[X:]` *X'ten sonuna kadar* kısayolu iken, `[:X]` *başından X'e kadar* kısayoludur. Diğer dillerden farklı olarak, Go negatif değerleri desteklemez. Sonuncusu hariç bir dilimin tüm değerlerini istiyorsak şunu yaparız:

```go
scores := []int{1, 2, 3, 4, 5}
scores = scores[:len(scores)-1]
```

Yukarıdaki, sıralanmamış bir dilimden bir değeri kaldırmanın etkili bir yolunun başlangıcıdır:

```go
func main() {
  scores := []int{1, 2, 3, 4, 5}
  scores = removeAtIndex(scores, 2)
  fmt.Println(scores) // [1 2 5 4]
}

// sıralamayı korumaz
func removeAtIndex(source []int, index int) []int {
  lastIndex := len(source) - 1
  //son değer ile çıkarmak istediğimiz değeri yer değiştirir
  source[index], source[lastIndex] = source[lastIndex], source[index]
  return source[:lastIndex]
}
```

Son olarak, dilimlerle ilgili bir çok şey bildiğimize göre, yaygın olarak kullanılan başka bir yerleşik işleve bakabiliriz: `copy` . `copy`, dilimlerin kodlama şeklimizi etkisini güçlendiren işlevlerden biridir. Normalde, değerleri bir diziden diğerine kopyalayan bir işlemde 5 parametreye ihtiyaç vardır: `source` , `sourceStart` , `count` , `destination` ve `destinationStart` . Dilimlerle ise sadece iki taneye ihtiyacımız var:

```go
import (
  "fmt"
  "math/rand"
  "sort"
)

func main() {
  scores := make([]int, 100)
  for i := 0; i < 100; i++ {
    scores[i] = int(rand.Int31n(1000))
  }
  sort.Ints(scores)

  worst := make([]int, 5)
  copy(worst, scores[:5])
  fmt.Println(worst)
}
```

Biraz zaman ayırın ve yukarıdaki kodla biraz oynayın. Varyasyonları deneyin. Kopyalamayı, şöyle bir şeyle değiştirirseniz `copy(worst[2:4], scores[:5])` ne olur veya `5`'ten fazla veya daha az değeri `worst` içine kopyalamaya çalışırsanız ne olur?

## Eşlemeler (Map)

Go'daki eşlemeler, diğer dillerde hashtable veya sözlük adı verilen şeylerdir. Beklediğiniz gibi çalışırlar: bir anahtar ve değer tanımlarsınız ve eşlemeden değer alabilir, değer atayabilir ve değer silebilirsiniz.

Haritalar, dilimler gibi, `make` işleviyle oluşturulur. Bir örneğe bakalım:

```go
func main() {
  lookup := make(map[string]int)
  lookup["goku"] = 9001
  power, exists := lookup["vegeta"]

  // ekran 0, false yazar
  // 0 ön tanımlı değerdir
  fmt.Println(power, exists)
}
```

Anahtar sayısını elde etmek için `len` kullanırız. Bir değeri anahtarına dayalı olarak kaldırmak için ise `delete` kullanırız:

```go
// 1 döner
total := len(lookup)

// bir dönüş değeri yoktur. Olmayan bir anahtar ile de çalışır
delete(lookup, "goku")
```

Eşlemeler dinamik olarak büyür. Ancak, başlangıç boyutunu ayarlamak için `make` işlevine ikinci bir argüman verebilirsiniz:

```go
lookup := make(map[string]int, 100)
```

Eşlemede kaç tane elemana sahip olacağına dair bir fikriniz varsa, bir başlangıç boyutu tanımlamak performansa yardımcı olabilir.

Bir yapının veri alanı olarak bir eşlemeye ihtiyacınız olduğunda, şöyle tanımlarsınız:

```go
type Saiyan struct {
  Name string
  Friends map[string]*Saiyan
}
```

Yukarıdakileri yapıya değer atamanın bir yolu şöyledir:

```go
goku := &Saiyan{
  Name: "Goku",
  Friends: make(map[string]*Saiyan),
}
goku.Friends["krillin"] = ... //todo load or create Krillin
```

Go'da değerleri tanımlamanın ve başlatmanın başka bir yolu daha var. `make` gibi, bu yaklaşım da eşlemelere ve dizilere özgüdür. Bileşik bir değişmez olarak ilan edebiliriz:

```go
lookup := map[string]int{
  "goku": 9001,
  "gohan": 2044,
}
```

`range` anahtar kelimesiyle birleştirilmiş bir `for` döngüsü kullanarak bir eşleme üzerinde döngü yapabiliriz:

```go
for key, value := range lookup {
  ...
}
```

Eşleme üzerinde yineleme sıralı değildir. Arama üzerindeki her yineleme, anahtar değer çiftini rastgele bir sırayla döndürür.

## İşaretçiler ve Değerler

İşaretçileri veya değerleri atamanız ve geçmeniz gerekip gerekmediğini tartışarak Bölüm 2'yi bitirdik. Şimdi dizi ve eşleme değerleri için de aynı tartışmayı yapacağız. Bunlardan hangilerini kullanmalısınız?

```go
a := make([]Saiyan, 10)
// ya da
b := make([]*Saiyan, 10)
```

Birçok geliştirici, `b`'i bir işleve geçmenin veya bir işlevden geri döndürmenin daha verimli olacağını düşünmektedir. Ancak, dilimin kendisi bir işaretçi olduğu için kopyası da bir işaretçidir. Dilimin kendisinin geçilmesinin ve ya geri döndürülmesinin bu açıdan bir farkı yoktur.

Farkı göreceğiniz yer, bir dilim veya eşlemenin değerlerini değiştirdiğiniz zamandır. Bu noktada, Bölüm 2'de gördüğümüz aynı mantık geçerlidir. Dolayısıyla, bir değer dizisine karşı bir işaretçi dizisinin tanımlanıp tanımlanmayacağına karar vermek, diziyi veya eşlemeyş nasıl kullandığınızla değil, tek tek değerleri nasıl kullandığınızla ilgilidir.

## Devam Etmeden Önce

Go'daki diziler ve eşlemeler diğer dillerde olduğu gibi çalışır. Dinamik dizilere alışkınsanız, küçük değişiklikler gerekebilir, ancak `append` rahatsızlığınızın çoğunu çözebilidir. Dizilerin yüzeysel sözdiziminin ve kullanımının ötesine bakarsak, dilimleri buluruz. Dilimler güçlüdür ve kodunuzun netliği üzerinde şaşırtıcı derecede büyük bir etkiye sahiptirler.

Bahsetmediğimiz bazı uç durumlar var, ancak bunlarla karşılaşmanız muhtemel değil. Ve eğer yaparsanız, umarım burada inşa ettiğimiz temeller neler olup bittiğini anlamanıza yardımcı olur.

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

# Sonuç

Son zamanlarda Go'nun *sıkıcı bir* dil olarak tanımlandığını duydum. Sıkıcı çünkü öğrenmesi kolay, yazması kolay ve en önemlisi okunması kolay. Belki de bu gerçeği göstererek bir kötülük yaptım. Üç bölümü tiplerden ve nasıl tanımlanacaklarından bahsederek *harcadık*.

Statik olarak yazılan bir dilde tecrübeniz varsa, gördüğümüz şeylerin çoğu muhtemelen en iyi ihtimalle size birer hatırlatma oldu. Go, işaretçileri görünür kılar ve diziler etrafındaki ince araçlar olarak verdiği dilimleri deneyimli Java veya C # geliştiricilerine zor gelmeyecektir.

Çoğunlukla dinamik dillerden faydalanıyorsanız, biraz farklı hissedebilirsiniz. Öğrenmesi *biraz* zor olabilir. Tanımlama ve değer atama ile ilgili çeşitli sözdizimi farkı göreceksiniz. Bir Go hayranı olmasına rağmen, basitliğe doğru tüm ilerlemeye rağman, bununla ilgili hala basit bir şey olmadığını düşünüyorum. Yine de, bazı temel kurallara (değişkenleri yalnızca bir kez bildirebileceğiniz ve `:=` değişkeni bildirdiğiniz gibi) ve temel anlayışa ( `new(X)` veya `&X{}` sadece bellek ayırır, ancak `make`  dilimler, eşlemeler ve kanallar için daha fazlasını gerektirir) değişiklik getirir.

Bunun ötesinde Go bize kodumuzu düzenlemenin basit ama etkili bir yolunu sunuyor. Arayüzler, dönüş değerine dayalı hata yönetimi, kaynak yönetimi için `defer` ve kompozisyon elde etmenin basit bir yolu gibi.

Son fakat bir o kadar önemli, eşzamanlılık için yerleşik desteğidir. Go rutinler hakkında etkili ve basit olmasından başka söylenecek çok az şey var (yine de kullanımı basit). İyi bir soyutlamadır. Kanallar daha karmaşıktır. Her zaman üst düzey sarmalayıcıları kullanmadan önce temel bilgileri anlamanın önemli olduğunu düşünüyorum. Ben kanallar olmadan eşzamanlı programlama öğrenmenin yararlı olduğunu *düşünüyorum*. Yine de, kanallar benim için basit bir soyutlama gibi hissetmeyecek şekilde uygulanmaktadır. Neredeyse kendi temel yapı taşlarıdır. Bunu söylüyorum çünkü eşzamanlı programlama hakkında yazma ve düşünme şeklinizi değiştiriyorlar. Eşzamanlı programlamanın ne kadar zor olabileceği göz önüne alındığında, bu kesinlikle iyi bir şeydir.
