package handlers

import "net/http"

func Routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/post/", postView)
	mux.HandleFunc("/post/create", postCreate)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/logout", logoutPost)

	return mux
}
