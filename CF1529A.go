package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1529A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1529A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1529A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1529A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1529A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1529A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1529A
Date: 11/10/21
User: epradere
URL: https://codeforces.com/problemset/problem/1529/A
Problem: CF1529A
**/
func (in *CF1529A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		m := map[int]int{}
		sum := 0
		for i := 0; i < n; i++ {
			num := in.NextInt()
			sum += num
			if _, ok := m[num]; !ok {
				m[num] = 0
			}
			m[num]++
		}
		sol := 0
		avg := sum / n
		for {
			deleted := 0
			for i, value := range m {
				if i > avg {
					sum -= i * value
					n -= value
					deleted += value
					sol += value
					delete(m, i)
				}
			}
			if deleted > 0 {
				avg = sum / n
			} else {
				break
			}

		}
		fmt.Println(sol)
	}
}

func NewCF1529A(r *bufio.Reader) *CF1529A {
	return &CF1529A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1529A(bufio.NewReader(os.Stdin)).Run()
}
