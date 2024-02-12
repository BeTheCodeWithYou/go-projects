package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func main() {
	fmt.Println("hello")

	img := loadImg("D:\\dev\\go-projects\\image-compressor\\manu.png")
	saveImg("D:\\dev\\go-projects\\image-compressor\\manu2.PNG", img)
}

func loadImg(filepath string) *image.RGBA64 {

	imgFile, err := os.Open(filepath)
	defer imgFile.Close()
	if err!=nil{
		log.Print("can't read file:", filepath)
	}

	img, _, err := image.Decode(imgFile)
	if err!=nil{
		log.Print("can't decode ", err)
	}

	return img.(*image.RGBA64)
}

func saveImg(filePath string, img *image.RGBA64) {

	 f, err := os.Create(filePath)
	 defer f.Close()
	 if err!=nil{
		log.Print("can't create file->", err)
		return
	 }

	 png.Encode(f, img)

}