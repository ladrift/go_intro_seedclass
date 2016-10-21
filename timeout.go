package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c <- "Oh, I am arrived."
	}()

	select {
	case msg := <-c:
		fmt.Printf("%s\n", msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("You are too late.")
	}
}
