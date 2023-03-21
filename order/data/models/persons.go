package models

type Persons struct {
	PersonID  int64  `json:"PersonID" db:"PersonID" `
	LastName  string `json:"LastName" db:"LastName" `
	FirstName string `json:"FirstName" db:"FirstName" `
	Address   string `json:"Address" db:"Address" `
	City      string `json:"City" db:"City" `
}
