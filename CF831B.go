package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF831B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF831B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF831B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF831B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF831B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF831B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF831B
Date: 19/02/22
User: epradere
URL: https://codeforces.com/problemset/problem/831/B
Problem: CF831B
**/
func (in *CF831B) Run() {
	a := in.NextString()
	b := in.NextString()
	mapa := make(map[uint8]uint8)

	for i := 0; i < len(a); i++ {
		mapa[a[i]] = b[i]
		mapa[strings.ToUpper(string(a[i]))[0]] = strings.ToUpper(string(b[i]))[0]
	}

	r := in.NextString()

	var ans strings.Builder

	for i := 0; i < len(r); i++ {
		if r[i] >= '0' && r[i] <= '9' {
			ans.WriteString(string(r[i]))
		} else {
			ans.WriteString(string(rune(mapa[r[i]])))
		}

	}
	fmt.Println(ans.String())
}

func NewCF831B(r *bufio.Reader) *CF831B {
	return &CF831B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF831B(bufio.NewReader(os.Stdin)).Run()
}
