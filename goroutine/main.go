package main

import (
	"fmt"
	"time"
)

func main() {

	// //parallel
	// ---------|-------|------ > process a or function a
	// ---------|-------|------ > process b or function b
	// ---------|-------|------ > process c or function c
	//  Cores
	// |1->main| 2 -> a|
	// |3 -> b | 4 -> c|

	// //Concurrency
	// ------- -----  | ---|------- - - ------ > process d or function d
	//        -     --|-   |       - - -   	> process e or function e
	//  Cores
	// |1->main or d| 2 -> a or d,e|
	// |3 -> b | 4 -> c|

	// a()
	// b()
	// c()
	// d()

	go a()
	go b()
	go c()
	go d()
	time.Sleep(time.Second * 3)
}

func a() {
	for i := 1 ; i <= 3 ; i++ {
		fmt.Println("Func a run hua hai , i = ",i)
	}

}
func b() {
	for i := 1 ; i <= 3 ; i++ {
		fmt.Println("Func b run hua hai , i = ",i)
	}
}
func c() {
	for i := 1 ; i <= 3 ; i++ {
		fmt.Println("Func c run hua hai , i = ",i)
	}
}
func d() {
	for i := 1 ; i <= 3 ; i++ {
		fmt.Println("Func d run hua hai , i = ",i)
	}
}
