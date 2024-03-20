package models

import (
	"forum/pkg/validator"
	"strconv"
	"time"
)

type Post struct {
	PostID     int
	UserID     int
	UserName   string
	Title      string
	Content    string
	ImageName  string
	Created    time.Time
	Like       int
	Dislike    int
	Comment    *[]Comment
	Categories map[int]string
	IsLiked    int
}

type Comment struct {
	CommentID int
	PostID    int
	UserID    int
	UserName  string
	Content   string
	Created   time.Time
	Like      string
	Dislike   string
}

type CommentForm struct {
	PostID  string
	UserID  string
	Content string
	validator.Validator
}

type ReactionForm struct {
	ID       string
	UserID   string
	Reaction bool
}

type PostForm struct {
	Title               string   `form:"title"`
	Content             string   `form:"content"`
	Categories          []int    `form:"category"`
	CategoriesString    []string `form:"category"`
	validator.Validator `form:"-"`
}

func (f *PostForm) ConverCategories() error {
	for _, str := range f.CategoriesString {
		nb, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		f.Categories = append(f.Categories, nb)
	}
	return nil
}
