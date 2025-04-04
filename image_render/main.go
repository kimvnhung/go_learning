package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
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

func drawText(img draw.Image, text string, x, y int, c color.Color, fontSize float64, fontWeight int) {
	// Load font file
	fontBytes, err := os.ReadFile("./Lato-Bold.ttf")
	if err != nil {
		panic(err)
	}

	// Parse font file
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	// Create freetype context
	fc := freetype.NewContext()
	fc.SetDPI(float64(fontWeight) * 0.7 / 3)
	fc.SetFontSize(fontSize)
	fc.SetFont(font)
	fc.SetSrc(image.NewUniform(c))
	fc.SetDst(img)
	fc.SetClip(img.Bounds())

	// Draw text
	pt := freetype.Pt(x, y+int(fc.PointToFixed(fontSize)>>6))
	_, err = fc.DrawString(text, pt)
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Hello")
	// resized cover : 1082x1453
	// background shared : 2160x2160
	start := time.Now().UnixMilli()
	background, _ := gg.LoadImage("background.png")
	unsplash, _ := gg.LoadImage("raw_image.jpg")

	unsplashResized := resize.Resize(1082, 1453, unsplash, resize.Lanczos2)

	gg.SavePNG("unsplashResized.png", unsplashResized)
	backgroundN := image.NewRGBA(background.Bounds())
	draw.Draw(backgroundN, background.Bounds(), background, image.Point{0, 0}, draw.Over)
	draw.DrawMask(backgroundN, image.Rect(526, 296, 1608, 1749), unsplashResized, image.Point{0, 0}, &RoundedRect{unsplashResized.Bounds().Dx(), unsplashResized.Bounds().Dy(), 8}, image.Point{0, 0}, draw.Over)

	rounedRect := image.Rect(0, 0, 1022, 1393)
	rounedRectLineBase := image.NewRGBA(rounedRect)
	draw.Draw(rounedRectLineBase, rounedRectLineBase.Bounds(), &image.Uniform{color.Transparent}, image.Point{0, 0}, draw.Over)
	//color.RGBA{255, 255, 255, uint8(0.2 * 255)}
	draw.DrawMask(rounedRectLineBase, rounedRectLineBase.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, uint8(0.2 * 255)}}, image.Point{0, 0}, &RoundedRect{rounedRect.Dx(), rounedRect.Dy(), 8}, image.Point{0, 0}, draw.Over)
	draw.DrawMask(rounedRectLineBase, image.Rect(4, 4, rounedRect.Dx()-4, rounedRect.Dy()-4), &image.Uniform{color.Transparent}, image.Point{0, 0}, &RoundedRect{rounedRect.Dx() - 4, rounedRect.Dy() - 4, 8}, image.Point{0, 0}, draw.Src)
	draw.Draw(backgroundN, image.Rect(526+30, 296+30, rounedRect.Dx()+526+30, rounedRect.Dy()+296+30), rounedRectLineBase, image.Point{0, 0}, draw.Over)

	// Postop
	// drawText(backgroundN, "JACKIE WILLIS", 526+63, 296+57, color.White, 22, 700)
	// drawText(backgroundN, "Empty Streets", 526+63, 296+114, color.White, 36, 900)

	// Pos bottom
	drawText(backgroundN, "JACKIE WILLIS", 526+63, 1350+57, color.White, 22, 700)
	drawText(backgroundN, "Empty Streets", 526+63, 1350+114, color.White, 36, 900)

	dc := gg.NewContextForRGBA(backgroundN)
	dc.SavePNG("shared.png")

	log.Printf("time : %d", time.Now().UnixMilli()-start)
}
