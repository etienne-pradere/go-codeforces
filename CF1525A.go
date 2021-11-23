package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1525A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1525A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1525A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1525A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1525A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1525A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1525A
Date: 22/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1525/A
Problem: CF1525A
**/
func (in *CF1525A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		magic := in.NextInt()
		water := 100 - magic
		div := Max(magic, water)
		ans := 0
		for {
			if magic%div == 0 && water%div == 0 {
				magic /= div
				water /= div
				div = Max(magic, water)
			} else {
				div--
			}
			if div == 1 {
				ans = magic + water
				break
			}
		}
		fmt.Println(ans)

	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewCF1525A(r *bufio.Reader) *CF1525A {
	return &CF1525A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1525A(bufio.NewReader(os.Stdin)).Run()
}
