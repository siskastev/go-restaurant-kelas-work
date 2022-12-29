package user

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/tracing"
	"gorm.io/gorm"
	"time"
)

type userRepo struct {
	db        *gorm.DB
	gcm       cipher.AEAD
	time      uint32
	memory    uint32
	threads   uint8
	keyLen    uint32
	signKey   *rsa.PrivateKey
	accessExp time.Duration
}

func GetUserRepository(
	db *gorm.DB,
	secret string,
	time uint32,
	memory uint32,
	threads uint8,
	keyLen uint32,
	signKey *rsa.PrivateKey,
	accessExp time.Duration) (Repository, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &userRepo{db: db,
		gcm:       gcm,
		time:      time,
		memory:    memory,
		threads:   threads,
		keyLen:    keyLen,
		signKey:   signKey,
		accessExp: accessExp,
	}, nil
}

func (u *userRepo) RegisterUser(ctx context.Context, user model.User) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()
	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil

}

func (u *userRepo) CheckRegister(ctx context.Context, username string) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckRegister")
	defer span.End()
	var user model.User
	if err := u.db.WithContext(ctx).Where(model.User{Username: username}).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return user.Username != "", nil
}

func (u *userRepo) GetUserData(ctx context.Context, username string) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetUserData")
	defer span.End()
	var user model.User
	if err := u.db.WithContext(ctx).Where(model.User{Username: username}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepo) VerifyLogin(ctx context.Context, username, password string, user model.User) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "VerifyLogin")
	defer span.End()
	if username != user.Username {
		return false, nil
	}
	verified, err := u.comparePassword(ctx, password, user.Hash)
	if err != nil {
		return false, err
	}
	return verified, nil
}
