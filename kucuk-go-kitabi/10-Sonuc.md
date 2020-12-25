# Sonuç

Son zamanlarda Go'nun *sıkıcı bir* dil olarak tanımlandığını duydum. Sıkıcı çünkü öğrenmesi kolay, yazması kolay ve en önemlisi okunması kolay. Belki de bu gerçeği göstererek bir kötülük yaptım. Üç bölümü tiplerden ve nasıl tanımlanacaklarından bahsederek *harcadık*.

Statik olarak yazılan bir dilde tecrübeniz varsa, gördüğümüz şeylerin çoğu muhtemelen en iyi ihtimalle size birer hatırlatma oldu. Go, işaretçileri görünür kılar ve diziler etrafındaki ince araçlar olarak verdiği dilimleri deneyimli Java veya C # geliştiricilerine zor gelmeyecektir.

Çoğunlukla dinamik dillerden faydalanıyorsanız, biraz farklı hissedebilirsiniz. Öğrenmesi *biraz* zor olabilir. Tanımlama ve değer atama ile ilgili çeşitli sözdizimi farkı göreceksiniz. Bir Go hayranı olmasına rağmen, basitliğe doğru tüm ilerlemeye rağman, bununla ilgili hala basit bir şey olmadığını düşünüyorum. Yine de, bazı temel kurallara (değişkenleri yalnızca bir kez bildirebileceğiniz ve `:=` değişkeni bildirdiğiniz gibi) ve temel anlayışa ( `new(X)` veya `&X{}` sadece bellek ayırır, ancak `make`  dilimler, eşlemeler ve kanallar için daha fazlasını gerektirir) değişiklik getirir.

Bunun ötesinde Go bize kodumuzu düzenlemenin basit ama etkili bir yolunu sunuyor. Arayüzler, dönüş değerine dayalı hata yönetimi, kaynak yönetimi için `defer` ve kompozisyon elde etmenin basit bir yolu gibi.

Son fakat bir o kadar önemli, eşzamanlılık için yerleşik desteğidir. Go rutinler hakkında etkili ve basit olmasından başka söylenecek çok az şey var (yine de kullanımı basit). İyi bir soyutlamadır. Kanallar daha karmaşıktır. Her zaman üst düzey sarmalayıcıları kullanmadan önce temel bilgileri anlamanın önemli olduğunu düşünüyorum. Ben kanallar olmadan eşzamanlı programlama öğrenmenin yararlı olduğunu *düşünüyorum*. Yine de, kanallar benim için basit bir soyutlama gibi hissetmeyecek şekilde uygulanmaktadır. Neredeyse kendi temel yapı taşlarıdır. Bunu söylüyorum çünkü eşzamanlı programlama hakkında yazma ve düşünme şeklinizi değiştiriyorlar. Eşzamanlı programlamanın ne kadar zor olabileceği göz önüne alındığında, bu kesinlikle iyi bir şeydir.

