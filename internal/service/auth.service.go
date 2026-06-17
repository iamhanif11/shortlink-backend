package service

import (
	"context"
	"errors"
	"log"

	"github.com/iamhanif11/shortlink-backend.git/internal/dto"
	"github.com/iamhanif11/shortlink-backend.git/internal/repository"
	"github.com/iamhanif11/shortlink-backend.git/pkg"
)

var ErrEmailAlreadyExists = errors.New("Email is already registered, please use another email")

type AuthService struct {
	authRepository *repository.AuthRepository
}

func NewAuthService(authRepository *repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (as *AuthService) RegisterUser(ctx context.Context, user dto.RegisterReq) (dto.RegisterRes, error) {
	existingUser, err := as.authRepository.GetUserByEmail(ctx, user.Email)

	if err == nil && existingUser.Email != "" {
		return dto.RegisterRes{}, ErrEmailAlreadyExists
	}
	var hc pkg.HashConfig
	hc.UseRecommended()
	hashPwd := hc.GenHash(user.Password)

	newUser, err := as.authRepository.AddUser(ctx, user.Email, hashPwd)
	if err != nil {
		return dto.RegisterRes{}, err
	}
	return dto.RegisterRes{
		Id:        newUser.Id,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

func (as *AuthService) LoginUser(ctx context.Context, user dto.LoginReq) (dto.LoginResponse, error) {
	log.Println(user)
	login, err := as.authRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Printf("repository error: %+v\n", err)
		return dto.LoginResponse{}, errors.New("Email or Password Invalid !")
	}

	var hash pkg.HashConfig
	if err := hash.Compare(user.Password, login.Password); err != nil {
		return dto.LoginResponse{}, errors.New("Password not match")
	}

	claims := pkg.NewClaims(login.Id, login.Email)
	token, err := claims.GenJWT()
	if err != nil {
		log.Println("GetUserByEmail Error:", err)
		return dto.LoginResponse{}, err
	}

	log.Println(token)

	return dto.LoginResponse{
		Token: token,
		User: dto.LoginUserDetail{
			Id:    login.Id,
			Email: login.Email,
		},
	}, nil
}
