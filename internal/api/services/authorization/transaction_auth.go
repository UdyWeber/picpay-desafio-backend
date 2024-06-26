package authorization

import (
	"desafio-pic-pay/internal/api/dtos/transaction"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TransactionAuthorizer struct {
	serviceUrl string
}

func NewTransactionAuthorizer() *TransactionAuthorizer {
	return &TransactionAuthorizer{
		"https://util.devi.tools/api/v2/authorize",
	}
}

func (ta *TransactionAuthorizer) getTransactionAuthorization() (*transaction.Authorization, error) {
	var authorization transaction.Authorization

	resp, err := http.Get(ta.serviceUrl)
	if err != nil {
		return &authorization, errors.New("Couldn't performe request because of: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &authorization, errors.New("Couldn't read response body because of: " + err.Error())
	}

	err = json.Unmarshal(body, &authorization)
	if err != nil {
		return &authorization, errors.New("Couldn't parse response body because of: " + err.Error())
	}

	return &authorization, nil
}

func (ta *TransactionAuthorizer) Authorize() error {
	transactionAuthorization, err := ta.getTransactionAuthorization()
	if err != nil {
		return err
	}

	if !transactionAuthorization.Data.Authorization {
		err = fmt.Errorf("[STATUS=%s | CODE=403] Transaction not authorized", transactionAuthorization.Status)
	}

	return err
}
