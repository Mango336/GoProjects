package pack7

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadFileData() {
	// inputFile 是*os.File类型(文件句柄) 下面代码：只读模式打开input.dat文件
	inputFile, inputError := os.Open("7_input_text.txt")
	if inputError != nil {
		fmt.Printf("读取文件错误\n" +
			"文件不存在?\n" +
			"无权限访问?\n")
		return // exit; os.Open the file is error.
	}
	defer inputFile.Close()                   // return后close file
	inputReader := bufio.NewReader(inputFile) // 读取器
	for {                                     // 在死循环中使用ReadString('\n')或ReadBytes('\n')将文件的内容逐行(行结束符'\n')读取出来
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		} // 读取到了文件末尾 return
	}
}

func ReadFileToAString() {
	fmt.Println("==1. 将整个文件的内容读到一个字符串里:==")
	inputFile := "7_input_text.txt"
	outputFile := "7_output_text.txt"
	buf, err := ioutil.ReadFile(inputFile) // return的是[]byte 里面存放读取到的内容和error(为nil就没错)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	//将读取的内容buf([]byte) 写入到文件outputFile
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("==2. 带缓存的读取 针对文件内容不按行划分或是二进制文件==")
	bufByte := make([]byte, 1024)
	input, inputErr := os.Open(inputFile)
	if inputErr != nil {
		fmt.Println("Open File Error!")
		return
	}
	defer input.Close()

	inputerReader := bufio.NewReader(input)
	n, err := inputerReader.Read(bufByte)
	if n == 0 {
		fmt.Println("Have no text!")
		return
	}
	fmt.Println(n) //读取到的字节数
	fmt.Println(string(bufByte[0:n]))

	fmt.Println("==3. 按列读取文件中的数据==")
	//数据按列排列并用空格分隔 使用fmt中的FScan_函数
	file, err2 := os.Open("7_test.txt")
	if err2 != nil {
		panic(err2.Error())
	}
	defer file.Close()
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

func ReadgzippedFile() {
	fmt.Println("==读取gzip文件==")
	fName := "MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

func WriteFileData() {
	fmt.Println("==写文件==")
	os.Stdout.WriteString("os.Stdout 输出到console!\n")
	//文件句柄;
	//只读模式打开文件 如果文件不存在则自创建
	//os.O_WRONLY 只写; os.O_RDONLY 只读;
	//OS.O_CREATE 创建: 指定文件不存在 就创建该文件
	//OS.O_TRUNC 截断: 指定文件已存在 就将该文件的长度截为0
	//在读文件时文件权限是被忽视的, 所以在使用OpenFile时传入第三个参数可以用0;
	//而在写文件时不管时Unix还是Windows 都需要使用0666
	outputFile, outputError := os.OpenFile("7_writeFileData.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("文件打开错误...")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile) //缓冲区: 写入器
	outputString := "Hello World!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString) //写入字符串
	}
	outputWriter.Flush() //缓冲区内容完全写入文件 **使用缓冲区必须加
	//写入简单string可以用下面该段代码
	//fmt包中F开头的Print函数可以直接写入任何io.Writer 包括文件
	fmt.Fprintf(outputFile, "Some test data.\n")
	fmt.Println("Write file success...")
}