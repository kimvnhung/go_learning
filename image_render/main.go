package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"time"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type Circle struct {
	radius int
	center image.Point
}

func isInsideCircle(point image.Point, circle Circle) bool {
	dd := math.Pow((float64(point.X-circle.center.X)), 2) + math.Pow((float64(point.Y-circle.center.Y)), 2)
	return dd < float64(circle.radius*circle.radius)
}

func isInsideRectangle(point image.Point, rect image.Rectangle) bool {
	return rect.Min.X <= point.X && point.X <= rect.Max.X && rect.Min.Y <= point.Y && point.Y <= rect.Max.Y
}

// circle is a custom shape (rounded rectangle) implementing the image.Image interface
type RoundedRect struct {
	width  int
	height int
	radius int
}

func (c *RoundedRect) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *RoundedRect) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.width, c.height)
}

func (c *RoundedRect) At(x, y int) color.Color {
	O1 := image.Point{c.radius, c.radius}
	O2 := image.Point{c.width - c.radius, c.radius}
	O3 := image.Point{c.width - c.radius, c.height - c.radius}
	O4 := image.Point{c.radius, c.height - c.radius}
	corner1 := image.Rect(0, 0, c.radius, c.radius)
	corner2 := image.Rect(c.width-c.radius, 0, c.width, c.radius)
	corner3 := image.Rect(c.width-c.radius, c.height-c.radius, c.width, c.height)
	corner4 := image.Rect(0, c.height-c.radius, c.radius, c.height)
	internal := image.Rect(O1.X, O1.Y, O3.X, O3.Y)
	curP := image.Point{x, y}
	circle1 := Circle{radius: c.radius, center: O1}
	circle2 := Circle{radius: c.radius, center: O2}
	circle3 := Circle{radius: c.radius, center: O3}
	circle4 := Circle{radius: c.radius, center: O4}
	if isInsideRectangle(curP, internal) || isInsideCircle(curP, circle1) || isInsideCircle(curP, circle2) || isInsideCircle(curP, circle3) || isInsideCircle(curP, circle4) {
		return color.Alpha{255}
	}

	if !(isInsideRectangle(curP, corner1) || isInsideRectangle(curP, corner2) || isInsideRectangle(curP, corner3) || isInsideRectangle(curP, corner4)) {
		return color.Alpha{255}
	}
	return color.Transparent
}

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
	backgroundN := image.NewRGBA(background.Bounds())
	// unsplashResizedRounedCorner := roundeRectangle(unsplashResized, 50)
	draw.Draw(backgroundN, background.Bounds(), background, image.Point{0, 0}, draw.Over)
	draw.DrawMask(backgroundN, image.Rect(526, 296, 1082+526, 1453+296), unsplashResized, image.Point{0, 0}, &RoundedRect{unsplashResized.Bounds().Dx(), unsplashResized.Bounds().Dy(), 50}, image.Point{0, 0}, draw.Over)
	// dc.DrawImage(unsplashResizedRounedCorner, 526, 296)

	dc := gg.NewContextForRGBA(backgroundN)
	dc.SavePNG("shared.png")

	log.Printf("time : %d", time.Now().UnixMilli()-start)
}

func roundeRectangle(input image.Image, radius int) image.Image {
	rounedRectBase := image.NewRGBA(input.Bounds())
	draw.Draw(rounedRectBase, rounedRectBase.Bounds(), &image.Uniform{color.Transparent}, image.Point{0, 0}, draw.Over)
	draw.DrawMask(rounedRectBase, rounedRectBase.Bounds(), input, image.Point{0, 0}, &RoundedRect{input.Bounds().Dx(), input.Bounds().Dy(), radius}, image.Point{0, 0}, draw.Over)
	return rounedRectBase
}
