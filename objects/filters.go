package objects

import (
	"gocv.io/x/gocv"
	"image"
)

type Filter interface {
	Apply(frame gocv.Mat) gocv.Mat
}

type GrayscaleFilter struct{}

func (f *GrayscaleFilter) Apply(frame gocv.Mat) gocv.Mat {
	result := gocv.NewMat()
	gocv.CvtColor(frame, &result, gocv.ColorBGRToGray)
	return result
}

type MirrorFilter struct{}

func (f *MirrorFilter) Apply(frame gocv.Mat) gocv.Mat {
	result := gocv.NewMat()
	gocv.Flip(frame, &result, 1)
	return result
}

type ResizeFilter struct {
	Width, Height int
}

func (f *ResizeFilter) Apply(frame gocv.Mat) gocv.Mat {
	result := gocv.NewMat()
	gocv.Resize(frame, &result, image.Pt(f.Width, f.Height), 0, 0, gocv.InterpolationLinear)
	return result
}

type BlurFilter struct {
	Size int
}

func (f *BlurFilter) Apply(frame gocv.Mat) gocv.Mat {
	result := gocv.NewMat()
	gocv.Blur(frame, &result, image.Point{
		X: f.Size,
		Y: f.Size,
	})
	return result
}
