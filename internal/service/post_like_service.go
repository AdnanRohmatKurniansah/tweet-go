package service

import (
	"errors"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"gorm.io/gorm"
)

type PostLikeService interface {
	LikeOrUnlike(postId, userId uint) (bool, int64, error)
	GetTotalLikes(postId uint) (int64, error)
}

type postLikeService struct {
	repo repository.PostLikeRepository
}

func NewPostLikeService(repo repository.PostLikeRepository) PostLikeService {
	return &postLikeService{repo}
}

func (s *postLikeService) GetTotalLikes(postId uint) (int64, error) {
	count, err := s.repo.Count(postId)
	if err != nil {
		return 0, err
	}

	return count, nil
}


func (s *postLikeService) LikeOrUnlike(postId, userId uint) (bool, int64, error) {
	existing, err := s.repo.FindByPostAndUser(postId, userId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, 0, err
	}

	if existing != nil {
		err := s.repo.Delete(existing)
		if err != nil {
			return false, 0, err
		}

		count, _ := s.repo.Count(postId)
		return false, count, nil
	}

	like := &model.PostLike{
		PostId: postId,
		UserId: userId,
	}

	err = s.repo.Create(like)
	if err != nil {
		return false, 0, err
	}

	count, _ := s.repo.Count(postId)

	return true, count, nil
}