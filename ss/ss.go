package ss

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"os"
)

func Screenshot(x0,y0,x1,y1 int) (*image.RGBA, error) {
	img, err := screenshot.Capture(x0,y0,x1,y1)
	if err != nil {
		panic(err)
		return img, err
	}
	return img, nil
}

func ScreenshotRect(bounds image.Rectangle) (*image.RGBA, error) {
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
		return img, err
	}
	return img, nil
}

func ScreenshotAllScreensAndSaveToFile(x int)  {
	n := screenshot.NumActiveDisplays()
	for j:=0; j<x; j++ {
		for i := 0; i < n; i++ {
			bounds := screenshot.GetDisplayBounds(i)
			//bounds = image.Rect(0,0,500,500)
			img, err := screenshot.CaptureRect(bounds)
			if err != nil {
				panic(err)
			}
			fileName := fmt.Sprintf("%d_%d_%dx%d.png", i, j, bounds.Dx(), bounds.Dy())
			file, _ := os.Create(fileName)
			defer file.Close()
			png.Encode(file, img)

			fmt.Printf("#%d %d : %v \"%s\"\n", i, j, bounds, fileName)
		}
	}
}