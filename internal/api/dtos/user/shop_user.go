package user

import (
	db "desafio-pic-pay/internal/storage/sqlc"
	"fmt"
)

type ShopUser struct {
	User
	CNPJ string `json:"cnpj"`
}

func (u ShopUser) SendTransaction(_ ApiUser) error {
	return fmt.Errorf("shops cannot send transactions to another user")
}

func NewShopUser(u *db.CommonUser) ApiUser {
	return &ShopUser{
		CNPJ: u.Cnpj.String,
		User: User{
			Id:       int(u.ID),
			FullName: u.FullName,
			CPF:      u.Cpf,
			Email:    u.Email,
		},
	}
}
