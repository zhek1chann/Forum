package handlers

import (
	"fmt"
	"forum/pkg/cookie"
	"net/http"
)

type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")

// func decorator(){

// }
func methodResolver(w http.ResponseWriter, r *http.Request, get, post func(w http.ResponseWriter, r *http.Request)) {
	fmt.Println(r.URL)
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	default:
		//error
	}
}

func (h *handler) requireAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the user is not authenticated, redirect them to the login page and
		// return from the middleware chain so that no subsequent handlers in
		// the chain are executed.
		cookie := cookie.GetSessionCookie(r)
		if cookie == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		// And call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

// func (h *handler) authenticate(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		id := context.Context.Value()
// 		//	a:=app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
// 		if id == 0 {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		exists, err := app.users.Exists(id)
// 		if err != nil {
// 			app.serverError(w, err)
// 			return
// 		}
// 		if exists {
// 			ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
// 			r = r.WithContext(ctx)
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
