package handlers

import (
	"fmt"
	"forum/models"
	"forum/pkg/cookie"
	"forum/pkg/validator"
	"net/http"
)

func (h *handler) postReaction(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	if r.Method != http.MethodPost {
		h.app.ClientError(w, http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		h.app.ServerError(w, err)
		return
	}
	reaction := r.FormValue("reaction")
	switch reaction {
	case "true":
		fmt.Println("true")
	case "false":
		fmt.Println("false")
	default:
		h.app.ClientError(w, http.StatusBadRequest)
	}
}

func (h *handler) commentPost(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Print(r.Method)
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
		h.app.Render(w, http.StatusUnprocessableEntity, "post.html", data)
	}
	err = h.service.CommentPost(form)
	if err != nil {
		h.app.ServerError(w, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/post/%d", form.PostID), http.StatusSeeOther)
}
