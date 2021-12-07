package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF984A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF984A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF984A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF984A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF984A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF984A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF984A
Date: 6/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/984/A
Problem: CF984A
**/
func (in *CF984A) Run() {
	n := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}
	sort.Ints(arr)
	a := 0
	if len(arr)%2 == 0 {
		a = -1
	}
	fmt.Println(arr[(len(arr)/2) + a])
}

func NewCF984A(r *bufio.Reader) *CF984A {
	return &CF984A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF984A(bufio.NewReader(os.Stdin)).Run()
}
