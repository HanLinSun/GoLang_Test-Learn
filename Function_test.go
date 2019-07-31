package GoProj01

import (
	"math/rand"
	"testing"
	"fmt"
)

//多返回值
func returnMultiValues()(int,int){
	return rand.Intn(10),rand.Intn(20)
}

func TestFunction(t *testing.T){
	a,_:=returnMultiValues()
	t.Log(a)
}
//可变长参数
func sum(ops ... int)int{
	ret:=0
	for _,op:=range ops{
		ret+=op
	}
	return ret
}
func TestVarParam(t *testing.T){
	t.Log(sum(1,2,3,4))
	t.Log(sum(1,2,3,4,5))
}

func Clear(){
	fmt.Println("Clear all Resources")
}
//Defer(用于安全释放资源和锁)
func TestDefer(t *testing.T){
	defer Clear()
	fmt.Println("start")//Defer 可延迟一段时间处理语句
	panic("expected error")
}
