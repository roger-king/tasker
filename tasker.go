// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package tasker

import (
	"os"
	"strings"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	cron "github.com/robfig/cron/v3"
	"github.com/roger-king/tasker/config"
	"github.com/roger-king/tasker/services"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// Tasker -
type Tasker struct {
	Config    *config.TaskerConfig
	DB        *mongo.Client
	Scheduler *cron.Cron
	Router    *mux.Router
}

var TaskerSet = wire.NewSet(services.ServiceSet, ProvideCron, ProivdeTasker)

func ProvideCron() *cron.Cron {
	return cron.New()
}

func ProivdeTasker(tc *config.TaskerConfig, c *cron.Cron, db *sqlx.DB) *Tasker {
	if tc.Auth {
		if len(tc.GithubClientID) == 0 && len(tc.GithubClientSecret) == 0 {
			log.Fatal("Authentication is enabled. Please provide the github client id and secret.")
			os.Exit(1)
		}
	}

	if len(tc.DBConnectionURL) > 0 {
		if !strings.Contains(tc.DBConnectionURL, "postgres") {
			log.Fatal("Please provide a valid postgres db connection")
			os.Exit(1)
		}
	} else {
		log.Fatal("DBConnectionURL is required")
		os.Exit(1)
	}

	return &Tasker{
		Config:    tc,
		Scheduler: c,
	}
}

// New - Creates a new instance of tasker
func New(tc *config.TaskerConfig) (*Tasker, error) {
	wire.Build(TaskerSet)
	return nil, nil
}

// Start - returns a mux router instance
func (t *Tasker) Start() *mux.Router {
	log.Info("Starting Tasker application")
	t.Scheduler.Start()

	return t.Router
}
