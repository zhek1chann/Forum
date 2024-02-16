package models

import "time"

type Post struct {
	PostID        int64
	CreatedUserID int64
	Title         string
	Content       string
	CreatedTime   time.Time
	LikeCounter    string
	DislikeCounter string
}

type Comment struct {
	CommentId      int64
	PostID         int64
	CreatedUserID  int64
	Content        string
	CreatedTime    time.Time
	LikeCounter    string
	DislikeCounter string
}
