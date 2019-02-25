package funding

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

func NewFund(initialBalance int) *Fund {
	return &Fund{
		balance: initialBalance,
	}
}

// Methods start with a receiver,  this sample uses Fund as the pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
