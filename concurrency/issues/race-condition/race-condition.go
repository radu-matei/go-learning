package main

import (
	"fmt"
	//"time"
)

func main() {
	var data int

	go func() {
		data++
	}()

	//time.Sleep(time.Second)
	if data == 0 {
		fmt.Printf("the value is: %v", data)
	}
}
