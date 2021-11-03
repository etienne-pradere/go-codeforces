package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	allLines := []string{}
	maxLength := 0

	for {

		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		text = strings.Replace(text, "\n", "", -1)

		maxLength = max(maxLength, len(text))
		allLines = append(allLines, text)
	}

	left := true
	printLine(maxLength)
	for _, line := range allLines {
		if printSentence(maxLength, line, left) {
			left = !left
		}
	}
	printLine(maxLength)
}

func printLine(length int) {
	for i := 0; i < length+2; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func printSentence(length int, sentence string, left bool) bool {

	spaces := (length - len(sentence)) / 2
	extraSpace := false
	if (length-len(sentence))%2 != 0 {
		extraSpace = true
	}

	fmt.Print("*")
	if extraSpace && !left {
		fmt.Print(" ")
	}

	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print(sentence)
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}

	if extraSpace && left {
		fmt.Print(" ")
	}
	fmt.Println("*")

	return extraSpace
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}