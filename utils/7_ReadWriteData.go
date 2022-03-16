package utils

import (
	"fmt"
	"go_test/pack7"
)

func ReadWriteData() {
	fmt.Println("====读取用户输入====")
	// 使用scan_读取console数据
	pack7.ReadInput1()
	// 使用bufio包提供的buffered reader来读取数据
	pack7.ReadInput2()
	// 使用switch
	pack7.SwitchInput()

	fmt.Println("====读取文件数据====")
	pack7.ReadFileData()

	fmt.Println("====读取文件数据到一个字符串====")
	pack7.ReadFileToAString()

	fmt.Println("====写入文件数据====")
	pack7.WriteFileData()

	fmt.Println("====复制文件====")
	pack7.CopyFile("data/7_target.txt", "data/7_source.txt")

	//练习测试: 复制每行第3-5字符到新文件中
	pack7.Remove_3till5Chars()

	//Json数据格式解析
	fmt.Println("JSON 解析: ")
	pack7.JsonDataFormat()
	//XML数据格式解析
	fmt.Println("XML 解析: ")
	pack7.XMLDataFormat()
	//Gob数据格式解析
	pack7.Gob1()  //gob encoder decoder
	pack7.Gob2()  //将数据存为gob文件
	pack7.Degob() //解码gob文件并打印
	pack7.GobHash_sha1()

}
