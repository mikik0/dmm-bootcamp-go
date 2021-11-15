package statuses

import (
	"encoding/json"
	"errors"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Status string `json:"status"`
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	status := new(object.Status)
	status.Content = req.Status

	statusRepo := h.app.Dao.Status() // domain/repository の取得

	status, err := statusRepo.FindByID(ctx, status.ID)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if status != nil {
		httperror.BadRequest(w, errors.New("account name is already in use"))
		return
	}

	if err := statusRepo.CreateStatus(ctx, status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
