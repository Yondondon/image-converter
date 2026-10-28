package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Yondondon/image-converter/converter"
	"github.com/spf13/cobra"
)

var quality int

var png2jpgCmd = &cobra.Command{
	Use:   "png2jpg <input.png>",
	Short: "Converts PNG image to JPG",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := args[0]

		outputPath, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		if outputPath == "" {
			dir := filepath.Dir(inputPath)
			name := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
			outputPath = filepath.Join(dir, name+".jpg")
		}

		if quality < 1 || quality > 100 {
			return fmt.Errorf("quality must be between 1 and 100, got %d", quality)
		}

		c := converter.PNG2JPG{
			InputPath:  inputPath,
			OutputPath: outputPath,
			Quality:    quality,
		}

		if err := c.Convert(); err != nil {
			return fmt.Errorf("conversion failed: %w", err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Converted %s to %s\n", inputPath, outputPath)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(png2jpgCmd)

	f := png2jpgCmd.Flags()
	f.StringP("output", "o", "", "Output file path. Converted image will be placed in the same directory if output isn't provided")
	f.IntVarP(&quality, "quality", "q", 85, "JPEG quality (1-100)")

}
