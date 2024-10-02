package objects

import "gocv.io/x/gocv"

type DisplayData struct{}

func (ds *DisplayData) Display(window *gocv.Window, frame gocv.Mat) {
	window.IMShow(frame)
}
