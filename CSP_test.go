package GoProj01

import (
	"time"
	"fmt"
	"testing"
)

//csp通过Channel完成，channel有容量限制
//单通道只有当接收者和发送者完成一次互动才能传递信息(一方缺席会导致另一方等待)
//消息通道缓冲区
func service () string{
	time.Sleep(50*time.Millisecond)
	return "Done"
}
func otherTask(){
	fmt.Println("I am doing something else")
	time.Sleep(time.Millisecond*50)
	fmt.Println("Task is Done")
}
func TestService(t *testing.T){
	fmt.Println(service())
	otherTask()
}
func AsyncService() chan string{//返回保存了ret的通道,该通道有5个缓冲区
	retch:=make(chan string,5)
	go func(){
		ret:=service()
		fmt.Println("returned result")
		retch<-ret
		fmt.Println("Service exited.")
	}()
	return retch
}
func TestAsyn(t *testing.T){
	select{//多路选择关键字
	case retch:=<-AsyncService():
	t.Log(retch)
	case <-time.After(time.Millisecond*200)://超时
	t.Error("Time out")
	}
}
//机制解释：