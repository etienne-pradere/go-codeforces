package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF344A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF344A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF344A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF344A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF344A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF344A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF344A
Date: 22/09/21
User: epradere
URL: http://codeforces.com/contest/344/problem/A
Problem: CF344A
**/
func (in *CF344A) Run() {
	n := in.NextInt()
	ans := 1
	var prev uint8
	for i := 0; i < n; i++ {
		num := in.NextString()
		if i == 0 {
			prev = num[1]
		} else {
			if prev == num[0] {
				ans++
			}
			prev = num[1]
		}
	}
	fmt.Println(ans)
}

func NewCF344A(r *bufio.Reader) *CF344A {
	return &CF344A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF344A(bufio.NewReader(os.Stdin)).Run()
}
