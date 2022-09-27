package repositories

import (
	"context"
	"errors"
	"log"

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

func (ur *usersRepo) FindOneUser(ctx context.Context, username string) (*entities.UsersCredentialsReq, error) {
	ctx = context.WithValue(ctx, entities.UsersRep, "Rep.FindOneUser")
	defer log.Println(ctx.Value(entities.UsersRep))

	query := `
	SELECT 
	COALESCE("username", '') AS "username", 
	COALESCE("password", '') AS "password"
	FROM "users" 
	WHERE "username" = $1`

	user := new(entities.UsersCredentialsReq)
	if err := ur.Db.Get(user, query, username); err != nil {
		log.Println(err.Error())
		return user, errors.New("error, user not found")
	}
	return user, nil
}

func (ur *usersRepo) GetUserInfo(ctx context.Context, reqType string, username string, refreshToken string) (*entities.UsersInfo, error) {
	ctx = context.WithValue(ctx, entities.UsersRep, "Rep.GetUserInfo")
	defer log.Println(ctx.Value(entities.UsersRep))

	switch reqType {
	case "username":
		query := `
		SELECT 
		COALESCE("id"::VARCHAR, '') AS "id",
		COALESCE("username", '') AS "username", 
		COALESCE("password", '') AS "password",
		COALESCE("role"::VARCHAR, '') AS "role"
		FROM "users" 
		WHERE "username" = $1`

		user := new(entities.UsersInfo)
		if err := ur.Db.Get(user, query, username); err != nil {
			log.Println(err.Error())
			return user, errors.New("error, user not found")
		}
		return user, nil
	case "refresh_token":
		query := `
		SELECT 
		COALESCE("id"::VARCHAR, '') AS "id",
		COALESCE("username", '') AS "username", 
		COALESCE("password", '') AS "password",
		COALESCE("role"::VARCHAR, '') AS "role"
		FROM "users" 
		WHERE "refresh_token" = $1`

		user := new(entities.UsersInfo)
		if err := ur.Db.Get(user, query, refreshToken); err != nil {
			log.Println(err.Error())
			return user, errors.New("error, user not found")
		}
		return user, nil
	default:
		return nil, errors.New("error, request type is invalid")
	}
}

func (ur *usersRepo) Register(ctx context.Context, req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	ctx = context.WithValue(ctx, entities.UsersRep, "Rep.Register")
	defer log.Println(ctx.Value(entities.UsersRep))

	query := `
	INSERT INTO "users"(
		"username",
		"password",
		"role"
	)
	VALUES (
		:username,
		:password,
		:role
	)
	RETURNING "id", "username", "created_at", "updated_at";
	`

	tx, err := ur.Db.BeginTxx(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error, transaction is not enable")
	}

	rows, err := tx.NamedQuery(query, req)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error, can't insert user")
	}

	user := new(entities.UsersRegisterRes)
	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			log.Println(err.Error())
			return nil, errors.New("error, can't parse user into the struct")
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, errors.New("error, can't commit query transaction")
	}
	return user, nil
}
