package GoProj01

import "testing"

func TestMapFactory(t *testing.T){
	m:=map[int]func(op int)int{}
	m[1]=func(op int)int{return op}
	m[2]=func(op int)int{return op*op}
	m[3]=func(op int)int{return op*op*op}
	t.Log(m[1](2))
	t.Log(m[2](2))
	t.Log(m[3](2))
}
//Set唯一集
func TestMapForSet(t *testing.T){
	mySet:=map[int]bool{}
	mySet[1]=true
	n:=3
	if mySet[n]{
		t.Log("1 is existing")
	} else{
		t.Log("1 is not existing")
	}
	mySet[n]=true
	t.Log(len(mySet))
}