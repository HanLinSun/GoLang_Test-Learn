package main
//Go语言基本操作练习
import "fmt"

type LinkList struct{
	Data int32 //数据域
	Next *LinkList //定义一个结构体类型指针
}//链表结构体

//Go语言 动态内存分配问题?
//头结点
type LinkHead struct{
	HeadNode *LinkList
}
//Head指头结点. (好像要先声明一个该结构体类型的变量作为接受者)
func (Head *LinkHead)IsEmpty()(bool){
	if Head.HeadNode==nil{
		return true
	} else {
		return false
	}
}
//返回链表长度(判断链表非空后)
func (Head *LinkHead) GetLength() int{
	count:=0
	NewNode:=Head.HeadNode//申请一个临时结点,其值为头结点的值
	for NewNode!=nil{
			count++
			NewNode=NewNode.Next
	}
	return count
}

//头插函数
func (Head *LinkHead)InsertHead(value int32) {
	NewNode:=&LinkList{Data:value}//申请一个新的结点,作为缓存
	NewNode.Next=Head.HeadNode
	Head.HeadNode= NewNode//把申请的新节点的值赋给Head
}

//指定位置index插入
func (Head *LinkHead)Insert(value int32,index int){
	NewNode:=&LinkList{Data:value}
	PresentNode:=Head.HeadNode
	if index<=0{//
		Head.InsertHead(value)
	}else if index > Head.GetLength(){
		Head.InsertEnd(value)

	}else {
		count:=0
		for count< index-1{
			PresentNode=PresentNode.Next
			count++
		}
		//插入结点
		PresentNode.Next=NewNode.Next
		PresentNode.Next=NewNode
	}
}

//尾插函数
func (Head *LinkHead)InsertEnd(value int32) {
	NewNode := &LinkList{Data: value}
	Node := Head.HeadNode //遍历用结点
	if Node == nil { //头结点为空
		Head.HeadNode = NewNode //直接把新节点赋给头结点
	} else {
		for Node.Next != nil {
			Node = Node.Next
		}//当Node.Next=nil,就是最后一个节点
		Node.Next = NewNode
	}
}
//指定位置删除节点
func (Head *LinkHead)Delete(index int){
	count:=0 //计数器
	Node:=Head.HeadNode
	if index<=0{//索引小于0，就删除头结点
		Head.HeadNode=Node.Next
	}else if index<Head.GetLength(){

		for count!=index-1&&Node.Next!=nil{
			Node=Node.Next
			count++
		}
		Node.Next=Node.Next.Next
	}else if index>Head.GetLength(){
		fmt.Print("超出链表长度")
	}

}
//打印链表
func (Head *LinkHead)PrintList(){

	if !Head.IsEmpty(){
		Node:=Head.HeadNode
		for{
			fmt.Printf("\t%v",Node.Data)
			if Node.Next!=nil{
				Node=Node.Next
			}else{
				break
			}
		}
	}

}
func main(){
	linkList:=LinkHead{}
	fmt.Println("Insert")
	linkList.InsertHead(2)
	linkList.InsertHead(4)
	linkList.InsertHead(10)
	linkList.InsertEnd(7)
	linkList.InsertEnd(221)
	linkList.PrintList()
	fmt.Println(" ")
	fmt.Println("Delete")
	linkList.Delete(1)
	linkList.PrintList()
}
