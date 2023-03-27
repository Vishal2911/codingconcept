package main

import "fmt"


func main() {
	for i := 1 ; i < 11 ; i++ {
		defer fmt.Println(i)
	}
// defer koi gadbadi ho to sambhalo
	defer fmt.Println("hello")
	defer fmt.Println("world")

	fmt.Println("koi bhi message ")
	defer fmt.Println("********")
	fmt.Println("koi bhi message ----------- ")
}