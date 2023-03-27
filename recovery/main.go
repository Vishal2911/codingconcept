package main

import (
	"fmt"
	"os"
)

func main() {

	dbconnection := true
	if !dbconnection {
		fatalError("database not connected")
	}

	// normal case
	a := []string{
		"a wala record save hua hai",
	}
	saveToDB(a)

	//panic case
	b := []string{}
	saveToDB(b)

	// normal case
	c := []string{
		"c wala record save hua hai",
	}
	saveToDB(c)

}

func recoverpanic() {
	r := recover()
	if r != nil {
		fmt.Println("kuch to jhol hua hai, panic hua code , message = ", r)
	}
}

func saveToDB(record []string) {
	// defer recoverpanic()
	if len(record) < 1 {
		panicError("no record found to save")
	} else {
		fmt.Println("record = ", record)
	}
}

func fatalError(anything interface{}) {
	fmt.Printf("Fatal error occured , message = %v", anything)
	os.Exit(1)
}

func panicError(anything interface{}) {
	defer recoverpanic()
	fmt.Println("Panic error occured ")
	panic(anything)
}

// 1/5
// 45454/0
// a[1,2,3,4]
//a[4] -> panic index out of bound
