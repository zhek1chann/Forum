package repository

import "forum/models"

type UserRepo interface {
	CreateUser(models.User) (int, error)
	GetUserByID(int) (models.User, error)
	GetUserByEmail(string) (models.User, error)
	UpdateUserByID(string) (models.User, error)
}

type PostRepo interface {
}

type CommentRepo interface {
}

type Storage interface {
}
