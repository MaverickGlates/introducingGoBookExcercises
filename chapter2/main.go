package main

import "fmt"

//DATA TYPES

//NUMBERS: in programming we have two types of number, we split in integers and floating-point numbers.

/*INTEGERS: Are numbers without decimal. Our system it's compuden of 10 numbres from 0 to 9, binary system have only 2 numbers 0 and 1.
So we represent binary system trought bits and bytes where a byte it's compouden by 8bits from "0000 0000" to "1111 1111" in binary system, or
"0" to "255" in decimal system" in GO integers are represent for uint and int.
where type uint means that the number don't contemple a signed minus, so if we have a integer of 8bits (int8 in GO), in decimal we have a range from -128 to 127,
but if we declare uint8 type value for a variable, the positive range up from 0 to 255
*/

/*FLOTATING-POINTS: Floats number are numbers that containt a decimal component (like 1.24; 2.66; 54.6, etc.)Their actual representation on a computer is
fairly complicated and it's not really necessary to know the particulars in order to use them. Go has two floating-point types: float32 and float64. Using
a larger-sized floating-point number increases its precision and this is why we referred to them as single precision and double precision, respectively.
It also has two additional types for representing complex numbers. complex64 and complex128.
Generally, we should stick with float64 when working with floating-point numbers.  */

//Float-Point precision example subtraction operation.
//We Statement 2 variables, a and b, assign a float64 type value and print (a - b) with Println func inside main func.
var a float64 = 1.01
var b float64 = 0.99

func main() {
	fmt.Println(a - b)              //terminal will return 0.020000000000000018 and not 0.02
	fmt.Println("1 + 1 =", 1+1)     //Example 1, we print the string 1+1 followed by the result of the expression 1+1 compouden by 2 numeric literal and 1 operator plus
	fmt.Println("1 + 1 =", 1.0+1.0) //Here we statment the floating-point using a .0 to indicate at GO the Type Value, the result in command prompt it's the same, 2.

	// Go supports the following standard arithmetic operators:
	// "+"Addition	 "-"Subtraction 	"*"Multiplication 	"/"Division		 "%"Remainder
	// ------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	/* STRINGS: String is a sequence of characters with a definite length used to represent the text. Strings are made up of individual bytes per characters
	(characters from some languages, such as Chinese, are represented by more than one byte)
	 Strings can be created using double quotes or backticks, the diff between both are that, the first statment allow to you special escape sequences like the next:
	\n gets replaced with a newline and \t gets replaced with a tab character  */

	fmt.Println(len("Hello, World")) // A space is also considered a character, so the string's length is 12, not 11
	fmt.Println("Hello, World"[0])   // Strings starting at 0 not 1 this why in print we return a value from H, if we wanna the value from "e" put 1 for l put 2 successively
	fmt.Println("Hello, " + "World")

	/* BOOLEANS:A boolean value is a special 1-bit integer type used to represent true and false (or on and off).
	Three logical operators are used with boolean values: && and || or ! not */

	fmt.Println(true && true)  // Return true
	fmt.Println(true && false) // Return false
	fmt.Println(true || true)  // Return true
	fmt.Println(true || false) // Return true
	fmt.Println(!true)         // Return false
	fmt.Println(!false)        // Return true
	//Usually it's good use truth tables to difine how these operator works

}

//EXERCISES

//1.- How are integers stored on a computer?
/*2.- We know that (in base 10) the largest one-digit number is 9 and the largest two-digit number is 99. Given that in binary the largest two-digit number is 11 (3),
the largest three-digit number is 111 (7) and the largest four-digit number is 1111 (15), what's the largest eight-digit number? (Hint: (10^1)−1 = 9 and (10^2)−1 = 99) */
/*3.- Although overpowered for the task, you can use Go as a calculator. Write a program that computes 32,132 × 42,452 and prints it to the terminal (use the * operator
for multiplication).*/
//4.- What is a string? How do you find its length?
//5.- What's the value of the expression (true && false) || (false && true) || !(false && false)?

//ANSWER

//1.- Integers are stored in byte form that is, using a base-2, binary number system.
//2.- The largest eight-digit number in binary is 255.

/*3.-

  package main

  import "fmt"

  func main (){
	  fmt.Println(32132 * 42452)
  }
*/

//4.-A string is a sequence of characters. You can find the length of a string using the len function (len("a string")).
/* 5.-

for (true && false) return false
for (false && true) return false
for !(false && false) return true
for expression (true && false) || (false && true) || !(false && false) return ever true even changing de values

*/
