package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1466A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1466A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1466A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1466A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1466A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1466A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1466A
Date: 22/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1466/A
Problem: CF1466A
**/
func (in *CF1466A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		mapa := make(map[float64]int)
		for i := 0; i < n; i++ {
			a1 := float64(arr[i]*1.0) / 2.0
			for j := i + 1; j < n; j++ {
				a2 := float64(arr[j]*1.0) / 2.0
				sol := math.Abs(a2 - a1)
				_, ok := mapa[sol]
				if !ok {
					mapa[sol] = 1
				}
			}
		}

		fmt.Println(len(mapa))

	}
}

func Max(a, b, c int) int {
	if a > b && a > c {
		return a
	}
	if b > a && b > c {
		return b
	}
	return c
}

func Dist(x1, x2, y1, y2 int) float64 {
	return math.Sqrt(float64(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1))))
}

func NewCF1466A(r *bufio.Reader) *CF1466A {
	return &CF1466A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1466A(bufio.NewReader(os.Stdin)).Run()
}
