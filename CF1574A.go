package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1574A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1574A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1574A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1574A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1574A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1574A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1574A
Date: 6/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1574/A
Problem: CF1574A
**/
func (in *CF1574A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		for i := 0; i < n; i++ {
			str := ""
			cont := 0
			ws := false
			for j := 0; j < 2*n; j++ {
				if cont <= i && !ws {
					str += "("
					cont++
				} else if ws && cont <= 0 {
					if -cont%2 == 0 {
						str += "("
					} else {
						str += ")"
					}
					cont--
				} else {
					ws = true
					str += ")"
					cont--
				}

			}
			fmt.Println(str)

		}
	}
}

func NewCF1574A(r *bufio.Reader) *CF1574A {
	return &CF1574A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1574A(bufio.NewReader(os.Stdin)).Run()
}
