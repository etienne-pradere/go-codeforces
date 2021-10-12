package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1547A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1547A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1547A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1547A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1547A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1547A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1547A
Date: 11/10/21
User: epradere
URL: https://codeforces.com/problemset/problem/1547/A
Problem: CF1547A
**/
func (in *CF1547A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		in.NextString()
		ax, ay := in.NextInt(), in.NextInt()
		bx, by := in.NextInt(), in.NextInt()
		fx, fy := in.NextInt(), in.NextInt()
		sol := 0
		if ax == bx {
			min := Min(ay, by)
			max := Max(ay, by)
			sol = max - min
			if ax == fx && min <= fy && fy <= max {
				sol += 2
			}
		} else if ay == by {
			min := Min(ax, bx)
			max := Max(ax, bx)
			sol = max - min
			if ay == fy && min <= fx && fx <= max {
				sol += 2
			}
		} else {
			sol = int(math.Abs(float64(ax-bx))) + int(math.Abs(float64(ay-by)))
		}
		fmt.Println(sol)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewCF1547A(r *bufio.Reader) *CF1547A {
	return &CF1547A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1547A(bufio.NewReader(os.Stdin)).Run()
}
