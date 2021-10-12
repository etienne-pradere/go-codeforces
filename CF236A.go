package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF236A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF236A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF236A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF236A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF236A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF236A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF236A
Date: 22/09/21
User: epradere
URL: http://codeforces.com/contest/236/problem/A
Problem: CF236A
**/
func (in *CF236A) Run() {

	name := in.NextString()

	letters := make([]int, 30)

	ans := 0
	for i := 0; i < len(name); i++ {
		if letters[name[i]-'a'] == 0 {
			ans++
		}
		letters[name[i]-'a']++
	}

	if ans%2!=0 {
		fmt.Println("IGNORE HIM!")
	} else {
		fmt.Println("CHAT WITH HER!")
	}

}

func NewCF236A(r *bufio.Reader) *CF236A {
	return &CF236A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF236A(bufio.NewReader(os.Stdin)).Run()
}
