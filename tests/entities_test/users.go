package entities_test

import "github.com/rayato159/manga-store/internals/entities"

type TestUsersContext string

const (
	TestUsersCon TestUsersContext = "TestUsersController"
	TestUsersUse TestUsersContext = "TestUsersUsecase"
	TestUsersRep TestUsersContext = "TestUsersRepository"
)

type UsersRegisterTest struct {
	Input  *entities.UsersRegisterReq
	Expect string
}
