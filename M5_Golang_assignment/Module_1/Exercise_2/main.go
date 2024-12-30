package main

import (
	"errors"
	"fmt"
)

// Account struct to store account details
type Account struct {
	ID              int
	Name            string
	Balance         float64
	TransactionHistory []string
}

var accounts []Account // Slice to store all accounts

// Deposit adds money to the account after validation
func Deposit(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	for i, acc := range accounts {
		if acc.ID == accountID {
			accounts[i].Balance += amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory, fmt.Sprintf("Deposited: %.2f", amount))
			return nil
		}
	}

	return errors.New("account not found")
}

// Withdraw subtracts money from the account after validation
func Withdraw(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}

	for i, acc := range accounts {
		if acc.ID == accountID {
			if acc.Balance < amount {
				return errors.New("insufficient balance")
			}
			accounts[i].Balance -= amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory, fmt.Sprintf("Withdrew: %.2f", amount))
			return nil
		}
	}

	return errors.New("account not found")
}

// DisplayTransactionHistory prints the transaction history of the account
func DisplayTransactionHistory(accountID int) error {
	for _, acc := range accounts {
		if acc.ID == accountID {
			fmt.Println("Transaction History for Account ID", accountID, ":")
			for _, transaction := range acc.TransactionHistory {
				fmt.Println(transaction)
			}
			return nil
		}
	}
	return errors.New("account not found")
}

// MenuSystem provides a menu-driven interface for the banking system
func MenuSystem() {
	for {
		fmt.Println("\n--- Banking System Menu ---")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Display Transaction History")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id int
			var amount float64
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Deposit Amount: ")
			fmt.Scanln(&amount)
			if err := Deposit(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful")
			}

		case 2:
			var id int
			var amount float64
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Withdrawal Amount: ")
			fmt.Scanln(&amount)
			if err := Withdraw(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful")
			}

		case 3:
			var id int
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			if err := DisplayTransactionHistory(id); err != nil {
				fmt.Println("Error:", err)
			}

		case 4:
			fmt.Println("Exiting Banking System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func main() {
	// Prepopulate some accounts
	accounts = append(accounts, Account{ID: 1, Name: "John Doe", Balance: 1000.00})
	accounts = append(accounts, Account{ID: 2, Name: "Jane Smith", Balance: 500.00})

	MenuSystem()
}
