package main

import "fmt"

func main() {
	var name string
	var age byte
	var height float32
	var salary float32
	var isPass bool
	// 方式一 fmt.Scanln()
	/*
		fmt.Println("请输入姓名：")
		fmt.Scanln(&name)
		fmt.Println("请输入年龄：")
		fmt.Scanln(&age)
		fmt.Println("请输入身高：")
		fmt.Scanln(&height)
		fmt.Println("请输入薪水：")
		fmt.Scanln(&salary)
		fmt.Println("请输入是否通过考试：")
		fmt.Scanln(&isPass)

		fmt.Printf("姓名：%v\n年龄：%v\n身高：%v\n薪水：%v\n是否通过考试：%v\n", name, age, height, salary, isPass
	*/

	// 方式二 fmt.Scanf()
	fmt.Println("请输入姓名,年龄,身高,薪水,是否通过考试,使用逗号分隔开")
	fmt.Scanf("%s %d %f %f %t", &name, &age, &height, &salary, &isPass)
	fmt.Printf("姓名：%v\n年龄：%v\n身高：%v\n薪水：%v\n是否通过考试：%v\n", name, age, height, salary, isPass)

}
