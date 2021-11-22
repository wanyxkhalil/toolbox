package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wanyxkhalil/toolbox/fortune"
)

var path string

// fortuneCmd represents the fortune command
var fortuneCmd = &cobra.Command{
	Use:   "fortune",
	Short: "displays a pseudorandom message from a database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fortune.Run(path)
	},
}

func init() {
	rootCmd.AddCommand(fortuneCmd)

	fortuneCmd.Flags().StringVarP(&path, "path", "p", "", "fortune data file path")
}
