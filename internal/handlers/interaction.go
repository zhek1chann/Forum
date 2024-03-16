package handlers

import (
	"errors"
	"fmt"
	"forum/models"
	"forum/pkg/cookie"
	"forum/pkg/validator"
	"net/http"
	"strconv"
)

func (h *handler) postReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.app.ClientError(w, http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		h.app.ServerError(w, err)
		return
	}
	url := r.FormValue("url")
	token := cookie.GetSessionCookie(r)
	form := models.PostReactionForm{
		PostID: r.FormValue("postID"),
		UserID: token.Value,
	}
	reaction := r.FormValue("reaction")

	switch reaction {
	case "true":
		form.Reaction = true
	case "false":
		form.Reaction = false
	default:
		h.app.ClientError(w, http.StatusBadRequest)
		return
	}
	err := h.service.PostReaction(form)
	if err != nil {
		h.app.ServerError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf(url), http.StatusSeeOther)
}

func (h *handler) commentPost(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodPost {
		h.app.ClientError(w, http.StatusBadRequest)
		return
	}
	token := cookie.GetSessionCookie(r)
	form := models.CommentForm{
		Content: r.FormValue("comment"),
		PostID:  r.FormValue("postID"),
		UserID:  token.Value,
	}

	form.CheckField(validator.NotBlank(form.Content), "comment", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Content, 2), "comment", "This field must be at least 2 characters long")
	form.CheckField(validator.MaxChars(form.Content, 50), "comment", "This field must be maximum 50 characters")

	if !form.Valid() {
		data := h.app.NewTemplateData(r)
		data.Form = form
		data.Categories, err = h.service.GetAllCategory()
		if err != nil {
			h.app.ServerError(w, err)
		}
		id, err := strconv.Atoi(form.PostID)
		if err!=nil{
			h.app.ServerError(w, err)
			return
		}
		post, err := h.service.GetPostByID(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				h.app.ClientError(w, 404)
				return
			} else {
				h.app.ServerError(w, err)
				return
			}
		}
		data.Post = post
		h.app.Render(w, http.StatusUnprocessableEntity, "post.html", data)
		return
	}

	err = h.service.CommentPost(form)
	if err != nil {
		h.app.ServerError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/post/%s", form.PostID), http.StatusSeeOther)
}
