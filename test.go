package main

import "fmt"
import "github.com/ibtisaambutt/golangexamples"

func main() {
	
	// Printing the main heading
	fmt.Printf("* Concurrent and Distributed Systems – Assignment 3 *\n")
	
	// Decaring test array
	testArray := make([]byte, 2)
	testArray[0] = 50
	testArray[1] = 51

	//Function 1
	fmt.Println("golangexamples.ConcatSlice: ", golangexamples.ConcatSlice(testArray))

	//Function 2
	golangexamples.Encrypt(testArray, 2)
	fmt.Println("golangexamples.Encrypt: ", testArray)

	//Function 3
	fmt.Println(golangexamples.EZGreetings("Ibtisaam"))
}