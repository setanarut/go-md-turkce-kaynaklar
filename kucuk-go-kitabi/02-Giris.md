- [Giriş](#giriş)
  - [Yazardan Bir Not](#yazardan-bir-not)

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

