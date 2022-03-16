/* 协程
一个并发程序可以在一个处理器或者内核上使用多个线程来执行任务，但是只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行
并行是一种通过使用多处理器以提高速度的能力。所以并发程序可以是并行的，也可以不是。
公认的，使用多线程的应用难以做到准确，最主要的问题是内存中的数据共享，它们会被多线程以无法预知的方式进行操作，导致一些无法重现或者随机的结果(称作竞态)

** 不要使用全局变量或者共享内存 在并发运算时带来危险 **
 =>解决方法：同步不同的线程 对数据加锁; 但会带来更高的复杂度，更容易使代码出错以及更低的性能，所以这个经典的方法明显不再适合现代多核/多处理器编程:thread-per-connection模型不够有效
 在Go中，应用程序并发处理的部分被称作goroutines（协程），它可以进行更有效地并发运算
	在协程和操作系统线程之间并无一对一的关系：协程是根据一个或多个线程的可用性，映射（多路复用，执行于）在他们之上的；
	协程调度器在Go运行时很好地完成了这个工作。
 协程工作在相同的地址空间中，所以共享内存的方式一定是同步的 (使用channels来同步协程)

协程是轻量的，比线程更轻
它们痕迹非常不明显（使用少量的内存和资源）：使用4K的栈内存就可以在堆中创建它们。
因为创建非常廉价，必要的时候可以轻松创建并运行大量的协程（在同一个地址空间中 100,000 个连续的协程）。
并且它们对栈进行了分割，从而动态的增加（或缩减）内存的使用；栈的管理是自动的，但不是由垃圾回收器管理的，而是在协程退出后自动释放。
*/
package utils

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

//GoroutineChannel main函数
func GoroutineChannel() {
	//goroutineForTest() //多核并行运行协程
	//goroutineUseChannel() //协程使用通道
	//channelBlock() //通道没有接收者 阻塞
	//blocking() //死锁 没有发送者协程 在通道两端互阻塞对方
	//channelIdiom()  //通道工厂模式
	//channelIdiom2() //通道迭代模式
	//channelDirect()
	//printPrime(20) //打印前n个素数 version1
	//printPrime2(20) //version2
	//goroutineSelect() //select切换协程
	timerGoroutine()
}

var numCores = flag.Int("n", 2, "number of cpu cores to use")

//goroutineForTest 使用通道的简单模拟例子
func goroutineForTest() {
	preTime := int64(time.Now().Nanosecond())
	fmt.Println("===Goroutine===")
	flag.Parse()
	runtime.GOMAXPROCS(*numCores) //使用2个cpu并行运行程序

	fmt.Println("In GoroutineMain()")
	//4,460,900ns
	go longWait() //协程通过使用关键字go 调用(或执行)一个函数或方法来实现的
	go shortWait()

	//24,085,000ns 未用协程 使用一个线程连续调用的情况(移除了go 关键字)
	//longWait()
	//shortWait()

	fmt.Println("About to sleep in Goroutine()")
	time.Sleep(10 * 1e9) // sleep works with a Duration in nanoseconds (ns) !
	fmt.Println("At the end of Goroutine()")
	curTime := int64(time.Now().Nanosecond())
	fmt.Println(curTime - preTime) //计算运行时间
}

//longWait use on GoroutineForTest
func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

//shortWait use on GoroutineForTest
func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

//goroutineUseChannel 协程使用通道传输数据
func goroutineUseChannel() {
	ch := make(chan string) //initial channel
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9) //sleep 1s; 如果没有该语句，主程序执行完了 但协程调用的函数方法没有执行完毕 该协程会被释放掉
	fmt.Println()
}

//sendData 发送数据
func sendData(ch chan string) {
	//<- 通信操作符 流向通道(发送)
	ch <- "123"
	ch <- "234"
	ch <- "345"
	ch <- "456"
	ch <- "567"
	ch <- "678"
}

//getData 接收数据
func getData(ch chan string) {
	var input string
	for {
		input = <-ch //从通道流出(接收)
		fmt.Printf("%s ", input)
	}
}

