package usecases

import (
	"context"

	"github.com/rayato159/manga-store/internals/entities"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

func NewUsersUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

func (uu *usersUse) Register(ctx context.Context) error {
	return nil
}
