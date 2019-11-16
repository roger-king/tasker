package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email                string `json:"email" bson:"email"`
	Name                 string `json:"name" bson:"name"`
	UserName             string `json:"username" bson:"username"`
	EncryptedAccessToken string `json:"accessToken" bson:"accessToken"`
	AccessToken          string `json:"-" bson:"-"`
	Bio                  string `json:"bio" bson:"bio"`
	GitHubURL            string `json:"githubURL" bson:"githubURL"`
}

func (u *User) BeforeCreate() error {
	bKey := []byte(u.AccessToken)

	hash, err := bcrypt.GenerateFromPassword(bKey, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.EncryptedAccessToken = string(hash)
	return nil
}
