package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "Dandi Irwanto"
		fmt.Println("Selesai Mengirim data ke Channel")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)

	// contoh kirim data ke channel
	// channel <- "Dandi"

	// contoh penerima data dari channel
	// kirim ke variable
	// data := <-channel

	// kirim data ke param
	// fmt.Println(<-channel)

	// menutup channel ini recomendasi untuk menghindari memory leak
	// close(channel)
	// better menggunakan defer
	// defer close(channel)
}
