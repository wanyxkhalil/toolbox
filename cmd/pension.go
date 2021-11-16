package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wanyxkhalil/toolbox/pension"
)

var (
	avgWages      int
	wageRatio     float64
	years         int
	personalWages int
	months        int
)

// pensionCmd represents the pension command
var pensionCmd = &cobra.Command{
	Use:   "pension",
	Short: "计算基本养老金",
	Long:  `基本养老金由基础养老金和个人账户养老金两部分组成。`,
	Run: func(cmd *cobra.Command, args []string) {
		pension.Run(avgWages, wageRatio, years, personalWages, months)
	},
}

func init() {
	rootCmd.AddCommand(pensionCmd)

	pensionCmd.Flags().IntVarP(&avgWages, "avgWages", "", 0, "退休上年度当地在岗职工月平均工资")
	pensionCmd.Flags().Float64VarP(&wageRatio, "wageRatio", "", 0, "本人历年缴费指数的平均值")
	pensionCmd.Flags().IntVarP(&years, "years", "", 0, "本人累计缴费年限")
	pensionCmd.Flags().IntVarP(&personalWages, "personalWages", "", 0, "本人缴费工资基数")
	pensionCmd.Flags().IntVarP(&months, "months", "", 0, "计发月数")
}
