package main

import (
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

func main() {
	log.Println("Hello")
	// resized cover : 1080x1440
	// background shared : 2160x2160
	start := time.Now().UnixMilli()
	background, _ := gg.LoadImage("background.png")
	// backgroundResized := resize.Resize(376, 376, background, resize.Lanczos2)
	unsplash, _ := gg.LoadImage("raw_image.jpg")
	unsplashResized := resize.Resize(1082, 1453, unsplash, resize.Lanczos2)
	gg.SavePNG("unsplashResized.png", unsplashResized)
	// gg.SavePNG("backgroundResized.png", backgroundResized)
	dc := gg.NewContextForImage(background)

	dc.DrawImage(unsplashResized, 526, 296)

	dc.SavePNG("shared.png")

	log.Printf("time : %d", time.Now().UnixMilli()-start)
}
