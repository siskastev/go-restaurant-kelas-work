package user

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"go-restaurant-kelas-work/internal/tracing"
	"golang.org/x/crypto/argon2"
	"golang.org/x/net/context"
	"math/rand"
	"strings"
)

const cryptFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"

func (u *userRepo) GenerateUserHash(ctx context.Context, password string) (hash string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "GenerateUserHash")
	defer span.End()
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	argonHash := argon2.IDKey([]byte(password), salt, u.time, u.memory, u.threads, u.keyLen)
	b64Hash := u.encrypt(ctx, argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	encodeHash := fmt.Sprintf(cryptFormat, argon2.Version, u.memory, u.time, u.threads, b64Salt, b64Hash)
	return encodeHash, nil
}

func (u *userRepo) encrypt(ctx context.Context, text []byte) string {
	_, span := tracing.CreateSpan(ctx, "encrypt")
	defer span.End()
	nonce := make([]byte, u.gcm.NonceSize())
	cipherText := u.gcm.Seal(nonce, nonce, text, nil)
	return base64.StdEncoding.EncodeToString(cipherText)
}

func (u *userRepo) decrypt(ctx context.Context, cipherText string) ([]byte, error) {
	_, span := tracing.CreateSpan(ctx, "decrypt")
	defer span.End()
	decode, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	if len(decode) < u.gcm.NonceSize() {
		return nil, errors.New("Invalid nonce size")
	}
	return u.gcm.Open(nil, decode[:u.gcm.NonceSize()], decode[u.gcm.NonceSize():], nil)
}

func (u *userRepo) comparePassword(ctx context.Context, password, hash string) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "comparePassword")
	defer span.End()
	parts := strings.Split(hash, "$")
	var memory, time uint32
	var parallelism uint8

	switch parts[1] {
	case "argon2id":
		_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &parallelism)
		if err != nil {
			return false, err
		}
		salt, err := base64.RawStdEncoding.DecodeString(parts[4])
		if err != nil {
			return false, err
		}

		hash := parts[5]
		decryptedHash, err := u.decrypt(ctx, hash)
		if err != nil {
			return false, err
		}

		var keyLen = uint32(len(decryptedHash))
		comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLen)
		return subtle.ConstantTimeCompare(comparisonHash, decryptedHash) == 1, nil
	}
	return false, nil
}
