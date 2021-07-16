package account

// Account is bank account entity
type Account struct {
}

// Open creates new bank account
func Open(initialDeposit int64) *Account {
	return &Account{}
}

// Balance returns current state of account
func (*Account) Balance() (balance int64, ok bool) {
	return 0, true
}

// Deposit increments balance
func (*Account) Deposit(amount int64) (newBalance int64, ok bool) {
	return 0, true
}

// Close closes account
func (*Account) Close() (payout int64, ok bool) {
	return 0, true
}
