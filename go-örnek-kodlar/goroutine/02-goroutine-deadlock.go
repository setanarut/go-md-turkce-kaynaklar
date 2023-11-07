package main

import (
	"fmt"
)

func main() {
	// Channel yarat
	mesajlar := make(chan string, 2)

	// goroutine olmadan mesaj gönder (deadlock)
	mesajlar <- "birmesaj"
	fmt.Println(mesajlar)
	mesajlar <- "birmesaj2"
	fmt.Println(mesajlar)

	// Mesajı aldığında onu msg değişkeninde sakla
	msj := <-mesajlar

	// print received message
	fmt.Println(msj)
}
