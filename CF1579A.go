package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1579A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1579A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1579A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1579A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1579A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1579A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1579A
Date: 11/10/21
User: epradere
URL: https://codeforces.com/problemset/problem/1579/A
Problem: CF1579A
**/
func (in *CF1579A) Run() {
	n := in.NextInt()
	for ; n > 0; n-- {
		s := in.NextString()
		a, b, c := 0, 0, 0
		for i := 0; i < len(s); i++ {
			if s[i] == 'A' {
				a++
			} else if s[i] == 'B' {
				b++
			} else {
				c++
			}
		}

		if a > 0 {
			b -= a
		}

		if c > 0 {
			b -= c
		}

		if b == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1579A(r *bufio.Reader) *CF1579A {
	return &CF1579A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1579A(bufio.NewReader(os.Stdin)).Run()
}
