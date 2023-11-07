package main

import (
	"fmt"
)

func main() {
	// Channel yarat
	mesajlar := make(chan string)

	// Mesajı göndermek için bir goroutine çalıştır
	go func() { mesajlar <- "birmesaj" }()

	// Mesajı aldığında onu msg değişkeninde sakla
	msj := <-mesajlar

	// print received message
	fmt.Println(msj)
}
