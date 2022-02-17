package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1440A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1440A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1440A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1440A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1440A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1440A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1440A
Date: 13/02/22
User: epradere
URL: https://codeforces.com/problemset/problem/1440/A
Problem: CF1440A
**/
func (in *CF1440A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, c0, c1, h := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
		str := in.NextString()
		currentCost := 0
		zeroes, ones := 0, 0
		for i := 0; i < n; i++ {
			if str[i] == '0' {
				zeroes++
				currentCost += c0
			} else {
				ones++
				currentCost += c1
			}
		}
		if c0 < c1 {
			c0Cost := (h * ones) + ((zeroes + ones) * c0)
			if c0Cost < currentCost {
				currentCost = c0Cost
			}
		} else if c1 < c0 {
			c1Cost := (h * zeroes) + ((zeroes + ones) * c1)
			if c1Cost < currentCost {
				currentCost = c1Cost
			}
		}

		fmt.Println(currentCost)

	}
}

func NewCF1440A(r *bufio.Reader) *CF1440A {
	return &CF1440A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1440A(bufio.NewReader(os.Stdin)).Run()
}
