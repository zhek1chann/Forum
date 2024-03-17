package service

import (
	"forum/models"
	"strconv"
)

func (s *service) CommentPost(form models.CommentForm) error {
	userIDint, err := s.repo.GetUserIDByToken(form.UserID)
	if err != nil {
		return err
	}
	form.UserID = strconv.Itoa(userIDint)
	return s.repo.CommentPost(form)

}



func (s *service) PostReaction(form models.ReactionForm) error {
	userIDint, err := s.repo.GetUserIDByToken(form.UserID)
	if err != nil {
		return err
	}
	form.UserID = strconv.Itoa(userIDint)
	exists, isLike, err := s.repo.CheckReactionPost(form)
	if err != nil {
		return err
	}
	if exists {
		err := s.repo.DeleteReactionPost(form, isLike)
		if err != nil {
			return err
		}
		if isLike == form.Reaction {
			return nil
		}
	}

	err = s.repo.AddReactionPost(form)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CommentReaction(form models.ReactionForm) error {
	userIDint, err := s.repo.GetUserIDByToken(form.UserID)
	if err != nil {
		return err
	}
	form.UserID = strconv.Itoa(userIDint)
	exists, isLike, err := s.repo.CheckReactionComment(form)
	if err != nil {
		return err
	}
	if exists {
		err := s.repo.DeleteReactionComment(form, isLike)
		if err != nil {
			return err
		}
		if isLike == form.Reaction {
			return nil
		}
	}

	err = s.repo.AddReactionComment(form)
	if err != nil {
		return err
	}

	return nil
}
