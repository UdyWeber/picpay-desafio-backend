package transaction

type Authorization struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Authorization bool `json:"authorization"`
}
