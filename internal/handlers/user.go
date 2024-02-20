package handlers

import "net/http"

func login(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, loginGet, loginPost)
}

func loginGet(w http.ResponseWriter, r *http.Request) {

}
func loginPost(w http.ResponseWriter, r *http.Request) {

}

func signup(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, signupGet, signupPost)
}

func signupGet(w http.ResponseWriter, r *http.Request) {

}

func signupPost(w http.ResponseWriter, r *http.Request) {

}

func logoutPost(w http.ResponseWriter, r *http.Request) {

}
