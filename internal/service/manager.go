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
	CategoryServiceI
	PostServiceI
	InteractionServiceI
}

type InteractionServiceI interface {
	CommentPost(models.CommentForm) error
}

type UserServiceI interface {
	GetUser(int) *models.User
	CreateUser(models.User) error
	Authenticate(string, string) (*models.Session, error)
	DeleteSession(string) error
}

type PostServiceI interface {
	CreatePost(string, string, string, []int) (int, error)
	GetPostByID(int) (*models.Post, error)
	GetAllPostPaginated(int, int) (*[]models.Post, error)
	GetAllPostByCategoryPaginated(curentPage, pageSize, category int) (*[]models.Post, error)
	GetPageNumber(int, int) (int, error)
	GetAllPostByCategory(category int) (*[]models.Post, error)
	GetAllPostByUser(token string) (*[]models.Post, error)
}

type CategoryServiceI interface {
	GetAllCategory() ([]string, error)
}

func New(r repo.RepoI) ServiceI {
	return &service{
		r,
	}
}
