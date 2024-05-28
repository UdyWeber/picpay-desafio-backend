package transaction

type Transaction struct {
	Value float64 `json:"value"`
	Payer int     `json:"payer"`
	Payee int     `json:"payee"`
}
