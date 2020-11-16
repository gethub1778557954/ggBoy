package main

import "fmt"
import "bytes"

func main(){
	var str1 bytes.Buffer

	var a string = "123:"
    var num int = 1

	str1.WriteString(a)
	
    
    fmt.Println("result:"+str1.String())
}