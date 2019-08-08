package Redis
import(
	"testing"
	"github.com/garyburd/redigo/redis"
	"fmt"
)


func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}

func Test01(t *testing.T){
	conn,err:=redis.Dial("tcp","124.156.206.94:6379",redis.DialDatabase(0))
	checkErr(err)
	defer conn.Close()
	//插入
	_,err=conn.Do("SET","name","nick")
	checkErr(err)
	//获取
	r,err:=redis.String(conn.Do("GET","name"))//写在redis_string里才能正常输出
	checkErr(err)
	fmt.Println(r)
	_,err=conn.Do("DEL","name")
	checkErr(err)

	_,err=conn.Do("MSet","name","nick","age","21")
	checkErr(err)
	r2,err:=redis.Strings(conn.Do("MGet","name","age"))
	checkErr(err)
	fmt.Println(r2)
	//Hash
	_,err=conn.Do("HSet","names","Galatians","Searching...")
	checkErr(err)
	r,err=redis.String(conn.Do("HGet","names","Galatians"))
	checkErr(err)
	fmt.Println(r)
	//队列
	_,err=conn.Do("lpush","Queue","nick","Dawn",9)
	checkErr(err)
	for{
		r,err=redis.String(conn.Do("lpop","Queue"))
		if err!=nil{
			fmt.Println(err)
			break
		}
		fmt.Println(r)
	}
	//r3,err:=redis.Int(conn.Do("llen","Queue")
}
//连接池
var pool *redis.Pool//声明一个全局类型的连接池对象

func Init(){
	    pool = &redis.Pool{
		MaxIdle:16,
		MaxActive:1024,
		IdleTimeout:300,
		Dial:func()(redis.Conn,error){
			return redis.Dial("tcp","124.156.206.94:6379")
		},//redis connect返回类型是redis.conn,连接方法只能是函数返回类型
	}
}
//获得连接:conn:=Pool.Get()