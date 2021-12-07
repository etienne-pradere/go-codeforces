package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1096A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1096A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1096A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1096A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1096A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1096A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1096A
Date: 6/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1096/A
Problem: CF1096A
**/
func (in *CF1096A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		l, r := in.NextInt(), in.NextInt()
		a, b := l, r-(r%l)
		if a == b {
			a++
			b = r-(r%a)
		}
		fmt.Println(a, b)
	}
}

func NewCF1096A(r *bufio.Reader) *CF1096A {
	return &CF1096A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1096A(bufio.NewReader(os.Stdin)).Run()
}
