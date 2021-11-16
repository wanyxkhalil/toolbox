// Package pension 参考人社通页面 https://m12333.cn/qa/isyw.html
package pension

import "fmt"

func Run(avgWages int, wageRatio float64, years int, personalWages int, months int) {
	basic := basicPension(avgWages, wageRatio, years)
	private := privatePension(personalWages, years, months)
	fmt.Printf("基础养老金：\t%.2f\n个人养老金：\t%.2f\n合计：\t\t%.2f\n", basic, private, basic+private)
}

// basicPension 基础养老金，计算公式 = P × (1 + i）÷ 2 × n × 1%
//P：退休上年度当地在岗职工月平均工资
//i：本人历年缴费指数的平均值（缴费指数=本人缴费工资基数÷社会平均工资）
//n：本人累计缴费年限（含视同缴费）
func basicPension(avgWages int, wageRatio float64, years int) float64 {
	return float64(avgWages) * (1 + wageRatio) / 2 * float64(years) * 0.01
}

// privatePension 个人账户养老金，计算公式 ＝ 个人账户储存额 ÷ 计发月数
// 个人账户存储额：主要来源于历年的个人缴费（费率8%）及其利息
// 计发月数：根据平均寿命计算，50岁退休按195个月，55岁退休按170个月，60岁退休按139个月
func privatePension(personalWages int, years int, months int) float64 {
	return float64(personalWages*12*years/months) * 0.08
}
