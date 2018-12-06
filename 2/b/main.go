package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("a.txt")
	sc := bufio.NewScanner(file)

	arr := make([]string, 0)
	for sc.Scan() {
		arr = append(arr, sc.Text())

	}

	// find ids that differ by one character
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			diffcount := 0
			diffIndex := 0
			for x := 0; x < len(arr[i]); x++ {
				if arr[i][x] != arr[j][x] {
					diffIndex = x
					diffcount++
				}
			}
			if diffcount == 1 {
				fmt.Println(arr[i][:diffIndex] + arr[i][diffIndex+1:])
			}
		}
	}
}
