package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int
	a = 10

	var b int16
	b = int16(a)
	fmt.Println("b = ",b)
	var c string
	c = fmt.Sprintf("%d",b)
	fmt.Println("c = ",c)

	num := "123123"
	numint , err := strconv.Atoi(num)
	fmt.Println("numint = ",numint , " err = ",err)

	num = "vishal"
	numint , err = strconv.Atoi(num)
	fmt.Println("numint = ",numint , " err = ",err)


}