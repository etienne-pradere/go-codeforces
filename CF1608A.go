package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1608A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1608A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1608A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1608A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1608A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1608A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1608A
Date: 11/12/21
User: epradere
URL: https://codeforces.com/contest/1608/problem/0
Problem: CF1608A
**/
func (in *CF1608A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		if n == 1 {
			fmt.Println(1)
		} else {
			for i := 2; i < n+2; i++ {
				fmt.Printf("%d ", i)
			}
			fmt.Println()
		}
	}
}

func NewCF1608A(r *bufio.Reader) *CF1608A {
	return &CF1608A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1608A(bufio.NewReader(os.Stdin)).Run()
}
