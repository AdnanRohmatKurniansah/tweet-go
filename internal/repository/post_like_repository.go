package repository

import (
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"gorm.io/gorm"
)

type PostLikeRepository interface {
	FindByPostAndUser(postId, userId uint) (*model.PostLike, error)
	Create(like *model.PostLike) error
	Delete(like *model.PostLike) error
	Count(postId uint) (int64, error)
}

type postLikeRepository struct {
	db *gorm.DB
}

func NewPostLikeRepository(db *gorm.DB) PostLikeRepository {
	return &postLikeRepository{db}
}

func (r *postLikeRepository) FindByPostAndUser(postId, userId uint) (*model.PostLike, error) {
	var like model.PostLike
	err := r.db.Where("post_id = ? AND user_id = ?", postId, userId).First(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func (r *postLikeRepository) Create(like *model.PostLike) error {
	return r.db.Create(like).Error
}

func (r *postLikeRepository) Delete(like *model.PostLike) error {
	return r.db.Delete(like).Error
}

func (r *postLikeRepository) Count(postId uint) (int64, error) {
	var total int64
	err := r.db.Model(&model.PostLike{}).
		Where("post_id = ?", postId).
		Count(&total).Error
	return total, err
}