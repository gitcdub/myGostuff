package funding

type Fund struct {
	// balance is lowercase and not exported
	balance int
}

func NewFund(initialBalance int) *Fund {
	/* returning pointer to new struct without worry about heaps/stacks, Go
	will work it out
	*/
	return &Fund{
		balance: initialBalance,
	}
}

// Methods start with a *receiver* in this example it's the Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
