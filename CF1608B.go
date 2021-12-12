package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1608B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1608B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1608B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1608B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1608B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1608B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1608B
Date: 11/12/21
User: epradere
URL: https://codeforces.com/contest/1608/problem/B
Problem: CF1608B
**/
func (in *CF1608B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, a, b := in.NextInt(), in.NextInt(), in.NextInt()
		diff := int(math.Abs(float64(a - b)))
		if diff >= 2 || a+b >= n {
			fmt.Println(-1)
		} else {
			arr := make([]int, n)
			if a > b {
				i, j, p := 0, 1, n
				for ; i < a+b && i < n; i++ {
					if i%2 == 0 {
						arr[i] = j
						j++
					} else {
						arr[i] = p
						p--
					}
				}
				for ; i < n; i++ {
					arr[i] = p
					p--
				}
			} else if a < b {
				i, j, p := 0, 1, n
				for ; i <= a+b && i < n; i++ {
					if i%2 != 0 {
						arr[i] = j
						j++
					} else {
						arr[i] = p
						p--
					}
				}
				for ; i < n; i++ {
					arr[i] = j
					j++
				}
			} else {
				i, j, p := 0, 1, n
				for ; i <= a+b && i < n; i++ {
					if i%2 != 0 {
						arr[i] = j
						j++
					} else {
						arr[i] = p
						p--
					}
				}
				for ; i < n; i++ {
					arr[i] = p
					p--
				}
			}

			max, min := 0, 0
			for i := 1; i < n-1; i++ {
				if arr[i-1] < arr[i] && arr[i+1] < arr[i] {
					max++
				} else if arr[i-1] > arr[i] && arr[i+1] > arr[i] {
					min++
				}
			}
			if max == a && min == b {
				for i := 0; i < n; i++ {
					fmt.Print(arr[i], " ")
				}
				fmt.Println()
			} else {
				fmt.Println(-1)
			}

		}

	}
}

func NewCF1608B(r *bufio.Reader) *CF1608B {
	return &CF1608B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1608B(bufio.NewReader(os.Stdin)).Run()
}
