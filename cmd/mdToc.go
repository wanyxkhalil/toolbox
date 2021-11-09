package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wanyxkhalil/toolbox/mdtoc"
)

// mdTocCmd represents the mdToc command
var mdTocCmd = &cobra.Command{
	Use:   "md-toc-github [FILE]",
	Short: "Generate markdown toc for github",
	Run: func(cmd *cobra.Command, args []string) {
		mdtoc.Run(args[0])
	},
}

func init() {
	rootCmd.AddCommand(mdTocCmd)

}
