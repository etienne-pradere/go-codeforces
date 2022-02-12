package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1618D struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1618D) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1618D) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1618D) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1618D) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1618D) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1618D
Date: 15/12/21
User: epradere
URL: https://codeforces.com/contest/1618/problem/D
Problem: CF1618D
**/
func (in *CF1618D) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}

		sort.Ints(arr)
		total := 0
		for i, j, p := 0, n-1, 0; p < k; p++ {
			total += int(math.Floor(float64(arr[i] / arr[j])))
			i++
			j--
		}

		for i := k; i < n-k; i++ {
			total += arr[i]
		}
		fmt.Println(total)
	}
}

func f(i, j, k int) int {

	return int(math.Min(float64(f(i, j+1))))
}

func NewCF1618D(r *bufio.Reader) *CF1618D {
	return &CF1618D{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1618D(bufio.NewReader(os.Stdin)).Run()
}
