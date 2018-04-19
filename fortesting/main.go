package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main enter")
	var sss string
	if sss == "" {
		fmt.Println("is empty")
	}

	now := time.Now()

	time.Sleep(2 * time.Second)

	fmt.Printf("%d\n", time.Since(now)/time.Second)
	fmt.Printf("sub %d\n", time.Now().Sub(now))

	var timeNil time.Time
	fmt.Println(timeNil)
}
