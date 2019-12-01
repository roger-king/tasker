package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/roger-king/tasker/utils"
	"github.com/sirupsen/logrus"
)

type User struct {
	Email                string `json:"email" bson:"email"`
	Name                 string `json:"name" bson:"name"`
	UserName             string `json:"username" bson:"username"`
	EncryptedAccessToken string `json:"-" bson:"accessToken"`
	AccessToken          string `json:"-" bson:"-"`
	Bio                  string `json:"bio" bson:"bio"`
	GitHubURL            string `json:"githubURL" bson:"githubURL"`
}

func (u *User) BeforeCreate() error {
	bKey := []byte(u.AccessToken)

	encryptedToken, err := encrypt(bKey, []byte(utils.TaskerSecret))

	if err != nil {
		logrus.Error("error", err)
		return err
	}

	u.EncryptedAccessToken = string(encryptedToken)
	return nil
}

func (u *User) GetAccessToken() (string, error) {
	secret := utils.TaskerSecret

	token, err := decrypt([]byte(u.EncryptedAccessToken), []byte(secret))

	return string(token), err
}

func encrypt(plaintext []byte, secret []byte) ([]byte, error) {
	c, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, secret []byte) ([]byte, error) {
	c, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
