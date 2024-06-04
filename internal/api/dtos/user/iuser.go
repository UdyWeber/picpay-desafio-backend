package user

type IUser interface {
	SendTransaction(IUser) error
}
