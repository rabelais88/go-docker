package main

import (
	"fmt"
	"myproject/concurrency"
	"myproject/employee"
	"myproject/shape"

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
	concurrency.GetEnvironment()

	c := shape.Circle{Radius: 5}
	fmt.Println(c.Area())
}
