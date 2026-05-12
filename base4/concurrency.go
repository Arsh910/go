package main

import (
	"fmt"
	"sync"
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

func doWork(i int, ch chan int, wg *sync.WaitGroup) {
	// for range
	defer wg.Done()

	if i == 2 {
		ch <- i
	} else {
		time.Sleep(time.Second * 10)
		ch <- i
	}
}

// for range

func closeRoute(ch chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(ch)
}

func fib(ch chan int, n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Microsecond)
	}
	close(ch)
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
	// go func() {
	// 	ch <- 40
	// }()
	// _, ok := <-ch
	// fmt.Println(ok)

	// bch := make(chan int, 100)
	// bch <- 10
	// _, ok := <-bch
	// fmt.Println(ok)

	// range

	// ch := make(chan int, 3)
	// var wg sync.WaitGroup

	// wg.Add(3)

	// go doWork(1, ch, &wg)
	// go doWork(4, ch, &wg)
	// go doWork(3, ch, &wg)

	// go closeRoute(ch, &wg)

	// for item := range ch {
	// 	fmt.Println(item)
	// }

	// concurrent fibonachi
	// ch := make(chan int) //initialize a channel always
	// go func(n int) {
	// 	fib(ch, n)
	// }(10)

	// for item := range ch {
	// 	fmt.Println(item)
	// }

	// select
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// // go func() { ch1 <- "one" }()
	// // go func() { ch2 <- "two" }()

	// select {
	// case i, ok := <-ch1:
	// 	fmt.Println("from channel 1 ", i, " ", ok)
	// case i, ok := <-ch2:
	// 	fmt.Println("from channel 2 ", i, " ", ok)
	// default:
	// 	fmt.Println("done select")
	// }

	// tickers

	// runs forever
	// tickT := time.Tick(500 * time.Microsecond)
	// for t := range tickT {
	// 	fmt.Println(t)
	// }

	// run once
	// tickA := time.After(500 * time.Microsecond)
	// fmt.Println(<-tickA)

	// Readonly and Writeonly channels
	// ch := make(chan int)
	// var readOnly <-chan int = ch
	// var writeOnly chan<- int = ch

	// go func() {
	// 	writeOnly <- 42
	// }()
	// val := <-readOnly
	// fmt.Println(val)
	fmt.Println("hello done")

	// mutexes

}
