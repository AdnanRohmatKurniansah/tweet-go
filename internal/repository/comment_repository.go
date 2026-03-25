package repository

import (
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAllPaginated(postId, page, limit int) ([]model.Comment, int64, error)
	FindById(id uint) (*model.Comment, error)
	FindByPostAndUser(postId, userId uint) (*model.Comment, error)
	CreateComment(comment *model.Comment) error
	UpdateComment(comment *model.Comment) error
	DeleteComment(comment *model.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) FindAllPaginated(postId, page, limit int) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64

	offset := (page - 1) * limit

	err := r.db.Model(&model.Comment{}).Where("post_id = ?", postId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").
		Where("post_id = ?", postId).
		Offset(offset).
		Limit(limit).
		Order("created_at desc").
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

func (r *commentRepository) FindById(id uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.Preload("User").First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) FindByPostAndUser(postId, userId uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.Where("post_id = ? AND user_id = ?", postId, userId).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) CreateComment(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) UpdateComment(comment *model.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) DeleteComment(comment *model.Comment) error {
	return r.db.Delete(comment).Error
}	