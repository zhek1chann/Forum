package service

import "forum/models"

func (s *service) GetUser(id int) *models.User {
	return nil
}

func (s *service) CreateUser(user *models.User) error {
	err := s.repo.CreateUser(user)
	return err
}
