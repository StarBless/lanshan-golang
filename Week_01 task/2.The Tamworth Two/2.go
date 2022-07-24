package main

import "fmt"

var (
	Map [10][10]byte
	x1  int
	x2  int
	y1  int
	y2  int
)
var dir = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func moni() {
	ren := 0
	niu := 0
	time := 0
	for {
		if x1 == x2 && y1 == y2 {
			fmt.Printf("%d", time)
			break
		}
		time += 1
		if time >= 99999 {
			fmt.Printf("0")
			break
		}
		ny1 := y1 + dir[ren][0]
		nx1 := x1 + dir[ren][1]
		ny2 := y2 + dir[niu][0]
		nx2 := x2 + dir[niu][1]
		if ny1 < 0 || ny1 > 9 || nx1 < 0 || nx1 > 9 || Map[ny1][nx1] == '*' {
			nx1 = x1
			ny1 = y1
			ren = (ren + 1) % 4
		}
		if ny2 < 0 || ny2 > 9 || nx2 < 0 || nx2 > 9 || Map[ny2][nx2] == '*' {
			nx2 = x2
			ny2 = y2
			niu = (niu + 1) % 4
		}
		x1 = nx1
		y1 = ny1
		x2 = nx2
		y2 = ny2
	}
}

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Scanf("%c", &Map[i][j])
			if Map[i][j] == 'F' {
				y1 = i
				x1 = j
			}
			if Map[i][j] == 'C' {
				y2 = i
				x2 = j
			}
		}
		fmt.Scanf("\n")
	}
	moni()
}
