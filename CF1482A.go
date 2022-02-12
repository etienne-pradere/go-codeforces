package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1482A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1482A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1482A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1482A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1482A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1482A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1482A
Date: 11/02/22
User: epradere
URL: https://codeforces.com/problemset/problem/1482/A
Problem: CF1482A
**/
func (in *CF1482A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b := in.NextInt(), in.NextInt()
		fmt.Println(a * b)
	}
}

func NewCF1482A(r *bufio.Reader) *CF1482A {
	return &CF1482A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1482A(bufio.NewReader(os.Stdin)).Run()
}
