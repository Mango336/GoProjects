package utils

import (
	"errors"
	"fmt"
	"go_test/pack8"
	"log"
	"math"
)

func ErrorsAndTest() {
	var errNotFound error = errors.New("not found error")
	fmt.Printf("error: %v\n", errNotFound)

	//Sqrt()参数测试
	if sf, err := Sqrt(-16); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sf)
	}

	/* panic 运行异常时直接从代码初始化 */
	//当错误条件（我们所测试的代码）很严苛且不可恢复，程序不能继续运行时，可以使用 panic 函数产生一个中止程序的运行时错误
	//panic 接收一个做任意类型的参数，通常是字符串，在程序死亡时被打印出来
	//Go 运行时负责中止程序并给出调试信息
	//fmt.Println("Starting the program")
	//panic("A severe error occurred: stopping the program!")
	//fmt.Println("Ending the program")

	/* panicking: */
	//在多层嵌套的函数调用中调用panic，可以马上中止当前函数的执行，所有的defer语句都会保证执行并把控制权交还给接收到panic的函数调用者
	//这样向上冒泡直到最顶层，并执行(每层的)defer，在栈顶处程序崩溃，并在命令行中用传给panic的值报告错误情况：这个终止过程就是 panicking。
	//不能随意地用 panic 中止程序，必须尽力补救错误让程序能继续执行。

	/*	从panic中恢复(Recover)
		总结: panic会导致栈被展开直到defer修饰的recover()被调用或者程序中止 */
	fmt.Println("Calling panic_recover")
	panicRecover()
	fmt.Println("panic_recover completed")

	protect(func() {
		fmt.Println("g")
		panic("g panic")
	})

	complexPanicRecover()
	fmt.Println("====panic+defer在闭包中处理错误的模式====")
	panicDefer()
}

// Sqrt 计算平方根函数的参数测试
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("math-square root of negative number")
		// fmt创建错误对象 fmt.Errorf() 可以接收有一个或多个格式占位符的格式化字符串和相应数量的占位变量
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return math.Sqrt(f), nil
}

func badCall() {
	panic("bad end")
}
func panicRecover() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n")
}

//protect函数调用函数参数g来保护调用者防止从g中抛出的运行时panic，并展示panic中的信息
func protect(g func()) {
	defer func() { //defer-panic-recover使用
		log.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g() //   possible runtime-error
}

//panic_recover
func complexPanicRecover() {
	var examples = []string{"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}
	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n", ex)
		num, err := pack8.Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(num)
	}
}

//panic+defer在闭包中处理错误的模式
func panicDefer() {
	fc()
	fmt.Println("Returned normally from f.")
	/*  运行结果
	Calling g.
	Printing in g 0
	Printing in g 1
	Printing in g 2
	Printing in g 3
	Panicking
		//panic=>i==4...
	Defer in g 3
	Defer in g 2
	Defer in g 1
	Defer in g 0
	Returned normally from g.   xxx
	Recovered in f 4
	Returned normally from f.
	*/
}
func fc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g.", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
