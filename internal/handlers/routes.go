package handlers

import "net/http"

func Routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/user", user)
	mux.HandleFunc("/post", post)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/logout", logout)

	return mux
}
