package main

import "fmt"

//using the learning until final of chapter 3, we go to writing a useful program.
// let's try with a counter program that counts to 10 from 1.

func main() {
	fmt.Println("1 secuence double quotes")
	fmt.Println("2 secuence double quotes")
	fmt.Println("3 secuence double quotes")
	fmt.Println("4 secuence double quotes")
	fmt.Println("5 secuence double quotes")
	fmt.Println("6 secuence double quotes")
	fmt.Println("7 secuence double quotes")
	fmt.Println("8 secuence double quotes")
	fmt.Println("9 secuence double quotes")
	fmt.Println("10 secuence double quotes")
	fmt.Println()
	// This probably is we can do with all we are learned til now, but it's not really efficient.

	fmt.Println(`1 secuence with backticks
2 secuence with backticks
3 secuence with backticks
4 secuence with backticks
5 secuence with backticks
6 secuence with backticks
7 secuence with backticks
8 secuence with backticks
9 secuence with backticks
10 secuence with backticks`) // or maybe this is another way that we would use
	fmt.Println()
	//Both programs are very tedious to write, so we need to use a short way, more advance, but which can do something multiple times

	//THE FOR STATEMENT

	//"for" is a type statement to allow us repeat the code, like the print function declared before, with this we have to write just one instruction for do 10 statements or the ones you want

	i := 1
	for i <= 10 {
		fmt.Println(i, "Secuence with for statement")
		i += 1
	}

	//this is a clear example of efficiency, with a bitly of code we get the same result.
	/* Explain: first we statement a variable called i, that is storing the value we want print, followed by for loop statement using the keyword for,
	providing a conditional expression that is either true or false and finally supplying a block to execute the code.
	So, for variable "i" we are saying, if "i" it's minus or equal to 10, execute the body instruction (the body just print the number assign to variable i and finally sum 1 to i)
	this create a bucle, cause initialy "i" have the value 1 assigned, after first execution "i" value is 2 cause we sum 1 in the last line of body sentences. this will repeat once and again
	til i get a value equal 10.
	It's very important have this in mind, cause if we don't define the rule for increment value "i += 1" or "i = i + 1", our for statment, would enter in a infinite loop
	cause our expression i <= 10 would be ever smaller than 10, this return true forever */

	// Other Programming langs, have a lot of different types of loops (while, do, until, foreach...) GO just has the for loop and it can be used in differents ways
	// Example variation to the above code:

	for i = 1; i <= 10; i++ {

		fmt.Println(i, "Secuence with for statement, second way")

	}
	fmt.Println()

	// here we have first the assign "i = 1", then "i <= 10" condition, next to operation "i++", finally the body. so we are printing a string and adding up 1 to i. this Til the "i" get 10 value

	//IF STATEMENT
	//let's modify the program we just wrote to show in each line when the number is even or odd, like this:

	/*
		1 odd
		2 even
		3 odd
		4 even
		5 odd
		6 even
		7 odd
		8 even
		9 odd
		10 even
	*/

	// First we need a way to know whether or not a number is even or odd, this is knowed like an algorithm, where we define the steps to determine when a number is ever o odd.
	//An easy way to find a even number is divide it in 2, if the remainder is equal to 0, so the number is even, otherwise, it's odd. Remember that remainder operator in GO it's called with %

	for i = 1; i <= 10; i++ {
		if i%2 == 0 {

			fmt.Println(i, "Secuence with for statement, second way, number is pair")

		} else {
			fmt.Println(i, "Secuence with for statement, second way, number is odd")
		}

	}
	/* Notice that we used a if statement,this is similar to a for statement in that it has a condition followed by a block code, in resume we choose to do an action based in a condition.
	Here we establish it the remain of i%2 should be equal to 0 for if statement to be true, otherwise, "else" block will executed, in recap the if the i reminder is 0 if block is true
	otherwise, else block is true. The conditions are checked top down and the first one to result in true will have its associated block executed.
	Example:

	if i % 2 == 0 {	"divisible by 2"}       Executed when a number is even
	else if i % 3 == 0 {"divisible by 3"}   Executed when a number is divisible by 3
	else if i % 4 == 0 {"divisible by 4"}   And executed when a number is divisible by 4

	NOTE: this is just a representation, the code structure is not really like that.
	has you can see, if we have a counter from 0 to 10, let's to execute the blocks we defined, but notice the number 8 is divisible by 2 and 4, so when the first condition to be true
	Go never will execute the last block cause the first condition is done

	*/

	//Switch Statment

	//Suppose you wanted to write a program that printed the English names for numbers.Using what you've learned so far, you might start by doing this:
	for i = 1; i <= 10; i++ {
		if i == 0 {
			fmt.Println("Zero")
		} else if i == 1 {
			fmt.Println("One")
		} else if i == 2 {
			fmt.Println("Two")
		} else if i == 3 {
			fmt.Println("Three")
		} else if i == 4 {
			fmt.Println("Four")
		} else if i == 5 {
			fmt.Println("Five")
		}

	}
	// Writing a program in this way would be pretty tedious, so another way of achieving the same result is to use the switch statement. We can rewrite our program to look like this:
	for i = 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknow Number")
		}

	}

	//A switch statment beging with the swtich keyword.
	//The value storage in i is compared to the expression following each case keyword if they are equivalent
	//then the statement following the ":" are executed.

	//like "if statement", "switch" check the code from top down, and the first match execute the block code. Switch also supports a default value, like "if" has "else", "switch" use "default"
	//for, if, and switch are the main control flow statements.

	//------------------------- 	Excercise answer   ---------------------------------------
	fmt.Println()
	counterOfNumberDivisibleBy3()
	fizzBuzzFunction()
}

//EXERCISE

/*
1. What does the following program print?

	i := 10
	if i > 10 {
	fmt.Println("Big")
	} else {
	fmt.Println("Small")
	}
2. Write a program that prints out all the numbers between 1 and 100 that are evenly divisible by 3 (i.e., 3, 6, 9, etc.).
3. Write a program that prints the numbers from 1 to 100, but for multiples of three, print "Fizz" instead of the number, and for the multiples of five, print
"Buzz." For numbers that are multiples of both three and five, print "FizzBuzz".
*/

//ANSWER

//1.- This program print "Small" cause i has assigned value 10, the condition is that if i is bigger than 10 print big, but i is equal to 10.
//2.-

func counterOfNumberDivisibleBy3() {
	for counter := 1; counter <= 100; counter++ {
		if counter%3 == 0 {
			fmt.Println(counter)
		}

	}
}

//3.-

func fizzBuzzFunction() {
	for fizzBuzzScan := 1; fizzBuzzScan <= 100; fizzBuzzScan++ {
		if fizzBuzzScan%5 == 0 && fizzBuzzScan%3 == 0 {
			fmt.Println(fizzBuzzScan, "Fizz Buzz")

		} else if fizzBuzzScan%3 == 0 {
			fmt.Println(fizzBuzzScan, "Fizz.")
		} else if fizzBuzzScan%5 == 0 {
			fmt.Println(fizzBuzzScan, "Buzz")
		} else {
			fmt.Println(fizzBuzzScan)
		}

	}
}
