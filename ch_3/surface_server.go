// Sufrace compates an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
    "io"
    "log"
	"math"
    "net/http"
    "strings"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

var functionMap = map[string]func(x,y float64) float64 {
    "sin" : sinGraph,
    "hill": hillGraph,
    "saddle": saddleGraph, 
}

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "image/svg+xml")
        generateSvg(w, r)
    }
    http.HandleFunc("/", handler) 
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generateSvg(out io.Writer, r *http.Request) {
    
    if err := r.ParseForm(); err != nil {
        log.Fatal(err)
    }

    width := strings.Join(r.Form["width"], "")
    height := strings.Join(r.Form["height"], "")
    color := strings.Join(r.Form["color"], "")

	fmt.Fprintf(out, "<svg xmlns='http:///www.w3.org/2000/svg' "+
        "xmlns:xlink='http://www.w3.org/1999/xlink' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%s' height='%s'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if ok != true {
				continue
			}

			bx, by, ok := corner(i, j)
			if ok != true {
				continue
			}

			cx, cy, ok := corner(i, j+1)
			if ok != true {
				continue
			}

			dx, dy, ok := corner(i+1, j+1)
			if ok != true {
				continue
			}

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
    // Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute the surface height z
	z := sinGraph(x, y)
	if math.IsNaN(z) {
		return 0.0, 0.0, false
	}

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func sinGraph(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func hillGraph(x, y float64) float64 {
    denominator := math.Pow(x, 2) + math.Pow(y,2) + 1
    return (-4 * y) / denominator
}

func saddleGraph(x, y float64) float64 {
    return (math.Pow(x, 2) - math.Pow(y, 2))
}
