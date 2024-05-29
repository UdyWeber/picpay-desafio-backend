package user

type User struct {
    Id       int    `json:"id"`
    FullName string `json:"full_name"`
    CPF      string `json:"cpf"`
    Email    string `json:"email"`
    CNPJ     string `json:"cnpj"`
}
