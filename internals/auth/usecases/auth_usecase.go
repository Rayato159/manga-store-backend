package usecases

import (
	"context"

	"github.com/rayato159/manga-store/internals/entities"
)

type authUse struct {
	AuthRepo entities.AuthRepository
}

func NewAuthUsecase(authRepo entities.AuthRepository) entities.AuthUsecase {
	return &authUse{
		AuthRepo: authRepo,
	}
}

func (au *authUse) Login(ctx context.Context, req *entities.UsersCredentialsReq) (*entities.UsersCredentialsRes, error) {
	return nil, nil
}
