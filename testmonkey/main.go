package main

import (
	"fmt"

	"time"

	"github.com/bouk/monkey"
)

func add(a, b int) int {
	fmt.Println("old add")
	return a + b
}

func main() {
	guard := monkey.Patch(add, func(a, b int) int {
		fmt.Println("new add")
		return a + b + 1
	})
	defer guard.Unpatch()

	fmt.Printf("test add : %d\n", add(10, 20))

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
