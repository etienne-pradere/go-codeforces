package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF231A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF231A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF231A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF231A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF231A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF231A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF231A
Date: 2/09/21
User: wotan
URL: https://codeforces.com/contest/231/problem/A
Problem: CF231A
**/
func (in *CF231A) Run() {
	n := in.NextInt()
	ans := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			sum += in.NextInt()
		}
		if sum >= 2 {
			ans++
		}
	}
	fmt.Println(ans)
}

func NewCF231A(r *bufio.Reader) *CF231A {
	return &CF231A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF231A(bufio.NewReader(os.Stdin)).Run()
}
