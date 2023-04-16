package main

import (
	"fmt"
	//"golang-note/advanced/go_progress/utils"
	//"errors"
)

type User struct {
	Id   int
	Name string
}

// func (u *User) SetName(name string) (int,bool,error) {
func (u *User) SetName(name string) {
	u.Name = name
}

func ValidIsLogin(is_loging int) bool {

	if is_loging == 1 {
		//return nil
		return true
	} else {
		return false

		//return errors.New(utils.Unlogin)
		//return "未登陆"
	}

}
func main() {

	day := 8

	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:

		fmt.Println("错误的天数")

	}

	is_login := ValidIsLogin(1)

	if is_login {
		fmt.Println("已登陆")
	} else {
		fmt.Println("未登陆")
	}
}
