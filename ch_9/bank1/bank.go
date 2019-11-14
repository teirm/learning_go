// Package bank provides a concurrency-safe bank with one account
package bank

type WithdrawlInfo struct {
	amount int
	result chan bool
}

var deposits = make(chan int)             // send ammount to deposit
var balances = make(chan int)             // receive balance
var withdrawls = make(chan WithdrawlInfo) // receive withdrawls

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawls <- WithdrawlInfo{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balance is confiend to the tell goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case info := <-withdrawls:
			if balance-info.amount < 0 {
				info.result <- false
			} else {
				balance -= info.amount
				info.result <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
