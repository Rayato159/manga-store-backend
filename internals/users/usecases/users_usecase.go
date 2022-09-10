package usecases

import (
	"context"
	"errors"
	"log"

	"github.com/rayato159/manga-store/internals/entities"
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
	ctx = context.WithValue(ctx, entities.UsersUse, "Use.Register")
	defer log.Println(ctx.Value(entities.UsersUse))

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
