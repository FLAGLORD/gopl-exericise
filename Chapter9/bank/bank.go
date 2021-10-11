package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan withDrawMessage)

type withDrawMessage struct {
	amount int
	ok     chan bool
}

func Deposit(amount int) {
	deposits <- amount
}
func Balance() int {
	return <-balances
}

func WithDraw(amount int) bool {
	ok := make(chan bool)
	withdraw <- withDrawMessage{
		amount: amount,
		ok:     ok,
	}
	return <-ok
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// Do nothing:
		case withdraw := <-withdraw:
			if balance >= withdraw.amount {
				balance -= withdraw.amount
				withdraw.ok <- true
			}
			withdraw.ok <- false
		}
	}
}

func init() {
	go teller()
}
