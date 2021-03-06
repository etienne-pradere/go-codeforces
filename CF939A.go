package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF939A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF939A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF939A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF939A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF939A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF939A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF939A
Date: 10/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/939/A
Problem: CF939A
**/
func (in *CF939A) Run() {
	n := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt() - 1
	}
	ans := false
	for i := 0; i < n; i++ {
		ini := i
		count := 0
		for j := 0; j < 3; j++ {
			ini = arr[ini]
			if ini != i {
				count++
			}
		}
		if ini == i && count == 2 {
			ans = true
			break
		}
	}
	if ans {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func NewCF939A(r *bufio.Reader) *CF939A {
	return &CF939A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF939A(bufio.NewReader(os.Stdin)).Run()
}
