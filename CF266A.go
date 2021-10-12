package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF266A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF266A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF266A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF266A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF266A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF266A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF266A
Date: 22/09/21
User: epradere
URL: http://codeforces.com/contest/266/problem/A
Problem: CF266A
**/
func (in *CF266A) Run() {
	in.NextInt()
	str := in.NextString()

	init := str[0]
	ans := 0
	for i := 1; i < len(str); i++ {
		if init == str[i] {
			ans++
		} else {
			init = str[i]
		}
	}
	fmt.Println(ans)

}

func NewCF266A(r *bufio.Reader) *CF266A {
	return &CF266A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF266A(bufio.NewReader(os.Stdin)).Run()
}
