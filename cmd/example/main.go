package main

import (
	"fmt"
	"os"
	"strings"
	. "github.com/AnastasiaYarema/design-lab-1/implementation"
	. "github.com/AnastasiaYarema/design-lab-1/build"
)

func main() {
  // print current build version
	fmt.Println("Build version: " + BuildVersion)

	// join command line arguments into the string by whitespace separator
	postfixExpression := strings.Join(os.Args[1:], " ")
	
	// calculate the result value of postfix expression
	result, err := CalculatePostfixExpression(postfixExpression)

	// check if error is not nil
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	// print the result
	fmt.Printf("Result: %f\n", result)
}
