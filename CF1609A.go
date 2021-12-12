package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1609A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1609A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1609A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1609A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1609A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1609A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1609A
Date: 10/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1609/A
Problem: CF1609A
**/
func (in *CF1609A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		all := int64(0)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			all += int64(arr[i])
		}
		mult := 1
		for i := 0; i < n; i++ {
			for arr[i]%2 == 0 {
				arr[i] /= 2
				mult *= 2
			}
		}
		sort.Ints(arr)
		arr[n-1] *= mult
		sum := int64(0)

		for i := 0; i < n; i++ {
			sum += int64(arr[i])
		}
		fmt.Println(sum)
	}
}

func NewCF1609A(r *bufio.Reader) *CF1609A {
	return &CF1609A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1609A(bufio.NewReader(os.Stdin)).Run()
}
