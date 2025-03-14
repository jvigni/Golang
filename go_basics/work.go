package main

import (
	"fmt"
	"gotest/calculator"
	"log"
	"strings"
	"sync"
	"time"
)

type IAnimalSpeak interface {
	Speak()
}

type animal struct {
	IAnimalSpeak
	name string
	age  int
}

type dog struct {
	animal
	breed string
}

func newDog(breed, name string, age int) *dog {
	d := &dog{
		animal: animal{
			name: name,
			age:  age,
		},
		breed: breed,
	}
	return d
}

func (d *dog) Speak() {
	fmt.Println("Woof! Woof!")
}

type cat struct {
	animal
	color string
}

func (d *cat) Speak() {
	fmt.Println("Miau")
}

func main() {
	defer fmt.Println("\nMAIN END") // Defer statement directly with function call
	Slice()
	fmt.Println("--------------------")
	Polymorphism()
	fmt.Println("--------------------")
	Calculator()
	fmt.Println("--------------------")
	Goroutines()

	//fmt.Scanln()
	//os.Exit(0)
}

func Goroutines() {
	fmt.Println("Sending Workers...")

	var wg sync.WaitGroup
	results := make(chan int)
	numGoroutines := 3
	startTotal := time.Now() // Record total start time

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("Work work! [worker:", id, "]")

			// Simulate some work
			time.Sleep(time.Duration(3) * time.Second)

			// Send result back through results channel
			results <- 40
		}(i)
	}

	// Use another goroutine to wait for all goroutines to finish
	go func() {
		wg.Wait()      // Wait for all goroutines to finish
		close(results) // Close the results channel after all goroutines are done
	}()

	// Collect results from channel
	sum := 0
	for result := range results {
		sum += result
	}

	totalElapsed := time.Since(startTotal) // Calculate total elapsed time
	fmt.Printf("All goroutines have finished! Total elapsed time: %v\n", totalElapsed)
	fmt.Printf("Total sum of results: %v\n", sum)
}

func Calculator() {
	div, err := calculator.Divide(4, 2)
	if err != nil {
		log.Fatal(err)
		//log.Panic(err)
	}
	fmt.Println("Division Result:", div)
}

func Polymorphism() {
	dog := newDog("Husky", "Chipi", 3)
	fmt.Printf("%v %v, %v, %v\n", "Dog:", dog.name, dog.age, dog.breed)
	dog.Speak()
	cat := cat{animal: animal{name: "Kitty", age: 2}, color: "White"}
	fmt.Printf("%v %v, %v, %v\n", "Cat:", cat.name, cat.age, cat.color)
	cat.Speak()
}

func Slice() {
	s := []string{"1", "2", "3", "4"}
	sliced := s[:2]
	fmt.Println("\nSliced:", strings.Join(sliced, " "))
}
