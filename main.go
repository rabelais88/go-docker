package main

import (
	"fmt"

	"myproject/employee"
)

func main() {
	fmt.Println(`hello world!`)
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
