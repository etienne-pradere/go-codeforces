package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1453A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1453A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1453A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1453A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1453A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1453A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1453A
Date: 6/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1453/A
Problem: CF1453A
**/
func (in *CF1453A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		mapa := make(map[int]int)
		for i := 0; i < n; i++ {
			mapa[in.NextInt()] = 1
		}
		ans := 0
		for i := 0; i < m; i++ {
			_, ok := mapa[in.NextInt()]
			if ok {
				ans++
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1453A(r *bufio.Reader) *CF1453A {
	return &CF1453A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1453A(bufio.NewReader(os.Stdin)).Run()
}
