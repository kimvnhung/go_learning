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
	start := time.Now().UnixMilli()
	scaleTemp, err := gg.LoadImage("raw_image.jpg")
	if err != nil {
		log.Printf("%v", err)
	}
	scaleCtx := gg.NewContextForImage(scaleTemp)
	img := resize.Resize(400, 300, scaleCtx.Image(), resize.Lanczos2)

	dc := gg.NewContextForImage(img)
	dc.SavePNG("out.png")
	log.Printf("time : %d", time.Now().UnixMilli()-start)
}
