package account

import "fmt"

type Account struct {
	FirstName string
	LastName  string
}

func (account *Account) ChangeName(firstName string, lastName string) {
	account.FirstName = firstName
	account.LastName = lastName
}

func (account Account) String() string {
	return fmt.Sprintf("%s, %s", account.LastName, account.FirstName)
}

type Employee struct {
	SourceAccount Account
	Credits       float32
}

func (emp *Employee) AddCredits(val float32) {
	emp.Credits += val
}

func (emp *Employee) RemoveCredits(val float32) {
	emp.Credits -= val
}

func (emp *Employee) CheckCredits(val float32) float32 {
	return emp.Credits
}
