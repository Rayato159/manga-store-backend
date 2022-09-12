package usecases

import "github.com/rayato159/manga-store/internals/entities"

type authUse struct {
	AuthRepo entities.AuthRepository
}

func NewAuthUsecase(authRepo entities.AuthRepository) entities.AuthUsecase {
	return &authUse{
		AuthRepo: authRepo,
	}
}
