package concurrency

import "fmt"

func logger(arg ...interface{}) {
	fmt.Println(`channel:`, arg)
}

func Channel() {
	c := make(chan int, 1)      // can access, put value inside
	c <- 1                      // works
	logger(`from channel`, <-c) // works

	// declaring unidirectional channel(send-only channel)
	d := make(chan<- int, 1) // can only put value inside, can't access
	d <- 1                   // works
	// fmt.Println(<-d) // does not work
}

// unidirectional channel

// send only, goroutine
func sendonly(c chan<- int) {
	c <- 10
}

// receive only, goroutine
func recvonly(c <-chan int) {
	logger(`receiving from channel`, <-c)
}

func DirectionalChannel() {
	c := make(chan int)
	go sendonly(c)
	go recvonly(c)

}

func loopOverChannel(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c) // closing channel to avoid deadlock
}

func RangeChannel() {
	c := make(chan int)
	go loopOverChannel(c)
	for v := range c {
		logger(v)
	}
	logger(<-c)
}

// select + channel

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0 // when this channel is finally used, the program automatically shuts because all the channels are used.
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			logger(`from even:`, v)
		case v := <-o:
			logger(`from odd`, v)
		case v := <-q:
			logger(`from the quit`, v)
			return
		}
	}
}

func SelectChannel() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(even, odd, quit)
	receive(even, odd, quit)
}
