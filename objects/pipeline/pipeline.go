package pipeline

import (
	"gocv.io/x/gocv"
	"gofilter/objects"
)

// Pipeline to sequentially pass frames through filters
type Pipeline struct {
	filters []objects.Filter
	pipes   []chan gocv.Mat
}

func (p *Pipeline) AddFilter(filter objects.Filter) {
	p.filters = append(p.filters, filter)
	p.pipes = append(p.pipes, make(chan gocv.Mat))
}

func (p *Pipeline) Process(frame gocv.Mat) gocv.Mat {
	p.pipes = append(p.pipes, make(chan gocv.Mat))
	for i, filter := range p.filters {
		go func() {
			WorkAsync(<-p.pipes[i], p.pipes[i+1], filter.Apply)
		}()
	}
	p.pipes[0] <- frame
	result := <-p.pipes[len(p.filters)]
	p.pipes = p.pipes[:len(p.filters)]
	return result
}
