package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("a.txt")
	sc := bufio.NewScanner(file)
	arr := make([]int64, 0)

	var sum int64
	m := make(map[int64]int)

	// parse input
	for sc.Scan() {
		i, _ := strconv.ParseInt(sc.Text(), 10, 64)
		arr = append(arr, i)
	}

	// loop until we see the same value twice
	for {
		for _, v := range arr {
			sum += v
			m[sum]++
			if m[sum] > 1 {
				fmt.Println(sum)
				return
			}
		}
	}

}
