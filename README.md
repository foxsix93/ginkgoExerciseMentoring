# ginkgoExerciseMentoring
Package, Import, Error Handling, and Unit Testing Problem
Problem: Bank Account Management System
Submission
Create github repository
Send link to form google (TBD)
Deadline 15 Nov 2024 - 19.00
Description
Create a simple bank account management system with two main functionalities: depositing and withdrawing money. Implement this in a custom package named bank. Use error handling to ensure that the balance does not go negative and that invalid amounts (like zero or negative values) cannot be deposited or withdrawn. Write unit tests to verify the correctness of your code.

Requirements
Create a package bank with a struct Account that has a Balance field.
Implement the following methods on Account:
Deposit(amount float64) error: Increases the balance by the given amount. Returns an error if the amount is zero or negative.
Withdraw(amount float64) error: Decreases the balance by the given amount. Returns an error if the amount is zero, negative, or greater than the current balance.
GetBalance() float64: Returns the current balance.
Use error handling to manage invalid transactions.
In the main package, import the bank package and create a few test cases for deposits and withdrawals to verify the functionality.
Write unit tests using Ginkgo for the Deposit, Withdraw, and GetBalance methods.
Example Usage
package main

import (
    "fmt"
    "bank"
)

func main() {
    account := bank.Account{}
    
    // Deposit money
    if err := account.Deposit(100.50); err != nil {
        fmt.Println("Error:", err)
    }
    
    // Attempt to withdraw money
    if err := account.Withdraw(50); err != nil {
        fmt.Println("Error:", err)
    }
    
    // Attempt to withdraw more than the balance
    if err := account.Withdraw(100); err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Println("Final Balance:", account.GetBalance())
}
Expected Output
Error: insufficient funds
Final Balance: 50.5
Example Unit Test (in bank_test.go)
func TestDeposit(t *testing.T) {
    account := bank.Account{}
    err := account.Deposit(100)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if account.GetBalance() != 100 {
        t.Errorf("Expected balance of 100, got %v", account.GetBalance())
    }
}
This problem demonstrates creating a custom package, handling errors, and writing unit tests to ensure functionality and proper error management in the bank account management system.