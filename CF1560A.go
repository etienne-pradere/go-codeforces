package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1560A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1560A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1560A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1560A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1560A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1560A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1560A
Date: 16/09/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1560/A
Problem: CF1560A
**/
func (in *CF1560A) Run() {

	arr := []int{}

	for i := 1; len(arr) < 1001; i++ {
		num := fmt.Sprintf("%d", i)
		if i%3 != 0 && num[len(num)-1] != '3' {
			arr = append(arr, i)
		}
	}

	t := in.NextInt()
	for ; t > 0; t-- {
		k := in.NextInt()
		fmt.Println(arr[k-1])
	}
}

func NewCF1560A(r *bufio.Reader) *CF1560A {
	return &CF1560A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1560A(bufio.NewReader(os.Stdin)).Run()
}
