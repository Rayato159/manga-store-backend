package repositories

import (
	"context"
	"errors"
	"log"

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

func (ar *authRepo) UpdateUserRefreshToken(ctx context.Context, userId string, token string) error {
	ctx = context.WithValue(ctx, entities.AuthRep, "Rep.UpdateUserRefreshToken")
	defer log.Println(ctx.Value(entities.AuthRep))

	query := `
	UPDATE "users"
	SET
	"refresh_token" = $2
	WHERE "id" = $1`

	if _, err := ar.Db.Exec(query, userId, token); err != nil {
		log.Println(err.Error())
		return errors.New("error, user not found")
	}
	return nil
}
