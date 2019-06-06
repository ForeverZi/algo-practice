package main

import (
	"fmt"
	"strings"
	"strconv"
)

const MAX_INT = int(^uint(0) >> 1)

func main() {
	adjMatrix := [][]int{
		{0, 3, 2, 5},
		{2, 0, 7, 8},
		{MAX_INT, MAX_INT, 0, 2},
		{7, 9, MAX_INT, 0},
	}
	dest := make(map[int]int)
	prev := make(map[int]int)
	target := 3
	source := 0
	for i := 0; i < 4; i++ {
		dest[i] = MAX_INT
	}
	dest[source] = 0
	visited := make(map[int]bool)
	index := getMinDestIndex(dest, visited)
	for index >= 0 {
		weight := dest[index]
		for i, v := range dest {
			if adjMatrix[index][i] < MAX_INT && weight+adjMatrix[index][i] < v {
				dest[i] = weight + adjMatrix[index][i]
				prev[i] = index
			}
		}
		index = getMinDestIndex(dest, visited)
	}
	fmt.Printf("[%v]->[%v], min:[%v]\n", source, target, dest[target])
	printPath(prev, target)
}

func reverse(s []string){
	for from, to := 0, len(s)-1; from < to; from, to = from+1,to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

func printPath(prev map[int]int, target int){
	minPath := []string{strconv.Itoa(target)}
	for {
		prevIndex, ok := prev[target]
		if !ok {
			break
		}
		minPath = append(minPath, strconv.Itoa(prevIndex))
		target = prevIndex 
	}
	reverse(minPath)
	fmt.Println("path:", strings.Join(minPath, "->"))
}

func getMinDestIndex(dest map[int]int, visited map[int]bool) int {
	min := MAX_INT
	index := -1
	for i, v := range dest {
		if !visited[i] && v < min {
			min = v
			index = i
		}
	}
	visited[index] = true
	return index
}
