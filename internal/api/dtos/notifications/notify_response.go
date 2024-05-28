package notifications

type NotifyResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
