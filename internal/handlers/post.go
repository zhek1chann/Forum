package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func (h *handler) postCreate(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, h.postCreateGet, h.postCreatePost)
}

func (h *handler) postCreateGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Create Get Page")
}

func (h *handler) postCreatePost(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) postView(w http.ResponseWriter, r *http.Request) {
	id, _ := strings.CutPrefix(r.URL.Path, "/post/")
	fmt.Fprintf(w, "View Post %s", id)
}
