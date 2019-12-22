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

// UserDTO -
type UserDTO struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	UserName  string `json:"username"`
	Bio       string `json:"bio"`
	GitHubURL string `json:"githubURL"`
}

// NewUserInput -
type NewUserInput struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	UserName             string `json:"username"`
	Bio                  string `json:"bio"`
	GitHubURL            string `json:"githubURL"`
	AccessToken          string `json:"accessToken"`
	EncryptedAccessToken string `json:"-"`
}

// User -
type User struct {
	ID                   int    `db:"id"`
	Email                string `db:"email"`
	Name                 string `db:"name"`
	UserName             string `db:"username"`
	EncryptedAccessToken string `db:"access_token"`
	AccessToken          string `db:"-"`
	Bio                  string `db:"bio"`
	GitHubURL            string `db:"githubURL"`
}

// BeforeCreate -
func (u *NewUserInput) BeforeCreate() error {
	bKey := []byte(u.AccessToken)

	encryptedToken, err := encrypt(bKey, []byte(utils.TaskerSecret))

	if err != nil {
		logrus.Error("error", err)
		return err
	}

	u.EncryptedAccessToken = string(encryptedToken)
	return nil
}

// GetAccessToken -
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
