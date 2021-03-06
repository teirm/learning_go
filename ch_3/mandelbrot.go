// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
        dx, dy                 = 0.25, 0.25
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y0 := float64(py)/height*(ymax-ymin) + ymin
        y1 := (float64(py)+dy)/height*(ymax-ymin) + ymin
        y2 := (float64(py)+2*dy)/height*(ymax-ymin) + ymin
        y3 := (float64(py)+3*dy)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := float64(px)/width*(xmax-xmin) + xmin
			x1 := (float64(px)+dx)/width*(xmax-xmin) + xmin
			x2 := (float64(px)+2*dx)/width*(xmax-xmin) + xmin
			x3 := (float64(px)+3*dx)/width*(xmax-xmin) + xmin
			z0 := complex(x0, y0)
			z1 := complex(x1, y1)
			z2 := complex(x2, y2)
			z3 := complex(x3, y3)

            z := (z0 + z1 + z2 + z3) / 4 

            // Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // Note: Ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
