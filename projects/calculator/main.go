package main

import (
	"fmt"
	"time"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for {
			select {
			case x := <-firstChan:
				output <- x * x
			case x := <-secondChan:
				output <- x * 3
			case <-stopChan:
				return
			}
		}
	}()

	return output
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	outputChan := calculator(firstChan, secondChan, stopChan)

	go func() {
		time.Sleep(5 * time.Second)
		close(stopChan)
	}()

	go func() {
		firstChan <- 5
		secondChan <- 10
		firstChan <- 21
		secondChan <- 13
	}()

	for result := range outputChan {
		fmt.Println(result)
	}
}
