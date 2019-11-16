package models

// GithubAccessTokenResponse -
type GithubAccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorURI         string `json:"error_uri"`
}

// GithubClientResponse -
type GithubClientResponse struct {
	ClientID string `json:"client_id"`
}
