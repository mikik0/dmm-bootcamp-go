package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddResponse struct {
	Username    string          `json:"username"`
	DisplayName *string         `json:"display_name"`
	Avatar      *string         `json:"avatar"`
	Header      *string         `json:"header"`
	Note        *string         `json:"note"`
	CreateAt    object.DateTime `json:"create_at"`
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	accountRepo := h.app.Dao.Account() // domain/repository の取得

	if err := accountRepo.CreateAccount(ctx, account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	user, err := accountRepo.FindByUsername(ctx, account.Username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if user == nil {
		httperror.BadRequest(w, errors.New("account name is already in use"))
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
