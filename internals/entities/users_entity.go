package entities

import (
	"context"
	"time"
)

type UsersContext string

const (
	UsersCon UsersContext = "UsersController"
	UsersUse UsersContext = "UsersUsecase"
	UsersRep UsersContext = "UsersRepository"
)

type UsersUsecase interface {
	Register(ctx context.Context, req *UsersRegisterReq) (*UsersRegisterRes, error)
}

type UsersRepository interface {
	Register(ctx context.Context, req *UsersRegisterReq) (*UsersRegisterRes, error)
	FindOneUser(ctx context.Context, username string) (*UsersCredentialsReq, error)
	GetUserInfo(ctx context.Context, reqType string, username string, refreshToken string) (*UsersInfo, error)
}

type UsersRole string

const (
	Admin   UsersRole = "admin"
	Manager UsersRole = "manager"
	User    UsersRole = "user"
)

type UsersInfo struct {
	Id       string    `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password"`
	Role     UsersRole `db:"role" json:"role"`
}

type UsersRegisterReq struct {
	Username        string    `db:"username" json:"username" form:"username" copier:"Username"`
	Password        string    `db:"password" json:"password" form:"password"`
	ConfirmPassword string    `json:"passowrd_confirm" form:"password_confirm"`
	Role            UsersRole `db:"role" json:"role" form:"role"`
	Key             string    `json:"key" form:"key"`
}

type UsersRegisterRes struct {
	Id        string    `db:"id" json:"id" copier:"Id"`
	Username  string    `db:"username" json:"username" copier:"Username"`
	CreatedAt time.Time `db:"created_at" json:"created_at" copier:"CreatedAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" copier:"UpdatedAt"`
}
