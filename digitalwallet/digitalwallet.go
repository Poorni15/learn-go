package digitalwallet

import (
	"errors"
	"fmt"
	"sync"
)

type DigitalWallet struct {
	mu   sync.Mutex
	user []*User
}

func GetDigitalWallet() *DigitalWallet {
	return &DigitalWallet{}
}

func (digitalWallet *DigitalWallet) AddUser(user *User) {
	digitalWallet.user = append(digitalWallet.user, user)

}

func (digitalWallet *DigitalWallet) AddAccount(account *Account) {
	users := digitalWallet.user
	var foundUser *User
	for i := range users {
		if users[i].name == account.user.name {
			foundUser = users[i]
			break
		}
	}
	foundUser.AddAccount(account)
}

func (digitalWallet *DigitalWallet) TransferFund(from *Account, to *Account, amount float64) error {
	defer digitalWallet.mu.Unlock()
	digitalWallet.mu.Lock()
	if from.Balance() < amount {
		return errors.New("Insufficient Fund")
	}
	from.Withdraw(amount)
	to.Deposit(amount)
	return nil
}

func (digitalWallet *DigitalWallet) GetTransactionHistory(account *Account) {
	for _, val := range account.transHistory {
		fmt.Printf("Transaction ID: %d\n", val.transactionId)
		fmt.Printf("Amount: %v \n", val.amount)
		fmt.Printf("Timestamp: %v\n\n", val.time)
	}
}

func main() {
	digitalWallet := GetDigitalWallet()
	user1 := CreateUser("U001", "John Doe", "john@example.com", "password123")
	user2 := CreateUser("U002", "Jane Smith", "jane@example.com", "password456")
	digitalWallet.AddUser(user1)
	digitalWallet.AddUser(user2)
	account1 := CreateAccount("A001", user1, "1234567890")
	account2 := CreateAccount("A002", user2, "9876543210")
	digitalWallet.AddAccount(account1)
	digitalWallet.AddAccount(account2)
	// creditCard := NewCreditCard("PM001", user1, "1234567890123456", "12/25", "123")
	// bankAccount := NewBankAccount("PM002", user2, "9876543210", "987654321")

	account1.Deposit(float64(1000.00))
	account2.Deposit(float64(500.00))

	amount := float64(100.00)

	if err := digitalWallet.TransferFund(account1, account2, amount); err != nil {
		fmt.Printf("Transfer failed: %v\n", err)
	}

	digitalWallet.GetTransactionHistory(account1)
	digitalWallet.GetTransactionHistory(account2)
}
