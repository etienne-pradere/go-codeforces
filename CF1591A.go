package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1591A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1591A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1591A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1591A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1591A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1591A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1591A
Date: 12/12/21
User: epradere
URL: https://codeforces.com/contest/1591/problem/0
Problem: CF1591A
**/
func (in *CF1591A) Run() {

	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		size := 1
		live := true
		lastDay := 0
		for i := 0; i < n; i++ {
			day := in.NextInt()
			if day > 0 {
				size++
			}
			if i > 0 {
				if lastDay == 1 && day == 1 {
					size += 4
				}
				if day == 0 && lastDay == 0 {
					live = false
				}
			}
			lastDay = day
		}

		if !live {
			fmt.Println(-1)
		} else {
			fmt.Println(size)
		}

	}
}

func NewCF1591A(r *bufio.Reader) *CF1591A {
	return &CF1591A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1591A(bufio.NewReader(os.Stdin)).Run()
}
