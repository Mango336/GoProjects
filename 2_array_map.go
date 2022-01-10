/*
数组与map
*/
package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("==========数组==========")
	// 数组长度也是数组类型的一部分 [5]int 和 [10]int属于不同类型
	var arr1 [5]int = [5]int{0}
	fmt.Println("==arr1: ")
	for i, j := range arr1 {
		fmt.Println(i, j)
	}

	str_arr := [...]string{"a", "b", "c", "d"}
	fmt.Println("==str_arr: ")
	for i := range str_arr {
		fmt.Println("Array item", i, "is", str_arr[i])
	}

	var arr2 = new([5]int) // 类型为*[5]int
	fmt.Println("==arr2: ")
	fmt.Println(arr2)  // 地址
	fmt.Println(*arr2) // 指针所指地址的内容

	arr3 := *arr2
	arr3[1] = 100
	fmt.Println(arr3[1], arr2[1])
	fmt.Println("==arr3: ")
	fmt.Println(arr3)
	f1(arr3)
	fmt.Println(arr3)
	f2(&arr3) // 引用数组
	fmt.Println(arr3)

	fmt.Println("==========slice切片==========")

	fmt.Println(cap(arr3))
	fmt.Println(arr3[:3])
	arr4 := arr3[:3] // 切片是连续片段的引用 指向原数组地址
	arr4[0] = 1
	fmt.Println(arr3, arr4)

	var arr5 [6]int
	var slice1 []int = arr5[2:5]
	// arr5 赋值
	for i := range arr5 {
		arr5[i] = i
	}
	// 切片
	for i, j := range slice1 {
		fmt.Printf("slice1[%d]=%d\t", i, j)
	}
	fmt.Println()
	fmt.Println(len(arr5), len(slice1), cap(slice1))

	slice1 = slice1[0:4]
	for i, j := range slice1 {
		fmt.Printf("slice1_[%d]=%d\t", i, j)
	}
	fmt.Println()
	fmt.Println(len(slice1), cap(slice1))

	str_arr2 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(str_arr2[1:4], str_arr2[:2], str_arr2[2:], str_arr2[:]) // 输出为ascii码

	fmt.Println("====切片传递给函数====")
	fmt.Println(arrSum(arr5[:])) // 传入数组arr
	fmt.Println("====make()创建一个切片====")
	arr6 := make([]int, 10)
	for i := range arr6 {
		arr6[i] = 5 * i
	}
	fmt.Println(arr6, cap(arr6), len(arr6))

	fmt.Println("====bytes包====")
	// bytes包中Buffer类型提供read和write方法 可以处理长度未知的bytes
	var buffer bytes.Buffer
	buffer.WriteString("buffer string")
	fmt.Println(buffer.String()) // 转为string
	buffer.WriteString(" add")   // 追加字符串  比使用+=更省内存和CPU
	fmt.Println(buffer.String())

	var buffer2 []byte = []byte("buffer ")
	var data []byte = []byte("data")
	fmt.Println(string(byteArrayAppend(buffer2, data)))

	buf1 := bytes.NewBuffer([]byte("buf1"))
	//buf1 := bytes.NewBufferString("buf1")
	//buf1 := bytes.NewBuffer([]byte{byte('b'), byte('u'), byte('f'), byte('1')})  // 这两种写法与上面等价
	fmt.Println(buf1.String())
	var sRead = make([]byte, 2)
	buf1.Read(sRead)
	fmt.Println(buf1.String()) // 前两个byte被读出
	fmt.Println(sRead)

	fmt.Println("====for-range结构====")
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, s := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, s)
	}
	// 多维切片下的for-range
	screen := [][]int{{1, 2, 3}, {4, 5, 6}}
	for row := range screen {
		for _, cs := range screen[row] {
			fmt.Printf("%d\t", cs)
		}
		fmt.Println()
	}

	fmt.Println(sum_array([]float32{1.2, 2.3, 3.4}))
	fmt.Println(SumAndAverage([]float32{1.2, 2.3, 3.4}))
	fmt.Println(min_max([]int{3, 4, 2, 1, 0, 3, 9}))

	fmt.Println("====切片reslice====")
	var arr7 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = arr7[5:7]
	fmt.Println(a, len(a), cap(a)) // [5 6] 2 5
	a = a[0:4]
	fmt.Println(a, len(a), cap(a)) // [5 6 7 8] 4 5

	fmt.Println("====切片copy append====")
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)                    // [1 2 3 0 0 0 0 0 0 0]
	fmt.Printf("Copied %d elements\n", n) // n==3
	slice2 := []int{4, 5, 6}
	slice2 = append(slice2, 7, 8, 9)
	fmt.Println(slice2)
	fmt.Println(AppendByte([]byte{1, 2, 3}, 4, 5, 6))

	fmt.Println("====字符串 数组 切片的应用====")
	str1 := "hello"
	fmt.Println(str1, str1[2:3])
	// string是不可变的 即str[index]不能放在等号左侧 否则报错
	//str1[0] = 'c'  // 报错
	// 修改string中某字符方法: string转为[]byte 然后修改 再转为string
	c := []byte(str1)
	c[0] = 'c'
	str2 := string(c)
	fmt.Println(str1, str2)
	arr8 := []int{1, 2, 1, 34, 5, 3, 687, 3, 0, -2}
	fmt.Println(sort.IntsAreSorted(arr8)) // 查看arr是否已经排序
	sort.Ints(arr8)
	fmt.Println(sort.IntsAreSorted(arr8), arr8)
	fmt.Println(sort.SearchInts(arr8, 5)) // 二分查找
	arr8 = append(arr8[:7], arr8[8:]...)  // 删除位于索引i的元素 这里i==7
	fmt.Println(arr8)

	fmt.Println("====map====")
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32) // map是引用类型的：内存用make方法来分配
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	mp1 := make(map[int][]int)
	mp1[0] = []int{1, 2, 3, 4, 5, 6}
	mp1[1] = []int{2, 3, 4}
	fmt.Println(mp1)
	val1, isPresent := mp1[0]
	fmt.Println(val1, isPresent) // [1 2 3 4 5 6] true
	val1, isPresent = mp1[2]     // map中key不存在 val1==空值 并且isPresent==false
	fmt.Println(val1, isPresent) // [] false
	//和 if 混合使用：
	//if _, ok := map1[key1]; ok {
	//	// ...
	//}
	delete(mp1, 1) // map中删除某个key
	fmt.Println(mp1)

	// map默认无序
	// 若需排序 则将key或value拷贝到一个切片 再对切片排序(使用sort包)
	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
	fmt.Println("unsorted: ", barVal)
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		//keys = append(keys, k)
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println(keys, len(keys))
	fmt.Println("sorted: ")
	for i := range keys {
		fmt.Printf("barVal[%s]: %d ", keys[i], barVal[keys[i]])
	}


}

func f1(a [5]int) {
	fmt.Println("===f1===")
	for i, j := range a {
		fmt.Println(i, j)
		a[i] = i * 2
	}
	fmt.Println(a)
}

func f2(a *[5]int) {
	fmt.Println("===f2 数组引用===")
	for i, j := range *a {
		fmt.Println(i, j)
		(*a)[i] = i * 2
	}
}

func arrSum(a []int) (s int) {
	s = 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return
}

func byteArrayAppend(slice, data []byte) []byte {
	for i := range data {
		slice = append(slice, data[i])
	}
	return slice
}

func sum_array(arrF []float32) (s float32) {
	for _, a := range arrF {
		s += a
	}
	return
}

func SumAndAverage(arrF []float32) (int, float32) {
	var sum float32 = 0
	for _, a := range arrF {
		sum += a
	}
	return int(sum), sum / float32(len(arrF))
}

func min_max(arr []int) (int, int) {
	var min int = math.MaxInt
	var max int = math.MinInt
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// append的具体实现方法
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // 如果需要 重分配内存
		newSlice := make([]byte, (n+1)*2) // 分配两倍的内存空间
		copy(newSlice, slice)
		slice = newSlice // 还想继续使用slice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
