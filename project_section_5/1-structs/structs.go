package main

import (
	"fmt"
	"time"
)

type userinfo struct {
	FName     string
	LName     string
	Birthdate string
	CreateAt  time.Time
}

func main() {
	var user1 userinfo = userinfo{"adel", "fahad", "1/1/2000", time.Now()}
	var user2 userinfo
	fNameInPut := userinput("Please enter your first name: ")
	lNameInPut := userinput("Please enter your last name: ")
	birthdateInPut := userinput("Please enter your birthdate: ")
	user2 = userinfo{
		FName:     fNameInPut,
		LName:     lNameInPut,
		Birthdate: birthdateInPut,
		CreateAt:  time.Now(),
	}
	PrintUserInfo(user1, &user2)
}
func PrintUserInfo(u1 userinfo, u2 *userinfo) {
	fmt.Println("user1 Info")
	fmt.Println("First Name: ", u1.FName)
	fmt.Println("Last Name: ", u1.LName)
	fmt.Println("birthdate: ", u1.Birthdate)
	fmt.Println("Account Create: ", u1.CreateAt)
	fmt.Println("user2 Info")
	fmt.Println("First Name: ", u2.FName)
	fmt.Println("Last Name: ", (*u2).LName)
	fmt.Println("birthdate: ", u2.Birthdate)
	fmt.Println("Account Create: ", u2.CreateAt)
}

func userinput(output string) string {
	fmt.Print(output)
	var value string
	fmt.Scan(&value)
	return value
}
