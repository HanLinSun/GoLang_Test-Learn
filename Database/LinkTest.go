package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	)

type User struct {
	Age int
	Name string
}

func Check(err error){//检查和报错函数
	if err!=nil{
		panic(err)
	}
}

func Search(db *sql.DB,age int) {//根据年龄查找特定元素并输出
    fmt.Println("start searching")
    stmt,err:=db.Prepare("select name from hello where age= ?")
    Check(err)
    rows,err:=stmt.Query(age)
    //rows:=QueryRow(age)
    //QueryRow只有一个返回值，只能返回满足条件的一个元素,而且不能够返回err
    //因此使用Query输出可以得到满足条件的所有值
    //同样，有几个?，Query中就有几个值，这里只有一个age。
    Check(err)
    user:=new(User)
    //err=rows.Scan(&user.Name)
    for rows.Next(){
	    err=rows.Scan(&user.Name)
     	Check(err)
     	fmt.Println(user.Name)
}
}

func SearchInRange(db *sql.DB,Minage int,Maxage int){
	fmt.Println("Begin Searching in Range")
	stmtout,err:=db.Prepare("select name from hello where age>=? and age<=?")
	Check(err)
	rows,err:=stmtout.Query(Minage,Maxage)
	Check(err)
	user:=new(User)
	for rows.Next(){
		err=rows.Scan(&user.Name)
		Check(err)
		fmt.Println(user.Name)
	}


}
func TestLink(db *sql.DB) {//查看表中所有元素
	fmt.Println("start")

	//stmt,err:=db.Prepare("select * from hello ")
	rows,err:=db.Query("select * from hello")
	Check(err)
		for rows.Next() {
		var age string
		var name string               //新建两个临时变量用于保存数据
		err := rows.Scan(&age, &name) //用scan遍历数据库，保存数据
		Check(err)
		fmt.Println(age, name)
	}
}
func Insert(db *sql.DB,age int,name string){//插入操作
	fmt.Println("start Insert")
	stmt,err:=db.Prepare("insert into hello (age,name) values (?,?)")//操作语句写在prepare里(有几个?，exec里就有几个值)
    Check(err)
	result,err:=stmt.Exec(age,name)//实际执行过程(可用于插入删除和更改)
	Check(err)
	id,err:=result.LastInsertId()
	fmt.Println("自增Id",id)
	defer stmt.Close()

}
func Delete(db *sql.DB,age int){//删除操作
	fmt.Println("Start Delete")
	stmt,err:=db.Prepare("Delete from hello where age = ?")
	Check(err)
	result,err:=stmt.Exec(age)
	Check(err)
	num,err:=result.RowsAffected()
	fmt.Println(num)
	defer stmt.Close()
}
func Update(db *sql.DB,age int,name string){//更新
    fmt.Println("Begin Update")
    stmt,err:=db.Prepare("Update hello set age=? where name=?")
    Check(err)
    result,err:=stmt.Exec(age,name)
    Check(err)
    num,err:=result.RowsAffected()
    fmt.Println(num)
    defer stmt.Close()
}
func CloseDataBase(db *sql.DB){
	db.Close()
}