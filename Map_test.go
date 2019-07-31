package GoProj01
//map:一种key和value对应的数据结构
import (
	"testing"
	)

func TestMap(t *testing.T){
	m1:=map[int]int{1:1,2:4,3:6,4:8}
	t.Log(m1[2])
	t.Logf("len m1=%d",len(m1))
	m2:=map[string]string{"France":"巴黎","Germany":"柏林","United Kingdom":"伦敦","Poland":"华沙","United States":"华盛顿"}
	t.Log(m2["Poland"])
	t.Logf("len m2=%d",len(m2))

}

//map数据结构存在的问题:如果一个键不存在，会自动将它的值设为0，和原本存在的键对应的0值无法区分,可能需要用循环另外判断
//不能通过返回nil判断
func TestMapElement(t *testing.T){
	m3:=make(map[int]int,10)
	t.Log(m3[2])
	if v,ok:=m3[2];ok{//v是m3[2]的值，ok赋值若m3[2]存在则为true，否则为false
		t.Logf("Key 2 is valid, its value is %d",v)
	} else {
		t.Log("Key 2 is not existed.")
	}
}

//map遍历
func TestMapTravel(t *testing.T){
	m1:=map[int]int{1:16,2:32,3:64,4:128,5:256}
	for k,v:=range m1{
		t.Log(k,v)
	}
}

