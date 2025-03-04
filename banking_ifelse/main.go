package main

import (
	"fmt"
	"os"
	"strconv"
)

const passbook string = "banking.txt"

func main() {

	banking_session := true

	// starts banking session
	for banking_session {
		userChoice := bankingChoices()
		banking_session = performBankOperations(userChoice)
	}
	fmt.Print("Thank you for Using Go-Bank")

}

func bankingChoices() (choice int) {
	fmt.Print(`


	Welcome to Go-Bank.
	Please select the action u want to perform:
	1. Get A/C Balance
	2. Deposit Money
	3. Withdraw Money
	4. Exit
	
	Choice : `)
	fmt.Scan(&choice)
	return
}

func performBankOperations(choice int) (continue_banking bool) {

	bankBalance, err := getBankBalance()
	fmt.Print("\n")
	if err != nil {
		fmt.Print("Banking server Down !!")
		return false
	}

	if choice == 1 {
		fmt.Printf("Account Balance : %v", bankBalance)
		return true
	} else if choice == 2 {
		var deposit float64
		fmt.Print("Please Enter the Amount to deposit : ")
		fmt.Scan(&deposit)
		if deposit < 0 {
			fmt.Print("Invalid amount")
			return true
		}
		bankBalance += deposit
		updateBankBalance(bankBalance)
		return true
	} else if choice == 3 {
		var debit float64
		fmt.Print("Please Enter the Amount to debit : ")
		fmt.Scan(&debit)
		if debit < 0 {
			fmt.Print("Invalid amount")
			return true
		}
		if debit > bankBalance {
			fmt.Print("Insufficient funds to make the transaction")
			return true
		}
		bankBalance -= debit
		updateBankBalance(bankBalance)
		return true
	} else if choice == 4 {
		return false
	} else {
		fmt.Print("Invalid Choice")
		return true
	}
}

func getBankBalance() (float64, error) {
	bal_bytes, err := os.ReadFile(passbook)
	if err != nil {
		fmt.Printf("Something went wrong while getting balance data. Error %v", err)
		return 0, err
	}
	bal_str := string(bal_bytes)
	balance, _ := strconv.ParseFloat(bal_str, 64)
	return balance, nil
}

func updateBankBalance(bal float64) {
	balanceText := fmt.Sprint(bal)
	os.WriteFile(passbook, []byte(balanceText), 0644)
}
