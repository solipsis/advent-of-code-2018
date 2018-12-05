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

	var sum int64

	// parse input
	for sc.Scan() {
		i, _ := strconv.ParseInt(sc.Text(), 10, 64)
		sum += i
	}
	fmt.Println(sum)
}
