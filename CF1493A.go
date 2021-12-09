package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1493A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1493A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1493A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1493A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1493A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1493A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1493A
Date: 8/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1493/A
Problem: CF1493A
**/
func (in *CF1493A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		var str strings.Builder
		size := 0
		for i := 1; i <= n; i++ {
			if i != k {
				if k-i > 0 && k-i < k && k-i > i {
					continue
				} else {
					size++
					str.WriteString(strconv.Itoa(i)+ " ")
				}
			}
		}
		fmt.Println(size)
		if size > 0 {
			fmt.Println(str.String())
		}
	}
}

func NewCF1493A(r *bufio.Reader) *CF1493A {
	return &CF1493A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1493A(bufio.NewReader(os.Stdin)).Run()
}
