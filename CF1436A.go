package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1436A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1436A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1436A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1436A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1436A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1436A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1436A
Date: 2/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1436/A
Problem: CF1436A
**/
func (in *CF1436A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		sum :=0
		for i := 0; i < n; i++ {
			sum += in.NextInt()
		}
		if sum == m {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1436A(r *bufio.Reader) *CF1436A {
	return &CF1436A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1436A(bufio.NewReader(os.Stdin)).Run()
}
