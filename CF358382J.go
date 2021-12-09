package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF358382J struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF358382J) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF358382J) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF358382J) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF358382J) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF358382J) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF358382J
Date: 8/12/21
User: epradere
URL: https://codeforces.com/gym/358382/problem/J
Problem: CF358382J
**/
var n, a, b, c int
var memo []int

func (in *CF358382J) Run() {
	n, a, b, c = in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
	memo = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -2
	}
	fmt.Println(f(n) )
}

func f(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}

	if memo[n] > -2 {
		return memo[n]
	}

	x := f(n - a)
	y := f(n - b)
	z := f(n - c)

	max := max(x, max(y, z))
	if max >= 0 {
		memo[n] = max + 1
		return max + 1
	}
	memo[n] = max
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewCF358382J(r *bufio.Reader) *CF358382J {
	return &CF358382J{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF358382J(bufio.NewReader(os.Stdin)).Run()
}
