package handlers

import "net/http"

func (h *handler) login(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, h.loginGet, h.loginPost)
}

func (h *handler) loginGet(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) loginPost(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) signup(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, h.signupGet, h.signupPost)
}

func (h *handler) signupGet(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) signupPost(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) logoutPost(w http.ResponseWriter, r *http.Request) {

}
