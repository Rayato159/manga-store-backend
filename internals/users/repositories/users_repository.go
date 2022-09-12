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

func (ur *usersRepo) FindOneUser(ctx context.Context, req string) (*entities.Credentials, error) {
	ctx = context.WithValue(ctx, entities.UsersRep, "Rep.FindOneUser")
	defer log.Println(ctx.Value(entities.UsersRep))

	query := `
	SELECT 
	COALESCE("username", '') AS "username", 
	COALESCE("password", '') AS "password"
	FROM "users" 
	WHERE "username" = $1`

	user := new(entities.Credentials)
	if err := ur.Db.Get(user, query, req); err != nil {
		log.Println(err.Error())
		return user, errors.New("error, user not found")
	}
	return user, nil
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
	RETURNING "username", "created_at", "updated_at";
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
