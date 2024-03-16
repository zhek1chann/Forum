package service

import (
	"fmt"
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

func (s *service) PostReaction(form models.PostReactionForm) error {
	userIDint, err := s.repo.GetUserIDByToken(form.UserID)
	if err != nil {
		return err
	}
	form.UserID = strconv.Itoa(userIDint)
	exists, isLike, err := s.repo.CheckReactionPost(form)
	if err != nil {
		fmt.Print("1")
		return err
	}
	if exists {
		err := s.repo.DeleteReactionPost(form, isLike)
		if err != nil {
			fmt.Print("2")
			return err
		}
		if isLike == form.Reaction {
			return nil
		}
	}

	err = s.repo.AddReactionPost(form)
	if err != nil {
		fmt.Print("3")
		return err
	}

	return nil
}
