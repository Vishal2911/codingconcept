package main

import (
	"fmt"

	"github.com/vishal2911/codingconcept/importableExecutable/calulator"
)


func main() {

	usr := calulator.User{}
	usr.Name = "gdfgdfgdf"
	// fmt.Println("sum = ", calulator.add(54646,5464))
	fmt.Println("Multiply = ", calulator.Multiply(54654,4564))
	fmt.Println("Divide = ", calulator.Divide(54646,5464))
	fmt.Println("sub = ", calulator.Sub(54646,5464))
}

