package pipeline

import "gocv.io/x/gocv"

func WorkAsync(frame gocv.Mat, pipe chan gocv.Mat, task func(frame gocv.Mat) gocv.Mat) {
	resChan := make(chan gocv.Mat)
	go func() {
		result := task(frame)
		resChan <- result
	}()
loop:
	for {
		select {
		case msg := <-resChan:
			pipe <- msg
			break loop
		}
	}
}
