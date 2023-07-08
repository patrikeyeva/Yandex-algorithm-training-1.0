//https://contest.yandex.ru/contest/27663/problems/J/

package main

import (
	"fmt"
	"math"
)

func intersection(area1, area2 *area) *area {
	upleft := chooseUp(area1.upleft, area2.upleft)
	upright := chooseUp(area1.upright, area2.upright)

	downleft := chooseDown(area1.downleft, area2.downleft)
	downright := chooseDown(area1.downright, area2.downright)

	return &area{upleft, upright, downleft, downright}
}

func chooseUp(p1, p2 int) int {
	if p1 < p2 {
		return p1
	}
	return p2
}

func chooseDown(p1, p2 int) int {
	if p1 > p2 {
		return p1
	}
	return p2
}

func expansion(newarea *area, t int) {
	newarea.upleft += t
	newarea.upright += t
	newarea.downleft -= t
	newarea.downright -= t
}

type area struct {
	upleft, upright, downleft, downright int
}

type point struct {
	x, y int
}

func main() {
	var t, d, n, xnew, ynew int
	fmt.Scan(&t, &d, &n)
	variaty1 := area{0, 0, 0, 0}

	for i := 0; i < n; i++ {
		expansion(&variaty1, t)
		fmt.Scan(&xnew, &ynew)
		variaty1 = *intersection(&variaty1, &area{d - xnew + ynew, d + xnew + ynew, -d + xnew + ynew, -d - xnew + ynew})
	}

	xmin := int(math.Ceil(float64(variaty1.downleft-variaty1.upleft) / 2.0))
	xmax := int(math.Floor(float64(variaty1.upright-variaty1.downright) / 2.0))
	xymin := float64(variaty1.downleft-variaty1.downright) / 2.0
	xymax := float64((variaty1.upright - variaty1.upleft)) / 2.0
	ymin := int(math.Ceil(xymin + float64(variaty1.downright)))
	ymax := int(math.Floor(xymax + float64(variaty1.upleft)))

	arr := make([]point, 0, (xmax-xmin+1)*(ymax-ymin+1))

	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			p1, p2 := y-x, y+x
			if p1 <= variaty1.upleft && p1 >= variaty1.downright && p2 <= variaty1.upright && p2 >= variaty1.downleft {
				arr = append(arr, point{x, y})
			}
		}
	}

	fmt.Println(len(arr))
	for _, p := range arr {
		fmt.Printf("%d %d\n", p.x, p.y)
	}
}
