package pipeline

import (
	"gocv.io/x/gocv"
)

type Pipeline struct {
	entry chan gocv.Mat
	in    chan gocv.Mat
	out   chan gocv.Mat
}

func New() *Pipeline {
	c := make(chan gocv.Mat)
	return &Pipeline{
		entry: c,
		in:    c,
	}
}

func (p *Pipeline) AddFilter(in chan gocv.Mat, out chan gocv.Mat) {
	go func(c chan gocv.Mat) {
		for f := range c {
			in <- f
		}
	}(p.in)
	p.in = out
	p.out = out
}

func (p *Pipeline) Process(frame gocv.Mat) gocv.Mat {
	p.entry <- frame
	return <-p.out
}
