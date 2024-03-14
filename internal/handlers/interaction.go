package handlers

import (
	"fmt"
	"net/http"
)

func (h *handler) postReaction(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	if r.Method != http.MethodPost {
		h.app.ClientError(w, http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		h.app.ServerError(w, err)
		return
	}
	reaction := r.FormValue("reaction")
	switch reaction {
	case "true":
		fmt.Println("true")
	case "false":
		fmt.Println("false")
	default:
		h.app.ClientError(w, http.StatusBadRequest)
	}
}

func (h *handler) commentPost(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	if r.Method != http.MethodPost {
		h.app.ClientError(w, http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		h.app.ServerError(w, err)
		return
	}
	comment := r.FormValue("comment")
	postID := r.FormValue("postID")
	fmt.Println(comment)
	fmt.Println(comment)
}
