package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF791A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF791A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF791A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF791A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF791A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF791A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF791A
Date: 2/09/21
User: pepradere
URL: https://codeforces.com/contest/791/problem/A
Problem: CF791A
**/
func (in *CF791A) Run() {
	a := in.NextInt()
	b := in.NextInt()
	sol := 0
	for ;a <= b; {
		a *= 3
		b *= 2
		sol++
	}
	fmt.Println(sol)
}

func NewCF791A(r *bufio.Reader) *CF791A {
	return &CF791A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF791A(bufio.NewReader(os.Stdin)).Run()
}
