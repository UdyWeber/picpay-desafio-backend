package authorization

type AuthService interface {
	Authorize() error
}
