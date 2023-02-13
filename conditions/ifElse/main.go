package main

import "fmt"


func main() {

	event := "ruk ja bhai aaram kar le"

	if event == "open door" { 
		fmt.Println("gate kholo")
	} else if event == "close door"{
		fmt.Println("door close")
	} else {
		fmt.Println("pata nahi kya karna hai, bola hai ",event)
	}

	num := 11
	if (num == 0 || num > 0) && (num % 2 == 0) { 
		fmt.Println("value positive hai aur even bhi hai +++++++++")
	}else if (num == 0 || num > 0) && (num % 2 != 0) { 
		fmt.Println("value positive hai aur odd bhi hai +++++++++")
	} else if num < 0{
		fmt.Println("negetive value hai ",num)
	} else {
		fmt.Println("positive value hai ---------",num)
	}


}


