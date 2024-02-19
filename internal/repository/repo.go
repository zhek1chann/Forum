package repository

import "forum/models"

type UserRepo interface {
	CreateUser(*models.User) (int, error)
	GetUserByID(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	UpdateUserByID(string) (*models.User, error)
}

type PostRepo interface {
	CreatePost(*models.Post) error
	GetAllPost() (*models.Post, error)
	// UpdatePost(string, *models.Post) error
	AddLikeAndDislike(bool, string, string) error
	DeleteLikeAndDislike(int, int) error
	GetAllPostByUserID(int) ([]*models.Post, error)
	GetAllPostByCategories([]int) ([]*models.Post, error)

	GetAllPostPaginated(int, int) ([]*models.Post, error)
}

type CategoryRepo interface {
	AddCategoryToPost(int, []int) error
	GetALLCategory() (map[int]string, error)
	CreateCategory(string) error
}

type CommentRepo interface {
	CreateComment(*models.Comment) error
	GetAllCommentByPostID(string) (*[]models.Post, error)
	GetAllCommentByUserID(string) (*[]models.Post, error)
	AddLikeAndDislike(bool, string, string) error
}

type Storage interface {
	UserRepo
	PostRepo
	CategoryRepo
	CommentRepo
}
