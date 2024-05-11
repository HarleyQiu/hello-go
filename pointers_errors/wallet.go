package pointers_errors

type Waller struct{}

func (w Waller) Deposit(amount int) {

}

func (w Waller) Balance() int {
	return 0
}
