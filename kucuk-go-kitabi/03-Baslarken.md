- [Başlarken](#başlarken)
  - [OSX / Linux](#osx--linux)
  - [Windows](#windows)
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

```shell
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

