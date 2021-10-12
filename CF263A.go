package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF263A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF263A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF263A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF263A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF263A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF263A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF263A
Date: 22/09/21
User: epradere
URL: http://codeforces.com/contest/263/problem/A
Problem: CF263A
**/
func (in *CF263A) Run() {
	posI, posJ := 0, 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			data := in.NextInt()
			if data == 1 {
				posI = i
				posJ = j
			}
		}
	}
	diff := int(math.Abs(float64(posI-2)) + math.Abs(float64(posJ-2)))
	fmt.Println(diff)
}

func NewCF263A(r *bufio.Reader) *CF263A {
	return &CF263A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF263A(bufio.NewReader(os.Stdin)).Run()
}
