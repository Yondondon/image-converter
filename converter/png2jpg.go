package converter

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
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

	outputDir := filepath.Dir(ptj.OutputPath)

	output, err := os.CreateTemp(outputDir, ".png2jpg-*.tmp")
	if err != nil {
		return err
	}

	tempPath := output.Name()
	defer os.Remove(tempPath)

	if err := jpeg.Encode(output, img, &jpeg.Options{
		Quality: ptj.Quality,
	}); err != nil {
		output.Close()
		return err
	}

	if err := output.Close(); err != nil {
		return err
	}

	if err := os.Rename(tempPath, ptj.OutputPath); err != nil {
		return err
	}

	return nil
}
