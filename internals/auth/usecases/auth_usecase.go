package usecases

import (
	"context"
	"errors"
	"log"

	"github.com/rayato159/manga-store/configs"
	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type authUse struct {
	AuthRepo  entities.AuthRepository
	UsersRepo entities.UsersRepository
}

func NewAuthUsecase(authRepo entities.AuthRepository, usersRepo entities.UsersRepository) entities.AuthUsecase {
	return &authUse{
		AuthRepo:  authRepo,
		UsersRepo: usersRepo,
	}
}

func (au *authUse) Login(ctx context.Context, cfg *configs.Configs, req *entities.UsersCredentialsReq) (*entities.UsersCredentialsRes, error) {
	ctx = context.WithValue(ctx, entities.AuthUse, "Con.Login")
	defer log.Println(ctx.Value(entities.AuthUse))

	// Find user by username
	user, err := au.UsersRepo.GetUserInfo(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// Password check
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("error, password is invalid")
	}

	claims := &entities.UsersJwtClaimsReq{
		UsersAccessToken: &entities.UsersAccessToken{
			Id:   user.Id,
			Role: string(user.Role),
		},
		UsersRefreshToken: &entities.UsersRefreshToken{
			Id:   user.Id,
			Role: string(user.Role),
		},
		UsersSessionToken: &entities.UsersSessionToken{
			Id:       user.Id,
			Username: user.Username,
			Role:     string(user.Role),
		},
	}

	accessTokenChan := make(chan string)
	accessTokenChanErr := make(chan error)
	refreshTokenChan := make(chan string)
	refreshTokenChanErr := make(chan error)
	sessionTokenChan := make(chan string)
	sessionTokenChanErr := make(chan error)

	go func() {
		accessToken, err := utils.JwtUsersClaims(ctx, cfg, au.AuthRepo, claims, entities.AccessToken)
		accessTokenChanErr <- err
		accessTokenChan <- accessToken
		close(accessTokenChanErr)
		close(accessTokenChan)
	}()
	go func() {
		refreshToken, err := utils.JwtUsersClaims(ctx, cfg, au.AuthRepo, claims, entities.RefreshToken)
		refreshTokenChanErr <- err
		refreshTokenChan <- refreshToken
		close(refreshTokenChanErr)
		close(refreshTokenChan)
	}()
	go func() {
		sessionToken, err := utils.JwtUsersClaims(ctx, cfg, au.AuthRepo, claims, entities.SessionToken)
		sessionTokenChanErr <- err
		sessionTokenChan <- sessionToken
		close(sessionTokenChanErr)
		close(sessionTokenChan)
	}()

	accessTokenErr := <-accessTokenChanErr
	refreshTokenErr := <-refreshTokenChanErr
	sessionTokenErr := <-sessionTokenChanErr

	if accessTokenErr != nil {
		return nil, accessTokenErr
	}
	if refreshTokenErr != nil {
		return nil, refreshTokenErr
	}
	if sessionTokenErr != nil {
		return nil, sessionTokenErr
	}

	accessToken := <-accessTokenChan
	refreshToken := <-refreshTokenChan
	sessionToken := <-sessionTokenChan

	res := &entities.UsersCredentialsRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		SessionToken: sessionToken,
	}
	return res, nil
}
