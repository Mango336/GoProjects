package utils

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const LIM = 41

var fibs [LIM]uint64

func LearnDoc() { // main()最好写前面
	randExam()
	fmt.Println("===================JoinExam==============================")
	var s string = "hello"
	s = strings.Join([]string{s, " world"}, "") // 比 + 要高效  更高效的是使用字节缓冲拼接 bytes.Buffer
	fmt.Println(s)
	strExam()
	timeExam()
	pointerExam()
	deferExam("Go")
	fmt.Println("===================fibonacciExam==============================")
	var map_fib = make(map[int]int)
	n := 5
	fibonacci(n, map_fib)
	for i := 0; i <= n; i++ {
		fmt.Printf("fibonacci(%d) is: %d\n", i, map_fib[i])
	}
	for k, v := range map_fib {
		fmt.Println(k, v)
	}
	fmt.Println("===================参数为func==============================")
	addTwoNums(1, callBackAdd)
	fmt.Println(f())
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3)) // b+2=3+2=5
	p3 := Adder(1)
	fmt.Printf("a+b= %v\n", p3(3)) // a+b=1+3=4
	fmt.Println("===================闭包Exam==============================")
	fmt.Println("=闭包调试func=")
	where := func() { // 打印闭包函数位置
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	var f2 = Adder2()
	fmt.Print(f2(1), "-")
	fmt.Print(f2(20), "-")
	fmt.Println(f2(300))
	fmt.Println("=fibonacciExam闭包=")
	f := fibonacci2()
	for i := 1; i < 10; i++ {
		where()
		fmt.Printf("%d: %d\n", i, f())
	}
	fmt.Println("=闭包例子=") // 在创建一系列相似的函数时运用 书写一个工厂函数
	addJpeg := makeAddSuffix(".jpeg")
	addBmp := makeAddSuffix(".bmp")
	fmt.Println(addJpeg("file"), "\t", addBmp("file"))
	fmt.Println("==内存缓存提高性能==")
	start := time.Now()
	fibonacci(LIM, map_fib)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)

	var result uint64 = 0
	start = time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci3(i)
		fmt.Printf("%v ", result)
	}
	end = time.Now()
	delta = end.Sub(start)
	fmt.Println(delta)
}

func randExam() {
	fmt.Println("===================randExam==============================")
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		a := rand.Intn(8) //rand.Intn(n) 返回介于[0, n)之间的伪随机数
		fmt.Printf("%d / ", a)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond()) // ns 纳秒
	fmt.Println(timens)
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
	fmt.Println()
	ch := 'a'
	fmt.Println(unicode.IsLetter(ch))
	fmt.Println(unicode.IsDigit(ch))
}

func strExam() {
	fmt.Println("====================strExam==============================")
	str1 := "asdfjhskjhfgiguigvissieur fhsduifyhsuidfh"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2)) // str2中部分字符不是utf-8

	var str string = "This is an example of a string"
	//fmt.Printf("str has 'Th': %t", strings.HasPrefix(str, "Th"))
	//fmt.Printf("str has 'Th': %v", strings.HasSuffix(str, "Th"))
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th") //  \ 转义字符
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Println(strings.Contains(str, "n ")) // 是否包含substr
	fmt.Println(strings.Index(str, "a"))
	fmt.Println(strings.LastIndex(str, "a"))
	fmt.Println(strings.Index(str, "ak"))
	fmt.Println(strings.Fields(str), len(strings.Fields(str)))
	s := strings.Fields(str)
	fmt.Println(strings.Join(s, " "))
	fmt.Println(strconv.IntSize) // 程序运行的操作系统平台下int类型所占位数
}

func timeExam() {
	fmt.Println("===================timeExam==============================")
	t := time.Now()
	fmt.Println(t)                                               // 2021-12-01 14:53:30.8929966 +0800 CST m=+0.004244201
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year()) // 01.12.0001
	t = time.Now().UTC()
	fmt.Println(t) // 2021-12-01 06:56:58.6196574 +0000 UTC  UTC表示通用协调世界时间
	// 计算时间
	var week time.Duration
	week = 60 * 60 * 24 * 7 * 1e9 // 1e9 将秒转为纳秒级别 这里Duration(周期)为一周
	week_from_now := t.Add(week)  // 加一个Duration周期
	fmt.Println(week_from_now)    // 2021-12-08 07:06:15.0359459 +0000 UTC
	// formatting times:
	fmt.Println(t.Format(time.ANSIC))  // Wed Dec  1 07:42:10 2021
	fmt.Println(t.Format(time.RFC822)) // 01 Dec 21 07:42 UTC
	s := t.Format("20060102")
	fmt.Println(t, "=>", s) // 2021-12-01 07:42:55.555498 +0000 UTC => 20211201
}

func pointerExam() { // 指针
	fmt.Println("===================PointerExam==============================")
	var i1 int = 5
	fmt.Printf("An integer %d, its' location in memory: %p\n", i1, &i1) // 取地址符 &
	var intP *int
	intP = &i1 // intP存储了i1的内存地址 指向了i1的位置
	// intP指向的i1地址, intP指针自己的地址, intP所指地址上所存储的值（间接引用 or 反引用 or 内容）
	fmt.Printf("intP指向的i1地址: %v, intP指针自己的地址: %v, intP所指地址上所存储的值: %v\n", intP, &intP, *intP)
	// 指针对string的例子
	s := "good bye"
	var p *string = &s
	fmt.Println(&s, p, &s == p) // 0xc000094030 0xc000094030 true
	fmt.Println(*p)             // good bye
	*p = "hello"                // 修改*p => s也跟着变换
	fmt.Println(*p, s, *p == s) // hello hello true

}

func deferExam(s string) (n int, err error) { // defer语句记录函数的参数与返回值
	fmt.Println("===================deferExam==============================")
	defer func() {
		fmt.Printf("func1(%q) = %d, %v\n", s, n, err)
	}()
	return 7, io.EOF
}

func fibonacci(n int, mp map[int]int) int {
	var res int
	if n == 0 || n == 1 {
		res = 1
	} else {
		res = fibonacci(n-1, mp) + fibonacci(n-2, mp)
	}
	mp[n] = res
	return res
} // 递归 斐波那契数列

func callBackAdd(a, b int) {
	fmt.Println(a, "+", b, "=", a+b)
}

func addTwoNums(y int, callBack func(int, int)) {
	callBack(y, 2)
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func Adder2() func(int) int {
	var x int // 多次调用该函数 变量x的值被保留
	// 闭包函数保存并累积其中变量的值 不管外部函数退出与否 它都能够继续操作外部函数中的局部变量
	return func(delta int) int {
		x += delta
		return x
	}
}

func fibonacci2() func() int { // 闭包实现斐波那契数列
	a := 0
	b := 1
	return func() int {
		a, b = b, a
		b += a
		return b
	}
}

func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func fibonacci3(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci3(n-1) + fibonacci3(n-2)
	}
	fibs[n] = res
	return
}
