package digitalwallet

type User struct {
	userId   string
	name     string
	emailId  string
	password string
	account  []*Account
}

func CreateUser(userId string, name string, emailId string, password string) *User {
	newUser := User{userId: userId, name: name, emailId: emailId, password: password}
	return &newUser
}

func (user *User) AddAccount(account *Account) {
	user.account = append(user.account, account)
}
