package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/internals/entities"
)

type usersRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) entities.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}
