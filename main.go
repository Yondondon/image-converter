package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"time"

	_ "image/gif"
	"image/jpeg"
	_ "image/png"
)

func main() {
	start := time.Now()

	file, err := os.Open("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create("output.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	err = jpeg.Encode(output, img, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
