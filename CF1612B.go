package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1612B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1612B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1612B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1612B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1612B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1612B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1612B
Date: 10/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1612/B
Problem: CF1612B
**/
func (in *CF1612B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, a, b := in.NextInt(), in.NextInt(), in.NextInt()
		if (a <= n/2 && b <= n/2) || (a > n/2 && b > n/2) {
			fmt.Println(-1)
		} else {
			arr := make([]int, n)
			arr[0] = a
			ans := true
			for i, j := n, 1; i >= 1; i-- {
				if i != a && i != b {
					arr[j] = i
					j++
					if i > n/2 && i <= a {
						ans = false
					} else if i <= n/2 && i >= b {
						ans = false
					}
				}
			}
			var sb strings.Builder
			arr[n-1] = b
			if ans {
				for i := 0; i < n; i++ {
					sb.WriteString(strconv.Itoa(arr[i]))
					sb.WriteString(" ")
				}
			} else {
				sb.WriteString("-1")
			}
			fmt.Println(sb.String())
		}
	}
}

func NewCF1612B(r *bufio.Reader) *CF1612B {
	return &CF1612B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1612B(bufio.NewReader(os.Stdin)).Run()
}
