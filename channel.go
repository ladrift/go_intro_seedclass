package main

import "fmt"

func main() {
	c := make(chan string)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			c <- fmt.Sprintf("I am Number %d", i)
		}()
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %s\n", <-c)
	}
	fmt.Println("I am leaving.")
}
