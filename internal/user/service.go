package user

import (
	"context"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NweService(s Service) *Service {
	return &s
}

func (s Service) Save(ctx context.Context, req UserRegisterRequest) (*User, error) {
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := gorm.G[User](s.DB).Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s Service) Find(ctx context.Context, req UserLoginRequest) (*User, error) {
	var user User
	user, err := gorm.G[User](s.DB).Where("name = ? AND password = ?", req.Name, req.Password).First(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
