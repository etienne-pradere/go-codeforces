package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1450A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1450A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1450A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1450A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1450A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1450A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1450A
Date: 8/12/21
User: epradere
URL: https://codeforces.com/problemset/problem/1450/A
Problem: CF1450A
**/
func (in *CF1450A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := make(map[string]int)
		str := in.NextString()
		for i := 0; i < n; i++ {
			val := mapa[string(str[i])]
			mapa[string(str[i])] = val + 1

		}
		nStr := "bugytr"
		var bf strings.Builder

		for i := 0; i < len(nStr); i++ {
			cant := mapa[string(nStr[i])]
			for j := 0; j < cant; j++ {
				bf.WriteString(string(nStr[i]))
			}
			mapa[string(nStr[i])] = 0
		}

		for k, v := range mapa {
			for i := 0; i < v; i++ {
				bf.WriteString(k)
			}
		}
		fmt.Println(bf.String())

	}
}

func NewCF1450A(r *bufio.Reader) *CF1450A {
	return &CF1450A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1450A(bufio.NewReader(os.Stdin)).Run()
}
