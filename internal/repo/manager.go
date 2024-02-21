package repo

import (
	"database/sql"
	"forum/models"
)

type UserRepo interface {
	CreateUser(*models.User) error
	GetUserByID(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	UpdateUserByID(string) (*models.User, error)
	Authenticate(email, password string) (int, error)
}

// type PostRepo interface {
// 	CreatePost(*models.Post) error
// 	GetAllPost() (*models.Post, error)
// 	// UpdatePost(string, *models.Post) error
// 	//AddLikeAndDislike(bool, string, string) error
// 	DeleteLikeAndDislike(int, int) error
// 	GetAllPostByUserID(int) ([]*models.Post, error)
// 	GetAllPostByCategories([]int) ([]*models.Post, error)
// 	GetAllPostPaginated(int, int) ([]*models.Post, error)
// }

// type CategoryRepo interface {
// 	AddCategoryToPost(int, []int) error
// 	GetALLCategory() (map[int]string, error)
// 	CreateCategory(string) error
// }

// type CommentRepo interface {
// 	CreateComment(*models.Comment) error
// 	GetAllCommentByPostID(string) (*[]models.Post, error)
// 	GetAllCommentByUserID(string) (*[]models.Post, error)
// 	AddLikeAndDislike(bool, string, string) error
// }

type RepoI interface {
	UserRepo
	//PostRepo
	// CategoryRepo
	// CommentRepo
}

type repo struct {
	db *sql.DB
}

func New(storagePath string) (RepoI, error) {
	db, err := NewDB(storagePath)
	if err != nil {
		return nil, err
	}
	return &repo{
		db: db,
	}, nil
}
