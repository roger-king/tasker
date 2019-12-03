// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package tasker

import (
	"os"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	cron "github.com/robfig/cron/v3"
	"github.com/roger-king/tasker/handlers"
	"github.com/roger-king/tasker/models"
	"github.com/roger-king/tasker/services"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tasker -
type Tasker struct {
	Config    models.TaskerConfig
	DB        *mongo.Client
	Scheduler *cron.Cron
	Router    *mux.Router
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

var TaskerSet = wire.NewSet(services.ServiceSet, handlers.RouterSet, ProvideCron, ProivdeTasker)

func ProvideCron() *cron.Cron {
	return cron.New()
}

func ProivdeTasker(tc models.TaskerConfig, r *mux.Router, c *cron.Cron) *Tasker {
	return &Tasker{
		Config:    tc,
		Scheduler: c,
		Router:    r,
	}
}

// New - Creates a new instance of tasker
func New(tc models.TaskerConfig) (*Tasker, error) {
	wire.Build(TaskerSet)
	return nil, nil
}

// Start - returns a mux router instance
func (t *Tasker) Start() *mux.Router {
	log.Info("Starting Tasker application")
	t.Scheduler.Start()

	return t.Router
}
