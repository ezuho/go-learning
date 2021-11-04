package main

import "fmt"

func main() {
	messgaes := make(chan string)

	go func() { messgaes <- "ping" }()

	msg := <-messgaes
	fmt.Println(msg)
}
