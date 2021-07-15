package account

// Account is bank account entity
type Account struct {

}

// Open creates new bank account
func Open(initialDeposit int64) *Account {
	return &Account{}
}

func (*Account) Balance() {

}