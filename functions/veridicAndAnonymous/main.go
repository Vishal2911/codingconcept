package main

import "fmt"

func main() {

	func (a int)int{
		fmt.Println("Anonyomus func is running...")
		return 0
	}(5)

	fmt.Println("Sum = ", add(5 , 6 , 7 , 1))

}

func add(b  ...int) int{
	sum := 0
	for i , val := range b {
		fmt.Println(" i = ", i , "val = ",val)
		sum += val
		//sum = sum + val // 1st case -> 0 + 5 = 5, 2nd case -> 5 +  6  = 11 , 3rd case  11 + 7 = 18
	}
	return sum

}