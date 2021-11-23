package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1391B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1391B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1391B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1391B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1391B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1391B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1391B
Date: 22/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1391/B
Problem: CF1391B
**/
func (in *CF1391B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		ans := 0
		for i := 0; i < n; i++ {
			arr := make([]rune, m)
			str := in.NextString()
			for j := 0; j < m; j++ {
				arr[j] = rune(str[j])
				if j == m-1 && arr[j] != 'D' && arr[j] != 'C' {
					ans++
				}
				if i == n-1 && arr[j] != 'R' && arr[j] != 'C' {
					ans++
				}
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1391B(r *bufio.Reader) *CF1391B {
	return &CF1391B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1391B(bufio.NewReader(os.Stdin)).Run()
}
