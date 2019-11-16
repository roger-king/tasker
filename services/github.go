package services

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/roger-king/tasker/utils"
)

type GithubService struct {
	Client        *resty.Client
	Req           *resty.Request
	LoginTokenURL string
	APIURL        string
}

type GithubAccessTokenBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

// NewService - Creates an instance of UserService
func NewGithubService() *GithubService {
	client := resty.New()
	req := client.R().SetHeaders(map[string]string{
		"Accept": "application/json",
	})

	return &GithubService{
		Client:        client,
		Req:           req,
		LoginTokenURL: "https://github.com/login",
		APIURL:        "https://api.github.com",
	}
}

func (g *GithubService) GetAccessToken(code string) (*interface{}, error) {
	var ghResponse interface{}
	url := fmt.Sprintf("%s/%s", g.LoginTokenURL, "oauth/access_token")

	resp, err := g.Req.SetQueryParams(map[string]string{
		"client_id":     utils.GithubClientID,
		"client_secret": utils.GithubClientSecret,
		"code":          code,
	}).Post(url)

	if err != nil {
		return nil, err
	}

	err = g.Client.JSONUnmarshal(resp.Body(), &ghResponse)

	if err != nil {
		return nil, err
	}

	return &ghResponse, nil
}
