package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/internals/entities"
)

type authRepo struct {
	Db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) entities.AuthRepotiory {
	return &authRepo{
		Db: db,
	}
}
