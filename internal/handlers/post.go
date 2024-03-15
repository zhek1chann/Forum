package handlers

import (
	"errors"
	"fmt"
	"forum/models"
	"forum/pkg/cookie"
	"forum/pkg/validator"
	"net/http"
	"strconv"
	"strings"
)

func (h *handler) postCreate(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, h.postCreateGet, h.postCreatePost)
}

func (h *handler) postCreateGet(w http.ResponseWriter, r *http.Request) {
	var err error
	data := h.app.NewTemplateData(r)

	data.Form = models.PostForm{}
	data.Categories, err = h.service.GetAllCategory()
	if err != nil {
		h.app.ServerError(w, err)
	}
	h.app.Render(w, http.StatusOK, "create.html", data)
}

func (h *handler) postCreatePost(w http.ResponseWriter, r *http.Request) {
	form := models.PostForm{
		Title:            r.FormValue("title"),
		Content:          r.FormValue("content"),
		CategoriesString: r.Form["categories"],
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.NotSelected(form.CategoriesString), "categories", "This field cannot be selected")
	form.CheckField(validator.IsError(form.ConverCategories()), "categories", "This field is incoreted")

	if !form.Valid() {
		data := h.app.NewTemplateData(r)
		data.Form = form
		categories, err := h.service.GetAllCategory()
		if err != nil {
			h.app.ServerError(w, err)
		}
		data.Categories = categories
		h.app.Render(w, http.StatusUnprocessableEntity, "create.html", data)
		return
	}
	cookie_ := cookie.GetSessionCookie(r)
	postID, err := h.service.CreatePost(form.Title, form.Content, cookie_.Value, form.Categories)
	if err != nil {
		h.app.ServerError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusSeeOther)
}

func (h *handler) postView(w http.ResponseWriter, r *http.Request) {
	id, _ := strings.CutPrefix(r.URL.Path, "/post/")
	ID, err := strconv.Atoi(id)
	if err != nil {
		h.app.ClientError(w, 400)
	}

	post, err := h.service.GetPostByID(ID)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			h.app.ClientError(w, 404)
		} else {
			h.app.ServerError(w, err)
		}
		return
	}

	data := h.app.NewTemplateData(r)
	data.Post = post
	data.Form = models.CommentForm{}
	data.Categories, err = h.service.GetAllCategory()
	if err != nil {
		h.app.ServerError(w, err)
	}
	h.app.Render(w, http.StatusOK, "post.html", data)
}
