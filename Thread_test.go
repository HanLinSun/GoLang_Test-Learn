package GoProj01

import (
	"testing"
	"fmt"
	"time"
	"sync"
)

//Go语言并发编程（协程）,更轻量级
func TestGoroutine(t *testing.T){
	for i:=0;i<10;i++ {
		go func(i int) {//这个函数能够正确执行，因为Go函数采用值传递，内部的i和外部的i内存地址不一致
			fmt.Printf("%d",i)
		}(i)
		//go func(){
		//fmt.Print(i)
	    //}()
	}
}
//共享内存并发(使用锁进行并发控制)
/*func TestCounter(t *testing.T){
	counter:=0
	for i:=0;i<5000;i++{
		go func() {
			counter++
		}()
	}
	time.Sleep(2*time.Second)
	fmt.Println(counter)
}*/
//上方的代码不是线程安全的，唯一的counter存在资源竞争，输出不是5000,会丢失一些写入操作
//用锁来控制,安全并发(输出和预期一致)
func TestCounterSafe(t *testing.T){
	var mut sync.Mutex //处理锁
	counter:=0
	for i:=0;i<5000;i++{
		go func(){
			defer func (){//在Go中一般使用Defer来处理锁
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(2*time.Second)
	fmt.Println(counter)
}
//WaitGroup sync（相当于Java的Join），只有当wait的所有事件完成才能继续向下执行
func TestCounterWaitGroup(t *testing.T){
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter:=0
	for i:=0;i<5000;i++{
		wg.Add(1)//每开启一个协程等待数就增加1
		go func(){
			defer func (){//在Go中一般使用Defer来处理锁,Defer可用于安全释放内存(在协程内容执行完要释放锁)
				mut.Unlock()
			}()
			mut.Lock()//每开启一个协程就为其增加锁
			counter++
			wg.Done()//协程完成后结束等待
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
//第一种方法是强制等待一个较长的时间使协程全部完成，第二种是明显省时间的方法。能够精确的等待每一个协程执行完毕
