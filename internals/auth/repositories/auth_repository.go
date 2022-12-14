package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/utils"
)

type authRepo struct {
	Db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) entities.AuthRepository {
	return &authRepo{
		Db: db,
	}
}

func (ar *authRepo) UpdateUserRefreshToken(ctx context.Context, reqType string, userId string, token string, newToken string) error {
	ctx = context.WithValue(ctx, entities.AuthRep, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.AuthRep).(int64)))

	switch reqType {
	case "user_id":
		_ = newToken
		query := `
		UPDATE "users"
		SET
		"refresh_token" = $2
		WHERE "id" = $1`

		if _, err := ar.Db.Exec(query, userId, token); err != nil {
			log.Println(err.Error())
			return errors.New("error, user not found")
		}
	case "refresh_toekn":
		query := `
		UPDATE "users"
		SET
		"refresh_token" = $2
		WHERE "refresh_token" = $1`

		if _, err := ar.Db.Exec(query, token, newToken); err != nil {
			log.Println(err.Error())
			return errors.New("error, user not found")
		}
	default:
		return errors.New("error, request type is invalid")
	}
	return nil
}
