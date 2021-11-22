package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1560C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1560C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1560C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1560C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1560C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1560C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1560C
Date: 21/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1560/C
Problem: CF1560C
**/
func (in *CF1560C) Run() {

	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		aux := int64(1)
		init := int64(1)
		x := int64(0)
		for ; init < n; init += aux {
			x++
			aux += 2
		}

		start := init - (aux - 2) - 1

		diff := n - start

		if diff <= int64(math.Ceil(float64(aux/2))) {
			fmt.Println(diff+1, x+1)
		} else {
			diff -= int64(math.Ceil(float64(aux / 2)))
			fmt.Println(x+1, x-diff+1)
		}

	}
}

func NewCF1560C(r *bufio.Reader) *CF1560C {
	return &CF1560C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1560C(bufio.NewReader(os.Stdin)).Run()
}
