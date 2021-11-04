package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wanyxkhalil/toolbox/mysqltogostruct"
	"log"
)

// mysqlToGostructCmd represents the mysqlToGostruct command
var mysqlToGostructCmd = &cobra.Command{
	Use:   "mysql-to-gostruct",
	Short: "mysql to go struct",
	Long:  `mysql to go struct`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatalf("Need source and dest path")
		}
		mysqltogostruct.Run(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(mysqlToGostructCmd)
}
