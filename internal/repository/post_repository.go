package repository

import (
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindAllPaginated(page, limit int) ([]model.Post, int64, error)
	FindById(id uint) (*model.Post, error)
	CreatePost(post *model.Post) error
	UpdatePost(post *model.Post) error
	DeletePost(post *model.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) CreatePost(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindAllPaginated(page, limit int) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	offset := (page - 1) * limit

	err := r.db.Model(&model.Post{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").
		Offset(offset).
		Limit(limit).
		Order("created_at desc").
		Find(&posts).Error

	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (r *postRepository) FindById(id uint) (*model.Post, error) {
	var post model.Post
	err := r.db.Preload("User").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) UpdatePost(post *model.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) DeletePost(post *model.Post) error {
	return r.db.Delete(post).Error
}
