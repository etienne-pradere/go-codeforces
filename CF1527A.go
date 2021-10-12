package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1527A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1527A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1527A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1527A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1527A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1527A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1527A
Date: 16/09/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1527/A
Problem: CF1527A
**/
func (in *CF1527A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt64()
		count := int64(1)
		for n > 1 {
			count++
			n /= 2
		}

		fmt.Println(1<<(count-1) - 1)
	}
}

func NewCF1527A(r *bufio.Reader) *CF1527A {
	return &CF1527A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1527A(bufio.NewReader(os.Stdin)).Run()
}
