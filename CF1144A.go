package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1144A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1144A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1144A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1144A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1144A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1144A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1144A
Date: 8/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1144/A
Problem: CF1144A
**/
func (in *CF1144A) Run() {
	for t := in.NextInt64(); t > 0; t-- {
		str := in.NextString()
		arr := make([]int, len(str))
		for i := 0; i < len(arr); i++ {
			arr[i] = int(str[i])
		}
		sort.Ints(arr)
		ans := true
		for i := 1; i < len(arr); i++ {
			if arr[i-1]+1 != arr[i]{
				ans = false
				break
			}
		}
		if ans {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func NewCF1144A(r *bufio.Reader) *CF1144A {
	return &CF1144A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1144A(bufio.NewReader(os.Stdin)).Run()
}
