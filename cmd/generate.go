/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the color palette generation command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the 5 color palette",
	Long: `Choose between (M)onochromatic, (A)nalogous, (C)omplementary, (T)riadic, or (S)quare
	using the --scheme flag.
	
	Ex. hexcli generate --scheme=M`,
	Run: func(cmd *cobra.Command, args []string) {
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
