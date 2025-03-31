package user

import (
	"context"
	"fullcycle-auction_go/internal/entity/user_entity"
)

func (ur *UserRepository) CreateUser(ctx context.Context, user *user_entity.User) error {
	_, err := ur.Collection.InsertOne(ctx, user)
	return err
}
