package accounts

import (
	"encoding/json"
	"errors"
	"net/http"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `GET /v1/accounts/{username}`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		httperror.BadRequest(w, errors.New("username is required"))
		return
	}

	ctx := r.Context()
	accountRepo := h.app.Dao.Account() // domain/repository の取得

	user, err := accountRepo.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if user == nil {
		httperror.BadRequest(w, errors.New("account does not exist"))
		return
	}

	res := AddResponse{
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Avatar:      user.Avatar,
		Header:      user.Header,
		Note:        user.Note,
		CreateAt:    user.CreateAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
