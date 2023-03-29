package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	accountmepaisakitnahai := 0
	wg := sync.WaitGroup{}
	// wg.Add(1000)
	for i := 1; i<= 10000 ; i++ {
		wg.Add(1)
		go paisabadhao(&accountmepaisakitnahai , &wg , &mutex)
	}
	wg.Wait()
	fmt.Println(accountmepaisakitnahai)
}

func paisabadhao(balance *int , wg *sync.WaitGroup , tala *sync.Mutex) {
	defer wg.Done()
	tala.Lock()
	*balance++
	tala.Unlock()
}