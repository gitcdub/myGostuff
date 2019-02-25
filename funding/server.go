package funding

type FundServer struct {
	Commands chan interface{} // Lowercase name, unexported
	fund     *Fund
}

func (s *FundServer) Balance() int {
	responseChan := make(chan int)
	s.Commands <- BalanceCommand{Response: responseChan}
	return <-responseChan
}

func (s *FundServer) Withdraw(amount int) {
	s.Commands <- WithdrawCommand{Amount: amount}
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		// make () creates builtins like channels, maps and slices
		Commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	go server.loop()
	return server
}

type WithdrawCommand struct {
	Amount int
}
type BalanceCommand struct {
	Response chan int
}

//Typedef the callback for readability
type Transactor func(fund *Fund)

//Add a new command type with a callback and a semaphore channel
type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

//Wrap in an API method like the other commands
func (s *FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.Commands <- command
	<-command.Done
}

func (s *FundServer) loop() {
	for command := range s.Commands {
		switch command.(type) {

		case WithdrawCommand:
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)

		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			balance := s.fund.Balance()
			getBalance.Response <- balance

		case TransactionCommand:
			transaction := command.(TransactionCommand)
			transaction.Transactor(s.fund)
			transaction.Done <- true

			//default:
			//panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}
