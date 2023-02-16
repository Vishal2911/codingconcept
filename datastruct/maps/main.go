package main

import "fmt"


func main() {

	// map decleration
	o := make(map[int]int)
	// save data in map
	o[5] = 10
	o[65465465] = 0

	// map decleration
	m := make(map[string]int)
	// save data in map
	m["xyz"] = 10
	m["pqr"] = 0

	// read data from map
	val , ok :=  m["pqr"]
	fmt.Println("val = ", val , " ok = ", ok)

	val , ok =  m["abc"]
	if ok {
		fmt.Println("value found , val = ",val)
	}else{
		fmt.Println("map me value hai hi nahi")
	}




}