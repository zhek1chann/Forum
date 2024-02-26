package cookie

import (
	"net/http"
	"time"
)

const cookieName = "session_id"

func GetSessionCookie(r *http.Request) *http.Cookie {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil
	}
	return cookie
}

func SetSessionCookie(w http.ResponseWriter, token string, expirationTime time.Time) {
	cookie := http.Cookie{
		Name:    cookieName,
		Value:   token,
		Path:    "/",
		Expires: expirationTime,
	}
	http.SetCookie(w, &cookie)
}
