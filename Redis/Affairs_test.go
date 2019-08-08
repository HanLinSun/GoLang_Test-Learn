package Redis
import(
	"testing"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func TestA(t *testing.T){
   conn,err:=redis.Dial("tcp","124.156.106.94:6379")
   checkErr(err)
   defer conn.Close()
   conn.Send("MULTI")//开启事务
   conn.Send("INCR","foo")
   conn.Send("INCR","bar")
   r,err:=conn.Do("EXEC")
   fmt.Println(r)
}