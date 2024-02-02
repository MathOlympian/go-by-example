// for is Go's ony looping construct. Here are some basic types of for loops
package main

import "fmt"

func main(){
// The most basic type, with a single condition.

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Print("------------- \n")
	q := 10
	for q <= 20 {
		fmt.Println(q)
		q = q + 2
	}

	fmt.Print("------------- \n")

	p := 3
	for p <= 18{
		// do something
		fmt.Println(p)

		// increment or decrement p
		p += 2
	}
// A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Print("------------- \n")

	for p := 2; p <= 20; p+=4 {
		fmt.Println(p)
	}

// for without a condition will loop repeatedly 
// until you break out of the loop or return from the enclosing function.

	for {
		fmt.Println("loop")
		break
	}

// or you can use it to break out of a loop once a specific condition is fulfilled
	y := 0
	for y <= 10 {
		
		y += 1
		if y >= 9 {
			break
		}	
		fmt.Println("loop go", y)	
	}

// You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for	z := 30; z >= 0; z--{
		if z % 4 != 0 {
			continue	
		}

		fmt.Println(z)
	}  

}
