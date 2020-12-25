- [Bölüm 6 - Paketler](#bölüm-6---paketler)
	- [Strings](#strings)
	- [os/exec \(Komut Satırına Erişim\)](#osexec-komut-satırına-erişim)
	- [Komut Satırı Argümanları \(Args\)](#komut-satırı-argümanları-args)
	- [Komut Satırı Bayrakları \(Flags\)](#komut-satırı-bayrakları-flags)
	- [description: 'https://golang.org/pkg/os/signal/ sayfasından çeviridir.'](#description-httpsgolangorgpkgossignal-sayfasından-çeviridir)
	- [os/signal](#ossignal)
		- [Sinyal Türleri](#sinyal-türleri)
		- [Go Programlarında Sinyallerin Varsayılan Davranışı](#go-programlarında-sinyallerin-varsayılan-davranışı)
		- [Go Programlarında Sinyallerin Davranışını Değiştirme](#go-programlarında-sinyallerin-davranışını-değiştirme)
		- [Windows](#windows)
		- [Plan9](#plan9)
		- [Örnek Uygulama](#örnek-uygulama)
		- [Örnek: Tüm Sinyalleri Yakalamak](#örnek-tüm-sinyalleri-yakalamak)
	- [Sort \(Sıralama\)](#sort-sıralama)
	- [Strconv \(String Çeviri\)](#strconv-string-çeviri)
	- [Log \(Kayıt\)](#log-kayıt)
	- [Paket \(Kütüphane\) Yazmak](#paket-kütüphane-yazmak)
			- [Bir Paketin Özellikleri](#bir-paketin-özellikleri)
			- [Proje Klasöründe Yerel Kütüphane Oluşturma](#proje-klasöründe-yerel-kütüphane-oluşturma)
			- [Git Sisteminde Kütüphane Paylaşımı](#git-sisteminde-kütüphane-paylaşımı)

# Bölüm 6 - Paketler

## Strings

Strings paketi ile **string** türünde değerler üzerinde işlemler yapabiliriz. Kısaca kullanımlarından bahsedelim.  


**Strings.Contains\(\) Fonksiyonu**  
**Contains\(\)** fonksiyonu ile istediğimiz bir string değerin içerisinde istediğimiz bir **string** değerin olup olmadığını kontrol edebiliriz. **Boolean** değer verir. Eğer varsa **true** değer döndürür. Yoksa **false** değer döndürür. Ufak bir uygulama örneği yapalım.

```go
package main
import (
 "fmt"
 "strings"
)
func main() {
 var eposta string
 fmt.Print("E-posta adresinizi giriniz: ")
 fmt.Scan(&eposta)
 if strings.Contains(eposta, "@") {
  fmt.Println("E-posta Adresiniz Onaylandı!")
 } else {
  fmt.Println("Geçerli Bir E-posta Adresi Giriniz!")
 }
}
```

**”strings”** paketini eklemeyi unutmuyoruz. Bu kodlar ile kullanıcıdan e-posta adresi isteyen ve e-posta adresi içinde **@** işareti var ise olumlu yanıt veren bir programcık oluşturduk. **Constains\(\)** fonksiyonunu açıklayacak olursak, **Contains** fonksiyonunun ilk parametresine kontrol edeceğimiz öğeyi giriyoruz. İkinci parametreye ise aranılacak **string** ifademizi giriyoruz. Gayet anlaşılır olduğunu düşünüyorum.  
**Strings.Count\(\) Fonksiyonu**  
**Count\(\)** fonksiyonu ile bir string değerin içinde istediğimiz bir string değerin kaç tane olduğunu öğrenebiliriz. Örneğimize geçelim.

```go
package main
import (
 "fmt"
 "strings"
)
func main() {
 fmt.Println(strings.Count("deneme", "e"))
}
```

**”strings”** paketini eklemeyi unutmuyoruz. Bu kodlar ile **Count\(\)** fonksiyonunda **“deneme”** stringi içerisinde **“e”** stringinin kaç tane geçtiğini öğreniyoruz. Çıktımız **3** olacaktır.  
**Strings.Index\(\) Fonksiyonu**  
**Index\(\)** fonksiyonu ile bir **string** değerin içindeki istediğimiz bir string değerin kaçıncı sırada yani **index**’te olduğunu öğrenebiliriz.Sıra sıfırdan başlar. Örneğimize geçelim.

```go
package main
import (
 "fmt"
 "strings"
)
func main() {
 fmt.Println(strings.Index("Merhaba Dünya", "h"))
}
```

Çıktımız **h** harfi **0**’dan başlayarak 3. sırada olduğu için, **3** olacaktır.  
**Strings.LastIndex\(\) Fonksiyonu**  
**LastIndex\(\)** fonksiyonu ile bir **string** değerin içinde istediğimiz bir string değerin sırasını **Index\(\)** fonksiyonunun tersine sağdan sola doğru kontrol eder. İlk çıkan sonucun index’ini seçer. Örnek:

```go
fmt.Println(strings.LastIndex("Merhaba Dünya", "a"))
```

**”Merhaba Dünya”** yazısının içinde **“a”** harfini aradık. **LastIndex\(\)** fonksiyonu sondan başa yani sağdan sola arama yaptığı için sondaki **“a”** harfini buldu. Yani **13** sonucunu ekrana bastırmış olduk.  
**Strings.Title\(\) Fonksiyonu**  
**Title\(\)** fonksiyonu ile içerisine küçük harflerle string türünde değer girdiğimizde baş harfleri büyük harf yapan bir fonksiyondur.

```go
fmt.Println(strings.Title("merhaba dünya"))
```

Çıktımız **“Merhaba Dünya”** olacaktır.  
**Strings.ToUpper\(\) Fonksiyonu**  
**ToUpper\(\)** fonksiyonu içerisine girilen string değerin tüm harflerini büyük harf yapar.

```go
fmt.Println(strings.ToUpper("merhaba dünya"))
```

Çıktımız **“MERHABA DÜNYA”** olacaktır.  
**Strings.ToLower\(\) Fonksiyonu**  
**ToLower\(\)** fonksiyonu içerisine girilen string değerin tüm harflerini küçük harf yapar.

```go
fmt.Println(strings.ToLower("Merhaba Dünya"))
```

Çıktımız **“merhaba dünya”** olacaktır.  
**Strings.ToUpperSpecial\(\) Fonksiyonu**  
**ToUpper\(\)** fonksiyonu ile string değeri büyük harf yaptığımız zaman Türkçe karakter sıkıntısı yaşarız. Örnek olarak **“i”** harfi büyüyünce **“I”** harfi olur. Bunun önüne **ToUpperSpecial\(\)** fonksiyonu ile geçebiliriz. Bu fonksiyonun ilk parametresine **karakter kodlamasını**, ikinci parametresine ise **string** değerimizi gireriz. Örnek olarak:

```go
fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "ıiüöç"))
```

Çıktımız **“IİÜÖÇ”** olacaktır.  
**Strings.ToLowerSpecial\(\) Fonksiyonu**  
**ToUpperSpecial\(\)** fonksiyonu ile aynı seçilde çalışır ;fakat harfleri belirlediğiniz karakter kodlamasına göre küçültür. Örnek kullanımı:

```go
fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "IİÜÖÇ"))
```

Çıktımız **“ıiüöç”** olacaktır.

## os/exec \(Komut Satırına Erişim\)

**os/exec** paketi komut satırına \(cmd, powershell, terminal\) komut göndermemizi sağlayan Golang ile bütünleşik gelen bir pakettir. Bu paket sayesinde oluşturacağımız programa sistem işlerini yaptırabiliriz. Örnek olarak dosya/klasör taşıma/silme/oluşturma/kopyalama gibi işlemleri yaptırabilir. Daha doğrusu komut satırı/terminal üzerinden yapabildiğimiz her işlemi yaptırabiliriz. Tabi kullandığımız işletim sistemine göre terminal komutları değiştiği için ona göre örnek vermeye çalışacağım.

  
**Örnek 1: Komut Satırına Komut Gönderme**  
Ufak bir örnek ile başlayalım.

```go
package main
import (
    "os"
    "os/exec"
)
func main() {
    cmd := exec.Command("mkdir", "klasörüm")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
```

**“mkdir klasörüm”** komutu programın çalıştırıldığı dizinde **“klasörüm”** adında bir klasör oluşturur. Komut girerken dikkat etmeniz geren çok önemi bir detay var. Yazacağınız komut birden fazla kelimeden oluşuyorsa mutlaka ayrı ayrı girmelisiniz. Eğer **exec.Command\(\)** fonksiyonuna direkt olarak **“mkdir klasörüm”** olarak girseydik, komutu tek kelime olarak algılayacaktı. Yani **string dizisi** mantığında çalışıyor bu olay. Sonuç olarak yukarıdaki gibi basit bir şekilde komut satırına komut yollayabilirsiniz.  
**Örnek 2: Komut Satırına Komut Gönderip Çıktısını Okuma**  
Yukarıda çok kolay bir şekilde komut göndermeyi gördük. Fakat iş komutun çıktısını okumaya gelince işler biraz karışıyor. Yavaştan vaziyetinizi alın  ![Wink](https://i1.wp.com/golangtr.org/forum/images/smilies/wink.png?w=770)  
Aslında korkulacak bir olay yok. Yeter ki mantığını anlayalım. Şimdi yapacağımız işlemleri 4 ana parçaya bölelim.

1. Komutun tanımlanması
2. Çıktı okuyucusunun tanımlanması
3. Komutun başlatılması
4. Komutun çalışması

Hemen kodlarımıza geçelim.

```go
package main
import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
)
func main() {
    //komutun tanımlanması
    cmd := exec.Command("go", "version")
    cmdOkuyucu, hata := cmd.StdoutPipe()
    if hata != nil {
        fmt.Fprintln(os.Stderr, "Çıktı okunurken hata oluştu:", hata)
        os.Exit(1)
    }
    //çıktı okuyucusunun tanımlanması
    çıktı := bufio.NewScanner(cmdOkuyucu)
    go func() {
        for çıktı.Scan() {
            fmt.Println(çıktı.Text())
        }
    }()
    //komutun başlatılması
    hata = cmd.Start()
    if hata != nil {
        fmt.Fprintln(os.Stderr, "Komut başlatılamadı:", hata)
        os.Exit(1)
    }
    //komutun çalışması
    hata = cmd.Wait()
    if hata != nil {
        fmt.Fprintln(os.Stderr, "Komut çalışırken hata oluştu:", hata)
        os.Exit(1)
    }
}
```

Gelelim yukarıdaki kodların açıklamasına…  
**cmd** adında bir değişken oluşturduk. Bu değişkenimiz sayesinde **exec.Command\(\)** fonksiyonuyla komutlarımızı girdik.  
**cmd.StdoutPipe\(\)** fonksiyonuyla gönderdiğimiz komutun çıktılarını alabiliyoruz. **cmdOkuyucu** değişkenine komut çıktımızı aldık. **hata** değişkenimize ise komut girildiğinde oluşan hata mesajını aldık.  
**hata** değişkeninin içi boş değilse ekrana bastırmasını ve **1** numaralı çıkış kodunu vermesini istedik. Bu arada 1 numaralı çıkış kodu hatalar için kullanılır. Golang programlarında görmüyoruz ama **0** numaralı çıkış kod da işler yolunda gittiği zaman kullanılır. C dili kodlayan arkadaşlarımız bilir, **int main** fonksiyonunun sonuna **return 0** ibaresi girilir. Buraya kadar olan işlemlerimiz komutun tanımlanması ile ilgiliydi.  
Çıktımızı okuyabilmemiz için birkaç işlem yapmamız gerekiyor. Ne yazık ki çıktımızı direkt olarak değişkene atayıp ekrana bastıramıyoruz. **çıktı** adında değişkenimizi oluşturuyoruz. Bu değişkenimiz **cmdOkuyucu** değişkenini taramaya yarayacak. Hemen aşağısında goroutine fonksiyonumuzda **çıktı.Scan\(\)** döngüsü ile çıktı sonucumuzu ekrana bastırıyoruz.  
Buraya kadar tanımlamalarımız yapmış bulunduk. Bundan sonra işlemlerimiz komutumuzun çalıştırılması ve sonucun beklenmesi olacak.  
**hata** değişkenimize **cmd.Start\(\)** fonksiyonunu atayarak komut başlatma işleminde hata oluşursa veriyi çekmesini sağladık. Hata var ise **error** tipindeki hata mesajımızı ekrana ve 1 numaralı hatayı ekrana bastıracak.  
Son işlemimiz ise komutun sonuçlanmasının beklenmesi. **hata** değişkenimize **cmd.Wait\(\)** fonksiyonunu ekleyerek bekleme işleminde oluşabilecek hatanın mesajını çekmiş olduk. Aşağısında eğer hata var ise ekrana bastırması için gerekli kodlarımızı girdik. Son olarak **1** numaralı çıkış işlemini yaptık.  
Gördüğünüz gibi çıktı alma işlemi biraz daha uzun. Ama mantığını anladıktan sonra kolay bir işlem olduğunu düşüyorum.  
**Örnek 3: Hata Detayı Çekmeden Komut Çıktısı Alma**  
Eğer ben hata çıktısının detayını almak istemiyorum, benim işim sadece çıktıyla diyorsanız yapacağımız işlemler gerçekten kolaylaşıyor. Hemen kodlarımızı görelim.

```go
package main
import (
    "fmt"
    "log"
    "os/exec"
)
func main() {
    cmd := exec.Command("go", "versison")
    çıktı, hata := cmd.CombinedOutput()
    if hata != nil {
        log.Fatalf("Komut hatası: %s\n", hata)
    }
    fmt.Printf(string(çıktı))
}
```

Kodlarımızın açıklamasına geçelim. **cmd** adında değişkenimizde **exec.Command\(\)** fonksiyonu ile komutlarımızı tanımladık. **çıktı** ve **hata** değişkenimize komut çıktılarımızı aldık. Burada **hata** değişkeni sadece hata numarasını verecektir. Detayları barındırmaz. Eğer hatamız var ise ekrana bastırmasını istedik. Aşağısında ise **çıktı** değişkenimiz **byte dizisi** tipinde olduğu için **string**‘e çevirip ekrana bastırdık.

## Komut Satırı Argümanları \(Args\)

Golang ile programlarımızın komut satırı üzerinden argümanlar ile çalışmasını sağlayabiliriz. İşte paketimiz:

```text
import "os"
```

`os` paketimizdeki `Args` fonksiyonu bize string dizi sunar. ****Bir örnek görelim.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, "=", arg)
	}
}
```

`for-range` ile `os.Args`'ın uzunluğu kadar işlem yapıyoruz ve içerisindekileri indeksi ile ekrana bastırıyoruz. Şöyle bir çıktımız oluyor:

> ./main naber nasılsın
>
> 0 = ./main  
> 1 = naber  
> 2 = nasılsın

## Komut Satırı Bayrakları \(Flags\)

Komut satırı bayrakları, örnek olarak;

> ./uygulamamız -h

Sondaki `-h` bir flag\(bayrak\)’dir. Örnek bir program yazalım.

```go
package main
import (
	"flag"
	"fmt"
)
func main() {
	kelime := flag.String("kelime", "varsayılan kelime", "metin tipinde")
	sayi := flag.Int("sayi", 1881, "sayı tipinde")
	mantiksal := flag.Bool("mantiksal", false, "boolean tipinde")
	flag.Parse()
	fmt.Println("kelime:", *kelime)
	fmt.Println("sayi:", *sayi)
	fmt.Println("mantiksal:", *mantiksal)
}
```

Gelelim açıklamasına;  
**kelime** isminde **string** tipinde bir flag oluşturduk. **flag.String\(\)** fonksiyonu içerisinde 1. parametre komut satırından **“-kelime”** argümanıyla gireceğimizi gösteriyor. Varsayılan değeri **“varsayılan kelime”** olacak ve açıklama bölümünde **“metin tipinde”** yazacak.  
**sayi** isminde **int** tipinde bir flag oluşturduk. **flag.Int\(\)** fonksiyonu içerisinde komut satırından **“-sayi”** argümanıyla gireceğimizi belirttik. Varsayılan değeri **1881** olacak ve açıklama bölümünde **“sayı tipinde”** yazacak.  
**mantiksal** isminde **bool** tipinde bir flag oluşturduk. **flag.Bool\(\)** fonksiyonunda **“-mantiksal”** argumanıyla çağırılacağını belirttik. Varsayılan değeri **false** olacak ve açıklama bölümünde **“boolean tipinde”** yazacak.  
Uygulamamızı build edelim ve ismi **uygulama** olsun.

> go build -o ./uygulama .

Windows için build ediyorsanız, `./uygulama` yerine `./uygulama.exe` yazarak build edin. \(Hatırlatma yapayım dedim\) Build ettikten sonra örnek bir kullanımını yapalım.

> ./uygulama -kelime=Atatürk -sayi=1881 -mantiksal=true

Çıktımız şu şekilde olacaktır.

> kelime: Atatürk  
> sayi: 1881  
> mantiksal: true

Peki bu girdiğimiz flag açıklamaları ne oluyor diye soracak olursanız eğer, onu da aşağıdaki komutu yazarak görebilirsiniz.

> ./uygulama -h

Çıktımız şu şekilde olacaktır.

> Usage of ./uygulama:  
>   -kelime string  
>         metin tipinde \(default "varsayılan kelime"\)  
>   -mantiksal  
>         boolean tipinde \(default false"\)  
>   -sayi int  
>         sayı tipinde \(default 1881\)

---
description: 'https://golang.org/pkg/os/signal/ sayfasından çeviridir.'
---

## os/signal

**os/signal** paketi gelen sinyallere erişim sağlar. Genellikle Unix-benzeri sistemlerde kullanılır. **Windows** ve **Plan9**'da kullanımı farklıdır.

### Sinyal Türleri

**SIGKILL** ve **SIGSTOP** sinyalleri bir program tarafından yakalanmayabilir ve bu nedenle bu paketten etkilenemez.

Senkron sinyaller, program yürütmedeki hatalarla tetiklenen sinyallerdir: **SIGBUS**, **SIGFPE** ve **SIGSEGV**. Bunlar, `os.Process.Kill` veya **kill** programı veya benzer bir mekanizma kullanılarak gönderildiklerinde değil, yalnızca program yürütülmesinden kaynaklandığında eşzamanlı olarak kabul edilir. Genel olarak, aşağıda tartışılanlar dışında, Go programları eşzamanlı bir sinyali çalışma zamanı paniğine dönüştürecektir.

Kalan sinyaller asenkron sinyallerdir. Program hataları tarafından tetiklenmezler, bunun yerine çekirdekten veya başka bir programdan gönderilirler.

Asenkron sinyallerden, **SIGHUP** sinyali, bir program kontrol terminalini kaybettiğinde gönderilir. **SIGINT** sinyali, kontrol terminalindeki kullanıcı, varsayılan olarak ^ C \(Kontrol-C\) olan kesme karakterine bastığında gönderilir. **SIGQUIT** sinyali, kontrol terminalindeki kullanıcı varsayılan olarak ^  \(Kontrol-Ters Taksim\) olan çıkış karakterine bastığında gönderilir. Genel olarak, bir programın ^ C'ye basarak çıkmasına neden olabilir ve ^  tuşuna basarak bir yığın dökümü ile çıkmasına neden olabilirsiniz.

### Go Programlarında Sinyallerin Varsayılan Davranışı

Varsayılan olarak, senkronize bir sinyal çalışma zamanı paniğine dönüştürülür. **SIGHUP**, **SIGINT** veya **SIGTERM** sinyali programın çıkmasına neden olur. **SIGQUIT**, **SIGILL**, **SIGTRAP**, **SIGABRT**, **SIGSTKFLT**, **SIGEMT** veya **SIGSYS** sinyali, programın yığın dökümü ile çıkmasına neden olur. Bir **SIGTSTP**, **SIGTTIN** veya **SIGTTOU** sinyali sistem varsayılan davranışını alır \(bu sinyaller kabuk tarafından iş kontrolü için kullanılır\). **SIGPROF** sinyali, `runtime.CPUProfile`'ı uygulamak için doğrudan Go çalışma zamanı tarafından işlenir. Diğer sinyaller yakalanacak ancak herhangi bir işlem yapılmayacaktır.

Go programı, **SIGHUP** veya **SIGINT** göz ardı edilerek başlatılırsa \(sinyal işleyici **SIG\_IGN**'a ayarlı\), bunlar ihmal edilmiş olarak kalacaktır.

Go programı boş olmayan bir sinyal maskesi ile başlatılırsa, bu genellikle kabul edilir. Bununla birlikte, bazı sinyaller açıkça engellenmiştir: eşzamanlı sinyaller, **SIGILL**, **SIGTRAP**, **SIGSTKFLT**, **SIGCHLD**, **SIGPROF** ve **GNU/Linux**'ta **32 \(SIGCANCEL\)** ve **33 \(SIGSETXID\)** **\(SIGCANCEL ve SIGSETXID\)** sinyalleri **glibc** tarafından dahili olarak kullanılır. `Os.Exec` veya `os/exec` paketi tarafından başlatılan alt işlemler, değiştirilmiş sinyal maskesini miras alır.

### Go Programlarında Sinyallerin Davranışını Değiştirme

Bu paketteki işlevler, bir programın Go programlarının sinyalleri işleme şeklini değiştirmesine izin verir.

**Notify**, belirli bir eşzamansız sinyal kümesi için varsayılan davranışı devre dışı bırakır ve bunun yerine bunları bir veya daha fazla kayıtlı kanal üzerinden iletir. Özellikle, **SIGHUP**, **SIGINT**, **SIGQUIT**, **SIGABRT** ve **SIGTERM** sinyalleri için geçerlidir. Bu aynı zamanda iş kontrol sinyalleri **SIGTSTP**, **SIGTTIN** ve **SIGTTOU** için de geçerlidir ve bu durumda sistem varsayılan davranışı oluşmaz. Aynı zamanda, başka şekilde hiçbir eyleme neden olmayan bazı sinyaller için de geçerlidir: **SIGUSR1**, **SIGUSR2**, **SIGPIPE**, **SIGALRM**, **SIGCHLD**, **SIGCONT**, **SIGURG**, **SIGXCPU**, **SIGXFSZ**, **SIGVTALRM**, **SIGWINCH**, **SIGIO**, **SIGPWR**, **SIGSIGTHEZW**, **SIGTHAW**, **SIGLOST**, **SIGXRES**, **SIGJVM1**, **SIGJVM2** ve sistemde kullanılan gerçek zamanlı sinyaller. Bu sinyallerin tümünün tüm sistemlerde mevcut olmadığını unutmayın.

Program **SIGHUP** veya **SIGINT** göz ardı edilerek başlatılmışsa ve her iki sinyal için de **Notify** çağrılırsa, bu sinyal için bir sinyal işleyici kurulacak ve artık göz ardı edilmeyecektir. Daha sonra bu sinyal için **Reset** veya **Ignore** çağrılırsa veya o sinyal için **Notify**'ye iletilen tüm kanallarda **Stop** çağrılırsa, sinyal bir kez daha yok sayılır. **Reset**, sinyal için sistemin varsayılan davranışını geri yüklerken, **Ignore**, sistemin sinyali tamamen yok saymasına neden olur.

Program boş olmayan bir sinyal maskesi ile başlatılırsa, bazı sinyallerin blokajı yukarıda açıklandığı gibi açıkça kaldırılacaktır. Engellenen bir sinyal için **Notify** çağrılırsa, engellemesi kaldırılır. Daha sonra bu sinyal için **Reset** çağrılırsa veya bu sinyal için **Notify**'ye iletilen tüm kanallarda **Stop** çağrılırsa, sinyal bir kez daha engellenecektir.

### Windows

Windows'ta a ^ C \(Control-C\) veya ^ BREAK \(Control-Break\) normalde programın çıkmasına neden olur. `Os.Interrupt` için **Notify** çağrılırsa, ^ C veya ^ BREAK `os.Interrupt`'ın kanala gönderilmesine neden olur ve program çıkmaz. **Notify**'ye geçen tüm kanallarda **Reset** çağrılırsa veya **Stop** çağrılırsa, varsayılan davranış geri yüklenir.

### Plan9

Plan 9'da, sinyaller bir dizge olan `syscall.Note` türüne sahiptir. **Notify** ile bir sistem çağrısı çağırmak, bu dize bir not olarak gönderildiğinde bu değerin kanala gönderilmesine neden olur.

### Örnek Uygulama

main.go
```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	//os.Signal tipinde değer taşıyan bir kanal oluşturduk
	signalKanalı := make(chan os.Signal, 1)

	/*
	* Ana iş parçacığı sonlanmaması için bir kanal
	* oluşturalım.
	 */
	programBitti := make(chan bool)

	/*
	* os.Interrupt sinyali ile programramın sonlanması
	* yerine sinyali signalKanalı'na yönlendirelim.
	 */
	signal.Notify(signalKanalı, os.Interrupt)

	/*
	* asenkron olarak signalKanalı'nı dinleyelim. Sinyal
	* geldiğinde yani CTRL + C'ye basıldığında for döngüsü
	* içerisindeki kodlar çalışacak.
	 */
	go func() {
		for range signalKanalı {
			fmt.Println("Kontrol + C 'ye basıldı")

			//5 sn bekleyelim.
			time.Sleep(time.Second * 5)

			/*
			* Burada bekleyerek size programın CTRL + C'ye
			* basıldığında kapanmadığını gösteriyorum :)
			 */
			fmt.Println("bitti")

			/*
			* Kanala değer göndererek ana iş parçacığındaki
			* programBitti kanalının bekleyişine son verelim.
			 */
			programBitti <- true
		}

	}()

	//Ana iş parçacığı sonlanmasın diye kanalı bekleyelim
	<-programBitti
}

```


### Örnek: Tüm Sinyalleri Yakalamak

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	//Sinyallerin gönderileceği kanalı oluşturalım.
	kanal1 := make(chan os.Signal, 1)

	//Gelen sinyalleri kanal1'e yönlendirelim
	signal.Notify(kanal1)

	// kanal1'e sinyal gelene kadar programı bekletelim.
	sinyalTürü := <-kanal1
	fmt.Println("Sinyal Türü:", sinyalTürü)
}
```



## Sort \(Sıralama\)

Golang’ta dizilerin içeriğini sıralaya bileceğimiz bütünleşik olarak gelen **“sort**” isminde bir paket mevcuttur. Bu paketin kullanımı oldukça kolaydır. Örneğimizi görelim.

```go
package main
import (
	"fmt"
	"sort"
)
func main() {
	yazilar := []string{"c", "a", "b"}
	sort.Strings(yazilar)
	fmt.Println("Yazılar:", yazilar)
	sayilar := []int{7, 2, 4}
	sort.Ints(sayilar)
	fmt.Println("Sayılar:", sayilar)
	yazisirali := sort.StringsAreSorted(yazilar)
	fmt.Println("Yazılar Sıralandı mı?: ", yazisirali)
	sayisirali := sort.IntsAreSorted(sayilar)
	fmt.Println("Sayılar Sıralandı mı?:", sayisirali)
}
```

Gelelim açıklamasına;  
Sıralama özelliğini kullanabilmek için **“sort”** paketini içe aktardık. **main\(\)** fonksiyonumuzun içini inceleyelim.  
**yazilar** isminde içerisinde rastgele harflereden oluşan bir **string** dizi oluşturduk. Hemen aşağısında **sort.Strings\(yazilar\)** diyerek sıralamanın **string** türünde olduğunu belirterek sıralamamızı yaptık. Altında **yazilar** değişkenimizi ekrana bastırdık.

**sayilar** isminde içerisinde rastgele sayılar olan **int** tipinde bir dizi oluşturduk. Hemen aşağısında **sort.Ints\(sayilar\)** diyerek **int** tipinde sıralamamızı yaptık. Altında **sayilar** değişkenimizi ekrana bastırdık.  
Dizilerin sıralı olup olmadığını öğrenmek için de aşağıdaki işlemleri yaptık.  
**yazisirali** değişkeninde **sort.StringsAreSorted\(yazilar\)** fonksiyonu ile **yazilar** dizisinin sıralı olup olmama durumuna göre **bool** değer aldık. Ve sonucu ekrana bastırdık.  


**sayisirali** değişkeninde **sort.IntsAreSorted\(sayilar\)** fonksiyonu ile **sayilar** dizisinin sıralı olup olmama durumuna göre **bool** değer aldık. Ve sonucu ekrana bastırdık.  
Yukarıdaki işlemlere göre çıktımız şu şekilde olacaktır.

> Yazılar: \[a b c\]  
> Sayılar: \[2 4 7\]  
> Yazılar Sıralandı mı?: true  
> Sayılar Sıralandı mı?: true

## Strconv \(String Çeviri\)

**strconv** paketi Golang ile bütünleşik gelen string tipi ve diğer tipler arasında çevirme işlemi yapabileceğimiz bir pakettir.  


İlk olarak `“strconv”` paketimizi içe aktarıyoruz.  
Aşağıda örnek kullanımılarını ve daha açıklayıcı olması için yanlarına kullanım amaçlarını yazdım.

```go
package main
import (
	"fmt"
	"strconv"
)
func main() {
	//basit string-int arası çevirme
	sayi, _ := strconv.Atoi("-42") //string > int
	yazi := strconv.Itoa(-42)      //int > string
	//string'ten diğerlerine çevirme
	b, _ := strconv.ParseBool("true")        //string > bool
	f, _ := strconv.ParseFloat("3.1415", 64) //string > float
	i, _ := strconv.ParseInt("-42", 10, 64)  //string > int
	u, _ := strconv.ParseUint("42", 10, 64)  //string > uint
	//diğerlerinden string'e çevirme
	s1 := strconv.FormatBool(true)                 //bool > string
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64) //float > string
	s3 := strconv.FormatInt(-42, 16)               //int > string
	s4 := strconv.FormatUint(42, 16)               //uint > string
	//Ekrana Yazdırma
	fmt.Printf("sayi: %d tip: %T\n", sayi, sayi)
	fmt.Printf("yazi: %s tip: %T\n", yazi, yazi)
	fmt.Printf("b: %t tip: %T\n", b, b)
	fmt.Printf("f: %f tip: %T\n", f, f)
	fmt.Printf("i: %d tip: %T\n", i, i)
	fmt.Printf("u: %d tip: %T\n", u, u)
	fmt.Printf("%T %T %T %T", s1, s2, s3, s4)
}
```

Çıktımız şu şekilde olacaktır.

> sayi: -42 tip: int  
> yazi: -42 tip: string  
> b: true tip: bool  
> f: 3.141500 tip: float64  
> i: -42 tip: int64  
> u: 42 tip: uint64  
> string string string string

## Log \(Kayıt\)

**Log** paketi standart Golang paketleri içerisinde gelir ve programdaki olayları kaydetmemizi yarayacak bir altyapı sunar. Log programcının gözü kulağıdır. Bize hataları \(bugs\) bulmamız için kolaylık sağlar. Örneğimize geçelim.

```go
package main
import (
    "log"      
)
func init(){
    log.SetPrefix("KAYIT: ")
    log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
    log.Println("ön yükleme tamamlandı")
}
func main() {
    log.Println("main fonksiyonu başladı")
 
    log.Fatalln("ölümcül hata")
 
	log.Panicln("panic mesajı")
}
```

Hemen açıklamasına geçelim. İlk olarak **log** paketimizi içe aktarıyoruz. **init\(\)** fonksiyonunda log paketimiz ile ilgili ön ayarları yapıyoruz.

../boeluem-5/init-fonksiyonu-oen-yuekleme.md

**init\(\)** fonksiyonumuzun içerisini dikkatlice inceleyelim. log paketimizin üzerine ayarlamalar yapıyoruz.  


**SetPrefix\(\)** fonksiyonu ile log çıktımızın satırının başında ne yazacağını belirleyebiliyoruz.  


**SetFlags \(\)** fonksiyonu ile log çıktımızın görünüşünü ayarlıyoruz. **log.Ldate** bize zamanını gösteriyor. **log.Lmicroseconds** mikrosaniyeyi ve **log.Llongfile** ise dosya ismini ve yapılan işlem ile ilgili satırı gösteriyor.  


log önyüklemizi yaptığımızı opsiyonel olarak **log.Println\(\)** ile belirtiyoruz.  


**main\(\)** fonksiyonumuzun içerisini incelediğimizde ise;  
l**og.Println\(\)** fonksiyonu ile klasik log çıktılama işlemini yapıyoruz. Fonksiyonun sonundaki **ln** bir alt satıra geçildğini gösteriyor.  


**log.Fatalln\(\)** fonksiyonu ile kritik hataları bildirir. log.Println\(\) fonksiyonundan farkı program **1 çıkış kodu** ile biter. Bu da programın hatalı bittiği anlamına gelir. Normalde sağlıklı çalışan bir program **0 çıkış kodu** ile biter. 0 çıkış kodunu Golang programlama da kullanmamıza gerek kalmaz. Fakat C gibi dillerde ana fonksiyonun sonunda **return 0** ibaresini yazmak zorundayız.  
**log.Panicln\(\)** fonksiyonunda ise ekrana çıktımızı verir ve aynı zamanda bunu normal **panic\(\)** fonksiyonu ile yapar.

  
Çıktımız ise şöyle olacaktır:

> KAYIT: 2019/10/10 20:29:14.107438 /home/ksc10/Desktop/deneme/main.go:10: ön yükleme tamamlandı  
> KAYIT: 2019/10/10 20:29:14.107529 /home/ksc10/Desktop/deneme/main.go:13: main fonksiyonu başladı  
> KAYIT: 2019/10/10 20:29:14.107539 /home/ksc10/Desktop/deneme/main.go:15: ölümcül hata  
> exit status 1

Gördüğünüz gibi son satırda **çıkış durumunun 1** olduğunu yazıyor.  
panic mesajı programı direkt sonlandırır. panic mesajını daha üste yazarak deneyebilirsiniz.

## Paket \(Kütüphane\) Yazmak

Bu bölümde Go üzerinde nasıl kendi paketimizi \(kütüphanemizi\) oluşturacağımıza bakacağız.

#### Bir Paketin Özellikleri

* İçerisinde `.go` dosyaları bulunan bir klasördür.
* Diğer projeler tarafından içe aktarılabilir.
* Dışa aktarılabilen veya aktarılamayan veriler içerir.
* Açık kaynaktır.
* `main()` fonksiyonu içermez.
* `package main` değildir.

Paket oluştururken dikkat etmemiz gereken prensipler vardır. Bunlar şık ve basit kod yazımı, dışarıdan kullanımı basit, mümkün olduğunca diğer paketlere bağımsız olmasıdır. Bu prensiplere dikkat ederek daha iyi bir paket yazabilirsiniz.

#### Proje Klasöründe Yerel Kütüphane Oluşturma

Öncelikle aşağıdaki gibi bir dosya düzenimiz olduğunu varsayalım.

![Proje Klas&#xF6;r&#xFC;m&#xFC;z&#xFC;n Yap&#x131;s&#x131;](./goruntuler/package.png)

Yukarıdaki gibi `paketim` klasörü içerisinde `paketim.go` dosyamız olsun.

`paketim.go` dosyamızın içi aşağıdaki gibi olsun.

```go
package paketim

import "fmt"

func Yaz() {
	fmt.Println("yazdım!")
}
```

`package paketim` ile paketimizin ismini belirledik. Bu isim paket klasörümüz ile aynı olmalıdır. Daha sonra projemizde kullanabilmemiz için dışa aktarılmış şekilde `Yaz()` fonksiyonu oluşturduk. Bu fonksiyonun ne işe yaradığı zaten belli.

`main.go` dosyamız ise aşağıdaki gibi olsun.

```go
package main

import p "./paketim"

func main() {
	p.Yaz()
}
```

`import p "./paketim"` yazarak özel paketimizin yerel konumunu belirterek, `p` lakabı \(alias\) ile çağırdık.

`Yaz()` fonksiyonumuzu ise `p.Yaz()` şeklinde kullandık.

#### Git Sisteminde Kütüphane Paylaşımı

Oluşturduğumuz kütüphaneyi Github, Gitlab, Bitbucket vb. sitelerde barındırarak diğer geliştiricilerinde kütüphanelerinizden faydalanmasını sağlayabilirsiniz.

Bunun için kütüphanenizin isminde bir repo oluşturup, içerisinde Go dosyalarınızı yükleyin. Daha sonra `go get github.com/id/repoismi` şeklinde projenize import edebilirsiniz.

