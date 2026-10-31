package main

import (
	"fmt"
	"sync"
)

func squares(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Calculating squares sum...")

	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i * i
	}

	fmt.Println("Squares sum calculation completed.")
	ch <- fmt.Sprintf("Squares Sum: %d", sum)
}

func cubes(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Calculating cubes sum...")

	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i * i * i
	}
	fmt.Println("Cubes sum calculation completed.")
	ch <- fmt.Sprintf("Cubes Sum: %d", sum)
}

func fibonacci(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Calculating Fibonacci sequence...")

	a, b := 0, 1
	result := ""
	for i := 0; i < 7; i++ {
		result += fmt.Sprintf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println("Fibonacci sequence calculation completed.")
	ch <- fmt.Sprintf("Fibonacci: %s", result)
}

func main() {
	ch := make(chan string, 3)
	var wg sync.WaitGroup

	wg.Add(3)

	go squares(ch, &wg)
	go cubes(ch, &wg)
	go fibonacci(ch, &wg)

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("All tasks completed!")
}
