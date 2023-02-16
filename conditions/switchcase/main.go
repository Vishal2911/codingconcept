package main

import "fmt"

func main() {

	daysofweek := 3

	switch daysofweek {
	case 1: //10 == 1 -> false
		{
			fmt.Println("Monday")
		}
	case 2: // 10 == 2 -> false
		{
			fmt.Println("Tus")
		}
	case 3:
		{
			fmt.Println("Wed")
		}
	case 4:
		{
			fmt.Println("Thus")
		}
	case 5:
		{
			fmt.Println("Fri")
		}
	case 6:
		{
			fmt.Println("Sat")
		}
	case 7:
		{
			fmt.Println("Sun")
		}
	default:
		{
			fmt.Println("not a week day")
		}

	}

}
