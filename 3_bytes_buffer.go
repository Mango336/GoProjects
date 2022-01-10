package main

import (
	"bytes"
	"fmt"
	"os"
)

//type Buff struct {
//	Buffer *bytes.Buffer
//	Writer *bufio.Writer
//}
//
//// 初始化
//func NewBuff() *Buff {
//	b := bytes.NewBuffer([]byte{})
//	return &Buff{
//		Buffer: b,
//		Writer: bufio.NewWriter(b),
//	}
//}

func main() {
	/* bytes.buffer是一个缓冲byte类型的缓冲器 存放的都是byte */
	/* 1. 创建bytes.buffer */
	// 下面三者等价
	buf1 := bytes.NewBufferString("hello buf1")
	buf2 := bytes.NewBuffer([]byte("hello buf2"))
	buf3 := bytes.NewBuffer([]byte{byte('h'), byte('e'), byte('l'), byte('l'), byte('o'), byte(' '), byte('b'), byte('u'), byte('f'), byte('3')})
	fmt.Println(buf1, buf2, buf3)
	// 下面两者等价
	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println(buf4, buf5)
	/* 2. bytes.buffer的数据写入 */
	// 写入string
	buf6 := bytes.NewBuffer([]byte{})
	buf6.WriteString("buf6 writes string")
	fmt.Println(buf6.String())
	// 写入[]byte
	buf7 := bytes.NewBuffer([]byte{})
	buf7.Write([]byte("buf7 writes []byte"))
	fmt.Println(buf7.String())
	// 写入byte
	var b byte = '?'
	buf7.WriteByte(b)
	fmt.Println(buf7)
	// 写入rune
	var r rune = '!'
	buf7.WriteRune(r)
	fmt.Println(buf7.String())

	/* 3. 从文件写入buffer */
	readFromFile()

	/* 4. 数据写入到文件 */
	writeToFile()

}

func readFromFile() {
	file, err := os.Open("./3_buffer_test.txt")
	if err != nil {
		fmt.Println("read from file error: ", err)
	}
	defer file.Close()
	fmt.Println(file.Sync())
	buf := bytes.NewBufferString("===writeToFile()===: hello ")
	buf.ReadFrom(file) // 将文件中的内容追加到缓冲器的尾部
	fmt.Println(buf.String())
}

func writeToFile() {
	buf := bytes.NewBuffer([]byte("3_buffer_text.txt"))
	file, err := os.Open("3_buffer_text.txt")
	if err != nil {
		fmt.Println("write to file error: ", err)
	}
	defer file.Close()
	buf.WriteTo(file)  // buf写入到file
	fmt.Println("write to file success")
}
