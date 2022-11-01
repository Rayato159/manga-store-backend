package usecases

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/rayato159/manga-store/internals/entities"
	"github.com/rayato159/manga-store/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

func NewUsersUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

func (uu *usersUse) Register(ctx context.Context, req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	ctx = context.WithValue(ctx, entities.UsersUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.UsersUse).(int64)))

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	user, _ := uu.UsersRepo.FindOneUser(ctx, req.Username)
	if user.Username == req.Username {
		log.Println("error, username has been already taken")
		return nil, errors.New("error, username has been already taken")
	}

	res, err := uu.UsersRepo.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uu *usersUse) ChangePassword(ctx context.Context, req *entities.ChangePasswordReq) error {
	ctx = context.WithValue(ctx, entities.UsersUse, time.Now().UnixMilli())
	log.Printf("called:\t%v", utils.Trace())
	defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.UsersUse).(int64)))

	userOldHashedPassword, err := uu.UsersRepo.GetUserPassword(ctx, req.UserId)
	if err != nil {
		return err
	}

	// Compare an old password
	if err := bcrypt.CompareHashAndPassword([]byte(userOldHashedPassword), []byte(req.OldPassword)); err != nil {
		log.Println(err.Error())
		return errors.New("error, old password is invalid")
	}

	// Hash a new password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	req.NewPassword = string(hashed)

	// Update a user's password
	if err := uu.UsersRepo.ChangePassword(ctx, req); err != nil {
		return err
	}
	return nil
}
