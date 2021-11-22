package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1512B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1512B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1512B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1512B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1512B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1512B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1512B
Date: 21/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1512/B
Problem: CF1512B
**/
func (in *CF1512B) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		arr := make([][]rune, n)
		pos := make([][]int, 0)
		for i := 0; i < n; i++ {
			str := in.NextString()
			line := make([]rune, n)
			for j := 0; j < n; j++ {
				line[j] = rune(str[j])
				if line[j] == '*' {
					pos = append(pos, []int{i, j})
				}
			}
			arr[i] = line
		}

		x1, y1 := pos[0][0], pos[0][1]
		x2, y2 := pos[1][0], pos[1][1]

		if x1 == x2 {
			if x1 == 0 {
				arr[x1+1][y1] = '*'
				arr[x2+1][y2] = '*'
			} else {
				arr[x1-1][y1] = '*'
				arr[x2-1][y2] = '*'
			}
		} else if y1 == y2 {
			if y1 == 0 {
				arr[x1][y1+1] = '*'
				arr[x2][y2+1] = '*'
			} else {
				arr[x1][y1-1] = '*'
				arr[x2][y2-1] = '*'
			}
		} else {
			arr[x1][y2] = '*'
			arr[x2][y1] = '*'
		}

		var sb strings.Builder

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				sb.WriteRune(arr[i][j])
			}
			sb.WriteString("\n")
		}
		fmt.Print(sb.String())
	}
}

func NewCF1512B(r *bufio.Reader) *CF1512B {
	return &CF1512B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1512B(bufio.NewReader(os.Stdin)).Run()
}
