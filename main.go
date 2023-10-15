package main

import (
	"fmt"
	"image"
	"net/http"
	"os"

	"golang.org/x/image/bmp"

	"github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Last-Modified", "*")
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	fmt.Println("Hello, World!")

	newWidth, newHeight := 840, 480
	blackImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	// blackImage := image.NewPaletted(image.Rect(0, 0, newWidth, newHeight), color.Palette{
	// 	color.RGBA{255, 0, 0, 255},     // Red
	// 	color.RGBA{0, 255, 0, 255},     // Green
	// 	color.RGBA{0, 0, 255, 255},     // Blue
	// 	color.RGBA{0, 0, 0, 255},       // White
	// 	color.RGBA{255, 255, 255, 255}, // Black
	// })
	// blackImage := image.NewPaletted(image.Rect(0, 0, newWidth, newHeight), palette.WebSafe)
	// draw.Draw(blackImage, blackImage.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)
	// addLabel(blackImage, 10, 10, "Hello, image!")

	fontPath, err := findfont.Find("arial.ttf")

	// dc := gg.NewContext(840, 480)
	dc := gg.NewContextForImage(blackImage)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(fontPath, 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", float64(newWidth/2), float64(newHeight/2), 0.5, 0.5)

	dc.SetRGB(0, 0, 0)
	dc.DrawRoundedRectangle(0, 0, 20, float64(newWidth), 0)
	dc.Fill()
	dc.DrawImage(blackImage, 0, 0)
	dc.Clip()
	image := dc.Image()

	outputFile, err := os.Create("output/go-black.bmp")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	err = bmp.Encode(outputFile, image)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fs := http.FileServer(http.Dir("output"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8090", nil)
}
