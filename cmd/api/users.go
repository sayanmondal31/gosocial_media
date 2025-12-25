package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 10)

	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	user, err := app.store.Users.Get(ctx, userId)
}
