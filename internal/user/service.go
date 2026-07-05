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

func (s Service) Save(ctx context.Context, user User) (*User, error) {
	err := gorm.G[User](s.DB).Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
