package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1530A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1530A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1530A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1530A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1530A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1530A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1530A
Date: 12/02/22
User: epradere
URL: https://codeforces.com/problemset/problem/1530/A
Problem: CF1530A
**/
func (in *CF1530A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		max := 0
		for i := 0; i < len(str); i++ {
			if max < int(str[i]-'0') {
				max = int(str[i] - '0')
			}
		}
		fmt.Println(max)
	}
}

func NewCF1530A(r *bufio.Reader) *CF1530A {
	return &CF1530A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1530A(bufio.NewReader(os.Stdin)).Run()
}
