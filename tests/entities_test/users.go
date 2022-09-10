package entities_test

import "context"

type TestUsersContext string

const (
	TestUsersCon TestUsersContext = "TestUsersController"
	TestUsersUse TestUsersContext = "TestUsersUsecase"
	TestUsersRep TestUsersContext = "TestUsersRepository"
)

type TestUsersUsecase interface {
	TestRegister(ctx context.Context) error
}

type TestUsersRepository interface{}
