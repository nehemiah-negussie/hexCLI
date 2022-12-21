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
	"github.com/gookit/color"
)

type HSVColor struct{
	h int
	s int
	v int
}

func HSVToRGB (color HSVColor) (r, g, b int){
	h := float64(color.h)
	s := float64(color.s) / 100
	v := float64(color.v) / 100

	M := 255 * v
	m := M * (1 - s)
	z := (M-m) * (1 - math.Abs(math.Mod((h/60), 2) - 1))

	q := int(M)
	w := int(m)
	e := int(z)

	switch {
	case h >= 0 && h < 60:
		r = q
		g = w + e
		b = w
	case h >= 60 && h < 120:
		r = w + e
		g = q
		b = w
	case h >= 120 && h < 180:
		r = w
		g = q
		b = w + e
	case h >= 180 && h < 240:
		r = w
		g = w + e
		b = q
	case h >= 240 && h < 300:
		r = w + e
		g = w
		b = q
	case h >= 300 && h < 360:
		r = q
		g = w
		b = w + e
	}
	
	return
}

func fixAngle (angle int) (fixed int) {
	fixed = angle
	// if the angle is between 0 and 360 its ok but if not
	if angle <= 360 && angle >= 0 {
		return angle
	}
	
	if angle < 0 {
		fixed = angle
		for fixed < 0{
			fixed += 360
		}
		return
	}

	if angle > 0 {
		fixed = angle
		for fixed >= 360{
			fixed -= 360
		}
		return
	}
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
		var palette[5]HSVColor
		// Switch statement for each color scheme (algorithms to be added in the future)
		switch scheme {
			case "":
				fmt.Println("Missing scheme choice")
			case "M":
				// Monochromatic

				// Pick a certain hue
				hue := rand.Intn(360)
				// Create an array to store the 5 color palette

				// Loop through and add all 5 colors with random saturation and values
				for i:=0; i < 5; i++ {
					palette[i] = HSVColor{hue, rand.Intn((100-30) + 1) + 30, rand.Intn((100-30) + 1) + 30}	
				}
				

			case "A":
				// Analgous

				palette[2] = HSVColor{rand.Intn(360), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				baseHue := palette[2].h 
				palette[0] = HSVColor{fixAngle(baseHue - 60), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				palette[1] = HSVColor{fixAngle(baseHue - 30), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				palette[3] = HSVColor{fixAngle(baseHue + 30), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				palette[4] = HSVColor{fixAngle(baseHue + 60), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	

				fmt.Println(palette)
				
			case "C":
				// Complementary
				palette[0] = HSVColor{rand.Intn(360), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				palette[3] = HSVColor{fixAngle(palette[0].h + 180), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[1] = HSVColor{palette[0].h, rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
				palette[2] = HSVColor{palette[0].h, rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
				palette[4] = HSVColor{palette[3].h, rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}

				fmt.Println(palette)
			case "T":
				// Triadic
				
				palette[0] = HSVColor{rand.Intn(360), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[1] = HSVColor{fixAngle(palette[0].h - 120), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[2] = HSVColor{fixAngle(palette[0].h + 120), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[3] = HSVColor{fixAngle(palette[rand.Intn(2)].h),  rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
				palette[4] = HSVColor{fixAngle(palette[rand.Intn(2)].h),  rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
				fmt.Println(palette)
			case "S":
				// Square
				var palette[4]HSVColor
				palette[0] = HSVColor{rand.Intn(360), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[1] = HSVColor{fixAngle(palette[0].h + 90), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[2] = HSVColor{fixAngle(palette[0].h - 90), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[3] = HSVColor{fixAngle(palette[0].h + 180), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				fmt.Println(palette)
			default:
				// Best looking color scheme out of the 5 is default
		}

		for i:=0; i < 5; i++ {
			r, g, b := HSVToRGB(palette[i])
			color.RGB(uint8(r), uint8(g), uint8(b)).Println("Color", i+1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().String("scheme", "", "Color scheme choice between (M)onochromatic, (A)nalogous, (C)omplementary, (T)riadic, or (S)quare")


}
