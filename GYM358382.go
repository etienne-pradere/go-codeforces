package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GYM358382 struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *GYM358382) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *GYM358382) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *GYM358382) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *GYM358382) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *GYM358382) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem GYM358382
Date: 7/12/21
User: epradere
URL: https://codeforces.com/gym/358382/problem/I
Problem: GYM358382
**/
func (in *GYM358382) Run() {

	n := in.NextInt()
	min := 1000000
	minIndex := -1

	max := 0
	maxIndex := -1

	inversions := 0

	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = in.NextInt()
	}
	for i := 0; i < len(array); i++ {

		if i > 0 && array[i] < array[i-1] {
			inversions++
		}

		if array[i] < min {
			min = array[i]
			minIndex = i
		}

		if array[i] > max {
			max = array[i]
			maxIndex = i
		}
	}

	switch {
	case inversions == 0:
		fmt.Println(0)
	case inversions == 1 && minIndex-1 == maxIndex:
		fmt.Println(len(array) - minIndex)
	default:
		fmt.Println(-1)
	}
}

func NewGYM358382(r *bufio.Reader) *GYM358382 {
	return &GYM358382{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewGYM358382(bufio.NewReader(os.Stdin)).Run()
}
