package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Sleep time simulation
func printNumber(func_name string) {
	fmt.Printf("start %s function.\n", func_name)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Syncronization with waitgroup
// waitgroup refer parameter
func wg_printNumber(func_name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("start %s function.\n", func_name)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Communication with channel
// Unbuffered channel
func ch_producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send data to the channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close the channel to siginal completion
}

// Goroutine with context for cancellation
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d : stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d: working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker_pull(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	// Base usage of goroutine
	// simple goroutine example
	go printNumber("Simple goroutine")

	// Anonymous functions with goroutine
	go func() {
		fmt.Println("Anonymous function Start")
		for i := 0; i < 3; i++ {
			fmt.Println("Anonymous function", i)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Anonymous function End")
	}()

	// Synchroniztion with WaitGroup
	// the `sync.WaitGroup` is used to wait for a collection of goroutines to finish
	// *sync package!
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go wg_printNumber("Wait_Group_goroutine", &wg)
	}
	wg.Wait()

	// Communication with channels
	// channel provide a way for goroutines to communicate and synchronize their execution.

	// unbuffered channel
	ch := make(chan int)
	go ch_producer(ch)
	for num := range ch {
		fmt.Println("Received : ", num)
	}

	// Buffered channel
	// buffered channels allow sending a limited number of values without blocking
	ch2 := make(chan int, 2) // Create a buffered channel with a capacity of 2

	ch2 <- 1
	ch2 <- 2
	// ch2 <- 3 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// Select Statement with Channels
	// the `select` statement lets you wait on multiple channel operations
	// Select 문을 통해 다중 채널의 입력을 들어온 순서대로 처리 할 수 있다.
	selected_ch00 := make(chan string)
	selected_ch01 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		selected_ch00 <- "One!"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		selected_ch01 <- "Two!"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-selected_ch00:
			fmt.Println("Received from ch1 :", msg1)
		case msg2 := <-selected_ch01:
			fmt.Println("Received from ch2 :", msg2)
		}
	}

	// Goroutine with context for cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	// Worker pull example
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start job
	var wg_pull sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg_pull.Add(1)
		go worker_pull(w, jobs, results, &wg_pull)
	}
	// send job
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channels to signal no more jobs
	wg_pull.Wait()
	close(results)
	for result := range results {
		fmt.Println("Result:", result, "(job * 2)")
	}

	// Let the main function wait to see the output of the goroutine
	time.Sleep(600 * time.Millisecond)
	// 만약 go routine function 보다 main function 이 빠르게 종료될 경우,
	// go routine은 강제 종료된다.
	fmt.Println("Main Function Ends.")
}
