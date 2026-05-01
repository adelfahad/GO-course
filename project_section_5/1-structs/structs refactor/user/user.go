package user

import (
	"errors"
	"fmt"
	"time"
)

type Userinfo struct {
	FName     string
	LName     string
	Birthdate string
	CreateAt  time.Time
}
type admininfo struct {
	email    string
	password string
	Userinfo
}

func NewAdminInfo(email, password string, u Userinfo) (*admininfo, error) {
	if email == "" || password == "" {
		return nil, errors.New("invalid admin info")
	}
	return &admininfo{
		email:    email,
		password: password,
		Userinfo: u,
	}, nil
}
func NewUserInfo(fNameInPut, lNameInPut, birthdateInPut string) (*Userinfo, error) {
	if fNameInPut == "" || lNameInPut == "" || birthdateInPut == "" {
		return nil, errors.New("invalid user info")
	}
	return &Userinfo{
		FName:     fNameInPut,
		LName:     lNameInPut,
		Birthdate: birthdateInPut,
		CreateAt:  time.Now(),
	}, nil
}
func (u Userinfo) PrintUserInfo() {
	fmt.Println("user Info")
	fmt.Println("First Name: ", u.FName)
	fmt.Println("Last Name: ", u.LName)
	fmt.Println("birthdate: ", u.Birthdate)
	fmt.Println("Account Create: ", u.CreateAt)
}
func (u *Userinfo) ChangeUser(newUser *Userinfo) {
	u.FName = newUser.FName
	u.LName = newUser.LName
	u.Birthdate = newUser.Birthdate
	u.CreateAt = newUser.CreateAt
}

func Userinput(output string) string {
	fmt.Print(output)
	var value string
	fmt.Scanln(&value)
	return value
}
