package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1452A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1452A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1452A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1452A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1452A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1452A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1452A
Date: 25/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1452/A
Problem: CF1452A
**/
func (in *CF1452A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		x, y := in.NextInt(), in.NextInt()
		if x == y {
			fmt.Println(x * 2)
		} else {
			if x > y {
				fmt.Println((x * 2) - 1)
			} else {
				fmt.Println((y * 2) - 1)
			}
		}
	}
}

func NewCF1452A(r *bufio.Reader) *CF1452A {
	return &CF1452A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1452A(bufio.NewReader(os.Stdin)).Run()
}
