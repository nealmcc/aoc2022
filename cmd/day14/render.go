package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

type RenderFunc func(points ...v.Point) error

type Renderer struct {
	cave     *Cavern
	min, max v.Point
	scale    int
	prefix   string
	frame    int
}

func NewRenderer(cave *Cavern, prefix string, min, max v.Point, scale int) *Renderer {
	return &Renderer{
		cave:   cave,
		min:    min,
		max:    max,
		prefix: prefix,
		scale:  scale,
	}
}

func (r *Renderer) SaveNext(points ...v.Point) error {
	if r.frame == 0 {
		points = nil
	}
	r.frame++

	file, err := os.Create(r.prefix + fmt.Sprintf("_%05d", r.frame) + ".png")
	if err != nil {
		return nil
	}
	defer file.Close()

	img := r.cave.Render(r.min, r.max, r.scale, points...)
	return png.Encode(file, img)
}

// Render renders the portion of this cavern as an image.
// The given bounding box will limit the portion of the cavern that is drawn.
// The bounding box combined with the scale will determine the size
// of the output image, in pixels.
//
// If any (optional) points are given, then only those squares will be rendered,
// and the rest will be transparent.
func (c *Cavern) Render(min, max v.Point, scale int, points ...v.Point) *image.RGBA {
	size := max.Sub(min).Times(scale)
	m := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	// project the given point in the cave onto the image.
	project := func(pcave v.Point) image.Rectangle {
		offset := min.Times(-1)
		p := pcave.Add(offset).Times(scale)
		x := p.X
		y := p.Y

		return image.Rect(x+1, y+1, x+scale-1, y+scale-1)
	}

	drawSquare := func(p v.Point, mat Material) {
		var src *image.Uniform
		switch mat {
		case Sand:
			src = &image.Uniform{color.RGBA{238, 206, 137, 255}}
		case Rock:
			src = &image.Uniform{color.RGBA{122, 111, 118, 255}}
		default:
			return
		}
		draw.Draw(m, project(p), src, image.Point{}, draw.Src)
	}

	if len(points) == 0 {
		background := color.RGBA{225, 226, 228, 255}
		draw.Draw(m, m.Bounds(), &image.Uniform{background}, image.Point{}, draw.Src)
		for x, col := range c.grid {
			for y, mat := range col {
				drawSquare(v.Point{X: x, Y: y}, mat)
			}
		}
		return m
	}

	draw.Draw(m, m.Bounds(), image.Transparent, image.Point{}, draw.Src)
	for _, p := range points {
		if p.X < min.X || p.X >= max.X || p.Y < min.Y || p.Y >= max.Y {
			continue
		}
		mat, ok := c.Get(p)
		if !ok {
			continue
		}
		drawSquare(p, mat)
	}

	return m
}

// Text renders the portion of this cavern as a string.
// All squares with x1 <= X < x2, y1 <= Y < y2 will be rendered.
func (c *Cavern) Text(x1, y1, x2, y2 int) string {
	height := y2 - y1
	width := x2 - x1
	rows := make([][]byte, height)
	for r := 0; r < height; r++ {
		row := make([]byte, width)
		for k := 0; k < width; k++ {
			switch mat := c.grid[x1+k][y1+r]; mat {
			case Air:
				row[k] = '.'
			case Sand:
				row[k] = 'o'
			case Rock:
				row[k] = '#'
			}
		}
		rows[r] = row
	}
	return string(bytes.Join(rows, []byte{'\n'}))
}
