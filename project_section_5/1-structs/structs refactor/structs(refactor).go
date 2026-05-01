package main

import (
	"fmt"
	"time"

	"adelf.com/structs/user"
)

func main() {
	var user1 user.Userinfo = user.Userinfo{FName: "adel", LName: "fahad", Birthdate: "1/1/2000", CreateAt: time.Now()}
	fNameInPut := user.Userinput("Please enter your first name: ")
	lNameInPut := user.Userinput("Please enter your last name: ")
	birthdateInPut := user.Userinput("Please enter your birthdate: ")
	user2, err := user.NewUserInfo(fNameInPut, lNameInPut, birthdateInPut)
	admin1, adminErr := user.NewAdminInfo("adel", "123456", user1)

	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	if adminErr != nil {
		fmt.Println("error : ", adminErr)
		return
	}
	user1.PrintUserInfo()
	user1.ChangeUser(user2)
	user1.PrintUserInfo()
	admin1.PrintUserInfo()
}
