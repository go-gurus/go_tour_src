package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func sendDataToChannel(input chan int) {
	defer close(input)
	for i := 0; i < 100; i++ {
		value := rand.Intn(50)
		input <- value
		time.Sleep(time.Duration(value))
	}
}

func processChannel(id int, input chan int, output chan string, done chan bool) {
	for {
		select {
		case value, ok := <-input:
			if !ok {
				done <- true
				return
			}
			result := value * 1000
			output <- "worker: " + strconv.Itoa(id) + ", input: " + strconv.Itoa(value) + ", result: " + strconv.Itoa(result)
		}
		time.Sleep(time.Duration(result))
	}
}

func main() {
	input := make(chan int)
	output := make(chan string)
	done := make(chan bool)

	defer close(output)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go processChannel(i, input, output, done)
	}
	go sendDataToChannel(input)

	go func() {
		for {
			select {
			case <-done:
				wg.Done()
			case x, ok := <-output:
				if ok {
					fmt.Println(x)
				} else {
					fmt.Println("Channel closed!")
				}
			}
		}
	}()

	wg.Wait()
}
