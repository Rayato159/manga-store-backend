package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/internals/entities"
)

type authRepo struct {
	Db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) entities.AuthRepository {
	return &authRepo{
		Db: db,
	}
}
