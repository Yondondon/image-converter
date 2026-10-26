package cmd

import (
	"fmt"

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

		// TODO: change from 'required' to output in the same directory if this flag
		// isn't provided
		if outputPath == "" {
			return fmt.Errorf("output path can't be an empty string")
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
	f.StringP("output", "o", "", "(required) Output file path")
	f.IntVarP(&quality, "quality", "q", 85, "JPEG quality (1-100)")

}
