package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF427A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF427A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF427A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF427A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF427A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF427A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF427A
Date: 23/09/21
User: epradere
URL: http://codeforces.com/contest/427/problem/A
Problem: CF427A
**/
func (in *CF427A) Run() {
	n := in.NextInt()
	ans := 0
	polices := 0
	for i := 0; i < n; i++ {
		num := in.NextInt()

		if num == -1 && polices == 0 {
			ans++
		} else if num == -1 && polices > 0 {
			polices--
		} else if num > 0 {
			polices += num
		}
	}
	fmt.Println(ans)
}

func NewCF427A(r *bufio.Reader) *CF427A {
	return &CF427A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF427A(bufio.NewReader(os.Stdin)).Run()
}
