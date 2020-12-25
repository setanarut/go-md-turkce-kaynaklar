- [Mikro-denetleyiciler](#mikro-denetleyiciler)
  - [Gobot ile Arduino Yanıp-Sönen LED Yapımı](#gobot-ile-arduino-yanıp-sönen-led-yapımı)
  - [Tinygo ile Küçük Yerler için Golang](#tinygo-ile-küçük-yerler-için-golang)
      - [GNU/Linux](#gnulinux)
      - [**Windows**](#windows)
      - [MacOS](#macos)
      - [Kurulum Sonrası](#kurulum-sonrası)

# Mikro-denetleyiciler

## Gobot ile Arduino Yanıp-Sönen LED Yapımı

Bu yazımda sizlere Golang için Robotik kütüphanesi olan **Gobot**‘tan bir örnek göstereceğim. Bu örneğimizde Arduino’da yanıp sönen LED yapacağız.  
İlk olarak Gobot kütüphanesini indiriyoruz.

> go get -d -u gobot.io/x/gobot/...

Daha sonra Arduino’muzla iletişim kurabilmemiz için **Gort**‘u yüklememiz gerekiyor.  
[https://gort.io/documentation/getting\_started/downloads/](https://gort.io/documentation/getting_started/downloads/)  
Bu örnekte **Arduino Uno** kullanacağız. Arduino’muzun bilgisayarımıza bağlıyoruz ve hangi porta bağlı olduğunu öğrenmek için komut satırına aşağıdakileri yazıyoruz.

> gort scan serial

Windows’ta **&lt;COM\*&gt;** benzeri, Linux’ta ise **/dev/ttyUSB\*** benzeri bir çıktı verecektir. Bu bizim Arduino’muzun bağlı olduğu portu gösteriyor.  
Aşağıdaki kodlar yanıp sönen LED için yazılmıştır. Kodları gördükten sonra açıklamasını yapacağım.

```go
package main
import (
        "time"
        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/firmata"
)
func main() {
        firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
        led := gpio.NewLedDriver(firmataAdaptor, "13")
        work := func() {
                gobot.Every(2*time.Second, func() {
                        led.Toggle()
                })
        }
        robot := gobot.NewRobot("bot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{led},
                work,
        )
        robot.Start()
}
```

Açıklamasına gelirsek;  
Gobot ile alakalı kütüphanelerimizi ekliyoruz. **firmataAdaptor** değişkenimizde Arduino’muzun portunu yazıyoruz. Ben Linux kullandığım için Linux’taki portunu yazdım. **led** değişkenimizde ledimizin **13**. dijital pinde yer aldığını belirttik. Yani LED’imizin artı ucunu **13. pine** eksi ucunu ise **GND** \(Ground-Toprak-Nötr\) girişine bağlayacağız.

Sıra geldi çalışma fonksiyonumuz olan **work**‘e. **work** değişkenine anonim bir fonksiyon tanımladık. Bu fonksiyonda **led.Toggle\(\)** fonksiyonu ile her **2** saniyede yanıp-sönmesini ayarladık. En sondaki **robot** değişkeninde ise firmataAdaptor değişkenimizdeki Arduino portuyla bağlantı kurmasını ve hemen altında **led** değişkenini cihaz olarak tanıttık. Son olarak **work** değişkenindeki olayları gerçekleştirip, **robot.Start\(\)** fonksiyonu ile çalışmasını sağladık.  
Yukarıda gördüğünüz üzere **Firmata** kelimesini kullandık. Firmata bizim Arduino cihazımız ile iletişimde bulunabilmemizi sağlayan bir yazılım. Yukarıdaki kodlarımızın çalışması için Arduino’muz içerisine Firmata yazılımını yüklememiz gerekir. Onu da yüklemesi aşağıdaki gibi çok kolay bir işlem.

> gort arduino upload firmata /dev/ttyUSB0

**/dev/ttyUSB0** yerine kendi Arduino portunuzu yazmayı unutmayın.  
Uygulamamızı başlatmak için ise aşağıdakileri yazıyoruz.

> go run main.go

**main.go** yerine ne isim verdiyseniz onu yazınız.

## Tinygo ile Küçük Yerler için Golang

![Tinygo Logosu](./goruntuler/tinygo-logo.png)

Tinygo, Golang kodları ile mikro-denetleyicilere program yazmamızı sağlayan bir derleyicidir.

Aynı zamanda yazdığımız kodları mikro-denetleyicinin beynine flash eder. Flash etme kelimesinden kastım, beyne çalışacak kodları yazdırmaktır.

gobot-ile-arduino-yanip-soenen-led-yapimi.md

Gobot ile Arduino Yanıp-Sönen LED Yapımı konusunda bahsettiğim. Gobot paketinden farkı, Gobot Firmata yazılımını Arduino’ya gömdükten sonra Arduino’ya çalıştırılabilir komutlar yolluyor. Yani kodlarımızı Arduino içine gömmediğinden, sadece Arduino USB veya TCP ile bağlı olduğundan çalışıyor.

Fakat Tinygo, Golang kodlarımızı Arduino’nun içerisine gömüyor. Bu sebeble Arduino’nun kodlarımızı çalıştırması için sadece bir elektrik kaynağına bağlı olması yetiyor.

**Gelelim Kuruluma**

#### GNU/Linux

Ubuntu/DebianBirinci Adım:  
`wget https://github.com/tinygo-org/tinygo/releases/download/v0.9.0/tinygo_0.9.0_amd64.deb`  
İkinci Adım:  
`sudo dpkg -i tinygo_0.9.0_amd64.deb`  
Üçüncü Adım:  
`export PATH=$PATH:/usr/local/tinygo/bin`RaspBerry PiBirinci Adım:  
`wget https://github.com/tinygo-org/tinygo/releases/download/v0.9.0/tinygo_0.9.0_armhf.deb`  
İkinci Adım:  
`sudo dpkg -i tinygo_0.9.0_armhf.deb`  
Üçüncü Adım:  
`export PATH=$PATH:/usr/local/tinygo/bin`Arch LinuxAUR deposundan [tinygo-bin](https://aur.archlinux.org/packages/tinygo-bin/) olarak aratabilirsiniz.Fedora Linux`sudo dnf install tinygo`

#### **Windows**

Öncelikle şuanda Windows üzerinde deneme aşamasında olduğunu söylemeliyim.

İlk olarak LLVM 8’i kurmalısınız.

[Buradan indirme sayfasına gidebilirsiniz.](http://releases.llvm.org/download.html#8.0.1)

LLVM 8 kurulumu esnasında “LLVM’yi ortam değişkenlerine ekle” seçeneğini seçmeyi unutmayın.

Daha sonra Tinygo arşiv dosyasını indirelim.

[Tinygo Arşiv Dosyası İndir](https://github.com/tinygo-org/tinygo/releases/download/v0.9.0/tinygo0.9.0.windows-amd64.zip)

Aşağıdaki komut ile Tinygo’yu kuralım.

`PowerShell Expand-Archive -Path "c:\Downloads\tinygo0.9.0.windows-amd64.zip" -DestinationPath "c:\tinygo"`

Aşağıdaki komut ile Tinygo’yu ortam değişkenlerine ekleyelim.

`set PATH=%PATH%;C:\tinygo\bin;`

#### MacOS

İlk adım:

​`brew tap tinygo-org/tools`

İkinci Adım:

`brew install tinygo`

#### Kurulum Sonrası

Kurulum işlemlerimiz tamamlandıktan sonra kontrol etme amaçlı komut satırına aşağıdaki komutları yazalım.

`tinygo version`

Kullandığınız işletim sistemine göre fark göstermekle birlikte aşapıdakine benzer bir çıktı alacaksınız.

`tinygo version 0.9.0 linux/amd64 (using go version go1.12.9)`

Bu işlemler sırasında elimde bulunan Arduino Uno kartı ile işlemler yapacağımız belirteyim. Diğer kartlar ile arasında çok bir işlem farkı bulunmamaktadır. Aynı veya benzer yollardan sizce bu işlemleri gerçekleştirebilirsiniz.

Öncelikle Arduino Uno kartımızın hangi USB portuna bağlı olduğunu bulalım.

Windows üzerinden **COM3** benzeri bir portta takılıdır. İnternet üzerinden detaylı araştırma yapabilirsiniz.

Unix-like sistemlerde \(Linux, MacOS\) ise genelde **/dev/ttyUSB** veya **/dev/ttyACM** portarından birinde takılı olabilir. Arduino’nun bağlı olduğu portu `ls /dev/ttyUSB*` komutu ile öğrenebilirsiniz.

Ben Arduino Uno kartımın **/dev/ttyUSB0** üzerinde olduğu için aşağıdaki işlemlerimi ona göre yapacağım. Kullandığım komutları kendi portunuza göre değiştirmeyi unutmayın.

Aşağıda Arduino UNO üzerindeki Built-In LED’i saniyede bir yanıp-söndürmeye yarayan Golang kodları yer alıyor.

```go
package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 1000)

		led.High()
		time.Sleep(time.Millisecond * 1000)
	}
}
```

Dosyamızın ismini main.go yapalım. Yukarıdaki Golang kodlarımızı kaydettikten sonra komut satırını main.go dosyasının bir üst klasöründe açalım.

Go kodlarımızı Arduino üzerine yazdırmak için aşağıdaki komutları kullanalım.

`tinygo flash -target=arduino -port=/dev/ttyUSB0 ./kodumuzunbulunduğuklasör`

Gördüğünüz gibi Tinygo ile flash etme işlemi çok basit.

