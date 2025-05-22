package digitalwallet

type PaymentMethod interface {
	AddPaymentMethod()
}

type CreditCard struct {
	paymentId    string
	user         *User
	creditCardNo string
	validThroug  string
	cvv          string
}

func NewCreditCard(paymentId string, user *User, creditcardNo string, validThrough string, cvv string) *CreditCard {
	credit := CreditCard{paymentId: paymentId, user: user, creditCardNo: creditcardNo, validThroug: validThrough, cvv: cvv}
	return &credit
}

type BankAccount struct {
	paymentId     string
	user          *User
	accountNumber string
	ifscCode      string
}

func NewBankAccount(paymentId string, user *User, accountNo string, ifscCode string) *BankAccount {
	bankAccount := BankAccount{paymentId: paymentId, user: user, accountNumber: accountNo, ifscCode: ifscCode}
	return &bankAccount
}
