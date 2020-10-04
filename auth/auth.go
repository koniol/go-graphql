package auth

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/o1egl/paseto/v2"
	"graphqltest/models"
	"strconv"
	"time"
)

const (
		EXPIRED = "EXPIRED"
		VALID = "VALID"
		INVALID = "INVALID"
	)

type TokenValidation struct {
	Valid bool
	Error string
}

type UserToken struct {
	Id int
	Email string
	IsAdmin bool
}


func CreateToken(user *UserToken) (string, error)  {
	b, _ := hex.DecodeString("b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	privateKey := ed25519.PrivateKey(b)

	// or create a new keypair
	// publicKey, privateKey, err := ed25519.GenerateKey(nil)

	jsonToken := paseto.JSONToken{
		Expiration: time.Now().Add(24 * time.Hour),
	}
fmt.Println("TOKEdadas", user)
	// Add custom claim    to the token
	jsonToken.Set("data", *user)
	footer := "some footer"

	// Sign data
	token, err := paseto.Sign(privateKey, jsonToken, footer)

	// token = "v2.public.eyJkYXRhIjoidGhpcyBpcyBhIHNpZ25lZCBtZXNzYWdlIiwiZXhwIjoiMjAxOC0wMy0xMlQxOTowODo1NCswMTowMCJ9Ojv0uXlUNXSFhR88KXb568LheLRdeGy2oILR3uyOM_-b7r7i_fX8aljFYUiF-MRr5IRHMBcWPtM0fmn9SOd6Aw.c29tZSBmb290ZXI"
	if err != nil {
		fmt.Println("Signature error", err)
	}
	return token, err
}

func IsValid(token string) TokenValidation  {
	b, _ := hex.DecodeString("1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	publicKey := ed25519.PublicKey(b)

	// Verify data
	var newJsonToken paseto.JSONToken
	var newFooter string
	err := paseto.Verify(token, publicKey, &newJsonToken, &newFooter)

	timeNow := time.Now()
	expireDate := newJsonToken.Expiration

	fmt.Println("expire 123", timeNow, "EXPIRE123dasassad", expireDate,  )

	// Token expired regenerate
	if timeNow.Before(expireDate) {
		return TokenValidation{
			Valid: false,
			Error: EXPIRED,
		}
	}

	var v UserToken
	err = newJsonToken.Get("data", &v)
	fmt.Println("TOKEN DATA", &newFooter, err)
	fmt.Println("TOKEN DATA123123", &v)

	if err != nil {
		fmt.Println("Cant verify token", err)
		return TokenValidation{
			Valid: false,
			Error: INVALID,
		}
	}

	return TokenValidation{
		Valid: true,
		Error: VALID,
	}
}

func GetTokenData(token string) (*UserToken, error)  {
	b, _ := hex.DecodeString("1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	publicKey := ed25519.PublicKey(b)

	// Verify data
	var newJsonToken paseto.JSONToken
	var newFooter string
	err := paseto.Verify(token, publicKey, &newJsonToken, &newFooter)

	if err != nil {
		fmt.Println("Cant verify token", err)
		return nil, errors.New("Cant verify token")
	}
	var u UserToken
	err = newJsonToken.Get("data", &u)

	if err != nil {
		fmt.Println("Cant verify token", err)
		return nil, errors.New("Cant verify token")
	}
	return &u, nil
}

func (d *Data) UserExists(userId string, email string) error  {
	var u models.User
	id, _ := strconv.Atoi(userId)

	err := d.db.Model(&u).
		Where("id = ?", id).
		Where("email = ?", email).Select()

	if err != nil {
		fmt.Println("User not exits", err)
		return errors.New("Cant authenticate")
	}

	return nil
}

type Data struct {
	db *pg.DB
}

func NewAuthData(db *pg.DB) *Data  {
	return &Data{
		db: db,
	}
}

