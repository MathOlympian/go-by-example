package main

import "fmt"

func main (){

	// var declares 1 or more variable
	var a = "initial"

	var q string  = "var string"
	var qu float64 = 123243.44255
	fmt.Println(q, qu)
	fmt.Println(a)	

	// You can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	var you, me string = "YOU", "ME"
	fmt.Println(you, me)
	
	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)
	
	// Variables declared without a corresponding initialization are zero-valued. For example for an int is 0
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable, e.g for var f string = :
	// "apple" in this case. This syntax is only available inside functions.
	
	f := "apple"
	fmt.Println(f) 
}