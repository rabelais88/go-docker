package concurrency

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 20; i++ {
		logger(`foo`, i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 20; i++ {
		logger(`bar`, i)
	}
}

// exported names should be started with a capital letter
func GetEnvironment() {
	logger("OS\t\t", runtime.GOOS)
	logger("ARCH\t\t", runtime.GOARCH)
	logger("CPUs\t\t", runtime.NumCPU())
	logger("Goroutines\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	logger("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}
