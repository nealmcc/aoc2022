package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/nealmcc/aoc2022/pkg/rope"
	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

type teeLogger []rope.Logger

func tee(logs ...rope.Logger) teeLogger {
	return teeLogger(logs)
}

var _ rope.Logger = teeLogger{}

// Log implements rope.Logger.
func (t teeLogger) Log(knots []v.Point) {
	for _, l := range t {
		l.Log(knots)
	}
}

// tracer is a rope logger that draws images
type tracer struct {
	knotRadius int
	frames     []*image.RGBA
	prefix     string
	numSaved   int
	count      int
}

func newTracer(prefix string) *tracer {
	return &tracer{
		knotRadius: 3,
		prefix:     prefix,
		frames:     make([]*image.RGBA, 0, 24),
	}
}

// Log implements rope.Logger.
func (t *tracer) Log(knots []v.Point) {
	(*t).count = t.count + 1
	if t.count > 1200 {
		// only write the first 1200 frames
		return
	}
	m := image.NewRGBA(image.Rect(-600, -600, 600, 600))
	draw.Draw(m, m.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)
	for _, knot := range knots {
		knot = knot.Times(3 * t.knotRadius)
		r := image.Rect(knot.X-t.knotRadius, knot.Y-t.knotRadius, knot.X+t.knotRadius, knot.Y+t.knotRadius)
		blue := color.NRGBA{0, 0, 255, 126}
		draw.Draw(m, r, &image.Uniform{C: blue}, image.Point{}, draw.Over)
	}
	t.frames = append(t.frames, m)
	if len(t.frames) >= 24 {
		t.Save()
	}
}

// Save the tracer's current image to the given output.
func (t *tracer) Save() error {
	for _, img := range t.frames {
		if err := save(fmt.Sprintf("%s_%04d.png", t.prefix, t.numSaved), img); err != nil {
			return err
		}
		t.numSaved++
	}
	(*t).frames = t.frames[:0]
	return nil
}

func save(filename string, img *image.RGBA) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}
