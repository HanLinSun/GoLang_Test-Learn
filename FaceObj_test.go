package GoProj01

import "testing"

//Go语言面向对象编程:通过struct
//Go语言不支持继承.
//Go语言接口

type Employee struct{
	ID int
	Skill string
	Name string
	Age int
}
func Test(t *testing.T){
	e:=Employee{4,"C++","Michael",22}
	t.Log(e)
}