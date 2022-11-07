// What it does:
//
// 	This program outputs the current OpenCV library version to the console.
//
// How to run:
//
// 		go run ./cmd/version/main.go
//

package main

import (
	"log"
	"time"

	// "github.com/kbinani/screenshot"
	"github.com/vova616/screenshot"
	"gocv.io/x/gocv"
)

// define variable to loop time

func main() {
	window := gocv.NewWindow("Game window")
	last_time := time.Now()
	for {
		img, err := screenshot.CaptureScreen()
		
		if err != nil {
			panic(err)
		}

		newImg, err := gocv.ImageToMatRGBA(img)
		if err != nil {
			panic(err)
		}

		log.Printf("Loop took %f seconds", time.Since(last_time).Seconds())
		last_time = time.Now()
		window.IMShow(newImg)
		window.WaitKey(1)
	}

	// for {
	// 	img, err := screenshot.CaptureScreen()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	myImg := image.Image(img)
	// 	newImg, err := gocv.ImageToMatRGB(myImg)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	window.IMShow(newImg)
	// }
}
