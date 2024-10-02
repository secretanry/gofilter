package objects

import (
	"fmt"
	"gocv.io/x/gocv"
)

type VideoSource struct {
	webcam *gocv.VideoCapture
}

func NewVideoSource(source int) *VideoSource {
	webcam, err := gocv.OpenVideoCapture(source)
	if err != nil {
		fmt.Println("Error opening video capture device")
		return nil
	}
	return &VideoSource{webcam: webcam}
}

func (vs *VideoSource) ReadFrame(frame *gocv.Mat) bool {
	return vs.webcam.Read(frame)
}

func (vs *VideoSource) Release() {
	vs.webcam.Close()
}
