package main

// import . "go_test/pack1"  // 不用通过包名来使用其中的方法 直接使用方法名即可
// import _ "go_test/pack1"  // 只执行pack1包的init函数并初始化其中的全局变量  为了具有更好的测试效果
// 导入外部的安装包: 使用go install在本地机器上安装它们 如：go install codesite.ext/author/goExample/goex
// 使用时：import goex "codesite.ext/author/goExample/goex" goex是这个包起的别名

import (
	"fmt"
	"go_test/pack1"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"sync"
)

type Info struct {
	mu sync.Mutex
	// ...其他field 如：str string
	str string
}

func main() {
	fmt.Println("====regexp包====")
	regexpPackage()
	fmt.Println("====精密计算和big包====")
	bigPackage()
	fmt.Println("====自定义包====")
	fmt.Println(pack1.ReturnStr()) // import pack1包 使用其Public方法
	fmt.Println(pack1.SP1Int, pack1.SP1Float)
	pack1.PrintStr()
}

func regexpPackage() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" // 目标字符串
	pat := "[0-9]+.[0-9]+"                                      // 正则
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {       // 对字符串进行正则表达式匹配 match成功return true (return第二项是error)
		fmt.Println("Match Found!")
	}
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#") // 将匹配到的部分替换为 '##.#'
	fmt.Println(str)
	// 参数为函数时
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)           // str convert to float32  bitsize=32:float32 64:float64
		return strconv.FormatFloat(v*2, 'f', 2, 32) // 操作：v*2  prec精度
	}
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

// 锁和sync包
func lockAndSyncPackage(info *Info) {
	// sync.Mutex 互斥锁 作用：守护在临界区入口来确保同一时间只能有一个线程进入临界区
	// 假设Info是需要上锁的放在共享内存中的变量
	info.mu.Lock()
	info.str = "new Value"
	info.mu.Unlock()
}

func bigPackage() {
	// 高精度计算
	// bigInts的一些计算:
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io) // ip = (im*in + im)/io
	fmt.Printf("Big Int: %v\n", ip)

	// bigRats的一些计算：(大有理数)  big.NewRat(N, D); N:分子 D:分母 =>都是int64型整数
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	//fmt.Println(rm, rn, ro, rp, rq)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp) // rq = (rm*rn + ro) * rp
	fmt.Printf("Big Rat: %v\n", rq)
}
