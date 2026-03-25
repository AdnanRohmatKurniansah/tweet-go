package service

import (
	"errors"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"gorm.io/gorm"
)

type CommentLikeService interface {
	LikeOrUnlike(commentId, userId uint) (bool, int64, error)
	GetTotalLikes(commentId uint) (int64, error)
}

type commentLikeService struct {
	repo repository.CommentLikeRepository
}

func NewCommentLikeService(repo repository.CommentLikeRepository) CommentLikeService {
	return &commentLikeService{repo}
}

func (s *commentLikeService) GetTotalLikes(commentId uint) (int64, error) {
	count, err := s.repo.Count(commentId)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *commentLikeService) LikeOrUnlike(commentId, userId uint) (bool, int64, error) {
	existing, err := s.repo.FindByCommentAndUser(commentId, userId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, 0, err
	}

	if existing != nil {
		err := s.repo.Delete(existing)
		if err != nil {
			return false, 0, err
		}

		count, _ := s.repo.Count(commentId)
		return false, count, nil
	}

	like := &model.CommentLike{
		CommentId: commentId,
		UserId: userId,
	}

	err = s.repo.Create(like)
	if err != nil {
		return false, 0, err
	}

	count, _ := s.repo.Count(commentId)

	return true, count, nil
}