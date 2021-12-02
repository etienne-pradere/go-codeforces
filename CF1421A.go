package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1421A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1421A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1421A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1421A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1421A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1421A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1421A
Date: 30/11/21
User: epradere
URL: http://codeforces.com/problemset/problem/1421/A
Problem: CF1421A
**/
func (in *CF1421A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b := in.NextInt(), in.NextInt()
		fmt.Println(a^b)
	}

}

func NewCF1421A(r *bufio.Reader) *CF1421A {
	return &CF1421A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1421A(bufio.NewReader(os.Stdin)).Run()
}
