package pointers_and_errors

import "fmt"

type Bitcoin int
func (b Bitcoin) String() string { return fmt.Sprintf("%d BTC", b) }