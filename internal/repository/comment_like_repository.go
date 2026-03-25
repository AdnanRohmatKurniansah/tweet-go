package repository

import (
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"gorm.io/gorm"
)

type CommentLikeRepository interface {
	FindByCommentAndUser(commentId, userId uint) (*model.CommentLike, error)
	Create(like *model.CommentLike) error
	Delete(like *model.CommentLike) error
	Count(commentId uint) (int64, error)
}

type commentLikeRepository struct {
	db *gorm.DB
}

func NewCommentLikeRepository(db *gorm.DB) CommentLikeRepository {
	return &commentLikeRepository{db}
}

func (r *commentLikeRepository) FindByCommentAndUser(commentId, userId uint) (*model.CommentLike, error) {
	var like model.CommentLike
	err := r.db.Where("comment_id = ? AND user_id = ?", commentId, userId).First(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func (r *commentLikeRepository) Create(like *model.CommentLike) error {
	return r.db.Create(like).Error
}

func (r *commentLikeRepository) Delete(like *model.CommentLike) error {
	return r.db.Delete(like).Error
}

func (r *commentLikeRepository) Count(commentId uint) (int64, error) {
	var total int64
	err := r.db.Model(&model.CommentLike{}).
		Where("comment_id = ?", commentId).
		Count(&total).Error
	return total, err
}