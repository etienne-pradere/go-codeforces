package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1618B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1618B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1618B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1618B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1618B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1618B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1618B
Date: 15/12/21
User: epradere
URL: https://codeforces.com/contest/1618/problem/B
Problem: CF1618B
**/
func (in *CF1618B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		var sb strings.Builder
		last := ""
		for i := 0; i < n-2; i++ {
			str := in.NextString()
			if i == 0 {
				sb.WriteString(str)
			} else {
				if last[1:] != str[0:1] {
					sb.WriteString(str)
				} else {
					sb.WriteString(str[1:])
				}
			}
			last = str
		}
		for sb.Len() < n {
			sb.WriteString("b")
		}
		fmt.Println(sb.String())
	}
}

func NewCF1618B(r *bufio.Reader) *CF1618B {
	return &CF1618B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1618B(bufio.NewReader(os.Stdin)).Run()
}
