package handlers

import (
	"net/http"
)

func (h *handler) home(w http.ResponseWriter, r *http.Request) {
	data := h.app.NewTemplateData(r)
	h.app.Render(w, http.StatusOK, "home.html", data)
}
