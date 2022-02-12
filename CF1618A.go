package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1618A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1618A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1618A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1618A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1618A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1618A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1618A
Date: 15/12/21
User: epradere
URL: https://codeforces.com/contest/1618/problem/A
Problem: CF1618A
**/
func (in *CF1618A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		arr := make([]int, 7)
		mapa := make(map[int]int, 0)
		for i := 0; i < 7; i++ {
			arr[i] = in.NextInt()
			mapa[arr[i]] = 1

		}

		sort.Ints(arr)
		var sb strings.Builder
		ans := false
		for i := 0; i < 7 && !ans; i++ {
			for j := 0; j < 7 && !ans; j++ {
				if i != j {
					for p := 0; p < 7 && !ans; p++ {
						if p != i && p != j {
							if arr[i]+arr[j]+arr[p] == arr[6] {
								sb.WriteString(strconv.Itoa(arr[i]) + " " + strconv.Itoa(arr[j]) + " " + strconv.Itoa(arr[p]))
								ans = true
							}
						}
					}
				}
			}
		}
		fmt.Println(sb.String())
	}
}

func NewCF1618A(r *bufio.Reader) *CF1618A {
	return &CF1618A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1618A(bufio.NewReader(os.Stdin)).Run()
}
