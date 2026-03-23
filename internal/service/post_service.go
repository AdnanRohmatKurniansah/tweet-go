package service

import (
	"errors"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/dto"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
)

type PostService interface {
	GetPosts(page, limit int) ([]model.Post, int64, error)
	GetPostById(id uint) (*model.Post, error)
	CreatePost(req dto.CreatePostRequest, userId uint, imageUrl string) (*model.Post, error)
	UpdatePost(id uint, req dto.UpdatePostRequest, newImageUrl string) (*model.Post, error)
	DeletePost(id uint) (*model.Post, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(req dto.CreatePostRequest, userId uint, imageUrl string) (*model.Post, error) {
	post := &model.Post{
		Title: req.Title,
		Content: req.Content,
		UserId: userId,
		ImageUrl: imageUrl,	
	}

	if err := s.repo.CreatePost(post); err != nil {
		return nil, err
	}

	return s.repo.FindById(post.Id)
}

func (s *postService) GetPosts(page, limit int) ([]model.Post, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return s.repo.FindAllPaginated(page, limit)
}

func (s *postService) GetPostById(id uint) (*model.Post, error) {
	return s.repo.FindById(id)
}

func (s *postService) UpdatePost(id uint, req dto.UpdatePostRequest, newImageUrl string) (*model.Post, error) {
	post, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("Post not found")
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	if newImageUrl != "" {
		utils.DeleteImage(post.ImageUrl)
		post.ImageUrl = newImageUrl
	}

	if err := s.repo.UpdatePost(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) DeletePost(id uint) (*model.Post, error) {
	post, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("Post not found")
	}

	if err := s.repo.DeletePost(post); err != nil {
		return nil, err
	}

	utils.DeleteImage(post.ImageUrl)

	return post, nil
}