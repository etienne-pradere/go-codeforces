package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF59A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF59A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF59A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF59A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF59A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF59A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF59A
Date: 22/09/21
User: epradere
URL: http://codeforces.com/contest/59/problem/A
Problem: CF59A
**/
func (in *CF59A) Run() {
	word := in.NextString()
	upper := 0
	for i := 0; i < len(word); i++ {
		if word[i] <= 'Z' && word[i] >= 'A' {
			upper++
		}
	}

	lower := len(word) - upper
	if lower == upper || lower > upper {
		fmt.Println(strings.ToLower(word))
	} else {
		fmt.Println(strings.ToUpper(word))
	}
}

func NewCF59A(r *bufio.Reader) *CF59A {
	return &CF59A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF59A(bufio.NewReader(os.Stdin)).Run()
}
