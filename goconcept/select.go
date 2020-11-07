package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		fmt.Printf("x = %d\n", x)
		select {
		case c <- x:
			x, y = y, x+y
		case quit_value := <-quit:
			fmt.Printf("quit_value = %d\n", quit_value)
			return
		}

	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("Goroutine started")
		for i := 0; i < 10; i++ {
			value := <-c
			fmt.Printf("received %d\n", value)
		}
		quit <- 999
	}()

	fmt.Printf("calling fibonacci()")
	fibonacci(c, quit)

}
