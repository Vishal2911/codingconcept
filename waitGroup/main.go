package main

import (
	"fmt"
	"sync"
)

func a(mg *sync.WaitGroup) {
	defer mg.Done()
	fmt.Println("a is running...")
	
}


func main() {
	wg := sync.WaitGroup{}
	// var wg sync.WaitGroup
	wg.Add(2)
	go a(&wg)
	go a(&wg)
	
	fmt.Println("*************")
	fmt.Println("-------------")
	wg.Wait()
}