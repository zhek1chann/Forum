package handlers

import (
	"errors"
	"forum/models"
	"forum/pkg/cookie"
	"net/http"
	"strconv"
	"strings"
)

const (
	pageSize    = 5
	defaultPage = 1
)

func (h *handler) home(w http.ResponseWriter, r *http.Request) {
	data := h.app.NewTemplateData(r)
	data, err := h.setUpPage(data, r)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			h.app.NotFound(w)
			return
		} else {
			h.app.ServerError(w, err)
			return
		}
	}
	if data.Category_id == 0 {
		posts, err := h.service.GetAllPostPaginated(data.CurrentPage, pageSize)
		if err != nil {
			h.app.ServerError(w, err)
			return
		}

		data.Posts = posts
	} else {
		posts, err := h.service.GetAllPostByCategoryPaginated(data.CurrentPage, pageSize, data.Category_id)
		if err != nil {
			h.app.ServerError(w, err)
			return
		}
		data.Posts = posts
		// for _, value := range *posts {
		// 	fmt.Println(value)
		// }
	}
	token := cookie.GetSessionCookie(r)
	if token != nil {
		reactions, err := h.service.GetReactionPosts(token.Value)
		if err != nil {
			h.app.ServerError(w, err)
			return
		}
		data.Posts = h.service.IsLikedPost(data.Posts, reactions)
	}
	h.app.Render(w, http.StatusOK, "home.html", data)
	return
}

func (h *handler) setUpPage(data *models.TemplateData, r *http.Request) (*models.TemplateData, error) {
	var err error
	currentPageStr := r.URL.Query().Get("page")
	data.Category = strings.Title(r.URL.Query().Get("category"))
	// fmt.Print(data.Category)
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
	data.NumberOfPage, err = h.service.GetPageNumber(pageSize, data.Category_id)
	if err != nil {
		return nil, err
	}
	data.CurrentPage, err = strconv.Atoi(currentPageStr)
	if err != nil || data.CurrentPage < 1 || data.CurrentPage > data.NumberOfPage {
		data.CurrentPage = defaultPage
	}

	return data, nil
}

// func (h *handler) homePost(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		h.app.ClientError(w, http.StatusBadRequest)
// 		return
// 	}

// 	filterCategoriesString := r.Form["categories"]
// 	filterCategories, err := ConverCategories(filterCategoriesString)
// 	if err != nil {
// 		h.app.ClientError(w, http.StatusBadRequest)
// 	}
// 	posts, err := h.service.GetAllPostByCategories(filterCategories)
// 	if err != nil {
// 		h.app.ServerError(w, err)
// 	}
// 	data := h.app.NewTemplateData(r)

// 	data.Categories, err = h.service.GetAllCategory()
// 	if err != nil {
// 		h.app.ServerError(w, err)
// 	}

// 	data.Posts = posts
// 	h.app.Render(w, http.StatusOK, "home.html", data)
// }

func ConverCategories(CategoriesString []string) ([]int, error) {
	categories := make([]int, len(CategoriesString))
	for i, str := range CategoriesString {
		nb, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		categories[i] = nb
	}

	return categories, nil
}
