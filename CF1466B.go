package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1466B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1466B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1466B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1466B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1466B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1466B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1466B
Date: 25/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1466/B
Problem: CF1466B
**/
func (in *CF1466B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := make(map[int]int)
		for i := 0; i < n; i++ {
			num := in.NextInt()
			_, ok := mapa[num]
			if !ok {
				mapa[num] = 1
			} else {
				mapa[num+1] = 1
			}

		}
		fmt.Println(len(mapa))
	}
}

func NewCF1466B(r *bufio.Reader) *CF1466B {
	return &CF1466B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1466B(bufio.NewReader(os.Stdin)).Run()
}
