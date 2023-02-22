package calulator

import "fmt"

type User struct{
	Name string
	contno int
}

type student struct{
	Name string
	RollNo int
}

func add(a , b int) int{
	fmt.Println("Add call hua")
	return a+b
}

func Multiply(a , b int) int{
	return a*b
}
func Divide(a , b int) int{
	return a/b
}
