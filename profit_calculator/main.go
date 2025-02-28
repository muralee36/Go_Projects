package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("What is your yearly revenue ? :")
	fmt.Scan(&revenue)
	fmt.Print("What was your expenses for the year ? :")
	fmt.Scan(&expenses)
	fmt.Print("What your yearly tax rate ? :")
	fmt.Scan(&taxRate)
	profit := (revenue * (1 - taxRate/100)) - expenses
	ebt := revenue - expenses
	ratio := ebt / profit
	fmt.Println("Earnings before tax : ", ebt)
	fmt.Println("Profit after tax : ", profit)
	fmt.Println("The ratio between ebt and profit : ", ratio)
}
