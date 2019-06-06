package main

import (
	"fmt"
	"strings"
)

const (
	NORMAL = iota
	BLOCK
)

var mapData = [][]uint8{
	{0, 0, 0, 1, 0, 0},
	{0, 0, 0, 1, 0, 0},
	{0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 1},
}

type Point struct{
	X	int
	Y 	int
}

type Record struct{
	GN	int
	HN	int
}

func main(){
	astar()
}

func astar(){
	source := Point{
		X: 0,
		Y: 0,
	}
	target := Point{
		X: 5,
		Y: 0,
	}
	open := make(map[Point]Record)
	closed := make(map[Point]Record)
	prev := make(map[Point]Point)
	open[source] = Record{
		GN: 0,
		HN: ManhattanDist(source, target),
	}
	p, r := ExtractMinFN(open)
	for p != nil {
		if *p == target {
			break
		}
		// fmt.Printf("enque close:[%v,%v]\n", p.X, p.Y)
		closed[*p] = r
		arounds := getAroundNormalPoint(mapData, *p)
		for _, ap := range arounds {
			if _,ok := closed[ap]; ok {
				continue
			}
			if ar,ok := open[ap]; ok {
				if ar.GN > r.GN+1 {
					ar.GN = r.GN + 1
					open[ap] = ar
					prev[ap] = *p
				}
			}else{
				open[ap] = Record{
					GN: r.GN + 1,
					HN: ManhattanDist(ap, target),
				}
				prev[ap] = *p
			}
		}
		p, r = ExtractMinFN(open)
	}
	pathPoints := []Point{target}
	for prevPoint, ok := prev[target]; ok; prevPoint, ok = prev[prevPoint] {
		pathPoints = append(pathPoints, prevPoint)
	}
	pathStrs := make([]string, len(pathPoints))
	for i, pathPoint := range pathPoints {
		pathStrs[len(pathStrs)-1-i] = fmt.Sprintf("[%v,%v]", pathPoint.X, pathPoint.Y)
	}
	fmt.Println(strings.Join(pathStrs, "->"))
}

func getAroundNormalPoint(m [][]uint8, p Point)[]Point{
	height := len(m)
	width := len(m[0])
	points := make([]Point, 0, 4)
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			if dy==0 && dx==0 || dx!=0 && dy!=0 {
				continue
			}
			if p.X+dx >= 0 && p.Y+dy >= 0 && p.X+dx < width && p.Y + dy < height && m[p.Y+dy][p.X+dx]==NORMAL {
				points = append(points, Point{
					X: p.X + dx,
					Y: p.Y + dy,
				})
			}
		}
	}
	return points
}

func ExtractMinFN(set map[Point]Record)(*Point, Record){
	min := int(^uint(0)>>1)
	var minPoint Point
	var record Record
	var hasPoint bool
	for p, r := range set {
		if r.GN + r.HN < min {
			min = r.GN + r.HN
			minPoint = p
			hasPoint = true
		}
	}
	if hasPoint {
		record = set[minPoint]
		delete(set, minPoint)
	}
	return &minPoint, record
}

func ManhattanDist(a, b Point)int{
	dx := b.X - a.X
	dy := b.Y - a.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx+dy
}