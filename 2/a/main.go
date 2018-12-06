package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("a.txt")
	sc := bufio.NewScanner(file)

	// count number of times characters appear 2 or three times
	c2 := 0
	c3 := 0
	for sc.Scan() {
		m := make(map[rune]int)
		for _, r := range sc.Text() {
			m[r]++
		}
		var two, three bool
		for _, v := range m {
			if v == 2 {
				two = true
			}
			if v == 3 {
				three = true
			}
		}
		if two {
			c2++
		}
		if three {
			c3++
		}

	}
	fmt.Println(c2 * c3)
}
