package main

import (
	"fmt"
	"sync"
)

// Since wg is a pointer, you need to put the reference of the pointer in the parameter or else you would
// be using the same value always (passing copies)
func one(wg *sync.WaitGroup) {
	defer wg.Done() //Decreamenting 1 to the counter
	fmt.Println("hola")
}

func two(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("ni hao")
}

func three(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}

func main() {

	var wg sync.WaitGroup //Tis is a pointer
	wg.Add(3)             //Adding value of 3 to the counter because there is 3 functions\
	go one(&wg)           //Dereferencing it
	go two(&wg)
	go three(&wg)
	/*
		func() {
			defer wg.Done() //Decreamenting 1 to the counter
			fmt.Println("hola")
		}()

		func() {
			defer wg.Done()
			fmt.Println("ni hao")
		}()

		func() {
			defer wg.Done()
			fmt.Println("hello")
		}()
	*/

	wg.Wait() //Once the counter is 0 or NULL, Continue with the main

}
