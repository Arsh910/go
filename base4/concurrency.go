package main

import (
	"fmt"
	"time"
)

// go
func doSomething() {
	for i := 0; i < 10; i++ {
		fmt.Println("sent")
		time.Sleep(time.Second)
	}
}

// channels

func doWork(i int, ch chan int) {
	if i == 2 {
		ch <- i
	} else {
		time.Sleep(time.Second * 10)
		ch <- i
	}
}

func main() {
	fmt.Println("hello")

	// concurency
	// go doSomething()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("mysent")
	// 	time.Sleep(time.Second)
	// }

	// channel
	// ch := make(chan int)

	// go doWork(1, ch)
	// go doWork(4, ch)
	// go doWork(3, ch)

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// deadlock , no reiever
	// ch <- 42

	// buffer channels
	// ch := make(chan int, 100)
	// ch <- 42

	// closing channel

	// ch := make(chan int)
	// // close(ch)
	// ch <- 40
	// _, ok := <-ch
	// fmt.Println(ok)

	// bch := make(chan int, 100)
	// bch <- 10
	// _, ok := <-bch
	// fmt.Println(ok)

	fmt.Println("hello done")

}
