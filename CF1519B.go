package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1519B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1519B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1519B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1519B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1519B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1519B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

var memo map[int]map[int]map[int]int

/**
Run solve the problem CF1519B
Date: 13/10/21
User: epradere
URL: https://codeforces.com/problemset/problem/1519/B
Problem: CF1519B
**/
func (in *CF1519B) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n, m, k := in.NextInt(), in.NextInt(), in.NextInt()

		memo = make(map[int]map[int]map[int]int)

		ans := f(n, m, k)
		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func f(i, j, n int) bool {
	if i == 1 && j == 1 && n == 0 {
		return true
	}
	if i < 1 || j < 1 || n < 0 {
		return false
	}
	if val, ok := memo[i][j][n]; ok && val != 0 {
		return memo[i][j][n] == 1
	}

	ans := f(i, j-1, n-i) || f(i-1, j, n-j)
	if ans {
		save(i, j, n, 1)
	} else {
		save(i, j , n, 2)
	}
	return memo[i][j][n] == 1

}

func save(i, j, n, val int) {
	if _, ok := memo[i]; !ok {
		memo[i] = make(map[int]map[int]int)
	}
	if _, ok := memo[i][j]; !ok {
		memo[i][j] = make(map[int]int)
	}
		memo[i][j][n] = val
}

func NewCF1519B(r *bufio.Reader) *CF1519B {
	return &CF1519B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1519B(bufio.NewReader(os.Stdin)).Run()
}
