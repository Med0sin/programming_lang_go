package main

import "fmt"

func main() {
	fmt.Println("Hello world! this is my first GO program ever.")

	var halfevil int = 333

	fmt.Print("The number for half evil is : ", halfevil, " The best fucking number in world or whateva.")
	fmt.Printf("\n")
	//looping thorugh and array
	number := []int{2, 4, 6, 8, 10}
	for _, num := range number {
		fmt.Print(num, " ")
	}
}
