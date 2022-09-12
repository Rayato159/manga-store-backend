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
	FindOneUser(ctx context.Context, req string) (*Credentials, error)
}

type UsersRole string

const (
	Admin UsersRole = "admin"
	User  UsersRole = "user"
)

type UsersRegisterReq struct {
	Username        string    `db:"username" json:"username" form:"username" copier:"Username"`
	Password        string    `db:"password" json:"password" form:"password"`
	ConfirmPassword string    `json:"passowrd_confirm" form:"password_confirm"`
	Role            UsersRole `db:"role" json:"role" form:"role"`
	AdminKey        string    `json:"admin_key" form:"admin_key"`
}

type UsersRegisterRes struct {
	Username  string    `db:"username" json:"username" copier:"Username"`
	CreatedAt time.Time `db:"created_at" json:"created_at" copier:"CreatedAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" copier:"UpdatedAt"`
}
