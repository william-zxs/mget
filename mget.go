package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {

	// 获取命令行参数
	args := os.Args
	if args == nil || len(args) < 2 {
		fmt.Println("your key word?")
		return
	}
	keyWord := args[1]

	inputFile, inputError := os.Open("/Users/william/.ssh/william")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	hasKeyWord := false
	paragraph := ""

	for {
		inputString, readerError := inputReader.ReadString('\n')

		// 判断是否有keyWord
		validKeyWord := regexp.MustCompile(keyWord)
		res := validKeyWord.MatchString(inputString)
		if res == true {
			hasKeyWord = true
		}

		if inputString == "\n" {
			// fmt.Println("==空格===")
			if hasKeyWord == true {
				// fmt.Println("===包含关键字的段落====")
				fmt.Println(paragraph)
				hasKeyWord = false
			}
			paragraph = ""

		} else {
			paragraph = paragraph + inputString
		}
		if readerError == io.EOF {
			return
		}
	}
}