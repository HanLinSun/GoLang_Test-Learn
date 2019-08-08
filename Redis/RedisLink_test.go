package Redis
import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"testing"
)
func Link() {
	c, err := redis.Dial("tcp", "124.156.206.94:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	//通过Do函数，发送redis命令
}
func Test(t *testing.T){
	Link()
}