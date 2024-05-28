package user

type CommonUser struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
