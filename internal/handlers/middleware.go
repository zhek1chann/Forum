package handlers

import (
	"fmt"
	"forum/models"
	"forum/pkg/cookie"
	"net/http"
	"strconv"
	"strings"
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
		// error
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

func GetIntForm(r *http.Request, form string) (int, error) {
	valueString := r.FormValue(form)
	value, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (h *handler) NewTemplateData(r *http.Request) (*models.TemplateData, error) {
	var TemplateData models.TemplateData

	TemplateData.IsAuthenticated = h.isAuthenticated(r)

	if TemplateData.IsAuthenticated {
		user, err := h.service.GetUser(r)
		if err != nil {
			return nil, err
		}
		TemplateData.User = user
	}
	return &TemplateData, nil
}

func (h *handler) isAuthenticated(r *http.Request) bool {
	cookie := cookie.GetSessionCookie(r)
	return cookie != nil && cookie.Value != ""
}

func (h *handler) setUpPage(data *models.TemplateData, r *http.Request) (*models.TemplateData, error) {
	var err error
	currentPageStr := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	data.Category = strings.Title(r.URL.Query().Get("category"))
	if len(data.Category) == 0 {
		data.Category = r.FormValue("category")
	}

	data.Categories, err = h.service.GetAllCategory()
	if err != nil {
		return nil, err
	}
	if data.Category != "" {
		for key, value := range data.Categories {
			if data.Category == value {
				data.Category_id = key + 1
				break
			}
		}
		if data.Category_id == 0 {
			return nil, models.ErrNoRecord
		}
	}
	data.Limit, err = strconv.Atoi(limit)
	if err != nil || data.Limit < 1 {
		data.Limit = pageSize
	}
	data.NumberOfPage, err = h.service.GetPageNumber(data.Limit, data.Category_id)
	if err != nil {
		return nil, err
	}
	data.CurrentPage, err = strconv.Atoi(currentPageStr)
	if err != nil || data.CurrentPage < 1 || data.CurrentPage > data.NumberOfPage {
		data.CurrentPage = defaultPage
	}

	return data, nil
}
