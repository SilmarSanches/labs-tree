package user_usecase

import (
	"context"
	"fullcycle-auction_go/internal/entity/user_entity"
)

func (u *UserUseCase) CreateUser(ctx context.Context, user *user_entity.User) error {
	return u.UserRepository.CreateUser(ctx, user)
}
