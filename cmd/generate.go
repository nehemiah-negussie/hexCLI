/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"github.com/spf13/cobra"
)

type HSVColor struct{
	h int
	s int
	v int
}

func HSVToRGB(h, s, v int) (r, g, b int) {
	c := v * s
	x := int(float64(c) * (1 - math.Abs(float64((h/60)%2-1))))
	m := v - c

	switch {
	case h >= 0 && h < 60:
		r = c
		g = x
		b = 0
	case h >= 60 && h < 120:
		r = x
		g = c
		b = 0
	case h >= 120 && h < 180:
		r = 0
		g = c
		b = x
	case h >= 180 && h < 240:
		r = 0
		g = x
		b = c
	case h >= 240 && h < 300:
		r = x
		g = 0
		b = c
	case h >= 300 && h < 360:
		r = c
		g = 0
		b = x
	}

	r = (r + m) * 255
	g = (g + m) * 255
	b = (b + m) * 255

	return
}


// generateCmd represents the color palette generation command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the 5 color palette",
	Long: `Choose between (M)onochromatic, (A)nalogous, (C)omplementary, (T)riadic, or (S)quare
	using the --scheme flag.
	
	Ex. hexcli generate --scheme=M`,
	Run: func(cmd *cobra.Command, args []string) {
		// Seed randomness with time
		rand.Seed(time.Now().UnixNano())
		// Get the value of the flag
		scheme, _ := cmd.Flags().GetString("scheme")
		// [testing] show if input was taken
		if scheme != "" {
			fmt.Println("Test")
		}
		// Switch statement for each color scheme (algorithms to be added in the future)
		switch scheme {
			case "":
				fmt.Println("Missing scheme choice")
			case "M":
				// Monochromatic

				// Pick a certain hue
				hue := rand.Intn(360)
				// Create an array to store the 5 color palette
				var palette[5]HSVColor

				// Loop through and add all 5 colors with random saturation and values
				for i:=0; i < 5; i++ {
					palette[i] = HSVColor{hue, rand.Intn((100-60) + 1) + 60, rand.Intn((100-80) + 1) + 80}	
				}
				
				fmt.Println(palette)

			case "A":
				// Analgous
			case "C":
				// Complementary
			case "T":
				// Triadic
			case "S":
				// Square
			default:
				// Best looking color scheme out of the 5 is default
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().String("scheme", "", "Color scheme choice between (M)onochromatic, (A)nalogous, (C)omplementary, (T)riadic, or (S)quare")


}
