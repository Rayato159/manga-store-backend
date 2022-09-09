package entities

import (
	"context"
)

type UsersContext string

const (
	usersCon UsersContext = "UsersController"
	UsersUse UsersContext = "UsersUsecase"
	UsersRep UsersContext = "UsersRepository"
)

type UsersUsecase interface {
	Register(ctx context.Context) error
}

type UsersRepository interface{}
