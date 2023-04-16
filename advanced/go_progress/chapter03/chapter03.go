package main

import (
	"fmt"
	"sort"
)


//var Name string = "hallen"
//
//var Name = "hallen"

type Student struct {
	Name string
}


var DB string

func Get()(a int,b int)  {

	a = 3
	b = 4

	return a,b

}


func main() {


	//var i = 3
	//i := 3
	//var i int = 3
	//fmt.Println(i)
	//a = 10   // 有问题


	var student *Student

	student = new(Student)
	*student = Student{Name:"hallen"}
	fmt.Println(student)

	fmt.Println(student.Name)
	fmt.Println((*student).Name)
	//fmt.Println((&student).Name)   // 有问题


	fmt.Println("hello"+"hallen")

	fmt.Sprintf("%s%s","hello","hallen")

	var ok bool

	//ok = true
	//ok = 1   // 错误
	//ok = bool(0) // 错误
	ok = (1==1)
	fmt.Println(ok)
	const name  = "hallen"
	//var a int = 4

	//switch a {
	//case 1:
	//	fmt.Println(1)
	//	fmt.Println(2)
	//	fallthrough
	//case 2:
	//
	//}


	var a *int
	a = new(int)
	*a = 6
	*a ++
	fmt.Println(*a)
	fmt.Println(&a)

	var name1 *string
	name1 = new(string)
	*name1 = "hallen"


	fmt.Println((*name1)[0])


	//slice := make([]int,0)

	if !ok {

	}

	var h int = 1

	if h == 2 {

	}

	map_data := map[string]string{"name":"hallen"}
	//json.Unmarshal(map_data,)

	delete(map_data,"name")
	fmt.Println(map_data)

	//var x error = nil

	//var name2 = "hallen"
	//fmt.Println(name2[0])
	//name2[0] = 'z'

	//var v_int int = 3
	//var u_int uint = -1
	//fmt.Println(v_int)
	//fmt.Println(u_int)

	//var u1_int uint = 3
	//var u2_int uint = 4
	//fmt.Println(u1_int-u2_int)


	//slice := make([]int,3)
	//
	//slice = append(slice, 1,2,3)
	//fmt.Println(slice)


	var slice1 []int
	var slice2 []int = []int{}

	if slice1 == nil {
		fmt.Println("slice1 == nil")
	}

	if slice2 == nil {
		fmt.Println("slice2 == nil")
	}


	var aa int

	//aa,_ = Get()
	//aa,_ := Get()  // aa已申明，不需要使用自动推导

	//aa,bb := Get()
	//aa,bb = Get()  // bb没有申明
	fmt.Println(aa)
	//fmt.Println(bb)
	//fmt.Fprintf()

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)


	map_data1 := map[string]interface{}{
		"name":"hallen",
		"age":18,
	}

	keys := []string{}
	for k,_ := range map_data1 {
		keys = append(keys,k)
	}


	sort.Strings(keys)

	for _,v := range keys {
		fmt.Println(map_data1[v])
	}






}
func funcMui(x,y int)(int,error){
	return x+y,nil
}


func GetName() string {
	name := "hallen"
	return name

}

