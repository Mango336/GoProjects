package utils

import (
	"fmt"
	"os"
	"runtime"
)

func Testmain() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("GOPATH")
	const a int = 1
	fmt.Printf("GOPath is %s, %v \n", path, a)
	fmt.Println("GOPath is %s, %v", path, a)
	q, w, e := 1, 4, "e" // 没声明的话 可以用:=简化
	fmt.Println(q, w, e)
}
