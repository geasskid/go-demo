package main //程序包名

//两种引入包的方式
import (
	"fmt"
	"time"
)

/**
import "fmt"
import "time"
*/

// main函数
func main() { //方法名后面必须接“{”，否则编译错误
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World")

}
