package user

type ShopkeeperUser struct {
	CommonUser
	CNPJ string `json:"cnpj"`
}
