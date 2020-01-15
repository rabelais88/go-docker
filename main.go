package main

import (
	"fmt"
	"myproject/concurrency"
	"myproject/employee"
	"myproject/shape"

	"github.com/rabelais88/fib"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println(`hello world!`)
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
	password := []byte(`mypassword`) // type casting/conversion
	bs, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(`hashed password:`, bs)

	c := shape.Circle{Radius: 5}
	fmt.Println(c.Area())
	concurrency.GetEnvironment()
	concurrency.Channel()
	concurrency.DirectionalChannel()
	concurrency.RangeChannel()
	concurrency.SelectChannel()
	fmt.Println(`fibonacci of 3`, fib.Fibonacci(3))
}
