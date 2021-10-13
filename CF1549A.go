package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1549A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1549A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1549A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1549A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1549A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1549A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1549A
Date: 12/10/21
User: epradere
URL: https://codeforces.com/problemset/problem/1549/A
Problem: CF1549A
**/
func (in *CF1549A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		P := in.NextInt64()
		fmt.Println(2, int(P/2)*2)
	}
}

func NewCF1549A(r *bufio.Reader) *CF1549A {
	return &CF1549A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1549A(bufio.NewReader(os.Stdin)).Run()
}
