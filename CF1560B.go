package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1560B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1560B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1560B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1560B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1560B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1560B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1560B
Date: 16/09/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1560/B
Problem: CF1560B
**/
func (in *CF1560B) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		a, b, c := in.NextInt64(), in.NextInt64(), in.NextInt64()

		max := Max(a, b) - Min(a, b)

		ans := int64(-1)

		if c <= max*2 && max > 1 && a <= max*2 && b <= max*2 {
			if c+max <= max*2 {
				ans = c + max
			} else if c-max >= 1 {
				ans = c - max
			}
		}
		fmt.Println(ans)
	}
}
func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func NewCF1560B(r *bufio.Reader) *CF1560B {
	return &CF1560B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1560B(bufio.NewReader(os.Stdin)).Run()
}
