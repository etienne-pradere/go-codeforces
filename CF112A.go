package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF112A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF112A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF112A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF112A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF112A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF112A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF112A
Date: 22/09/21
User: epradere
URL: http://codeforces.com/contest/112/problem/A
Problem: CF112A
**/
func (in *CF112A) Run() {
	a := strings.ToLower(in.NextString())
	b := strings.ToLower(in.NextString())
	fmt.Println(strings.Compare(a,b))
}

func NewCF112A(r *bufio.Reader) *CF112A {
	return &CF112A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF112A(bufio.NewReader(os.Stdin)).Run()
}
