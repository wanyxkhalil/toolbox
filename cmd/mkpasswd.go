package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wanyxkhalil/toolbox/mkpasswd"
)

var (
	length, lower, upper, digit, special uint
)

// mkpasswdCmd represents the mkpasswd command
var mkpasswdCmd = &cobra.Command{
	Use:   "mkpasswd",
	Short: "Generate password",
	Long:  `A tool for generating random passwords`,
	Run: func(cmd *cobra.Command, args []string) {
		mkpasswd.Run(length, lower, upper, digit, special)
	},
}

func init() {
	rootCmd.AddCommand(mkpasswdCmd)

	mkpasswdCmd.Flags().UintVarP(&length, "length", "l", 9, "Length in chars")
	mkpasswdCmd.Flags().UintVarP(&lower, "lower", "c", 2, "Number of lowercase chars")
	mkpasswdCmd.Flags().UintVarP(&upper, "upper", "C", 2, "Number of uppercase chars")
	mkpasswdCmd.Flags().UintVarP(&digit, "digit", "d", 2, "Number of digits")
	mkpasswdCmd.Flags().UintVarP(&special, "special", "s", 0, "Number of special chars")
}
