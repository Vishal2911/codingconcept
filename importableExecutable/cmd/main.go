package main

import (
	"fmt"

	"github.com/vishal2911/codingconcept/importableExecutable/calulator"
	"github.com/vishal2911/codingconcept/importableExecutable/store"
)


func main() {

	usr := calulator.User{}
	usr.Name = "gdfgdfgdf"
	// fmt.Println("sum = ", calulator.add(54646,5464))
	fmt.Println("Multiply = ", calulator.Multiply(54654,4564))
	fmt.Println("Divide = ", calulator.Divide(54646,5464))
	fmt.Println("sub = ", calulator.Sub(54646,5464))

	//inport using interface
	var dboprs store.DBOperations
	dboprs = store.Store{}
	dboprs.SaveRecord("just print this as record.....")
}

