package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

// Goroutine merupakan jawaban dari proses thread dan pararel programming yang sifatnya berjalan bersamaan
// goroutine berjalan di atas thread di sebuah proses dan merupakan gabungan dari pararel dan concurency programming
// dengan goroutine proses menjadi async dan berjalan bergantian

// Go Scheduler: G: Goroutine M: Machine (Thread) P: Prossesor [sudah berjalan otomatis dari go runtimenya]

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	// untuk menunggu goroutinenya selesai eksekusi
	time.Sleep(1 * time.Second) // default menggunakan nano

	// tidak recomend menggunakan goroutine kalo dari func yang menggunakan return value
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
