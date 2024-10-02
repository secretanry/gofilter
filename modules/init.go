package modules

import (
	"fmt"
	"gocv.io/x/gocv"
	"gofilter/objects"
	"gofilter/objects/pipeline"
)

func initPipeline() *pipeline.Pipeline {
	pipelineInstance := &pipeline.Pipeline{}
	pipelineInstance.AddFilter(&objects.GrayscaleFilter{})
	pipelineInstance.AddFilter(&objects.MirrorFilter{})
	pipelineInstance.AddFilter(&objects.ResizeFilter{Width: 400, Height: 300})
	pipelineInstance.AddFilter(&objects.BlurFilter{Size: 10})
	return pipelineInstance
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

		resultChan := make(chan gocv.Mat)
		go func() {
			processedFrame := pipeline.Process(frame)
			resultChan <- processedFrame
		}()
		sink.Display(sinkWindow, <-resultChan)

		if sinkWindow.WaitKey(1) == 'q' {
			break
		}

		if sourceWindow.WaitKey(1) == 'q' {
			break
		}
	}
}
