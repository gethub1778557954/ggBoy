package operator

import (
	"testing"
)

func TestOperator(t *testing.T) {
	//运算符  +  -  *  /  %  --  ++
	// 不支持前自增  ++a  这种情况

	//比较运算符 ==  !=  >  <  >=  <=  * 数组比较  切片比较 *

	//逻辑运算符  &&  ||  ！

	//位运算符 &  |  ！  <<  >>
	// 按位清零 &^ 

	//数组构建
	a := [...] int{1, 2, 3, 4, 5}
	b := [...] int{1, 2, 3, 4, 5}
	c := [...] int{1, 2, 3, 4, 5, 6}
	//for key,value := range a{
	//	fmt.Println(key," : ",value)
	//}

	//strArr := [...] string {"1","2","3","4","5","6","7","8","9","10"}


}
