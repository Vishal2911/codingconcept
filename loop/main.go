package main

import "fmt"

func main() {

	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println(4)
	// fmt.Println(5)
	// fmt.Println(6)
	// fmt.Println(7)
	// fmt.Println(8)
	// fmt.Println(9)
	// fmt.Println(10)
	var i int
	// start  ; end  ; change 
	for i = -100 ; i <= 100 ; i+=1 {
		fmt.Println(i)
	}

	// second example
	i = 10
	condition := true
	for condition {
		if i >=101 {
			condition = false
		}else{
			i++
			fmt.Println("dusara wala for loop hai",i)
		}
	}

	// range

	var list [10]int
	list[0] = 10
	list[1] = 11
	list[3] = 13
	list[4] = 14
	list[5] = 15
	list[6] = 16
	list[7] = 17
	list[8] = 18
	list[9] = 19
	list[2] = 12

	for i := 0; i < len(list) ; i++ {
		fmt.Println(list[i])
	}
	for _ , j := range list {
		fmt.Println(" value is ",j)
	}

		// map decleration
		m := make(map[string]string)
		// save data in map
		m["xyz"] = "XYZ"
		m["pqr"] = "PQR"
		m["A"] = "a"
		m["B"] = "b"
		m["C"] = "c"

		for key , vcxvcxv := range m {
			fmt.Println("key = ",key , " value = ",vcxvcxv)
		}
}