package handlers

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Page")
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Page")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sogin Page")
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Signup Page")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Page")
}
