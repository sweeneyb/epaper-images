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

func createRedImageData() image.Image {
	newWidth, newHeight := 800, 480
	redImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	dc := gg.NewContextForImage(redImage)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	// left
	dc.DrawRectangle(0, 0, 40, float64(newHeight))
	dc.Fill()
	//top
	dc.DrawRectangle(0, 0, float64(newWidth), 20)
	dc.Fill()
	//right
	dc.DrawRectangle(float64(newWidth-40), 0, 40, float64(newHeight))
	dc.Fill()
	//bottom
	dc.DrawRectangle(0, float64(newHeight-20), float64(newWidth), 20)
	dc.Fill()
	dc.DrawImage(redImage, 0, 0)
	dc.Clip()
	return dc.Image()
}

func createBlackImageData() image.Image {
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
	return dc.Image()
}

func greyScale(imageData image.Image) []color.Gray {
	var raw []color.Gray
	result := image.NewGray(imageData.Bounds())
	draw.Draw(result, result.Bounds(), imageData, imageData.Bounds().Min, draw.Src)

	var i = 0
	for y := result.Bounds().Min.Y; y < result.Bounds().Max.Y; y++ {
		for x := result.Bounds().Min.X; x < result.Bounds().Max.X; x++ {
			result.At(x, y)
			// fmt.Print(resultImage.At(x, y))
			// (*buffer)[i] = result.At(x, y).(color.Gray)
			raw = append(raw, result.At(x, y).(color.Gray))
			i++
		}
	}
	return raw
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

func redLayer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Last-Modified", "*")
	var raw []color.Gray
	redImageData := createRedImageData()
	raw = greyScale(redImageData)

	outputFile, err := os.Create("output/go-red-01.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	err = png.Encode(outputFile, redImageData)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	var xbm []uint8
	compressToXBM(&xbm, &raw)
	outputFile, err = os.Create("output/go-red-01.xbm")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	for _, pixel := range xbm {
		// Assuming little-endian byte order for simplicity
		binary.Write(outputFile, binary.LittleEndian, pixel)
	}
	// fmt.Fprintf(w, "", raw)
	for _, pixel := range xbm {
		// Assuming little-endian byte order for simplicity
		binary.Write(w, binary.LittleEndian, pixel)
	}
	// fmt.Println("data: %v", raw)
}

func blackLayer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Last-Modified", "*")
	var raw []color.Gray
	blackImageData := createBlackImageData()
	raw = greyScale(blackImageData)
	outputFile, err := os.Create("output/go-black-01.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	err = png.Encode(outputFile, blackImageData)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	var xbm []uint8
	compressToXBM(&xbm, &raw)
	outputFile, err = os.Create("output/go-black-01.xbm")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	for _, pixel := range xbm {
		// Assuming little-endian byte order for simplicity
		binary.Write(outputFile, binary.LittleEndian, pixel)
	}
	// fmt.Fprintf(w, "", raw)
	for _, pixel := range xbm {
		// Assuming little-endian byte order for simplicity
		binary.Write(w, binary.LittleEndian, pixel)
	}
	// fmt.Println("data: %v", raw)
}

func main() {
	// createImages()
	fmt.Println("Starting Server on 8090")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fs := http.FileServer(http.Dir("output"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/blackLayer", blackLayer)
	http.HandleFunc("/redLayer", redLayer)

	http.ListenAndServe(":8090", nil)
}
