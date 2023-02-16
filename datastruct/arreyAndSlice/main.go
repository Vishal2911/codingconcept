package main

import "fmt"

func main() {
	// p := 1
	// q := 2
	// r := 3
	// var a , b,  c,  d, f ,g , h , i int
	// a= 10
	var list [10]int
	list[0] = 10
	list[1] = 11
	list[3] = 13
	list[4] = 14
	list[5] = 15
	list[6] = 16
	list[7] = 17
	list[8] = 18
	list[9] = 19
	list[2] = 12

	fmt.Println("list = ",list)

	list[5] = 454

	fmt.Println("list = ",list)

	var names [10]string // [],[],[],[],[],[],[],[],[],[]
	names[0] = "10"

	fmt.Println("4465464545464 ++++++++ names = ",names[1])
	names[1] = "11"
	names[3] = "13"
	names[4] = "14"
	names[5] = "15"
	names[6] = "16"
	names[7] = "17"
	names[8] = "18"
	names[9] = "19"
	names[2] = "12"
	fmt.Println("names = ",names)
	names[5] = "454"
	fmt.Println("names = ",names)

	var students []string // ->
// [],[],[],[],[],[],[],[]
//  0  1  2  3  4  5  6  7


	students = append(students,"vishal") // []
	students = append(students,"xyz") // [],[]

	fmt.Println("students = ",students)

}