package objects

import (
	"image"

	"gocv.io/x/gocv"
)

func NewGrayscaleFilter() (chan gocv.Mat, chan gocv.Mat) {
	in, out := make(chan gocv.Mat), make(chan gocv.Mat)
	f := &GrayscaleFilter{
		in:  in,
		out: out,
	}
	go f.run()
	return in, out
}

type GrayscaleFilter struct {
	in, out chan gocv.Mat
}

func (f *GrayscaleFilter) run() {
	for frame := range f.in {
		result := gocv.NewMat()
		gocv.CvtColor(frame, &result, gocv.ColorBGRToGray)
		f.out <- result
	}
}

func NewMirrorFilter() (chan gocv.Mat, chan gocv.Mat) {
	in, out := make(chan gocv.Mat), make(chan gocv.Mat)
	f := &MirrorFilter{
		in:  in,
		out: out,
	}
	go f.run()
	return in, out
}

type MirrorFilter struct {
	in, out chan gocv.Mat
}

func (f *MirrorFilter) run() {
	for frame := range f.in {
		result := gocv.NewMat()
		gocv.Flip(frame, &result, 1)
		f.out <- result
	}
}

func NewResizeFilter(width, height int) (chan gocv.Mat, chan gocv.Mat) {
	in, out := make(chan gocv.Mat), make(chan gocv.Mat)
	f := &ResizeFilter{
		in:     in,
		out:    out,
		Width:  width,
		Height: height,
	}
	go f.run()
	return in, out
}

type ResizeFilter struct {
	in, out       chan gocv.Mat
	Width, Height int
}

func (f *ResizeFilter) run() {
	for frame := range f.in {
		result := gocv.NewMat()
		gocv.Resize(frame, &result, image.Pt(f.Width, f.Height), 0, 0, gocv.InterpolationLinear)
		f.out <- result
	}
}

func NewBlurFilter(size uint) (chan gocv.Mat, chan gocv.Mat) {
	in, out := make(chan gocv.Mat), make(chan gocv.Mat)
	f := &BlurFilter{
		in:  in,
		out: out,
	}
	go f.run()
	return in, out
}

type BlurFilter struct {
	in, out chan gocv.Mat
	Size    int
}

func (f *BlurFilter) run() {
	for frame := range f.in {
		result := gocv.NewMat()
		gocv.Blur(frame, &result, image.Point{
			X: f.Size,
			Y: f.Size,
		})
		f.out <- result
	}
}
