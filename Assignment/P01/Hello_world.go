package main

//need to import the "fmt" module to allow printing
import "fmt"

func main() {
	// this is a small message for my first go program
	fmt.Println("Hello world! this is my first GO program ever.")

	//created my first int varliable in go, when creating varibales in Go you need to define "var"
	var halfevil int = 333

	//fmt.print and fmt.printf statment need to research what they mean
	fmt.Print("The number for half evil is : ", halfevil, " The best fucking number in world or whateva.")

	fmt.Print("This program is used to demonstracte my useage of go the language.")
	fmt.Printf("\n")

	//looping thorugh and array
	number := []int{2, 4, 6, 8, 10}
	for _, num := range number {
		fmt.Print(num, " ", "\n")
	}

	//this is a pointer demonstraton

	i := 489

	p := &i

	j := 500

	t := &j
	*t = 50

	fmt.Println(*p)
	fmt.Println(*t)
}
