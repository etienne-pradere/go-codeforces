package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1541A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1541A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1541A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1541A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1541A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1541A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1541A
Date: 11/10/21
User: epradere
URL: https://codeforces.com/problemset/problem/1541/A
Problem: CF1541A
**/
func (in *CF1541A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		sol := make([]int, n)

		for i := 1; i < n; i += 2 {
			sol[i] = i
			sol[i-1] = i + 1
		}

		if n % 2 != 0 {
			sol[n-2], sol[n-1] = n, sol[n-2]
		}

		for i := 0; i < n; i++ {
			fmt.Print(sol[i], " ")
		}
		fmt.Println()
	}
}

func NewCF1541A(r *bufio.Reader) *CF1541A {
	return &CF1541A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1541A(bufio.NewReader(os.Stdin)).Run()
}
