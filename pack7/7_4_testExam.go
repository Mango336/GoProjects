package pack7

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Remove_3till5Chars() {
	inputFile, _ := os.Open("data/7_input_text.txt")
	outputFile, _ := os.OpenFile("data/7_output_text2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		outputString := string(inputString[2:5]) + "\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	outputWriter.Flush()
	fmt.Println("Conversion done")
}
