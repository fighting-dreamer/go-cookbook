package entity

type User struct {
	Name              string
	Email             string
	Contact           string
	Age               int32
	Wallet            int64
	TransactionListId string
	Address           Address
}
