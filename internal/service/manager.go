package service

import (
	"forum/internal/repo"
	"forum/models"
)

type service struct {
	repo repo.RepoI
}

type ServiceI interface {
	UserServiceI
}

type UserServiceI interface {
	GetUser(int) *models.User
	CreateUser(*models.User) error
}

func New(r repo.RepoI) ServiceI {
	return &service{
		r,
	}
}
