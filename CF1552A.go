package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1552A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1552A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1552A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1552A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1552A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1552A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1552A
Date: 25/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1552/A
Problem: CF1552A
**/
func (in *CF1552A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		ord := make([]int, n)
		str := in.NextString()
		for i := 0; i < n; i++ {
			arr[i] = int(str[i])
			ord[i] = int(str[i])
		}

		sort.Ints(ord)

		ans := 0
		for i := 0; i < n; i++ {
			if arr[i] != ord[i] {
				ans++
			}
		}

		fmt.Println(ans)
	}
}

func NewCF1552A(r *bufio.Reader) *CF1552A {
	return &CF1552A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1552A(bufio.NewReader(os.Stdin)).Run()
}
