package service

import (
	"errors"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/dto"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
	"gorm.io/gorm"
)

type CommentService interface {
	GetComments(postId, page, limit int) ([]model.Comment, int64, error)
	GetCommentById(id uint) (*model.Comment, error)
	CreateComment(req dto.CreateCommentRequest, userId uint) (*model.Comment, error)
	UpdateComment(id uint, req dto.UpdateCommentRequest, userId uint) (*model.Comment, error)
	DeleteComment(id uint, userId uint) (*model.Comment, error)
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) GetComments(postId, page, limit int) ([]model.Comment, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return s.repo.FindAllPaginated(postId, page, limit)
}

func (s *commentService) GetCommentById(id uint) (*model.Comment, error) {
	comment, err := s.repo.FindById(id)
	if err != nil {
		return nil, utils.ErrNotFound
	}
	return comment, nil
}

func (s *commentService) CreateComment(req dto.CreateCommentRequest, userId uint) (*model.Comment, error) {
	existing, err := s.repo.FindByPostAndUser(req.PostId, userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, utils.ErrAlreadyExists
	}

	comment := &model.Comment{
		PostId:  req.PostId,
		UserId:  userId,
		Content: req.Content,
	}

	if err := s.repo.CreateComment(comment); err != nil {
		return nil, err
	}

	return s.repo.FindById(comment.Id)
}

func (s *commentService) UpdateComment(id uint, req dto.UpdateCommentRequest, userId uint) (*model.Comment, error) {
	comment, err := s.repo.FindById(id)
	if err != nil {
		return nil, utils.ErrNotFound
	}

	if comment.UserId != userId {
		return nil, utils.ErrForbidden
	}

	if req.Content != "" {
		comment.Content = req.Content
	}

	if err := s.repo.UpdateComment(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) DeleteComment(id uint, userId uint) (*model.Comment, error) {
	comment, err := s.repo.FindById(id)
	if err != nil {
		return nil, utils.ErrNotFound
	}

	if comment.UserId != userId {
		return nil, utils.ErrForbidden
	}

	if err := s.repo.DeleteComment(comment); err != nil {
		return nil, err
	}

	return comment, nil
}