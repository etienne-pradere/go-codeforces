package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF746A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF746A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF746A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF746A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF746A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF746A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF746A
Date: 8/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/746/A
Problem: CF746A
**/
func (in *CF746A) Run() {
	a, b, c := in.NextInt(), in.NextInt(), in.NextInt()

	sol := 0
	for i := 1; i <= a; i++ {
		if i*2 <= b && i*4 <= c {
			sol = i + (2 * i) + (4 * i)
		}
	}
	fmt.Println(sol)
}

func NewCF746A(r *bufio.Reader) *CF746A {
	return &CF746A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF746A(bufio.NewReader(os.Stdin)).Run()
}
