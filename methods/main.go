package main

import "fmt"

type UserA struct {
	Name      string
	Address   string
	ContactNo int
}

type UserB struct {
	Name      string
	Address   string
	ContactNo int
}

func main() {

	var vishal UserA
	vishal.Name = "Vishal Singh"
	vishal.Address = "Varanasi"
	vishal.ContactNo = 15612312311
	shivam := UserB{
		Name:      "Shivam Singh",
		Address:   "same hai address Varanasi",
		ContactNo: 515656156,
	}
	vishal.addUser(5)
	fmt.Println( "after addUser vishal = ", vishal)

	value := vishal.addUserDub(54545454)
	fmt.Println("value from addDub  =  ", value , " vishal = ", vishal)



	shivam.addUser(5)
	add(656565)

}

func add(a int) {
	fmt.Println("function run ho raha hai , aur a  =  ", a)
}
func (user UserA) addUser(a ...int) {
	user.Name = "454564564654564"
	fmt.Println("User a wala method char raha hai , user = ", user)
}

func (u *UserA) addUserDub(a ...int) int {
	u.Name = "Name value change kar raha hu"
	fmt.Println("User a dub wala method char raha hai , user = ", u)
	return 0 
}
func (user UserB) addUser(a int) {
	fmt.Println("User B wala method char raha hai , user = ", user)
}
