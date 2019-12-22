package config

// TaskerConfig -
type TaskerConfig struct {
	Auth               bool
	DBConnectionURL    string
	GithubClientID     string
	GithubClientSecret string
	Migrate            bool
	MongoConnectionURL string
	Secret             string
}
