package notification

type NotifierService interface {
	Notify() error
}
