package dataType

import (
	"testing"
	"fmt"
)
type MyInt int

func TestDatatype(t *testing.T) {
	// bool 枚举
	// string 字符串
	// int  int8  int16  int32  int64  多种环境下的 int 类型
	// uint uint8  uint16  uint32  int64  uintptr  无符号 int 型
	// byte alias for uint8 uint8 的字节别名
	// rune alias for int32 ,represents a Unicode code point  int32的别名，表示Unicode代码 中文 日文 适合这个类型
	// float32 float64  浮点数
	// complex64 complex128  复数

	//a := cmplx.Pow(10,5)
	//fmt.Println(a)

	var c uint = 8
	fmt.Println("c: ", c)

	// MyInt 是在上面给 int 设置的别名
	var d MyInt = 8
	fmt.Println("d: ", d)

	//fmt.Println(c == d)  // 报错 类型不匹配
	var e int = 8
	fmt.Println("e: ", e)

	//fmt.Println(e==d)  // 报错  类型不匹配  别名类型 和 原类型 也不能进行比较

	var f int = 8
	fmt.Println("f: ",f)

	fmt.Println("e == f ? ",e == f)

	// 打印 各个值的类型
	//t.Log("%T %T %T",c,d,e)
	fmt.Printf("%T %T %T",c,d,e)
}