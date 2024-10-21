package modules

import (
	"fmt"
	"gofilter/objects"
	"gofilter/objects/pipeline"

	"gocv.io/x/gocv"
)

func initPipeline() *pipeline.Pipeline {
	pipe := pipeline.New()
	pipe.AddFilter(objects.NewGrayscaleFilter())
	pipe.AddFilter(objects.NewMirrorFilter())
	pipe.AddFilter(objects.NewResizeFilter(400, 300))
	pipe.AddFilter(objects.NewBlurFilter(10))
	return pipe
}

func StartApp() {
	sourceData := objects.NewVideoSource(0)
	if sourceData == nil {
		return
	}
	defer sourceData.Release()

	sourceWindow := gocv.NewWindow("Initial Video")
	defer sourceWindow.Close()
	sinkWindow := gocv.NewWindow("Processed Video")
	defer sinkWindow.Close()

	frame := gocv.NewMat()
	defer frame.Close()

	source := &objects.DisplayData{}
	sink := &objects.DisplayData{}

	pipeline := initPipeline()

	for {
		if ok := sourceData.ReadFrame(&frame); !ok {
			fmt.Println("Unable to read frame from video sourceData")
			break
		}

		if frame.Empty() {
			continue
		}

		source.Display(sourceWindow, frame)

		sink.Display(sinkWindow, pipeline.Process(frame))

		if sinkWindow.WaitKey(1) == 'q' {
			break
		}

		if sourceWindow.WaitKey(1) == 'q' {
			break
		}
	}
}
