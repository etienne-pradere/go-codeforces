package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1392A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1392A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1392A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1392A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1392A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1392A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1392A
Date: 8/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1392/A
Problem: CF1392A
**/
func (in *CF1392A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := map[int]int{}
		ans := 1
		for i := 0; i < n; i++ {
			num := in.NextInt()
			val := mapa[num]
			mapa[num] = val + 1
			if val+1 == n {
				ans = n
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1392A(r *bufio.Reader) *CF1392A {
	return &CF1392A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1392A(bufio.NewReader(os.Stdin)).Run()
}
