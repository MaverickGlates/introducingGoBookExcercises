package main

import (
	"fmt"
)

//this is a comment

// //Exercises
// 1. What is whitespace?
// 2. What is a comment? What are the two ways of writing a comment?
// 3. Our program began with package main. What would the files in the fmt package
// begin with?
// 4. We used the Println function defined in the fmt package. If you wanted to use
// the Exit function from the os package, what would you need to do?
// 5. Modify the program we wrote so that instead of printing Hello, World it prints
// Hello, my name is followed by your name.

//1.-Are Special characters, that we can't see. whitespace characters, made up of space, tab. and newline do nothing in the code, we can use it to make programs easier to read
/*2.-Comments are the form to explain a part of code, this is a good practice for the future cause you can read like help the comment or your partners that are working with you.
you can commento a line using // or more than one line using the form to comment this answer*/
//3.-The files in the fmt package would begin with package fmt
//4.-We have to import os package then statement the function Exit with "os.Exit()"
//5.- package main import ("fmt) func main (){fmt.Println("Hello my name is Alexander")}
func main() {
	fmt.Println("Hello, World")

}
