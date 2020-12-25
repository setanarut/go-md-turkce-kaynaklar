- [Bölüm 4 - Eşzamanlılık](#bölüm-4---eşzamanlılık)
	- [Goroutine](#goroutine)
		- [Terimler](#terimler)
			- [Ana iş parçacığı](#ana-iş-parçacığı)
			- [Eşzamanlılık](#eşzamanlılık)
		- [Eşzamanlı Bir İşlem Oluşturalım](#eşzamanlı-bir-i̇şlem-oluşturalım)
	- [Kanallar \(Channels\)](#kanallar-channels)
		- [Boyutlu Kanal Oluşturma](#boyutlu-kanal-oluşturma)
	- [Anonim Goroutine Fonksiyonlar](#anonim-goroutine-fonksiyonlar)
	- [WaitGroup ile Asenkron İşlemleri Beklemek](#waitgroup-ile-asenkron-i̇şlemleri-beklemek)
	- [Mutex ile Asenkron İşlem Sırası](#mutex-ile-asenkron-i̇şlem-sırası)
	- [Zamanlayıcılar \(Tickers\)](#zamanlayıcılar-tickers)
	- [Select](#select)

# Bölüm 4 - Eşzamanlılık

## Goroutine

**Goroutine**’ler **Go Runtime** tarafından yönetilen hafif bir sistemdir. Bir işlemi eşzamanlı olarak yapmak istiyorsak, Goroutine'den faydalanabiliriz. Bu sayede aynı çalışma-zamanı içerisinde birden fazla iş parçacığı oluşturabiliriz.

### Terimler

#### Ana iş parçacığı

`Main()` fonksiyonu içerisine yazdığımız, asenkron olmayan kodlardır. Varsayılan olarak Go Runtime bu iş parçacığını izler. Programımız asenkron işlemlerin tamamlanmasını beklemiyorsa, ana iş parçacığı tamamlandığında program sona erer.

#### Eşzamanlılık

Eşzamanlılık, programlamada bir işlem gerçekleşirken, aynı zamanda başka işlemlerin de gerçekleşmesidir.

### Eşzamanlı Bir İşlem Oluşturalım

Eşzamanlı bir işlem oluşturmak için `go` anahtar kelimesinden faydalanabiliriz. Bunun için eşzamanlı çalışacak işlemin başına `go` yazmamız yeterli olacaktır.

![Asenkron &#x130;&#x15F;lem &#xD6;rne&#x11F;i](../.gitbook/assets/2020-11-09_23-38.png)

Aslında yukarıdaki örnekte `time.Sleep()` kullanarak 2 saniye bekletmemizin bir sebebi. Eğer `time.Sleep()` eklememiş olsaydık, ekrana _"Merhaba Dünya!"_ yazıldıktan sonra programımız sonlanacaktı. Bunun sebebi Go Runtime'ının Sadece Ana iş parçacığını beklemesi. Ana iş parçacığındaki işlemler sonlandıktan sonra, diğer işlemleri beklemiyor. Yukarıdaki örnekte bunu engellemek için `time.Sleep()` kullandık. Böylece program 2 saniye beklerken eşzamanlı işlemimiz de tamamlandı. Tabii `time.Sleep()` kullanarak beklemek mantıklı bir yöntem değil. İşlemin ne kadar süreceğini bilmediğimiz durumlar olacaktır. Bunun için Kanalları kullanabiliriz.

## Kanallar \(Channels\)

**Kanallar**, Go dilinde asenkron programlama yaparken değer aktarımı yapabileceğimiz hatlardır. Kanala değer atanması iş parçacığı tarafından bekleneceği için asenkron işlemler arasındaki senkronizasyonu ayarlayabiliriz. Kanallar `make()` fonksiyonu ile oluşturulur.

{% code title="Örnek" %}
```go
k := make(chan bool)
```
{% endcode %}

Yukarıdaki örnekte `make()` fonksiyonu ile `k` isminde bir kanal oluşturduk. Bu kanalın özelliği `bool` tipinde değer taşımasıdır. Yani bu kanal ile `true` veya `false` değerlerini taşıyabiliriz. Kanala değer göndermek için `<-` işaretini kullanırız. Yani bir nevi atama işlemi yapıyoruz. Atama işleminden farkı, kanala atama işlemi yapılana kadar iş parçacığının devam etmemesidir.

{% code title="Örnek Atama" %}
```go
k <- true
```
{% endcode %}

Atama işlemi ile kanalımıza değer yolladık. Bir de bu kanalın çıkış noktası olması gerekir. Bu çıkış noktasında, ister kanaldan gelen veriyi bir değişkene atayabiliriz, istersek de sadece kanala veri gelmesini bekleyebiliriz.

{% code title="Kanaldan gelen değeri değişkene atama" %}
```go
a := <-k
```
{% endcode %}

Yukarıdaki örnekte `a` isimli değişkene `k` kanalından gelen `bool` tipinde değer atadık. `a` değişkenine atama işlemi `k` kanalına değer gönderildiği zaman yapılacaktır. Yani `k` kanalına değer gelene kadar iş parçacığı duraklatılacaktır. _\(Program `k` kanalına gelecek değeri bekler.\)_

{% code title="Sadece kanala değer gelmesini beklemek" %}
```go
<- k
```
{% endcode %}

Yukarıdaki anlatılanlardan yola çıkarak bir örnek oluşturalım.

{% code title="Örnek kanal işlemleri" %}
```go
package main

import (
	"time"
)

func main() {

	//bir kanal oluşturalım
	k := make(chan bool)
	//bu kanalımız bool değer taşıyacak

	//asenkron bir iş parçacığı oluşturalım
	go func() {
		
		//bu iş parçacığı 5 sn beklesin
		time.Sleep(time.Second * 5)

		//k kanalına bool bir değer gönderelim
		k <- true
	}()

	//ana iş parçacığı k kanalına değer gelene kadar bekleyecek
	<-k
	//değer geldiğinde program sonlanacaktır.
}
```
{% endcode %}

### Boyutlu Kanal Oluşturma

Oluşturduğumuz kanala boyut vermek de mümkün. Yani kanalımıza birden fazla değer yollayabiliyoruz. Bunun için kanalı oluştururken `make()` fonksiyonunda boyutu da belirtelim.

{% code title="Örnek" %}
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	//2 adet bool değer taşıyan bir kanal oluşturalım
	k := make(chan bool, 2)
	

	//asenkron bir iş parçacığı oluşturalım
	go func() {

		//5 sn beklesin
		time.Sleep(time.Second * 5)

		//k kanalına bool bir değer gönderelim
		k <- true

		//tekrardan 2 sn beklesin
		time.Sleep(time.Second * 2)

		//ve k kanalına 2. değer de gönderilsin.
		k <- false
	}()

	//ana iş parçacığı k kanalına 2 değer gelene kadar bekleyecek
	fmt.Println(<-k, <-k) //çıktı: true false
	//iki bool değeri de baştırmak için k kanalını 2 defa yazdık
}
```
{% endcode %}

Ana iş parçacığı _\(`main()` içerisine yazılan kodlar\)_ devam etmek için `k` kanalına gelen 2 değeri de bekleyecektir.

`fmt.Println()` içerisine sadece bir defa `<-k` yazsaydık, `k` kanalına ilk gelen değeri ekrana bastıracaktı.

## Anonim Goroutine Fonksiyonlar

Bu yazımız **Goroutine** ve **Kanallar** dersi için biraz alıştırma tadında olacak.  
   
Daha önceki yazılarımızda belirli bir fonksiyonu **Goroutine** ile **asenkron** \(eş zamanlı\) olarak çalıştırmayı gördük. Bu yazımızda da **anonim** bir Goroutine fonksiyonunu göreceğiz. Bu fonksiyonun özelliği bir ismi olmaması ve asenkron olarak çalışmasıdır. Örneğimizi görelim.

```go
package main
import (
	"fmt"
	"time"
)
func main() {
	go func() {
		time.Sleep(time.Second * 2)
                fmt.Println("İlk yazımız")
	}()
	fmt.Println("İkinci yazımız")
}
```

Açıklamasına gelirsek **go func\(\)** ile anonim bir fonksiyon oluşturduk. Bu tür fonksiyonda fonksiyonumuzun sonuna **\( \)** parantezlerimizi yerleştirmek zorundayız. Çünkü fonksiyonumuza parametreleri bu parantezler içerisinde yolluyoruz. Şuanlık parametre yollamadığımzın için boş kalacak. Bu fonksiyonumuz programın geri kalanı ile aynı zamanda çalışacak. Hatta programın geri kalanı ile bağlantısı bile olmayacak. Bu sebepten ötürü mantıken 2 saniye sonra işlem yapmasını belirttiğimiz için **“İkinci yazımız”** metni gözüktükten sonra **“İlk yazımız”** metni gözükeceğini tahmin etsekte **go func\(\)** fonksiyonu yapısı gereği zaman bağımsız çalışacağı için **fmt.Println\(“İkinci yazımız”\)** fonksiyonu tamamlandıktan sonra **“İlk yazımız”** metni ekrana bastırılmayacaktır bile. İsterseniz programı çalıştırıp deneyebilirsiniz.  
Bunun önüne geçebilmein yolu **go func\(\)** fonksiyonundaki işlemlerin programın çalışma zamanı içerisinde sonuç vermesidir.

```go
package main
import (
	"fmt"
	"time"
)
func main() {
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("İlk yazımız")
	}()
	fmt.Println("İkinci yazımız")
	time.Sleep(time.Second * 3)
}
```

Yukarıdaki mantıkla çalışması için zamanı böyle ayarlamamız gerekir. Ama bu yöntem çok boş \(gereksiz\) bir yöntemdir. Her zaman böyle zamanı tahmin edemeyiz. Örnek olarak, **go func\(\)** fonksiyonunda internet üzerinden bir dosyanın inmesini bekleyecek olsaydık tahmini bir zaman belirleyemezdik. Ki koskoca Windows bile belirleyemiyor. Çünkü bu internet hızımız ile alaklı bir şeydir. Bu yüzden garanti bir yöntem değildir.

  
Bundan %100 daha garantili olan yöntem **kanallar** üzerinden haberleşmektir. Çünkü bir yerde kanal ataması yapıldığında program akışının devam edebilmesi için mutlaka kanaldan gelecek verinin beklenmesi gerekir. Bu sayede zaman ile alakalı işlerde tahmin yürütmemize gerek kalmaz. Biraz uzun bir açıklama oldu ama örneğimizi görünce mantığını anlayacaksınız.

```go
package main
import (
	"fmt"
	"time"
)
func main() {
	kanal := make(chan string) //kanal oluşturuyoruz
	go func() {
		time.Sleep(time.Second * 2) //2 saniye uyku
	        kanal <- "Kanal bitti" //İletişime geçiriyoruz
		fmt.Println("Anonim fonksiyon yazısı")
	}()
	fmt.Println("Öylesine bir yazı")
	fmt.Println(<-kanal) //kanaldan gelen veri bekleniyor
}
```

Öncelikle kanal ile ilgili işlemler yapabilmek için **make** fonksiyonu ile kanal oluşturduk. Hemen altında kanalımızı iletişime sokmak için öylesine bir **string** değer yolladım.

  
**go func\(\)** fonksiyonumuz yukarıdaki örnekler ile aynıdır. Bu fonksiyonumuzun 2 saniye beklemesi olduğundan dolayı fonksiyonumuzun altındaki **“Öylesine bir yazı”** daha önce görüntülenecek. Buraya kadar ilk örnek ile aynı olayla sonuçlanıyor. Programın sonlanmasını engellemek için **&lt;- kanal** içinden değeri bastırarak kanal iletişimini beklemesini ve bundan dolayı **“Anonim fonksiyonu yazısı”**‘nı da beklemiş oluyoruz.  
   
Anonim Goroutine fonksiyonları bu şekilde kullanabiliriz.

## WaitGroup ile Asenkron İşlemleri Beklemek

Goroutine’leri Asenkron programlama yaparken kullanırız. Böylece aynı anda birden fazla işlem gerçekleştirebiliriz. Peki programımızın belirttiğimiz asenkron işlemleri bekleme gibi bir ihtiyacı olsaydı, ne yapmamız gerekirdi? Bu durumlarda WaitGroup'lardan faydalanabiliriz. Örneğin projemizde 3 adet asenkron işlem bulunuyorsa, WaitGroup'a 3 değerini ekleriz. Her asenkron işlem tamamlandığında WaitGroup -1 azalır ve sıfıra geldiğinde WaitGroup tamamlanmış olur. WaitGroup'u kullanmak için ise "sync" paketini projemize dahil ediyoruz. Kodlar üzerinde açıklamasını görelim.

{% code title="main.go" %}
```go
package main

import (
	"fmt"
	"sync" //WaitGroup'u kullammak için
	"time" //bekleme işlemleri için
)

/*
* waitgroup nesnesini işaretçi olarak parametre veriyoruz.
* işaretçi olarak vermemizin sebebi, programın bekleme işlemi için
* asıl waitgroup nesnesini kontrol etmesidir.
 */
func fonksiyon1(wg *sync.WaitGroup) {

	//fonksiyonun 2 sn beklemesini istiyoruz.
	time.Sleep(2 * time.Second)
	fmt.Println("Fonk1 tamamlandı")

	//wg.Done() fonksiyonu ile waitgroup nesnesini -1 azalttık.
	wg.Done()
}

//bu fonksiyonumuza da wg nesnesini işaretçi ile parametre olarak verdik.
func fonksiyon2(wg *sync.WaitGroup) {
	//fonksiyonu 3 sn uyuttuk.
	time.Sleep(3 * time.Second)
	fmt.Println("Fonk2 tamamlandı")

	//-1 daha eksilttik.
	wg.Done()
}

func main() {
	/*
	* Öncelikle waitgroup'u kullanabilmek için  bir waitgroup
	* nesnesi oluşturuyoruz.
	 */
	var wg sync.WaitGroup

	/*
	* waitgroup'a 2 ekliyoruz. Yani 2 tane işlemden yanıt gelmesini
	* beklemesini istiyoruz. Aslında burada beklemeyecek. Sadece
	* işlem sayısını belirttik.
	 */
	wg.Add(2)

	/*
	* fonksiyon1 ve fonksiyon2'ye oluşturduğumuz wg örneğinin
	* bellekteki adresinin veriyoruz.
	 */
	go fonksiyon1(&wg)
	go fonksiyon2(&wg)
	fmt.Println("Merhaba Dünya!")

	/*
	* Burada wg.Wait() fonksiyonu ile asenkron işlemleri beklemesini
	* sağlıyoruz. yani waitgroup'un 0'a düşmesini bekliyoruz.
	* Eğer waitgroup olmadan yapsaydık. asenkron fonksiyonlarımızın tamamlanmasını
	* beklemeden program kendini sonlandırırdı.
	 */
	wg.Wait()

	//waitgroup tamamlandığında ekrana yazı bastıralım.
	fmt.Println("WaitGroup'lar tamamlandı.")
}
```
{% endcode %}

## Mutex ile Asenkron İşlem Sırası

Size konu başlığını şöyle açıklayayım. Örneğin bir banka uygulaması para çekme ve yatırma gibi özelliklere sahiptir. Programlama mantığında para yatırmak ve çekmek için mevcut para miktarını bilmemiz gerekir. Banka uygulamasının mantığı en basit derecede bu şekilde çalışır.

Banka hesabımızda asenkron işlem yapıldığını varsayalım. Yani bir hesaptan aynı anda birden fazla kullanıcı işlem yapıyor olsun.

Örneğin hesabımızda 100₺ olsun. Birinci kullanıcı 20₺ yatırsın. Aynı anda ikinci kullanıcı 50₺ çeksin. Bu iki kullanıcınında kullandığı program işlem yapmaya başladığında önce para miktarını alıyor. Daha sonra yapılacak işleme göre ya ekleme ya da çıkarma işlemi yapıyor. Fakat birden fazla kullanıcı aynı anda bu işlemi yaparsa hesaptaki parada yanlışlık olacaktır.

Basit bir görsel ile inceleyelim.

İşlemlere aynı anda başlandığını varsayalım.

![&#xD6;rnek asenkron i&#x15F;lem](../.gitbook/assets/mutex.png)

Bu işlemin sonuncunda hangi kullancının işlemi sonuncu olarak biterse para miktarı onun sonucu olur. Yani kullanıcı 2'nin işlemi kullanıcı 1'den sonra biterse yeni para miktarı 50₺ olur.

Bu gibi örneklerde asenkron işlemlere sıra verilmesi gerekir. Mutex tam olarak bu işi yapıyor. Bunun için bir Mutex nesnesi oluşturuyoruz. İşlemlerimizi bu nesne üzerinden yapıyoruz. Bu nesne aynı anda sadece bir işlemi gerçekleştiriyor. Bu yüzden sıra işlemi sağlıyor. Önce başlayan asenkron işlem  ilk sırada oluyor. Tamamlanınca diğerine sıra geçiyor. Böyle düşündüğümüz zaman "bunun senkron programlamadan ne farkı var?" diyebilirsiniz. Farkı asenkron fonksiyonların içindeki istediğimiz kısımları senkron çalıştırmamız.

Örnek bir para yatırma-çekme uygulaması yazalım. İşlemin sağlık çalışması için, para miktarıyla aynı anda sadece bir kişi işlem yapabilmelidir.

```go
package main

import (
	"fmt"
	"sync" // mutex'i kullanmak için
)
//global olarak mutex nesnesi oluşturalım.
var mt sync.Mutex


func paraÇek(bakiye *float64, çekilecekMiktar float64, wg *sync.WaitGroup) {
	/*
	* mt isimli mutex'i bu işlem yapılırken kilitliyoruz.
	* bu sayede mt mutex'ini başka işlemler kullanamıyor.
	*/
	mt.Lock()

	/*
	bu kısımda asenkron olmasını istemediğimiz işlemi yapalım.
	*/
	*bakiye -= 15
	fmt.Printf("Yeni Bakiye: %.2f\n", *bakiye)

	/*
	* diğer işlemlerinde kullanabilmesi için mutex'i tekrardan açalım.
	* mt mutex açılınca diğer asenkron işlemdeki mt mutex'i çalışmaya başlar. 
	*/
	mt.Unlock()
	fmt.Println("Çekme işlemi tamamlandı.")

	/*
	* waitgroup ile işlemin tamamlandığını belirttik.
	* böylece wg havuzu 2'den 1'e düştü
	*/
	wg.Done()
}

//bu fonksiyonda yukarıdaki ile aynı mantıkta
func paraYatır(bakiye *float64, yatırılacakMiktar float64, wg *sync.WaitGroup) {
	mt.Lock()
	*bakiye += 65
	fmt.Printf("Yeni Bakiye: %.2f\n", *bakiye)
	mt.Unlock()
	fmt.Println("Yatırma işlemi tamamlandı.")
	wg.Done()
}

func main() {

	/*
	* asenkron işlemlerimizin, ana iş parçacığında tamamlanmasını
	* beklemek için waitgroup nesnesi oluşturalım
	*/
	var wg sync.WaitGroup

	//2 fonksiyonu da bekleyeceğimiz için Add'e 2 yazalım
	wg.Add(2)

	//fonksiyonlarımızın kullancağı bakiye değişkenimiz
	var bakiye float64 = 100
	fmt.Printf("İlk Bakiye: %.2f\n", bakiye)
	
	/*
	* paraÇek ve paraYatır fonksiyonlarımızı aynı anda başlatıyoruz.
	* hangisi daha önce başlarsa mutex sırasına ilk o girer. bu esnada diğer
	* fonksiyon mutex'in açılmasını bekler.
	*/
	go paraÇek(&bakiye, 25, &wg)
	go paraYatır(&bakiye, 65, &wg)

	/*
	* ana iş parçacığı tamamlandığında asenkron çalışan fonksiyonları beklemez.
	* beklemediğinde de asenkron fonksiyonlar çalışmadan program sonlanır.
	* ana iş parçacığının asenkron işlemleri beklemesi için waitgroup sonucunun 0 olmasını bekleriz.
	* wg.Add(2) yazarak 2 adet wg.Done() fonksiyonu çalıştığında wg.Add(0) olur ve
	* wg.Wait() tamamlanır ve program başka işlemler yapılmıyor ise sonlanır.
	*/
	wg.Wait()
}
```

Çıktımız aşağıdaki gibi olacaktır.

> İlk Bakiye: 100.00  
> Yeni Bakiye: 165.00  
> Yatırma işlemi tamamlandı.  
> Yeni Bakiye: 150.00  
> Çekme işlemi tamamlandı.

Yukarıdaki çıktıya göre, `paraYatır()` fonksiyonu `paraÇek()` fonkisyonundan önce çalışmıştır.

## Zamanlayıcılar \(Tickers\)

Golang’de **zamanlayıcılar**, belirli sürede bir tekrar etme işlemi için kullanılır. Zamanlayıcılar programın çalışma süresince veya durdurulana kadar çalışabilir. Örneğimizi görelim:

```go
package main
import (
	"fmt"
	"time"
)
func main() {
	tekrar := time.NewTicker(500 * time.Millisecond) // her yarım saniyede 1
	bitti := make(chan bool)
	go func() {
		for {
			select {
			case <-bitti:
				return
			case zaman := <-tekrar.C:
				fmt.Println("Tekrar zamanı:", zaman)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond) // 1,6 saniye programı uyut
	tekrar.Stop()                       // Durdurduk
	bitti <- true                       // for döngüsünü sonlandırdık.
	fmt.Println("Tekrarlayıcı durdu!")
}
```

Açıklaması şöyledir:  
   
**tekrar** adında bir zamanlayıcı oluşturduk ve bu zamanlayıcının özelliği her **yarım saniyede bir** tetiklenmesi.  
   
**bitti** adında, boolean değer taşıyan bir kanal oluşturduk. Bu kanalın mantığı ileride anlayacaksınız.  
   
Anonim Goroutine fonksiyonunun içine, yani **go func\(\)**, sınırsız döngü çeviren bir **for** oluşturduk. Bu döngünün içerisinde **select** ile kanal iletişimlerimizi dinledik. Döngümüzün sonlanması için **bitti** kanalına herhangi bir veri gelmesi gerekiyor. Aşağısındaki **case**‘de zaman değişkenimize tekrar zamanlayıcımız tetiklendikçe bu durum çalışacak. \(tekrar.C ile zaman bilgisini alıyoruz.\) Yani yarım saniyede bir zaman kanalına veri gelecek.  
   
Anonim Goroutine fonksiyonu, **main\(\)** fonksiyonundan ayrı olarak çalıştığından bu fonksiyonumuzun çalışması için ona zaman aralığı vermemiz gerekiyor. **time.Sleep\(1600 \* time.Millisecond\)** ile **main\(\)** fonksiyonumuzu 1,6 saniye bekletiyoruz. Bu bekleme süresi içinde tekrar zamanlayıcımız 3 kere tetikleniyor. \(500 \* x &lt; 1600  \| x = 3\) Haliyle de 3 kere ekrana çıktımızı bastırıyor. 1,6 saniye geçtikten sonra tekrar zamanlayımızı **tekrar.Stop\(\)** ile durduruyoruz.  
   
**bitti** kanalına değer yollayarak, yukarıdaki **for** döngümüzü **return** ile sonlandırmış oluyoruz.  
   
Ve en son ekranımıza **“Tekrarlayıcı durdu!”** yazımızı bastırıyoruz.  
Çıktımız aşağıdaki gibi olacaktır:

> Tekrar zamanı: 2019-10-15 14:08:02.002909142 +0300 +03 m=+0.500235484  
> Tekrar zamanı: 2019-10-15 14:08:02.502993622 +0300 +03 m=+1.000319851  
> Tekrar zamanı: 2019-10-15 14:08:03.002952074 +0300 +03 m=+1.500278387  
> Tekrarlayıcı durdu!

## Select

**Select** ile çoklu goroutine işlemlerinin iletişimini bekleyebiliriz. Örneğimizi görelim:

```go
package main
import (
	"fmt"
	"time"
)
func main() {
	k1 := make(chan string)
	k2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		k1 <- "video"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		k2 <- "ses"
	}()
	for i := 0; i < 2; i++ {
		select {
		case mesaj1 := <-k1:
			fmt.Println("Mesaj 1:", mesaj1)
		case mesaj2 := <-k2:
			fmt.Println("Mesaj 2:", mesaj2)
		}
	}
}
```

Yukarıdaki kodların bize **ses** ve **video** verisi sağlayacak bir programdan parça olduğu senaryosunu kuralım. Bu programda işlem yapabilmemiz için bize bu 2 verinin gelmesini beklememiz lazım. Verileri bekleme işlemini **select** ile yapıyoruz. Burada dikkat etmemiz gereken nokta **2** tane veri beklediğimiz için  **for** atamalarında **i &lt; 2** olarak girmeliyiz. Çünkü **i := 0** olduğu için **i 2** olana kadar arada **2** sayı var. Bu sayı boşluğu da 2 veri almayı beklememizi sağlıyor. Örnek olarak **i &lt; 1** girip 2 veri almaya kalksak **k2**‘den gelen veriyi beklemeyecek bile. Tam tersi olarak 2 veri alacağımız halde **i &lt; 4** girsek program **deadlock**‘a girecektir. Yani başarısız bir program olacaktır.

