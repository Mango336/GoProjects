package pack7

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput1() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)
	fmt.Println("==readinput1 使用scan_==")
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// or: fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi! %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}

func ReadInput2() {
	fmt.Println("==readinput2 使用bufio==")
	var inputReader *bufio.Reader //指向bufio.Reader的指针
	//创建一个读取器 并将其与标准输入绑定
	//返回一个新的带缓存的io.Reader对象 它将从指定读取器(如:os.Stdin)中读取内容
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

func SwitchInput() {
	fmt.Println("==switch_input 使用bufio==")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!")
		return
	}
	fmt.Println("Your name is ", input)  // 这里input后面有 \n
	// version 1
	switch input {
	case "Tom\n": fmt.Println("Version1: Hello Tom!")
	case "Jonney\n": fmt.Println("Version1: Hello Jonney!")
	case "Mike\n":fmt.Println("Version1: Hello Mike!")
	default: fmt.Println("Version1: Goodbye!")
	}
	// version 2
	switch input {
	case "Tom\n": fallthrough
	case "Jonney\n": fallthrough
	case "Mike\n": fmt.Printf("Version2: Hello %s", input)
	default: fmt.Println("Version2: Goodbye!")
	}
	// version 3
	switch input {
	case "Tom\n", "Jonney\n", "Mike\n":
		fmt.Printf("Version3: Hello %s", input)
	default: fmt.Println("Version3: Goodbye!")
	}

}
