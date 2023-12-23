package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
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

func createImages() {
	newWidth, newHeight := 800, 480
	// newWidth, newHeight := 210, 120
	blackImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	redImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
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
	dc.DrawStringAnchored("On Air!!", float64(newWidth/2), float64(newHeight/2), 0.5, 0.5)

	dc.SetRGB(0, 0, 0)
	// left
	// dc.DrawRectangle(0, 0, 20, float64(newHeight))
	// dc.Fill()
	// //top
	// dc.DrawRectangle(0, 0, float64(newWidth), 20)
	// dc.Fill()
	// //right
	// dc.DrawRectangle(float64(newWidth-20), 0, 20, float64(newHeight))
	// dc.Fill()
	// //bottom
	// dc.DrawRectangle(0, float64(newHeight-20), float64(newWidth), 20)
	// dc.Fill()
	dc.DrawImage(blackImage, 0, 0)
	dc.Clip()
	resultImage := dc.Image()

	outputFile, err := os.Create("output/go-black3.bmp")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	err = bmp.Encode(outputFile, resultImage)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	outputFile, err = os.Create("output/go-black3.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	err = png.Encode(outputFile, resultImage)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	dc = gg.NewContextForImage(redImage)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	// left
	dc.DrawRectangle(0, 0, 20, float64(newHeight))
	dc.Fill()
	//top
	dc.DrawRectangle(0, 0, float64(newWidth), 20)
	dc.Fill()
	//right
	dc.DrawRectangle(float64(newWidth-20), 0, 20, float64(newHeight))
	dc.Fill()
	//bottom
	dc.DrawRectangle(0, float64(newHeight-20), float64(newWidth), 20)
	dc.Fill()
	dc.DrawImage(blackImage, 0, 0)
	dc.Clip()
	resultImage = dc.Image()

	outputFile, err = os.Create("output/go-red3.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	err = png.Encode(outputFile, resultImage)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	result := image.NewGray(resultImage.Bounds())
	draw.Draw(result, result.Bounds(), resultImage, resultImage.Bounds().Min, draw.Src)

	for y := result.Bounds().Min.Y; y < result.Bounds().Max.Y; y++ {
		for x := result.Bounds().Min.X; x < result.Bounds().Max.X; x++ {
			result.At(x, y)
			// fmt.Print(result.At(x, y))
		}
	}
	fmt.Println("")

	outputFile, err = os.Create("output/go-red-grey1.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	err = png.Encode(outputFile, result)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}
}

func createBlackImage(buffer *[]color.Gray) {
	newWidth, newHeight := 800, 480
	blackImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	fontPath, _ := findfont.Find("arial.ttf")

	dc := gg.NewContextForImage(blackImage)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(fontPath, 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("On Air!!", float64(newWidth/2), float64(newHeight/2), 0.5, 0.5)

	dc.SetRGB(0, 0, 0)
	dc.DrawImage(blackImage, 0, 0)
	dc.Clip()
	resultImage := dc.Image()
	result := image.NewGray(resultImage.Bounds())
	draw.Draw(result, result.Bounds(), resultImage, resultImage.Bounds().Min, draw.Src)

	var i = 0
	for y := result.Bounds().Min.Y; y < result.Bounds().Max.Y; y++ {
		for x := result.Bounds().Min.X; x < result.Bounds().Max.X; x++ {
			result.At(x, y)
			// fmt.Print(resultImage.At(x, y))
			// (*buffer)[i] = result.At(x, y).(color.Gray)
			*buffer = append(*buffer, result.At(x, y).(color.Gray))
			i++
		}
	}
	fmt.Println("")
}

func compressToXBM(xbm *[]uint8, buffer *[]color.Gray) {
	for i := 0; i < len(*buffer); i++ {
		if i%8 == 0 {
			*xbm = append(*xbm, 0) // add a byte if needed
		}
		// fmt.Printf("i: %v %02b \n", i, (*buffer)[i].Y)
		if (*buffer)[i].Y > 127 {
			// Set the corresponding bit in the buffer
			(*xbm)[i/8] |= 1 << uint(7-(i%8))
		}
	}
}

func blackLayer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Last-Modified", "*")
	var raw []color.Gray
	createBlackImage(&raw)
	var xbm []uint8
	compressToXBM(&xbm, &raw)
	// fmt.Fprintf(w, "", raw)
	for _, pixel := range xbm {
		// Assuming little-endian byte order for simplicity
		binary.Write(w, binary.LittleEndian, pixel)
	}
	// fmt.Println("data: %v", raw)
}

func main() {
	createImages()
	fmt.Println("Starting Server on 8090")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fs := http.FileServer(http.Dir("output"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/blackLayer", blackLayer)

	http.ListenAndServe(":8090", nil)
}
