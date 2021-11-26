package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1492A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1492A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1492A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1492A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1492A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1492A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1492A
Date: 25/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1492/A
Problem: CF1492A
**/
func (in *CF1492A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		p, a, b, c := in.NextInt64(), in.NextInt64(), in.NextInt64(), in.NextInt64()

		ans := a - (p % a)
		aux := a - (p % a)
		if aux == a {
			ans = 0
		}
		ans = Min(ans, aux)
		aux = b - (p % b)
		if aux == b {
			ans = 0
		}
		ans = Min(ans, aux)
		aux = c - (p % c)
		if aux == c {
			ans = 0
		}
		ans = Min(ans, aux)
		fmt.Println(ans)

	}

}
func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func NewCF1492A(r *bufio.Reader) *CF1492A {
	return &CF1492A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1492A(bufio.NewReader(os.Stdin)).Run()
}
