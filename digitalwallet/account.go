package digitalwallet

import "math/rand"

const curr string = "INR"
const CREDIT_TRANSACTION = "CREDIT"
const DEBIT_TRANSACTION = "DEBIT"

type Account struct {
	accountNo    string
	user         *User
	accountId    string
	currency     string
	amount       float64
	transHistory []TransactionHistory
}

func CreateAccount(accountId string, user *User, accountNo string) *Account {
	account := Account{accountId: accountId, user: user, accountNo: accountNo, currency: curr}
	return &account
}

func (account *Account) Deposit(amount float64) {
	account.amount += amount
	account.transHistory = append(account.transHistory, *CreateTransactionHistory(rand.Int(), account.amount, CREDIT_TRANSACTION))
}

func (account *Account) Withdraw(amount float64) {
	account.amount -= amount
	account.transHistory = append(account.transHistory, *CreateTransactionHistory(rand.Int(), account.amount, DEBIT_TRANSACTION))
}

func (account Account) Balance() float64 {
	return float64(account.amount)
}

func (account Account) GetTransactionHistory() []TransactionHistory {
	return account.transHistory
}
