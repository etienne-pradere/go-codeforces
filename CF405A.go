package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF405A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF405A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF405A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF405A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF405A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF405A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF405A
Date: 22/09/21
User: epradere
URL: http://codeforces.com/contest/405/problem/A
Problem: CF405A
**/
func (in *CF405A) Run() {
	n := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		val := in.NextInt()
		arr[i] = val
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if arr[i] > arr[j]{
				aux := arr[i] - arr[j]
				arr[j] += aux
				arr[i] -= aux
			}
		}
	}

	for i := range arr {
		if i == 0 {
			fmt.Print(arr[i])
		} else {
			fmt.Print(" ", arr[i])
		}

	}

}

func NewCF405A(r *bufio.Reader) *CF405A {
	return &CF405A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF405A(bufio.NewReader(os.Stdin)).Run()
}
