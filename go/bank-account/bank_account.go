package account

import "sync"

// Account is bank account entity
type Account struct {
	balance  int64
	isClosed bool
	mu       sync.Mutex
}

// Open creates new bank account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	var a = &Account{isClosed: false}
	a.Deposit(initialDeposit)
	return a
}

// Balance returns current state of account
func (a *Account) Balance() (balance int64, ok bool) {
	// make sure that the balance is not changing at the moment
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.isClosed {
		return a.balance, false
	}
	return a.balance, true
}

// Deposit increments balance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	// committing changes to balance is blocking op
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.isClosed {
		return a.balance, false
	}
	if a.balance+amount < 0 {
		return a.balance, false
	}

	a.balance += amount
	return a.balance, true
}

// Close closes account
func (a *Account) Close() (payout int64, ok bool) {
	// closing account is blocking op
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.isClosed {
		return a.balance, false
	}

	// closing and committing changes to balance
	a.isClosed = true
	payout = a.balance
	a.balance = 0

	return payout, true
}
