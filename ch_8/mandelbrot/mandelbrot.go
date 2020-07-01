// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

// package for all image fields
type MandelbrotData struct {
	x     int
	y     int
	color color.Color
}

// struct for worker data
type WorkerData struct {
	startX int
	endX   int
	startY int
	endY   int
}

// divide work breaks up the image
// into "worker" number of stripes
// each of height "height"
func divideWork(width int, height int, workers int) []WorkerData {

	var result []WorkerData

	startX := 0
	dX := width / workers
	for i := 0; i < workers; i++ {
		result = append(result, WorkerData{startX, startX + dX, 0, height})
		startX += dX
	}
	return result
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		dx, dy                 = 0.25, 0.25
	)

	var worker_count = flag.Int("workers", 4, "Number of goroutines to create")
	flag.Parse()

	workers := divideWork(width, height, *worker_count)

	ch := make(chan MandelbrotData)
	var wg sync.WaitGroup
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, worker := range workers {
		wg.Add(1)
		go func(worker WorkerData) {
			defer wg.Done()
			for py := worker.startY; py < worker.endY; py++ {
				y0 := float64(py)/height*(ymax-ymin) + ymin
				y1 := (float64(py)+dy)/height*(ymax-ymin) + ymin
				y2 := (float64(py)+2*dy)/height*(ymax-ymin) + ymin
				y3 := (float64(py)+3*dy)/height*(ymax-ymin) + ymin
				for px := worker.startX; px < worker.endX; px++ {
					x0 := float64(px)/width*(xmax-xmin) + xmin
					x1 := (float64(px)+dx)/width*(xmax-xmin) + xmin
					x2 := (float64(px)+2*dx)/width*(xmax-xmin) + xmin
					x3 := (float64(px)+3*dx)/width*(xmax-xmin) + xmin
					z0 := complex(x0, y0)
					z1 := complex(x1, y1)
					z2 := complex(x2, y2)
					z3 := complex(x3, y3)

					z := (z0 + z1 + z2 + z3) / 4

					ch <- MandelbrotData{px, py, mandelbrot(z)}
				}
			}
		}(worker)
	}

	// closer
	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		// Image point (px, py) represents complex value z.
		img.Set(data.x, data.y, data.color)
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
