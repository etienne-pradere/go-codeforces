package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1562A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1562A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1562A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1562A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1562A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1562A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1562A
Date: 22/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1562/A
Problem: CF1562A
**/
func (in *CF1562A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		l, r := in.NextInt64(), in.NextInt64()
		ans := int64(math.Ceil(float64(r)/2.0) - 1)
		if ans < l {
			ans = r % l
			if l+1 < r {
				ans = int64(math.Max(float64(ans), float64(r%(l+1))))
			}
		}
		fmt.Println(ans)
	}

}

func NewCF1562A(r *bufio.Reader) *CF1562A {
	return &CF1562A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1562A(bufio.NewReader(os.Stdin)).Run()
}
