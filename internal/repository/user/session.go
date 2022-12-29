package user

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/tracing"
	"golang.org/x/net/context"
	"time"
)

type Claims struct {
	jwt.StandardClaims
}

func (u *userRepo) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateUserSession")
	defer span.End()
	accessToken, err := u.generateAccessToken(ctx, userID)
	if err != nil {
		return model.UserSession{}, err
	}
	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (u *userRepo) generateAccessToken(ctx context.Context, userID string) (string, error) {
	ctx, span := tracing.CreateSpan(ctx, "generateAccessToken")
	defer span.End()
	accessTokenExpired := time.Now().Add(u.accessExp).Unix()
	accessClaims := Claims{
		jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: accessTokenExpired,
		},
	}

	accessJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaims)
	return accessJwt.SignedString(u.signKey)
}

func (u *userRepo) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()
	accessToken, err := jwt.ParseWithClaims(data.JWTToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return &u.signKey.PublicKey, nil
	})
	if err != nil {
		return "", err
	}

	accessTokenClaims, ok := accessToken.Claims.(*Claims)
	if !ok {
		return "", errors.New("unauthorized")
	}
	if accessToken.Valid {
		return accessTokenClaims.Subject, nil
	}
	return "", errors.New("unauthorized")

}
