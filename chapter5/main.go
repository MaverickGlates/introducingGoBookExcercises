package main

import "fmt"

//In Chapter 2, we learned about GO's basic types. In this chapter, we will look at three more built-in types: arrays, slices, and maps

//ARRAYS
//An Array is a numbered sequence of elements of a single type with a fixed lenght. in GO they look like this:
//(var x [5]int) x is an example of an array that is composed of five ints. Try running the following program (Line 12, Line 16):

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	//You should see this: [0 0 0 0 100]

	/*x[4] = 100 should be read "set the fifth element of the array x to 100." It might seem strange that x[4] represents the fifth element instead of the fourth,
	but like strings, arrays are indexed starting from 0. Arrays are accessed in a similar way. We could change fmt.Println(x) to fmt.Println(x[4]) and we would get 100.*/
	fmt.Println(x[4]) //Return 100

	//Example program that use array (Line 24, Line 35)

	var z [5]float64
	z[0] = 98
	z[1] = 93
	z[2] = 77
	z[3] = 82
	z[4] = 83

	var totalScore float64 = 0
	for i := 0; i < 5; i++ {
		totalScore += z[i] // this is the same which (total = total + z[i])
	}
	fmt.Println(totalScore / 5)
	//This program computes the average of a series of test scores. If you run it, you should see 86.6. Let's walk through the program:
	//1.- First we create an array of lenght 5 to storage our test score, then we fill up each element with a grade.
	//2.- Next, we set up a for loop to compute the total score adding up all five elements
	//3.- Finally, we divide the total score by numbers of elements to find the average.
	/*This program works, but Go provides some features we can use to improve it. Specifically, i < 5 and total / 5 should throw up a red flag for us. Say we changed the
	number of grades from 5 to 6. We would also need to change both of these parts. It would be better to use the length of the array instead*/

	// So let's to fix our program (Line 45, Line 49):

	totalScore = 0                // assign 0 value again to compute the average, cause in the above code, totalScore storaged the average like a last value
	for i := 0; i < len(z); i++ { // here we are saying that as long as i is less than the length of z we adding up one unit to i
		totalScore += z[i]
	}
	fmt.Println(totalScore / float64(len(z))) //Finally, we divide the total score by the length of z, but totalScore and len(z) are differents types value.

	// totalScore like variable is a float64 type, while len(z) is an integer, if we statement the operation without assign the same value type (float64) to len(z) we should be get an error.
	// Error says # command-line-arguments chapter5\main.go:49:25: invalid operation: totalScore / len(z) (mismatched types float64 and int)
	// So, to solution the problem we used a type convertion. In general, to convert between types, you use the type name like a function.

	//let's do another change, using a special form of the for loop (Line 57, Line 61)

	/*	totalScore = 0
		for i, value := range z {
			totalScore += value
		}
		fmt.Println(totalScore / float64(len(z)))
	*/

	//In this for loop, i represents the current position in the array and value is the same as x[i]. We use the keyword range followed by the name of the variable we want to loop over.
	//Running this program will result in another error: # command-line arguments "i" declared and not used.
	//The Go compiler won't allow you to create variables that you never use. Because we don't use i inside of our loop, we need to change it to this:

	totalScore = 0
	for _, value := range z {
		totalScore += value
	}
	fmt.Println(totalScore / float64(len(z)))

	//A single underscore (_) is used to tell the compiler that we don't need this (in this case, we don't need the iterator variable)
	//Go also provides a shorter syntax for creating arrays, by example:

	y := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(y) // Sometimes arrays	like this can get too long to fit on one line, so Go allows you to break it up like the next:

	az := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println(az)

	// Notice the extra trailing , after 83. This is required by Go and it allows us to easily remove an element from the array by commenting out the line

	az = [5]float64{
		98,
		93,
		77,
		82,
		//83,
	}

	/* Because the length of an array is part of its type name, working with arrays can be a little cumbersome. Adding or removing elements as we did here requires also changing
	the length inside the brackets. Because of this and other limitations, you rarely see arrays used directly in Go code. Instead, you will usually use a slice, which is a type built on
	top of an array.*/

}
