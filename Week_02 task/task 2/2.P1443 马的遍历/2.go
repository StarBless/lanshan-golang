package main

import "fmt"

var Map [450][450]int
var vis [450][450]int
var n, m, x, y int
var dir = [][]int{
	{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

func main() {
	fmt.Scanf("%d %d %d %d\n", &n, &m, &x, &y)
}
