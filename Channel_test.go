package GoProj01

import (
	"sync"
	"fmt"
	"testing"
)

//通道的关闭和广播机制
//简易生产者消费者
func DataProducer(ch chan int, wg *sync.WaitGroup){//wg用于控制线程
	go func(){
		for i:=0;i<9;i++{
			ch<-i
		}
		close(ch)
		wg.Done()
	}()
}
func DataReceiver(ch chan int,wg *sync.WaitGroup){
	go func(){
		for{
			if data,ok:= <- ch;ok{//用Bool值ok来判断通道是否关闭
				fmt.Println(data)
			}else{
				break
			}
		}
		wg.Done()
	}()
}
func TestChannel(t *testing.T){
	var wg sync.WaitGroup
	ch:=make(chan int)
	wg.Add(1)
	DataProducer(ch,&wg)
	wg.Add(1)
	DataReceiver(ch,&wg)
	wg.Add(1)
	DataReceiver(ch,&wg)
}
//向关闭的通道发送消息会导致Panic，从关闭的通道接收值会默认返回定义类型的0值