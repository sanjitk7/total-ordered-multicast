package main

import (
	"errors"
	"fmt"
	"sort"
)

// we wrap our map around with a struct for encapsulation and good programming practice - methods associated with the accounts map would be in the class
type AccountStore struct {
	accounts map[string]float64
}

// init AccountStore
func NewAccountStore() *AccountStore {
	return &AccountStore{
		accounts: make(map[string]float64),
	}
}

// deposit money
func (as *AccountStore) Deposit(accountName string, depositAmount float64) {
	if depositAmount <= 0 {
		fmt.Println("Deposit Failed: Deposit Value is -ve")
	}

	as.accounts[accountName] += depositAmount // this already handles the edge case where the account doesnt exist - works like python default dict
}

// get balance

func (as *AccountStore) GetBalance(accountName string) {
	val, exists := as.accounts[accountName]

	if exists {
		fmt.Printf("%s balance : %f\n", accountName, val)
	} else {
		fmt.Printf("Value Error: %s does not exist")
	}

}

// get all accounts balance alphabetically

func (as *AccountStore) GetAllBalances() {
	acc := make([]string, 0, len(as.accounts))

	for account := range as.accounts {
		acc = append(acc, account)
	}

	sort.Strings(acc)

	for _, account := range acc {
		fmt.Printf("%s: %f\n", account, as.accounts[account])
	}

}

// transfer money

func (as *AccountStore) Transfer(fromAccount string, toAccount string, amount float64) error {
	if amount <= 0 {
		return errors.New("Transfer Failed: Transfer Value -ve")
	}
	if as.accounts[fromAccount] < amount {
		return errors.New("Transfer Failed: Not enough funds in source account")
	}

	as.accounts[fromAccount] -= amount
	as.accounts[toAccount] += amount

	return nil
}

func main() {
	curAccountStore := NewAccountStore()
	curAccountStore.Deposit("A", 100)
	curAccountStore.Deposit("B", 200)
	curAccountStore.Deposit("C", 100)

	curAccountStore.Transfer("B", "A", 10)
	curAccountStore.Transfer("B", "C", 20)

	curAccountStore.GetAllBalances()

}
