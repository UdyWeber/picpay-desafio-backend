package user

import db "desafio-pic-pay/internal/storage/sqlc"

type User struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
}

func (u *User) SendTransaction(to ApiUser) error {
	return nil
}

func NewUser(u *db.CommonUser) ApiUser {
	if u.Cnpj.Valid {
		return NewShopUser(u)
	}

	return &User{
		Id:       int(u.ID),
		FullName: u.FullName,
		CPF:      u.Cpf,
		Email:    u.Email,
	}
}
