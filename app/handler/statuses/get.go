package statuses

// Get handles the GET: /v1/statuses/{id}.
// func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	if err != nil {
// 		httperror.Response(w, http.StatusBadRequest, "invalid id")
// 		return
// 	}

// 	status, err := h.statusRepo.Get(id)
// 	if err != nil {
// 		httperror.Response(w, http.StatusInternalServerError, "failed to get status")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(status)
// }
