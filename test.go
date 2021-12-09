package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)

	line, _ := sc.ReadString('\n')
	line = strings.TrimSpace(line[:len(line)-1])
	n, _ := strconv.Atoi(line)

	line, _ = sc.ReadString('\n')
	nums := strings.Split(line[:len(line)-1], " ")
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(strings.TrimSpace(nums[i]))
	}

	min := 1000000
	minIndex := -1

	max := 0
	maxIndex := -1

	inversions := 0

	for i := 0; i < len(array); i++ {

		if i > 0 && array[i] < array[i-1] {
			inversions++
		}

		if array[i] < min {
			min = array[i]
			minIndex = i
		}

		if array[i] > max {
			max = array[i]
			maxIndex = i
		}
	}

	switch {
	case inversions == 0:
		fmt.Println(0)
	case inversions == 1 && minIndex-1 == maxIndex:
		fmt.Println(len(array) - minIndex)
	default:
		fmt.Println(-1)
	}
}
