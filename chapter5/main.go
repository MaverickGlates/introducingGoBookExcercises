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

	/* Because the length of an array is part of it's type name, working with arrays can be a little cumbersome. Adding or removing elements as we did here requires also changing
	the length inside the brackets. Because of this and other limitations, you rarely see arrays used directly in Go code. Instead, you will usually use a slice, which is a type built on
	top of an array.*/

	//SLICES

	// A slice is a segment of an array. Like arrays, slices are indexable and have a length. Unlike arrays, this length is allowed to change. Here's an example of a slice:

	var ax []float64
	fmt.Println(len(ax)) //The only difference between slices and an arrays is the missing length between the brackets. In this case, x has been created with a length of zero.
	// have this in mind, all that you can do with an array, you can do it with a slices with one or two exceptions.

	//If you want to create a slice, you should use the built-in make function like this:

	ay := make([]float64, 5)
	fmt.Println(len(ay)) /* This creates a slice that is associated with an underlying float64 array of length 5. Slices are always associated with some array, and although they can never
	be longer than the array, they can be smaller. The make function also allows a third parameter: */

	thirdParameter := make([]float64, 5, 10)
	fmt.Println(len(thirdParameter))
	fmt.Println(cap(thirdParameter))
	/* How we said before a slice basically is a underlying from an array, now there's an additional function that we have available with slices, the capacity, that's because the the number
	of elements in the slices doesn't necessarily match the size of the backing array, because the slice is really a projection of that underlaying array, this mean that we work with an
	array indirectly, imagine we have a slice length 5 and back it an array length 10, this is exactly example of the representation before in line 117 code. In resume, we can have statement
	a slice and define a length that not necessarily is his max length, because we can define a capacity that matching with the underlying array length*/

	//Another way to create slices is to use the [low : high] expression:

	/* low is the index of where to start the slice, high is the index of where to end it (but not including the index itself). For example, while arr[0:5] returns
	[1,2,3,4,5], arr[1:4] returns [2,3,4] */

	arr := [5]float64{1, 2, 3, 4, 5}
	xArr := arr[1:4] //[Low expression include his element: High expression don't include it] we see 1 is the second element (2) and 4 is the fifth element (5), the print just show us [2 3 4]
	fmt.Println(xArr)

	/*For convenience, we are also allowed to omit low, high, or even both low and high. arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5], and
	  arr[:] is the same as arr[0:len(arr)].*/

	//xArr = arr[0:len(arr)]
	fmt.Println(xArr)

	xArr = arr[:]
	fmt.Println(xArr)

	//In addition to the indexing operator, Go includes two built-in functions to assist with slices: append and copy

	//APPEND
	/*append adds elements onto the end of a slice. if there's sufficient capacity in the underlying array, the element is placed after the last element and the length
	is incremented. However, if there is not sufficient capacity, a new array is created, all of the existing elements are copied over, the new element is added onto the
	end, and the new slice is returned.*/

	//The definition of append can be a bit confusing but it's easier to grasp once you see it used. Here is an exampler:

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	fmt.Println(len(slice1), len(slice2))

	/*After running this program, slice1 has [1,2,3] and slice2 has [1,2,3,4,5]. append creates a new slice by taking an existing slice (the first argument)
	and appending all the following arguments to it*/
}
