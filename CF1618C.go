package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1618C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1618C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1618C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1618C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1618C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1618C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
WA
Run solve the problem CF1618C
Date: 15/12/21
User: epradere
URL: https://codeforces.com/contest/1618/problem/C
Problem: CF1618C
**/
func (in *CF1618C) Run() {
	fmt.Println(int(math.Sqrt(float64(1000000000000000000))))
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int64, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt64()
		}
		ans := int64(-1)
		for i := 0; i < n && ans == -1; i++ {
			all := true
			for j := 0; j < n; j++ {
				//fmt.Println(i, ",", j, " j;", j%2 == 0, " i:", i%2 == 0, " mod:", arr[j]%arr[i])
				if j%2 == 0 {
					if i%2 == 0 && arr[j]%arr[i] == 0 {
						continue
					} else if i%2 != 0 && arr[j]%arr[i] != 0 {
						continue
					} else {
						all = false
					}
				} else {
					if i%2 == 0 && arr[j]%arr[i] != 0 {
						continue
					} else if i%2 != 0 && arr[j]%arr[i] == 0 {
						continue
					} else {
						all = false
					}
				}
			}
			//fmt.Println(all)
			if all {
				ans = arr[i]
			}
		}
		if ans == -1 {
			fmt.Println(0)
		} else {
			fmt.Println(ans)
		}
	}
}

func NewCF1618C(r *bufio.Reader) *CF1618C {
	return &CF1618C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1618C(bufio.NewReader(os.Stdin)).Run()
}
