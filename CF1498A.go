package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1498A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1498A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1498A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1498A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1498A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1498A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1498A
Date: 11/02/22
User: epradere
URL: https://codeforces.com/problemset/problem/1498/A
Problem: CF1498A
**/
func (in *CF1498A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		num, _ := strconv.ParseInt(str, 10, 64)

		ans := 1
		sol := num
		for ans == 1 {
			sum := 0
			for i := 0; i < len(str); i++ {
				sum += int(str[i]) - '0'
			}
			for i := sum; i >= 1; i-- {
				if sum%i == 0 && num%int64(i) == 0 {
					ans = i
					sol = num
					break
				}
			}
			num++
			str = fmt.Sprintf("%v", num)

		}

		fmt.Println(sol)

	}
}

func NewCF1498A(r *bufio.Reader) *CF1498A {
	return &CF1498A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1498A(bufio.NewReader(os.Stdin)).Run()
}
