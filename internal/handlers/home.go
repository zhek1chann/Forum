package handlers

import (
	"fmt"
	"forum/models"
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
		h.app.ServerError(w, err)
	}
	posts, err := h.service.GetAllPostPaginated(data.CurrentPage, pageSize)
	if err != nil {
		h.app.ServerError(w, err)
	}

	data.Posts = posts
	h.app.Render(w, http.StatusOK, "home.html", data)
}

func (h *handler) category(w http.ResponseWriter, r *http.Request) {
	category := strings.TrimPrefix(r.URL.Path, "/category/")
	categories, err := h.service.GetAllCategory()
	if err != nil {
		h.app.ServerError(w, err)
	}
	var category_id int
	for key, value := range categories {
		if category == value {
			category_id = key + 1
			break
		}
	}
	fmt.Print(category)
	data := h.app.NewTemplateData(r)
	data, err = h.setUpPage(data, r)
	if err != nil {
		h.app.ServerError(w, err)
	}
	posts, err := h.service.GetAllPostByCategoryPaginated(data.CurrentPage, pageSize, category_id)
	if err != nil {
		h.app.ServerError(w, err)
	}

	data.Posts = posts
	h.app.Render(w, http.StatusOK, "home.html", data)
}

func (h *handler) setUpPage(data *models.TemplateData, r *http.Request) (*models.TemplateData, error) {
	currentPageStr := r.URL.Query().Get("page")
	var err error
	data.NumberOfPage, err = h.service.GetPageNumber(pageSize, data.Category)
	if err != nil {
		return nil, err
	}
	data.CurrentPage, err = strconv.Atoi(currentPageStr)
	if err != nil || data.CurrentPage < 1 || data.CurrentPage > data.NumberOfPage {
		data.CurrentPage = defaultPage
	}

	data.Categories, err = h.service.GetAllCategory()
	if err != nil {
		return nil, err
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
