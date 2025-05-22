package digitalwallet

import "time"

type TransactionHistory struct {
	transactionId   int
	amount          float64
	time            time.Time
	transactionType string
}

func CreateTransactionHistory(tranactionId int, amount float64, transactionType string) *TransactionHistory {
	transactionHistory := TransactionHistory{transactionId: tranactionId, amount: amount, time: time.Now(), transactionType: transactionType}
	return &transactionHistory
}
