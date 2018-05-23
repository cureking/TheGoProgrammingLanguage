package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 1400, 720
	cells         = 150
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * .4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	//一定要注意上面凭借字体末尾处的空格。切记切记。

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon style='fill: rgba(%d,%d,50,0.7);' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", int(cHeight(i, j)*128+128), 255-int(cHeight(i, j)*128+128), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}
func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) / 2

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	//写出需要的三维体的z计算函数，返回即可。
}

func cHeight(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) / 2
	return z
}
