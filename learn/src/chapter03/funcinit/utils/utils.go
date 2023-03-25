package utils

var Age int
var Name string

//如果需要在其他文件里面引用Age和Name，需要先定义(初始化)这两个
//可以使用init函数进行初始化
func init() {
	Age = 200
	Name = "tomtao626"
}
