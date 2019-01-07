package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type block struct {
	startH, startM int
	endH, endM     int
	num            int
}

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	input := make([]string, 0)
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	//sort
	sort.Slice(input, func(i, j int) bool {
		ti := strings.Split(input[i], " ")[0]
		tj := strings.Split(input[j], " ")[0]
		ti2 := strings.Split(input[i], " ")[1]
		tj2 := strings.Split(input[j], " ")[1]
		if ti == tj {
			return ti2 < tj2
		}
		return ti < tj
	})

	//guard starts #determins guard num
	// guard falls asleep #deturmins start time
	// while guard asleep increment
	// guard wakes up

	guards := make(map[int][]int, 0)

	guardNum := -1
	start := -1
	var err error
	// TODO: cleaner parsing
	for _, v := range input {
		v = strings.Replace(v, "[", "", -1)
		v = strings.Replace(v, "]", "", -1)
		line := strings.Split(v, " ")
		if strings.Contains(v, "Guard") {
			guardNum, _ = strconv.Atoi(line[3][1:])
			continue
		}
		if strings.Contains(v, "asleep") {
			start, err = strconv.Atoi(strings.Split(line[1], ":")[1])
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		if strings.Contains(v, "wakes") {
			end, _ := strconv.Atoi(strings.Split(line[1], ":")[1])
			if guards[guardNum] == nil {
				guards[guardNum] = make([]int, 60)
			}
			for start < end {
				guards[guardNum][start]++
				start++
			}
			continue
		}
		panic("invalid line")
	}

	max := 0
	maxID := -1
	for id, g := range guards {
		sum := 0
		for _, v := range g {
			sum += v
		}
		if sum > max {
			max = sum
			maxID = id
		}
	}

	lm := 0
	lmID := -1
	for n, v := range guards[maxID] {
		if v > lm {
			lm = v
			lmID = n
		}
	}

	fmt.Println(maxID * lmID)
}
