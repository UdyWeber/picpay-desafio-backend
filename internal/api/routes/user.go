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
		WriteJSON(w, http.StatusUnprocessableEntity, apiError)
	}

	err = json.Unmarshal(bodyBytes, &newUser)
	if err != nil {
		apiError = errors.NewUnprocessableEntityError(
			err.Error(),
			"Couldn't parse JSON body",
			nil,
		)
		WriteJSON(w, http.StatusUnprocessableEntity, apiError)
		return
	}

	apiError = newUser.Validate()
	if apiError != nil {
		WriteJSON(w, http.StatusUnprocessableEntity, apiError)
		return
	}
}
