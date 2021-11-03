package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF691B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF691B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF691B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF691B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF691B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF691B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF691B
Date: 26/10/21
User: epradere
URL: https://codeforces.com/problemset/problem/691/B
Problem: CF691B
**/
func (in *CF691B) Run() {
	sym := "AHIMOoTUVvWwXxY"
	m := map[string]string{
		"p": "q",
		"q": "p",
		"b": "d",
		"d": "b",
	}
	str := in.NextString()
	ws := true
	for i := 0; i <= len(str)/2; i++ {
		left := str[i : i+1]
		right := str[len(str)-i-1 : len(str)-i]
		op, exist := m[left]
		if exist && op != right {
			ws = false
			break
		} else if exist && op == right {
			continue
		} else if !strings.Contains(sym, left) {
			ws = false
			break
		} else if strings.Contains(sym, left) && str[i] != str[len(str)-i-1] {
			ws = false
			break
		}
	}
	if ws {
		fmt.Println("TAK")
	} else {
		fmt.Println("NIE")
	}
}

func NewCF691B(r *bufio.Reader) *CF691B {
	return &CF691B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF691B(bufio.NewReader(os.Stdin)).Run()
}
