package while

import (
	"testing"
	"fmt"
)

func TestWhileTest(t *testing.T) {
	jobWhileMoni()
}
//输出10次hello,world(使用类似while循环形式，先判断后做)
func jobWhileMoni() {
	var count = 0
	for {
		if count >= 10 {
			break //如果count>=10则退出
		}
		fmt.Println("hello,world", count)
		count++
	}
}

//模拟do……while实现输出10次hello,world（先做后判断）
func jobDowhileMoni(){
	var i = 0
	for{
		fmt.Println("hello,world",i)
		i++
		if(i>=10){
			break
		}
	}
}
