package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 // 画布大小(以像素为单位)
	cells = 100 // 网格细胞的数量
	xyrange = 30.0 // 轴范围
	syscale = width / 2 /xyrange // pixels per x or y unit
	zscale = height * 0.4 // pixels per  unit
	angle = math.Pi / 6 // x和y轴的夹角
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>",width, height )
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				  ax, ay, bx, by, cx, cy, dx, dy)
		}

	}
	fmt.Println("</svg>")
}

// corner ...
func corner(i, j, int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x,y)

	sx := width / 2 + (x - y ) * cos30 * syscale
	sy := height / 2 + (x + y) * sin30 * syscale - z * zscale
	return sx, sy
}

// f ...
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}
