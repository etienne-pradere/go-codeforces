package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF381A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF381A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF381A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF381A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF381A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF381A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF381A
Date: 22/09/21
User: epradere
URL: http://codeforces.com/contest/381/problem/A
Problem: CF381A
**/
func (in *CF381A) Run() {
	n := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}

	a, b := 0, 0
	for i, j, p := 0, n-1, 0; i <= j; p++ {
		if arr[i] > arr[j] {
			if p%2 == 0 {
				a += arr[i]
			} else {
				b += arr[i]
			}
			i++
		} else {
			if p%2 == 0 {
				a += arr[j]
			} else {
				b += arr[j]
			}
			j--
		}
	}

	fmt.Println(a, b)
}

func NewCF381A(r *bufio.Reader) *CF381A {
	return &CF381A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF381A(bufio.NewReader(os.Stdin)).Run()
}
