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
		for fixed < 0{
			fixed += 360
		}
		return
	}

	if angle > 0 {
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
	
	Ex. hexcli generate --scheme=M
	
	Use --info for more info on the color schemes and common applications.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		info, _ := cmd.Flags().GetBool ("info")
		if info {
			fmt.Println(`
			(M)onochromatic: Creates a cohesive, unified look by using variations of a single color. Great for single page design or design around a dominant theme.
			
			(A)nalogous: Uses colors that are adjacent to each other on the color wheel, creating a harmonious and cohesive look.
			
			(C)omplementary: Uses colors that are opposite each other on the color wheel, creating a bold, high-contrast look.
			
			(T)riadic: Uses colors that are evenly spaced around the color wheel. Great for high contrast designs.
			
			(S)quare: Uses four colors that are evenly spaced around the color wheel. Useful for finding a base color to work off of.`)
			return
		}

		// Seed randomness with time
		rand.Seed(time.Now().UnixNano())

		// Get the value of the flag
		scheme, _ := cmd.Flags().GetString("scheme")

		// Define palette array
		var palette[5]HSVColor

		// Switch statement for each color scheme
		switch scheme {
			case "":
				fmt.Println("Missing scheme choice")
			case "M":
				// Monochromatic

				// Pick a certain hue
				hue := rand.Intn(360)

				// Loop through and add all 5 colors with same hue and random saturation and values (30%-100%)
				for i:=0; i < 5; i++ {
					palette[i] = HSVColor{hue, rand.Intn((100-30) + 1) + 30, rand.Intn((100-30) + 1) + 30}	
				}
				

			case "A":
				// Analgous
				
				// Set the middle color as the base color
				palette[2] = HSVColor{rand.Intn(360), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				baseHue := palette[2].h 
				// The next 4 colors are 30 degrees away from base hue or 30 degrees away from the adjacent color
				palette[0] = HSVColor{fixAngle(baseHue - 60), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				palette[1] = HSVColor{fixAngle(baseHue - 30), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				palette[3] = HSVColor{fixAngle(baseHue + 30), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				palette[4] = HSVColor{fixAngle(baseHue + 60), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}	
				
			case "C":
				// Complementary

				// Pick a random color
				palette[0] = HSVColor{rand.Intn(360), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				// Pick the hue directly opposite to it
				palette[3] = HSVColor{fixAngle(palette[0].h + 180), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				// Fill the rest of the color palette with 2 of the base hues and 1 of the opposite hues
				palette[1] = HSVColor{palette[0].h, rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
				palette[2] = HSVColor{palette[0].h, rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
				palette[4] = HSVColor{palette[3].h, rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
			case "T":
				// Triadic

				// Find a base color
				palette[0] = HSVColor{rand.Intn(360), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				// Add 120 degrees and subtract 120 degrees from the base hue to get the triadic colors
				palette[1] = HSVColor{fixAngle(palette[0].h - 120), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[2] = HSVColor{fixAngle(palette[0].h + 120), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				// Choose 1 of the three hues randomly and vary the saturation and value
				palette[3] = HSVColor{fixAngle(palette[rand.Intn(2)].h),  rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
				palette[4] = HSVColor{fixAngle(palette[rand.Intn(2)].h),  rand.Intn((100-40) + 1) + 40, rand.Intn((100-40) + 1) + 40}
				fmt.Println(palette)
			case "S":
				// Square

				// Find a base color
				palette[0] = HSVColor{rand.Intn(360), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				// Add 90 degrees unitl you find all four colors
				palette[1] = HSVColor{fixAngle(palette[0].h + 90), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[2] = HSVColor{fixAngle(palette[0].h - 90), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				palette[3] = HSVColor{fixAngle(palette[0].h + 180), rand.Intn((100-80) + 1) + 80, rand.Intn((100-80) + 1) + 80}
				fmt.Println(palette)
			default:
				fmt.Println("Incorrect scheme choice, please choose from (M)onochromatic, (A)nalogous, (C)omplementary, (T)riadic, or (S)quare")
		}

		// Print out the palette with colored output
		for i:=0; i < 5; i++ {
			r, g, b := HSVToRGB(palette[i])
			if r == 0 && g == 0 && b == 0{
				continue
			}
			output := fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
			color.RGB(uint8(r), uint8(g), uint8(b)).Println(output)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().String("scheme", "", "Color scheme choice between (M)onochromatic, (A)nalogous, (C)omplementary, (T)riadic, or (S)quare")
	generateCmd.Flags().Bool("info", false, "Give more info on the applications of each color scheme")

}
