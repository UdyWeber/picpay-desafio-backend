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

func HandleCreateNewUser(w http.ResponseWriter, r *http.Request, queries *db.Queries, _ context.Context) *errors.APIErrorWrapper {
	var newUser user.CreateNewUser

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.NewErrorWrapper(
			http.StatusUnprocessableEntity,
			errors.NewUnprocessableEntityError(err.Error(), "Couldn't read response body properly", nil),
		)
	}

	if err = json.Unmarshal(bodyBytes, &newUser); err != nil {
		return errors.NewErrorWrapper(
			http.StatusUnprocessableEntity,
			errors.NewUnprocessableEntityError(
				err.Error(),
				"Couldn't parse JSON body",
				nil,
			),
		)
	}

	if err = newUser.Validate(); err != nil {
		return errors.NewErrorWrapper(
			http.StatusUnprocessableEntity,
			err,
		)
	}

	exists, err := queries.UserExists(context.Background(), db.UserExistsParams{
		Cpf:   newUser.CPF,
		Email: newUser.Email,
	})

	if err != nil {
		return errors.NewErrorWrapper(
			http.StatusInternalServerError,
			errors.NewBaseError(err.Error(), "Couldn't check user existence"),
		)
	}

	if exists > 0 {
		return errors.NewErrorWrapper(
			http.StatusConflict,
			errors.NewBaseError("", "User already exists"),
		)

	}

	dbUser, err := queries.CreateCommonUser(context.Background(), *newUser.ToDbArgs())

	if err != nil {
		return errors.NewErrorWrapper(
			http.StatusInternalServerError,
			errors.NewBaseError(err.Error(), "Couldn't create new user"),
		)
	}

	responseUser := user.NewUser(&dbUser)

	if err = WriteJSON(w, http.StatusCreated, responseUser); err != nil {
		return errors.NewErrorWrapper(
			http.StatusInternalServerError,
			errors.NewBaseError(err.Error(), "Couldn't write response"),
		)
	}

	return nil
}
