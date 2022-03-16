package utils

import (
	"fmt"
	"runtime"
	"sync"
)

//GoroutineChannel main函数
func GoroutineChannelBook() {
	//goroutine
	//单核运行 两个协程 打印大小写字母表
	//listing01()
	//listing04()

	//竞争状态 => go build -race 竞争检测器 检测哪个goroutine引发数据竞争
	listing09()
	//lock共享资源 => atomic, sync包

}

//单核运行 两个协程 打印大小写字母表
/*展示如何创建goroutine 以及调度器的行为
第一个goroutine完成所有打印需要花费的时间太短了
以至于调度器切换到第二个goroutine之前, 就完成了所有任务
*/
func listing01() {
	runtime.GOMAXPROCS(1) //分配一个逻辑处理器CPU给调度器使用
	//分配两个CPU(逻辑处理器) 多核并行 每个协程在独自运行在自己的核上
	//只有在有多个逻辑处理器且可以同时让每个goroutine运行在一个可用的物理处理器上的时候 goroutine才会并行运行
	runtime.GOMAXPROCS(2)
	//wg用来等待程序完成
	var wg sync.WaitGroup //计数信号量(用来记录并维护运行的goroutine)
	wg.Add(2)             //add(2) 表示要等待两个goroutine

	fmt.Println("Start Goroutine...")

	go func() { //匿名函数 并创建一个协程
		defer wg.Done() //在函数退出时 调用Done来通知main函数 工作已经完成
		//print 小写字母表三次
		for count := 0; count < 5; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() { //匿名函数 并创建一个协程
		defer wg.Done() //在函数退出时 调用Done来通知main函数 工作已经完成
		//print 大写字母表三次
		for count := 0; count < 5; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	//等待goroutine
	fmt.Println("Waiting To Finish...")
	wg.Wait() //阻塞 知道所有goroutine都wg.Done()
	fmt.Println("\nTerminating Program")
}

//展示goroutine调度器是如何再单个线程上切分时间片
var wgListing04 sync.WaitGroup

//单核运行 两个协程 打印0~5000以内的所有素数
/*
下面程序中创建了两个goroutine 分别打印1~5000内的素数
查找并打印会消耗不少时间 这使得调度器有机会在第一个goroutine找到所有素数之前 切换该goroutine的时间片
*/
func listing04() {
	runtime.GOMAXPROCS(1)
	wgListing04.Add(2)

	fmt.Println("Create Goroutine...")
	go printPrime3("A")
	go printPrime3("B")

	fmt.Println("Waiting To Finish")
	wgListing04.Wait()

	fmt.Println("Terminating Program")
}

//printPrime3 打印5000以内的素数值
func printPrime3(prefix string) {
	defer wgListing04.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

var (
	counter     int            //counter是所有goroutine都要增加其值的变量
	wgListing09 sync.WaitGroup //wg等待程序结束
)

//listing09 展示如何在程序里造成竞争状态
func listing09() {
	wgListing09.Add(2)
	go inCounter(1)
	go inCounter(2)
	wgListing09.Wait()
	fmt.Println("Final Counter: ", counter)
	//输出应该为 Final Counter: 4 因为进行了4次读写操作 但结果为2
	//每个goroutine创造一个counter变量副本 之后就切换到另一个goroutine,
	//而当这个goroutine再次运行时 counter的值已经改变 但goroutine并没有更新自己的那个副本的值 而是继续使用这个副本的值 用该值递增 并存回counter
	//结果覆盖了另一个goroutine完成工作
}

//inCounter 增加包里counter变量的值
func inCounter(id int) {
	defer wgListing09.Done()
	for count := 0; count < 2; count++ {
		value := counter  //获取counter
		runtime.Gosched() //当前goroutine从线程退出 并放回到队列
		value++           //增加本地value值
		counter = value   //将value放回counter
	}
}
