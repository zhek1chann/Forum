package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func postCreate(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, postCreateGet, postCreatePost)
}

func postCreateGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Create Get Page")
}

func postCreatePost(w http.ResponseWriter, r *http.Request) {

}

func postView(w http.ResponseWriter, r *http.Request) {
	id, _ := strings.CutPrefix(r.URL.Path, "/post/")
	fmt.Fprintf(w, "View Post %s", id)
}
