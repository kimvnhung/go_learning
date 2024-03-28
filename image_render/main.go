package main

import (
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

func main() {
	log.Println("Hello")
	// background, err := gg.LoadImage("output_temp.png")
	// if err != nil {
	// 	log.Printf("%v", err)
	// }

	// dc := gg.NewContextForImage(background)
	// resized cover : 300x400
	// background shared : 376x376
	start := time.Now().UnixMilli()
	// scaleTemp, err := gg.LoadImage("background.png")
	// if err != nil {
	// 	log.Printf("%v", err)
	// }
	// scaleCtx := gg.NewContextForImage(scaleTemp)
	// img := resize.Resize(400, 300, scaleCtx.Image(), resize.Lanczos2)

	// dc := gg.NewContextForImage(img)
	// dc.SavePNG("out.png")
	background, _ := gg.LoadImage("background.png")
	backgroundResized := resize.Resize(376, 376, background, resize.Lanczos2)
	unsplash, _ := gg.LoadImage("raw_image.jpg")
	unsplashResized := resize.Resize(300, 400, unsplash, resize.Lanczos2)
	gg.SavePNG("unsplashResized.png", unsplashResized)
	// gg.SavePNG("backgroundResized.png", backgroundResized)
	dc := gg.NewContextForImage(backgroundResized)

	log.Printf("time : %d", time.Now().UnixMilli()-start)
}
