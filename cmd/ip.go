package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wanyxkhalil/toolbox/ip"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip [domain]",
	Short: "Get IP Local or Remote",
	Long:  `Get IP Local or Remote`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			ip.Domain(args[0])
		} else {
			ip.Local()
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)

	ipCmd.UsageString()
}
