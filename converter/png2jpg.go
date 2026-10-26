package converter

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
)

type PNG2JPG struct {
	InputPath  string
	OutputPath string
	Quality    int
}

func (ptj PNG2JPG) Convert() error {
	input, err := os.Open(ptj.InputPath)
	if err != nil {
		return err
	}
	defer input.Close()

	img, _, err := image.Decode(input)
	if err != nil {
		return err
	}

	//TODO: write to temp file and rename after successful conversion to prevent
	// partial file
	output, err := os.Create(ptj.OutputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	if err := jpeg.Encode(output, img, &jpeg.Options{Quality: ptj.Quality}); err != nil {
		return err
	}

	return nil
}
