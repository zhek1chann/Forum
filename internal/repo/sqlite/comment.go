package sqlite

import "forum/models"

func (s *Sqlite) CommentPost(form models.CommentForm) error {
	return nil
}

func (s *Sqlite) GetCommentByPostID() error {
	return nil
}

// like system

func (s *Sqlite) AddLikeComment() error {
	return nil
}

func (s *Sqlite) AddDisLikeComment() error {
	return nil
}

func (s *Sqlite) DeleteLikeStatusComment() error {
	return nil
}
