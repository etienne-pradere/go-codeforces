package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF431A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF431A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF431A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF431A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF431A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF431A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF431A
Date: 23/09/21
User: epradere
URL: http://codeforces.com/contest/431/problem/A
Problem: CF431A
**/
func (in *CF431A) Run() {
	arr := []int{
		in.NextInt(),
		in.NextInt(),
		in.NextInt(),
		in.NextInt(),
	}
	str := in.NextString()
	ans := 0
	for i := 0; i < len(str); i++ {
		ans += arr[str[i]-'0'-1]
	}
	fmt.Println(ans)
}

func NewCF431A(r *bufio.Reader) *CF431A {
	return &CF431A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF431A(bufio.NewReader(os.Stdin)).Run()
}
