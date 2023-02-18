package main

import "fmt"

func main() {

f := 5
g := 6

sum := add(f , g)
fmt.Println(" add = ", sum)

fmt.Println(" sub = ", sub(f , g))

fmt.Println(" multiply = ", multiply(f , g))

// failing case for divide
divide , err := divid(f , 0)
if err == nil {
	fmt.Println(" divide = ",divide)
}else {
	fmt.Println(" divide hi nahi hua , koi error aa jaya , aur error ki value hai ",err)
}
// passing case for divide
divide , err = divid(f , 2)
if err == nil {
	fmt.Println(" divide = ",divide)
}else {
	fmt.Println(" divide hi nahi hua , koi error aa jaya , aur error ki value hai ",err)
}

fmt.Println("f = ", f , " g = ", g)
}

// for adding two numbers
func add(a , b int) int {
	c := a + b
	return c 
}

// for subtractinc two numbers
func sub(a , b int) int {
	return a - b
}

// for multiply two numbers
func multiply(a , b int) int {
	return a * b
}

// for dividing two numbers
func divid(d , e int) (int , error) { // d = 5 , e = 6 
	if e != 0 {
		return d / e , nil  // 0
	}else {
		return 0 , fmt.Errorf("dividing with 0 is not possible") // d = 5654654 , e = 0 -> 0
	}
}


