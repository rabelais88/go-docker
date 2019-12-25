package main

import (
	"fmt"
	"myproject/employee"

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
}
