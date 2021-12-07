package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1504A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1504A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1504A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1504A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1504A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1504A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1504A
Date: 6/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1504/A
Problem: CF1504A
**/
func (in *CF1504A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()

		ws := p(string('a') + str)
		if ws {
			ws = p(str + string('a'))
			str = str + string('a')
		} else {
			str = string('a') + str
		}
		if !ws {
			fmt.Println("YES")
			fmt.Println(str)
		} else {
			fmt.Println("NO")
		}
	}
}

func p(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func NewCF1504A(r *bufio.Reader) *CF1504A {
	return &CF1504A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1504A(bufio.NewReader(os.Stdin)).Run()
}
