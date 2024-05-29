package routes

import (
	"desafio-pic-pay/internal/api/dtos/user"
	"desafio-pic-pay/internal/api/errors"
	"encoding/json"
	"io"
	"net/http"
)

func HandleCreateNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser user.CreateNewUser
	var apiError errors.IBaseError

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		apiError = errors.NewBaseError(err.Error(), "Couldn't read response body properly")
		w.Write(apiError.ToResponse())
	}

	err = json.Unmarshal(bodyBytes, &newUser)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		apiError = errors.NewUnprocessableEntityError(
			err.Error(),
			"Couldn't parse JSON body",
			nil,
		)
		w.Write(apiError.ToResponse())
		return
	}

	apiError = newUser.Validate()
	if apiError != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(apiError.ToResponse())
		return
	}
}
