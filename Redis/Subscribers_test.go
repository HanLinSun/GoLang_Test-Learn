package Redis
import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"testing"
	"time"
)
func checkerr(err error){
	if err!=nil{
		fmt.Println("error is:" ,err)

	}
}
func Subscriber(){
	conn,err:=redis.Dial("tcp","124.156.206.94:6379")
	checkerr(err)
	defer conn.Close()
	psc:=redis.PubSubConn{conn}
	psc.Subscribe("channel01")//订阅channel01频道
	for{
		switch v:=psc.Receive().(type){
		case redis.Message:
			fmt.Printf("%s:message:%s\n",v.Channel,v.Data)
		case redis.Subscription:
			fmt.Printf("%s:%s %d\n",v.Channel,v.Kind,v.Count)
		case error:
			fmt.Println(v)
		return
		}
	}
}
func Pusher(message string){
	conn,_:=redis.Dial("tcp","124.156.206.94:6379")
	_,err:=conn.Do("PUBLISH","channel01",message)
	checkerr(err)
}
func Test03(t *testing.T){
	go Subscriber()
	go Pusher("Channel test")
	time.Sleep(time.Second*3)
}