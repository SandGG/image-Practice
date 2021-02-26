package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type img struct {
	width  int
	height int
}

func main() {
	var imageVar = img{width: 200, height: 200}

	//Min: {0,0}, Max: {200, 200}
	var img = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{imageVar.width, imageVar.height}})

	//Colors
	var pink = color.RGBA{241, 193, 212, 0xff}
	var purple = color.RGBA{199, 176, 225, 0xff}
	var skyBlue = color.RGBA{187, 231, 235, 0xff}
	var lemon = color.RGBA{216, 245, 158, 0xff}

	for x := 0; x < imageVar.width; x++ {
		for y := 0; y < imageVar.height; y++ {
			switch {
			case x < imageVar.width/2 && y <= imageVar.height/2: //Top left
				img.Set(x, y, pink)
			case x >= imageVar.width/2 && y <= imageVar.height/2: //Top right
				img.Set(x, y, lemon)
			case x < imageVar.width/2 && y > imageVar.height/2: //Bott left
				img.Set(x, y, skyBlue)
			case x >= imageVar.width/2 && y >= imageVar.height/2: //Bott right
				img.Set(x, y, purple)
			}
		}
	}

	// Encode as PNG.
	var file, err = os.Create("./image.png")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, img)
	fmt.Println("File PNG Updated Successfully")
}
