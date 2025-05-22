package main

import (
	"fmt"
	"sync"
	"time"
)

func printName(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("Hello %s!\n", name)
}

// func main() {
// 	fmt.Println("Inside main")

// 	var wg sync.WaitGroup

// 	names := []string{"John", "Paul", "Doe", "Walker", "Alice"}

// 	for _, name := range names {
// 		wg.Add(1)
// 		go printName(name, &wg)
// 	}

// 	wg.Wait()

// 	// SUM
// 	var slice []int
// 	slice = append(slice, 10)
// 	slice = append(slice, 20)
// 	slice = append(slice, 30)
// 	slice = append(slice, 40)
// 	slice = append(slice, 50)
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	firstSlice, secondSlice := SplitSlice(slice)
// 	fmt.Println(firstSlice)
// 	fmt.Println(secondSlice)
// 	go CalculateSum(firstSlice, ch1)
// 	go CalculateSum(secondSlice, ch2)

// 	val1 := <-ch1
// 	val2 := <-ch2

// 	fmt.Printf("Sum is %d", val1+val2)
// 	fmt.Println()

// 	fmt.Println("main over")
// }
