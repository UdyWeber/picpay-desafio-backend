package user

import (
	"desafio-pic-pay/internal/api/errors"
	"fmt"
	"strings"
)

type CreateNewUser struct {
	FullName string `json:"full_name"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CNPJ     string `json:"cnpj"`
}

func (u *CreateNewUser) Validate() *errors.UnprocessableEntityError {
	fields := make(map[string]string)

	trimmedName := strings.TrimSpace(u.FullName)
	if trimmedName == "" || len(trimmedName) < 10 {
		fields["full_name"] = "Full name is blank or lower than 10 characters"
	}

	trimmedCPF := strings.NewReplacer("-", "", ".", "").Replace(u.CPF)
	if trimmedCPF == "" || len(trimmedCPF) != 11 {
		fields["cpf"] = fmt.Sprintf("CPF=%s is not a valid cpf, must have 11 characters", u.CPF)
	}

	if strings.TrimSpace(u.Email) == "" {
		fields["email"] = "Email is required"
	}

	trimmedPass := strings.TrimSpace(u.Password)
	if trimmedPass == "" || len(trimmedPass) < 12 {
		fields["password"] = "Password is blank or too short, must be 12 characters long"
	}

	trimmedCNPJ := strings.NewReplacer(".", "", "-", "", "/", "").Replace(u.CNPJ)
	if !(trimmedCNPJ == "") {
		if len(trimmedCNPJ) != 14 {
			fields["cnpj"] = fmt.Sprintf("CNPJ=%s is not a valid cpnj, must have 14 characters", u.CNPJ)
		}
	}

	if len(fields) > 0 {
		return errors.NewUnprocessableEntityError(
			"",
			"Error while validating request body fields",
			fields,
		)
	}

	return nil
}
