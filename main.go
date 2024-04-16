package main

import (
	"fmt"
	"time"
)

func sender(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Sending:", i)
		c <- i
		time.Sleep(time.Second) // Menambahkan delay satu detik
	}
	close(c)
}

func main() {
	c1 := make(chan string, 1)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- "every 500 millisecond"
			time.Sleep(500 * time.Millisecond)

		}
		close(c1)
	}()

	for msg := range c1 {
		fmt.Println(msg)
	}
}
