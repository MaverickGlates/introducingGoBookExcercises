package main

import "fmt"

//A variable is a storage location, with a specific type and an associated name.

var outMainFunctionVar string = " 30"

func main() {

	var x string = "Hello, World" /* Variables in Go are created by first using the var keyword, then specifying the variable name (x) and the type (string),
	and finally, assigning a value to the variable (Hello, World)*/
	fmt.Println(x)

	var z string
	z = "Hello, World" //  Assigning a value is optional, so we could use two statements,like this. first  statement for create z, second statement for assign value to z.
	fmt.Println(z)
	//It's better read variables like z or x takes a "Hello, World" string value, and not to read saying, this variable x is equal to hello world, casue them are variables finally
	x = "second value"
	fmt.Println(x) // This is an example that the value of a variable can change
	x = "third value"
	fmt.Println(x) // Here, the value that we are printing it's changing again, so the order value for x is, "Hello, World", then "second value" and finally "third value"
	x = x + ", string"
	fmt.Println(x) //this is another way where the value can change, cause your variable it's concatenated, fact is that variables are not equal to just one value

	/* Notice that variables in GO are similar to Algebra, them has subtle differences, the first it's related with sign, for example if we a have a variable x = "a string"
	this variable it's not really equal to the value mentioned, the true is that x it's taking that value, this is why in GO exist an double equal operator == cause when we need
	statement a equality, we have to use == to stablish a equality between two values. in recap = it's for assing value, double equal == it's used to compare equality
	and also == it's an operator, like +, or *, etc, but this operator returns a boolean. */

	var string1 string = "hello"
	var string2 string = "world"
	fmt.Println(string1 == string2) // A print func for do a operation ==, the return is obviously false cause "hello" it's not equal to "world"
	string1 = "hello"
	string2 = "hello"
	fmt.Println(string1 == string2) // This case, the return from operation it's true, because the variables string1 and string2 are the same assigned value
	// In GO is common do a statement with a type value and starting value, this is why GO also supports shorter statements, to do it fastly

	y := "Hello World, y variable" //  in this case we statement y, then define value assign. notice that for this declaration you need put : before =, GO infer the type value
	fmt.Println(y)

	l := 2 //same type statement like before, and GO infer that value is an integer, use this statement type whenever possible
	fmt.Println(l)

	/* Languages often have a set of informal, conventional rules. For example, in English, there are rules that govern the order of adjectives in a sentence.
	If, instead of "my small black cat", you were to say, "my black small cat", it would strike native speakers as very strange.

	Go programmers (as well as programmers in other communities) often refer to this as the idiomatic usage of the language. Learning idiomatic Go is a worthwhile pursuit,
	but at this stage, you should probably focus on simply writing correct programs, as that's challenging enough. Don't let the best be the enemy of the good.*/

	//HOW TO NAME VARIABLES

	/*Naming a variable properly is an important part of software development, all names must start with a letter and may content letters, numbers or underscore symbol.
	The GO compiler doesn't care the name you assign to a variable,  but you should choose names that clearly describe the variable's purpose*/

	x = "Max"      // in this case we will change value of x by a name like string value, then print it. (Max is a dog)
	fmt.Println(x) // if we execute de programm with "go run main.go" command, we print the name of a dog called Max, but this is accesing to variable x. not very good name for a variable.

	x = "variable x for math" // so ignore x and we declare a new variable
	fmt.Println(x)

	name := "Max" // Here an example of good name variable for a name value
	fmt.Println(name)

	dogsName := "Max"
	fmt.Println(dogsName) /* or well,  A better name would be dogsName for dogs name. How you can see, this variable is using camelCase, which is a style for writing compound
	words in which the first letter of each new word or phrase is capitalized */

	// SCOPE
	/* Let's take another look at the program we saw at the beginning of the chapter 1:
	--------------------------------
		package main

			import "fmt"
			func main() {
			var x string = "Hello, World"
			fmt.Println(x)
		}
	--------------------------------
		Another way of writing this program would be like this:
	--------------------------------
		package main

		import "fmt"

		var x string = "Hello, World"

		func main() {
		 fmt.Println(x)
		}
	--------------------------------
		 so let's do it in code */
	fmt.Println(outMainFunctionVar) // here we are printing an integer that's declared outside of the main func, this means that other functions can access this variable
	var insideMainFunctionVar string = "asd"
	fmt.Println(insideMainFunctionVar)

	f()

}

func f() {
	fmt.Println(outMainFunctionVar) // here func f is accesing to variable outMainFunctionVar, until here all's good.
	//fmt.Println(insideMainFunctionVar) // here we should see an error, "undeclared name: insideMainFunctionVar" or maybe "Undefine: insideMainFunctionVar"
	//(comment the line 101, to skip error )
	// The range of places where you are allowed to use x is called the scope of the variable

	//CONSTANT

	/* Constant are other type of value, special to assign values that cannot be changed later
	   They are created in the same way you create variables, but instead of using the var keyword we use the const keyword*/

	const excampleConstVariable string = "Hello World, Constant"
	fmt.Println(excampleConstVariable)

	// excampleConstVariable = "random string" // this statement result in a compile-time error
	//Constants are a good way to reuse common values in a program without writing them out each time. For example, Pi in the math package is defined as a constant.

	//GO also can statement multiple variables like shorthand
	var (
		a = 5
		b = 3
		c = 6
	)
	fmt.Println(a, b, c)
	inputValueByUser()    // the function executed ask to user the value and then return the same value multiplied by 2
	fahrenheitToCelsius() // Fahrenheit to Celsius calculator for exercise number 5 from chapter 3.
	feetToMeters()        // Feet to Meters calculator for exercise number 6 from chapter 3.

}

//An Example Program

func inputValueByUser() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

	/* This func is made up to a entry and one exit, where we have the same func to print characters in console and then a function to request data at user.
	it's here where the compile-time is stopping, casue Scanf is another function from the fmt package to read the user input and fill with the number we enter */
	//in this example we called at the function inputValueByUser at the final of main function (line 125)
}

// EXERCISES
/*
1. What are two ways to create a new variable?
2. What is the value of x after running x := 5; x += 1?
3. What is scope? How do you determine the scope of a variable in Go?
4. What is the difference between var and const?
5. Using the example program as a starting point, write a program that converts
from Fahrenheit into Celsius (C = (F − 32) * 5/9).
6. Write another program that converts from feet into meters (1 ft = 0.3048 m)
*/

//ANSWER

//1.- can creat a variable trought long wat or short way, using a explicit statement with the var keyword, then nameVariable, next to type value and finally the value assign or use := for go infer the type value
//2.- The value is 6, cause we are additioning 1 to the intial value of x.
//3.- Scope is the range of place where can be used a variable. the scope is define by blocks, so is important looking for curly brackets, if a variable is inside a function this var it's private.
//4.- var is a keyword to statement a memory location with a changing value, a constant it's the same, but with the restriction that value can't change.
//5.-

func fahrenheitToCelsius() {
	fmt.Println("Enter a number: ")
	var fahrenheitInput float64
	fmt.Scanf("%f", &fahrenheitInput)
	celsiusOutput := (fahrenheitInput - 32) * 5 / 9
	fmt.Println(celsiusOutput)

} // This function use de same structure which last function called "inputValueByUser" but oriented to a life problem about unit conversions

//6.-
func feetToMeters() {
	fmt.Println("Enter a number: ")
	var feetInput float64
	fmt.Scanf("%f", &feetInput)
	metersOutput := feetInput * 0.3048
	fmt.Println(metersOutput)
}