//channelBlock 通道阻塞
func channelBlock() {
	ch := make(chan int)
	go pump4ChannelBlock(ch) //pump hangs
	fmt.Println(<-ch)        //输出0; 因为没有接收者协程接收通道中的数据 所以通道阻塞 (新的输入无法在通道非空的情况下传入)
}
func pump4ChannelBlock(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// blocking  死锁 没有发送者协程 在通道两端互阻塞对方
func blocking() {
	out := make(chan int)
	out <- 2
	go f1Blocking(out)
}
func f1Blocking(in chan int) {
	fmt.Println(<-in)
}

//channelIdiom 习惯用法 通道工厂模式=>
//不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）；函数内有个匿名函数被协程调用。
func channelIdiom() {
	stream := pump4channelIdiom() //return一个channel
	go suck4ChannelIdiom(stream)
	time.Sleep(1e9)
}

//pump4channelIdiom 用于channelIdiom案例中 send value
func pump4channelIdiom() chan int {
	ch := make(chan int)
	go func() { //lambda作为协程
		for i := 0; ; i++ {
			ch <- i //给channel send value操作
		}
	}()
	return ch
}

//suck4ChannelIdiom 用于channelIdiom案例中 receive value
func suck4ChannelIdiom(ch chan int) {
	for {
		fmt.Println(<-ch) //从指定channel中读取数据 直至channel关闭
	}
}

//channelIdiom2 习惯用法 通道迭代模式=>
func channelIdiom2() {
	suck4ChannelIdiom2(pump4channelIdiom())
	time.Sleep(1e9)
}
func suck4ChannelIdiom2(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v) //for-range用于通道
		}
	}()
}

/* channelDirect 通道方向
var send_only chan<- int //通道只能接收数据 协程发送数据
var recv_only <-chan int //通道只能发送数据 协程接收数据
*/
func channelDirect() {
	var c = make(chan int)
	go source4ChannelDirect(c)
	go sink4ChannelDirect(c)
	time.Sleep(1e6)
}
func source4ChannelDirect(ch chan<- int) {
	for {
		ch <- 1
	}
}
func sink4ChannelDirect(ch <-chan int) {
	for {
		fmt.Println(<-ch)
	}
}

//printPrime 打印素数 version1
func printPrime(n int) {
	ch := make(chan int)
	go generate4PrintPrime(ch)
	for i := 0; i < n; i++ {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter4PrintPrime(ch, ch1, prime)
		ch = ch1 //包含已被筛选掉一轮的数的输出通道 作为下一轮的输入 每个prime又开启了一个新的协程 生成器选择器并发请求
	}
}

//generate4PrintPrime 生成器generate
func generate4PrintPrime(ch chan int) {
	for i := 2; ; i++ {
		ch <- i //send i to channel
	}
}

//filter4PrintPrime 选择器filter
func filter4PrintPrime(in, out chan int, prime int) {
	for {
		i := <-in // receive new variable i from channel in
		if i%prime != 0 {
			out <- i //拷贝不能被prime整除的数字到输出通道
		}
	}
}

//printPrime2 打印素数 version2
func printPrime2(n int) {
	primes := sieve4PrintPrime2()
	for i := 0; i < n; i++ {
		fmt.Print(<-primes, " ")
	}
}
func generate4PrintPrime2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func filter4PrintPrime2(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 { //为啥是i%prime？
				out <- i
			}
		}
	}()
	return out
}
func sieve4PrintPrime2() chan int { //sieve 返回了素数通道
	out := make(chan int)
	go func() {
		ch := generate4PrintPrime2()
		for {
			prime := <-ch //？为啥这边出来就是prime了
			ch = filter4PrintPrime2(ch, prime)
			out <- prime
		}
	}()
	return out
}

//closeChannel close(chan) 关闭通道
func closeChannel() {
	ch := make(chan string)
	//sendData是协程 getData和main(closeChannel)在同一个线程里
	go sendData4CloseChannel(ch)
	getData4CloseChannel(ch)
}
func sendData4CloseChannel(ch chan string) {
	for i := 97; i < 108; i++ {
		ch <- string(byte(i))
	}
	close(ch) //关闭通道
}
func getData4CloseChannel(ch chan string) {
	for {
		input, open := <-ch
		if !open { //通道关闭
			break
		}
		fmt.Printf("%s ", input)
	}
	// version2 for-range读取通道能自动检测通道是否关闭
	//for input := range ch {
	//	process(input) //处理input
	//}
}

//goroutineSelect select 切换协程进行操作
func goroutineSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump14Select(ch1)
	go pump24Select(ch2)
	go suck4Select(ch1, ch2)
	time.Sleep(1e9)
}
func pump14Select(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}
func pump24Select(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}
func suck4Select(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

//timerGoroutine 通道、超时和计时器(Ticker)
func timerGoroutine() {
	tick := time.Tick(1e8)  //以d为周期给返回的通道发送时间, d是纳秒数;(使用情况：想返回一个通道而不必关闭它的时候)
	boom := time.After(5e8) //After()只发送一次时间
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
