package GoProj01

import (
	"testing"
	"fmt"
)

//斐波那契数列
func TestFibTry(t *testing.T){
	var a int = 1
	var b int = 1
	fmt.Print(a)
	for i:=0;i<5;i++{
		fmt.Print("  ",b)
		temp:=a
		a=b
		b=temp+a
	}
}
