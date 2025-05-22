package digitalwallet

import "time"

type TransactionHistory struct {
	transactionId int
	amount        float64
	time          time.Time
}

func CreateTransactionHistory(tranactionId int, amount float64) *TransactionHistory {
	transactionHistory := TransactionHistory{transactionId: tranactionId, amount: amount, time: time.Now()}
	return &transactionHistory
}
