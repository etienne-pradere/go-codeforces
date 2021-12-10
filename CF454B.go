package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF454B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF454B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF454B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF454B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF454B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF454B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF454B
Date: 9/12/21
User: epradere
URL: https://codeforces.com/contest/454/problem/B
Problem: CF454B
**/
func (in *CF454B) Run() {
	n := in.NextInt()
	arr := make([]int, n)
	indexMin := 0
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
		if arr[indexMin] >= arr[i] && i > 0 && arr[i-1] > arr[i] {
			indexMin = i
		}
	}

	ord := true
	for i, p := (indexMin+1)%n, 1; p < n; p++ {
		if arr[(i - 1 + n)%n] > arr[i] && arr[i] != arr[(i+1)%n] {
			ord = false
			break
		}
		i = (i + 1) % n
	}

	if !ord {
		fmt.Println(-1)
	} else {
		if indexMin != 0 {
			fmt.Println(n - indexMin)
		} else {
			fmt.Println(0)
		}
	}

}

func NewCF454B(r *bufio.Reader) *CF454B {
	return &CF454B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF454B(bufio.NewReader(os.Stdin)).Run()
}
