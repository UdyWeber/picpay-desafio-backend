package user

type ApiUser interface {
	SendTransaction(ApiUser) error
}
