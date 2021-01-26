package model

//Transaction struct ...
type Transaction struct {
	Id           int
	Amount       float32
	Direction    string
	Reason       string
	Organization string
	Datetime     string
	Bill         string
}

//GetUserTransactions ...
func GetUserTransactions() (transactions []Transaction, err error) {
	transactions = []Transaction{
		{1, 28, "Расход", "Проезд", "ISET TRANSPORT", "11.12.2020 13:52", "Tinkoff"},
		{2, 28, "Расход", "Проезд", "ISET TRANSPORT", "11.12.2020 18:52", "Ubrir"},
	}
	return
}
