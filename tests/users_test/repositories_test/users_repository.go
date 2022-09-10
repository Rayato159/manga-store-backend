package repositories_test

import (
	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/tests/entities_test"
)

type testUsersRepo struct {
	TestDb *sqlx.DB
}

func NewTestUsersRepository(db *sqlx.DB) entities_test.TestUsersRepository {
	return &testUsersRepo{
		TestDb: db,
	}
}
