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

type UserOperations interface{
	addUser()
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

	// vishal.addUser()
	// shivam.addUser()

	var userOperations UserOperations
	userOperations = vishal // UserA
	userOperations.addUser()

	userOperations = shivam
	userOperations.addUser()

}


func (user UserA) addUser() {
	user.Name = "454564564654564"
	fmt.Println("User a wala method char raha hai , user = ", user)
}

func (user UserB) addUser() {
	fmt.Println("User B wala method char raha hai , user = ", user)
}
