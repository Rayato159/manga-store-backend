package entities

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rayato159/manga-store/configs"
)

type AuthContext string

const (
	AuthCon AuthContext = "AuthController"
	AuthUse AuthContext = "AuthUsecase"
	AuthRep AuthContext = "AuthRepository"
)

type ClaimsType string

const (
	AccessToken  ClaimsType = "access_token"
	RefreshToken ClaimsType = "refresh_token"
	SessionToken ClaimsType = "session_token"
)

type AuthRepository interface {
	UpdateUserRefreshToken(ctx context.Context, reqType string, userId string, token string, newToken string) error
}

type AuthUsecase interface {
	Login(ctx context.Context, cfg *configs.Configs, rdb *redis.Client, req *UsersCredentialsReq) (*UsersCredentialsRes, error)
	RefreshToken(ctx context.Context, cfg *configs.Configs, rdb *redis.Client, refreshToken string) (*UsersCredentialsRes, error)
}

type UsersCredentialsReq struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password" form:"password"`
}

type UsersCredentialsRes struct {
	AccessToken  string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
	SessionToken string `db:"session_token" json:"session_token"`
}

type UsersAccessToken struct {
	Id   string `db:"id" json:"id"`
	Role string `db:"role" json:"role"`
}

type UsersRefreshToken struct {
	Id        string `db:"id" json:"id"`
	Role      string `db:"role" json:"role"`
	ExpiresAt *time.Time
	IssuedAt  *time.Time
}

type UsersSessionToken struct {
	Id        string `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Role      string `db:"role" json:"role"`
	ExpiresAt *time.Time
	IssuedAt  *time.Time
}

type UsersJwtClaimsReq struct {
	UsersAccessToken  *UsersAccessToken
	UsersRefreshToken *UsersRefreshToken
	UsersSessionToken *UsersSessionToken
}

type UsersJwtTokenMapClaims struct {
	Id   string `db:"id" json:"id"`
	Role string `db:"role" json:"role"`
	jwt.RegisteredClaims
}

type UsersJwtSessionMapClaims struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Role     string `db:"role" json:"role"`
	jwt.RegisteredClaims
}

type RefreshTokenReq struct {
	RefreshToken string `db:"refresh_token" json:"refresh_token" form:"refresh_token"`
}
