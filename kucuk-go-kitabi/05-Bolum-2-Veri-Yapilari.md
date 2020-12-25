- [Bölüm 2 - Veri Yapıları](#bölüm-2---veri-yapıları)
  - [Tanımlamalar ve Başlangıç Değerleri](#tanımlamalar-ve-başlangıç-değerleri)
  - [Veri Yapılarındaki İşlevler](#veri-yapılarındaki-i̇şlevler)
  - [Oluşturucu Yöntemler](#oluşturucu-yöntemler)
  - [New](#new)
  - [Yapıların Alanları](#yapıların-alanları)
  - [İçerme](#i̇çerme)
    - [Yeni Görev Yükleme](#yeni-görev-yükleme)
  - [İşaretçiler mi, Değerler mi?](#i̇şaretçiler-mi-değerler-mi)
  - [Devam Etmeden Önce](#devam-etmeden-önce)

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

