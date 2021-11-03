package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wanyxkhalil/toolbox/httpsexpired"
)

// httpsExpiredCmd represents the httpsExpired command
var httpsExpiredCmd = &cobra.Command{
	Use:   "https-expired",
	Short: "Show the cert expiration time",
	Long:  `Show the cert expiration time`,
	Run: func(cmd *cobra.Command, args []string) {
		httpsexpired.Run(args[0])
	},
}

func init() {
	rootCmd.AddCommand(httpsExpiredCmd)
}
