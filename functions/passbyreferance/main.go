package main

import "fmt"


func main() {

	f := 5
	g := 5

	addbyone( &f, g)

	fmt.Println("f = ",f , " g = ", g)

}


func addbyone(a *int , b int) {
// db connection break
	*a = *a + 1 // reconnect
	b = b + 1

}