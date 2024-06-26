package notification

import (
	"bytes"
	"desafio-pic-pay/internal/api/dtos/notifications"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MockNotifier struct {
	serviceUrl string
}

func NewMockNotifier() *MockNotifier {
	return &MockNotifier{
		serviceUrl: "https://util.devi.tools/api/v1/notify",
	}
}

func (mn *MockNotifier) Notify() error {
	var notifyResponse notifications.NotifyResponse

	resp, err := http.Post(
		mn.serviceUrl,
		"application/json",
		&bytes.Buffer{},
	)

	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("couldn't read response bytes due to: %e", err)
	}

	err = json.Unmarshal(body, &notifyResponse)
	if err != nil {
		return fmt.Errorf("couldn't parse response as JSON due to: %e", err)
	}

	return fmt.Errorf(
		"[CODE=%d | STATUS=%s] Could not notify user because of error in third party service: %s",
		resp.StatusCode, notifyResponse.Status, notifyResponse.Message,
	)
}
