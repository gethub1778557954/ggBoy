package main

import "fmt"
import (
	"bytes"
	"strconv"
	"time"
)

func main() {
	var str1 bytes.Buffer

	var a string = "123:"
	var num int = 1

	var strr string = strconv.Itoa(num)
	str1.WriteString(a)
	str1.WriteString(strr)
	fmt.Println("result:" + str1.String())
	startTime := time.Now()
	fmt.Println(startTime)
	//os.Exit(-1)
}