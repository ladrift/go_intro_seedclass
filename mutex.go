package main

import (
	"fmt"
	"sync"
)

func main() {
	// main START OMIT
	var a string
	// Initial state is un-locked
	var m sync.Mutex

	m.Lock()
	go func() {
		a = "hello, world"
		fmt.Println("sub routine exit")
		m.Unlock()
	}()

	m.Lock()
	fmt.Println("main routine exit")
	fmt.Println(a)
	// main END OMIT
}
