package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1080A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1080A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1080A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1080A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1080A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1080A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1080A
Date: 6/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1080/A
Problem: CF1080A
**/
func (in *CF1080A) Run() {
	n, k := in.NextInt(), in.NextInt()
	r := n * 2.0
	g := n * 5.0
	b := n * 8.0

	ans := math.Ceil(float64(r)/float64(k)) + math.Ceil(float64(g)/float64(k)) + math.Ceil(float64(b)/float64(k))
	fmt.Println(int(ans))
}

func NewCF1080A(r *bufio.Reader) *CF1080A {
	return &CF1080A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1080A(bufio.NewReader(os.Stdin)).Run()
}
