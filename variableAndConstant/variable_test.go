package ist_test

import (
	"testing"
	"fmt"
)

func TestFib(t *testing.T)  {
	a := 1
	b := 1
	fmt.Println("  ",a)
	for i := 0;i<=10;i++{
		fmt.Println("  ",b)
		temp := a
		a = b
		b = temp + a
	}
}

func TestLeft(t *testing.T)  {
	var a int = 1
	var b int = 2

	t.Log("将a的值赋值给b  将b的值赋值给a")
	t.Log("简单写法，同一赋值语句 给多个变量赋值")
	a,b = b,a
	fmt.Println("a:",a,"b:",b)
}