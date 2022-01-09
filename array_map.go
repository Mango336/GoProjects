/*
数组与map
*/
package main

import "fmt"

func main() {
	fmt.Println("==========数组==========")
	// 数组长度也是数组类型的一部分 [5]int 和 [10]int属于不同类型
	var arr1 [5]int = [5]int{0}
	for i, j := range arr1 {
		fmt.Println(i, j)
	}

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	var arr2 = new([5]int)  // 类型为*[5]int
	fmt.Println(arr2)
	fmt.Println(*arr2)

	arr3 := *arr2
	arr3[1] = 100
	fmt.Println(arr3[1], arr2[1])

	fmt.Println(arr3)
	f1(arr3)
	fmt.Println(arr3)
	f2(&arr3)  // 引用数组
	fmt.Println(arr3)

	fmt.Println("==========slice切片==========")

	fmt.Println(cap(arr3))
	fmt.Println(arr3[:3])
	arr4 := arr3[:3]  // 切片是连续片段的引用 指向原数组地址
	arr4[0] = 1
	fmt.Println(arr3, arr4)

	var arr6 [6]int
	var slice1 []int = arr6[2:5]
	for i := range arr6 {
		arr6[i] = i
	}
	for i, j:=range slice1 {
		fmt.Printf("slice1[%d]=%d\t", i, j)
	}
	fmt.Println()
	fmt.Println(len(arr6), len(slice1), cap(slice1))

	slice1 = slice1[0:4]
	for i, j:=range slice1 {
		fmt.Printf("slice1_[%d]=%d\t", i, j)
	}
	fmt.Println()
	fmt.Println(len(slice1), cap(slice1))

	b := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Println(b[1:4], b[:2], b[2:], b[:])  // ascii码

}

func f1(a [5]int)  {
	fmt.Println("===f1===")
	for i,j:=range a{
		fmt.Println(i,j)
		a[i] = i*2
	}
	fmt.Println(a)
}

func f2(a *[5]int)  {
	fmt.Println("===f2===")
	for i,j:=range *a{
		fmt.Println(i,j)
		(*a)[i] = i*2
	}
}
