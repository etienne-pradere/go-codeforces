package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1139A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1139A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1139A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1139A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1139A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1139A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1139A
Date: 6/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1139/A
Problem: CF1139A
**/
func (in *CF1139A) Run() {
	n := in.NextInt()
	str := in.NextString()
	ans := 0
	for i := 0; i < n; i++ {
		if (str[i]-'0')%2 == 0 {
			ans += i+1
		}
	}
	fmt.Println(ans)
}

func NewCF1139A(r *bufio.Reader) *CF1139A {
	return &CF1139A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1139A(bufio.NewReader(os.Stdin)).Run()
}
