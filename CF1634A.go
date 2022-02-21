package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1634A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1634A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1634A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1634A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1634A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1634A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1634A
Date: 19/02/22
User: epradere
URL: https://codeforces.com/problemset/problem/1634/A
Problem: CF1634A
**/
func (in *CF1634A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		str := in.NextString()

		mapa := make(map[uint8]int)
		for i := 0; i < n; i++ {
			val := mapa[str[i]]
			mapa[str[i]] = val + 1
		}
		min := len(mapa)

		if min > k {
			min = k
		}
		if min == 0 {
			min = 1
		}

		fmt.Println(min)
	}
}

func NewCF1634A(r *bufio.Reader) *CF1634A {
	return &CF1634A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1634A(bufio.NewReader(os.Stdin)).Run()
}
