package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var taxRate float64
	InputCollector("Revenue", &revenue)
	InputCollector("Expense", &expense)
	InputCollector("Tax Rate", &taxRate)

	ebt, eat, ratio := Earnings(revenue, expense, taxRate)
	fmt.Printf("Earnings Before Tax : %.2f\nEarning After Tax : %.2f\nRatio : %.2f\n", ebt, eat, ratio)

}

func InputCollector(humanText string, valueVariable *float64) {
	fmt.Print(humanText, " : ")
	fmt.Scan(valueVariable)

}

func Earnings(revenue, expense, tax float64) (float64, float64, float64) {
	ebt := revenue - expense
	eat := ebt * (1 - tax/100)
	ratio := ebt / eat
	return ebt, eat, ratio
}
