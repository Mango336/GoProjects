package utils

import (
	"fmt"
	"go_test/even"
)

func EvenMain() {
	for i := 0; i < 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
