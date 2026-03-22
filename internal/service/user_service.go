package service

import (
	"errors"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/dto"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/model"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
)

type UserService interface {
	Register(req dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
	Refresh(req dto.RefreshRequest) (*dto.RefreshResponse, error)
}

type userService struct {
	cfg  *config.Config
	repo repository.UserRepository
}

func NewService(cfg *config.Config, repo repository.UserRepository) UserService {
	return &userService{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *userService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}

	existsEmail, _ := s.repo.GetUserByEmail(req.Email)
	if existsEmail != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email: req.Email,
		Name: req.Name,
		Phone: req.Phone,
		Password: hashedPassword,
	}

	return &dto.RegisterResponse{
		User: dto.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
	}, nil
}

func (s *userService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.Id, user.Email, user.Name, user.Phone, s.cfg.JWT_SECRET)
	
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		User: dto.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) Refresh(req dto.RefreshRequest) (*dto.RefreshResponse, error) {
	claims, err := utils.ValidateJWT(req.RefreshToken, s.cfg.JWT_SECRET)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	accessToken, refreshToken, err := utils.GenerateTokens(claims.ID, claims.Email, claims.Name, claims.Phone, s.cfg.JWT_SECRET)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}