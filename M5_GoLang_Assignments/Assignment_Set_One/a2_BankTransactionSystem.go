package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	ID                 int
	Name               string
	Balance            float64
	TransactionHistory []string
}

const (
	DespostOption                = 1
	WithdrawOption               = 2
	ViewBalanceOption            = 3
	ViewTransactionHistoryOption = 4
	Exit                         = 5
)

var BankAccounts []BankAccount

func FindAccount(id int) (*BankAccount, error) {
	for i := range BankAccounts {
		if BankAccounts[i].ID == id {
			return &BankAccounts[i], nil
		}
	}
	return nil, errors.New("Account not found")
}

func Deposit(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("Amount should be greater than 0")
	}
	account, err := FindAccount(id)
	if err != nil {
		return err
	}
	account.Balance += amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Deposited %.2f", amount))
	return nil
}

func Withdraw(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("Amount should be greater than 0")
	}

	account, err := FindAccount(id)
	if err != nil {
		return err
	}

	if account.Balance < amount {
		return errors.New("Insufficient balance")
	}
	account.Balance -= amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Withdrawn %.2f", amount))
	return nil
}

func ViewBalance(id int) (float64, error) {
	account, err := FindAccount(id)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

func ViewTransactionHistory(id int) error {
	account, err := FindAccount(id)
	if err != nil {
		return err
	}
	fmt.Println("Transaction History for Account ID: ", account.ID)
	for _, transaction := range account.TransactionHistory {
		fmt.Println(transaction)
	}
	return nil
}

// func main() {
// 	BankAccounts = append(BankAccounts, BankAccount{ID: 1, Name: "John Doe", Balance: 1000.0})
// 	BankAccounts = append(BankAccounts, BankAccount{ID: 2, Name: "Jane Doe", Balance: 2000.0})

// 	var choice int
// 	for {
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Deposit")
// 		fmt.Println("2. Withdraw")
// 		fmt.Println("3. View Balance")
// 		fmt.Println("4. View Transaction History")
// 		fmt.Println("5. Exit")
// 		fmt.Print("Enter your choice: ")
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case DespostOption:
// 			var id int
// 			var amount float64
// 			fmt.Print("Enter Account ID: ")
// 			fmt.Scanln(&id)
// 			fmt.Print("Enter Amount: ")
// 			fmt.Scanln(&amount)
// 			err := Deposit(id, amount)
// 			if err != nil {
// 				fmt.Println("Error: ", err)
// 			} else {
// 				fmt.Println("Amount deposited successfully")
// 			}
// 		case WithdrawOption:
// 			var id int
// 			var amount float64
// 			fmt.Print("Enter Account ID: ")
// 			fmt.Scanln(&id)
// 			fmt.Print("Enter Amount: ")
// 			fmt.Scanln(&amount)
// 			err := Withdraw(id, amount)
// 			if err != nil {
// 				fmt.Println("Error: ", err)
// 			} else {
// 				fmt.Println("Amount withdrawn successfully")
// 			}
// 		case ViewBalanceOption:
// 			var id int
// 			fmt.Print("Enter Account ID: ")
// 			fmt.Scanln(&id)
// 			balance, err := ViewBalance(id)
// 			if err != nil {
// 				fmt.Println("Error: ", err)
// 			} else {
// 				fmt.Printf("\nBalance: %.2f\n", balance)
// 			}
// 		case ViewTransactionHistoryOption:
// 			var id int
// 			fmt.Print("Enter Account ID: ")
// 			fmt.Scanln(&id)
// 			err := ViewTransactionHistory(id)
// 			if err != nil {
// 				fmt.Println("Error: ", err)
// 			}
// 		case Exit:
// 			fmt.Println("Exiting...")
// 			return
// 		default:
// 			fmt.Println("Invalid choice")
// 		}
// 		if choice == Exit {
// 			break
// 		}
// 	}
// }
