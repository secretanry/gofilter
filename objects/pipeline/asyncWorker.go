package pipeline

import "gocv.io/x/gocv"

func BlockWork(frame gocv.Mat, pipe chan gocv.Mat, task func(frame gocv.Mat) gocv.Mat) {
	pipe <- task(frame)
}
