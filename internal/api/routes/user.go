package routes

import (
	"context"
	"desafio-pic-pay/internal/api/dtos/user"
	"desafio-pic-pay/internal/api/errors"
	db "desafio-pic-pay/internal/storage/sqlc"
	"encoding/json"
	"io"
	"net/http"
)

// TODO: Make this functions accept: http.ResponseWriter, context.Context, db.Querier, http.Request
func HandleCreateNewUser(queries db.Querier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser user.CreateNewUser
		var apiError errors.IBaseError

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			apiError = errors.NewBaseError(err.Error(), "Couldn't read response body properly")
			WriteJSON(w, http.StatusUnprocessableEntity, apiError)
			return
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

		exists, err := queries.UserExists(context.Background(), db.UserExistsParams{
			Cpf:   newUser.CPF,
			Email: newUser.Email,
		})
		if err != nil {
			apiError = errors.NewBaseError(err.Error(), "Couldn't check user existence")
			WriteJSON(w, http.StatusInternalServerError, apiError)
			return
		}

		if exists > 0 {
			apiError = errors.NewBaseError("", "User already exists")
			WriteJSON(w, http.StatusConflict, apiError)
			return
		}

		dbUser, err := queries.CreateCommonUser(context.Background(), *newUser.ToDbArgs())

		if err != nil {
			apiError = errors.NewBaseError(err.Error(), "Couldn't create new user")
			WriteJSON(w, http.StatusInternalServerError, apiError)
			return
		}

		responseUser := user.NewUser(&dbUser)
		WriteJSON(w, http.StatusCreated, responseUser)
	}
}
