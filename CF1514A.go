package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1514A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1514A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1514A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1514A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1514A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1514A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1514A
Ejemplo de producto entre cuadrados perfectos
2*2 = 4
3*3 = 9
4*4 = 16

4   ,   9
(2*2)  (3*3)
 a b    a b

a*a = 6
b*b = 6
6*6 = 36

Date: 2/11/21
User: epradere
URL: https://codeforces.com/problemset/problem/1514/A
Problem: CF1514A
**/
func (in *CF1514A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		ws := false
		for i := 0; i < n; i++ {
			num := in.NextInt()
			root := int(math.Sqrt(float64(num)))
			if root*root == num {
				continue
			} else {
				ws = true

			}
		}
		if ws {
		 	fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}

func NewCF1514A(r *bufio.Reader) *CF1514A {
	return &CF1514A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1514A(bufio.NewReader(os.Stdin)).Run()
}
