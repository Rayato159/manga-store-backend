package usecases

import "github.com/rayato159/manga-store/internals/entities"

type authUse struct {
	AuthRepo entities.AuthRepotiory
}

func NewAuthUsecase(authRepo entities.AuthRepotiory) entities.AuthUsecase {
	return &authUse{
		AuthRepo: authRepo,
	}
}
