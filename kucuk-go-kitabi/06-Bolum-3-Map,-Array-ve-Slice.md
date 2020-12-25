- [Bölüm 3 - Map, Array ve Slice](#bölüm-3---map-array-ve-slice)
  - [Diziler](#diziler)
  - [Dilimler](#dilimler)
  - [Eşlemeler (Map)](#eşlemeler-map)
  - [İşaretçiler ve Değerler](#i̇şaretçiler-ve-değerler)
  - [Devam Etmeden Önce](#devam-etmeden-önce)

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

