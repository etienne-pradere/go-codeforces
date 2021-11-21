package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1511A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1511A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1511A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1511A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1511A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1511A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1511A
Date: 21/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1511/A
Problem: CF1511A
**/
func (in *CF1511A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		s1 := 0
		s2 := 0
		for i := 0; i < n; i++ {
			tp := in.NextInt()
			if tp == 1 {
				s1++
			} else if tp == 2 {
				s2--
			} else {
				if s1 >= 0 {
					s1++
				} else {
					s2--
				}
			}
		}
		fmt.Println(s1)
	}
}

func NewCF1511A(r *bufio.Reader) *CF1511A {
	return &CF1511A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1511A(bufio.NewReader(os.Stdin)).Run()
}
