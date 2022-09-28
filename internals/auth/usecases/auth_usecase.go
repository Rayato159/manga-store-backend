package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
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

func (au *authUse) getUserInfoCache(ctx context.Context, rdb *redis.Client, key string) (*entities.UsersInfo, error) {
	ctx = context.WithValue(ctx, entities.AuthCon, "Con.getLoginCache")
	defer log.Println(ctx.Value(entities.AuthCon))

	val, err := rdb.Get(ctx, key).Result()
	if err != nil || val == "" {
		log.Println(err.Error())
		return nil, fmt.Errorf("error, can't get value of key: %v", key)
	}

	user := new(entities.UsersInfo)
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error, can't parse: %v to object", val)
	}
	return user, nil
}

func (au *authUse) setUserInfoCache(ctx context.Context, rdb *redis.Client, key string, data any) error {
	ctx = context.WithValue(ctx, entities.AuthCon, "Con.setLoginCache")
	defer log.Println(ctx.Value(entities.AuthCon))

	dataStringtify, err := json.Marshal(&data)
	if err != nil {
		return fmt.Errorf("error, can't parse: %v to stringtify", data)
	}

	if err := rdb.Set(ctx, key, string(dataStringtify), time.Hour*24).Err(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("error, can't set value of key: %v", data)
	}
	return nil
}

func (au *authUse) Login(ctx context.Context, cfg *configs.Configs, rdb *redis.Client, req *entities.UsersCredentialsReq) (*entities.UsersCredentialsRes, error) {
	ctx = context.WithValue(ctx, entities.AuthUse, "Use.Login")
	defer log.Println(ctx.Value(entities.AuthUse))

	user := new(entities.UsersInfo)
	var err error
	if rdb != nil {
		user, err = au.getUserInfoCache(ctx, rdb, req.Username)
		if err != nil {
			// Find user by username
			user, err = au.UsersRepo.GetUserInfo(ctx, "username", req.Username, "")
			if err != nil {
				return nil, err
			}
			if err = au.setUserInfoCache(ctx, rdb, req.Username, user); err != nil {
				return nil, err
			}
		}
	} else {
		user, err = au.UsersRepo.GetUserInfo(ctx, "username", req.Username, "")
		if err != nil {
			return nil, err
		}
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

func (au *authUse) RefreshToken(ctx context.Context, cfg *configs.Configs, rdb *redis.Client, resfreshToken string) (*entities.UsersCredentialsRes, error) {
	ctx = context.WithValue(ctx, entities.AuthUse, "Use.RefreshToken")
	defer log.Println(ctx.Value(entities.AuthUse))

	user := new(entities.UsersInfo)
	var err error
	if rdb != nil {
		user, err = au.getUserInfoCache(ctx, rdb, resfreshToken)
		if err != nil {
			// Find user by username
			user, err = au.UsersRepo.GetUserInfo(ctx, "refresh_token", "", resfreshToken)
			if err != nil {
				return nil, err
			}
			if err = au.setUserInfoCache(ctx, rdb, resfreshToken, user); err != nil {
				return nil, err
			}
		}
	} else {
		user, err = au.UsersRepo.GetUserInfo(ctx, "refresh_token", "", resfreshToken)
		if err != nil {
			return nil, err
		}
	}

	expiresAtString, err := utils.JwtExtractPayload(ctx, cfg, "exp", resfreshToken)
	if err != nil {
		return nil, err
	}
	expiresAtInt, err := strconv.ParseInt(strconv.Itoa(int(expiresAtString.(float64))), 10, 64)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error, can't convert expires to int")
	}
	expiresAt := time.Unix(expiresAtInt, 0)
	issueAt := time.Now()

	claims := &entities.UsersJwtClaimsReq{
		UsersAccessToken: &entities.UsersAccessToken{
			Id:   user.Id,
			Role: string(user.Role),
		},
		UsersRefreshToken: &entities.UsersRefreshToken{
			Id:        user.Id,
			Role:      string(user.Role),
			IssuedAt:  &issueAt,
			ExpiresAt: &expiresAt,
		},
		UsersSessionToken: &entities.UsersSessionToken{
			Id:        user.Id,
			Username:  user.Username,
			Role:      string(user.Role),
			IssuedAt:  &issueAt,
			ExpiresAt: &expiresAt,
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
