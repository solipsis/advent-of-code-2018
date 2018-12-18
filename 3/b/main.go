package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type entry struct {
	x, y int
	w, h int
}

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	input := make([]entry, 0)

	// parse
	for sc.Scan() {
		// Make parsing easier and skip to first @ sign
		replacer := strings.NewReplacer(",", " ", ":", "", "x", " ")
		in := replacer.Replace(sc.Text())
		in = in[strings.Index(in, "@")+1:]

		e := entry{}
		fmt.Sscanf(in, "%d %d %d %d", &e.x, &e.y, &e.w, &e.h)

		input = append(input, e)
	}

	// add squares to grid
	var grid [1000][1000]int
	for _, e := range input {
		for i := e.x; i < e.x+e.w; i++ {
			for j := e.y; j < e.y+e.h; j++ {
				grid[i][j]++
			}
		}
	}

	// count overlaps
	overlap := 0
	for _, row := range grid {
		for _, v := range row {
			if v > 1 {
				overlap++
			}
		}
	}
	fmt.Println(overlap)

	for n, e := range input {
		fail := false
		for i := e.x; i < e.x+e.w; i++ {
			for j := e.y; j < e.y+e.h; j++ {
				if grid[i][j] > 1 {
					fail = true
					break
				}
			}
		}
		if !fail {
			fmt.Println(n+1, e)
			break
		}
	}

}
