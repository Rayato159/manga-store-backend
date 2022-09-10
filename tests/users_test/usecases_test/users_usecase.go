package usecases_test

import (
	"context"

	"github.com/rayato159/manga-store/tests/entities_test"
)

type testUsersUsecase struct {
	TestUsersRepo entities_test.TestUsersRepository
}

func NewTestUsersUsecase(testUsersRepo entities_test.TestUsersRepository) entities_test.TestUsersUsecase {
	return &testUsersUsecase{
		TestUsersRepo: testUsersRepo,
	}
}

func (tuu *testUsersUsecase) TestRegister(ctx context.Context) error {
	return nil
}
