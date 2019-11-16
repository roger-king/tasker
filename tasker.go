package tasker

import (
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	cron "github.com/robfig/cron/v3"
	"github.com/roger-king/tasker/handlers"
	"github.com/roger-king/tasker/services"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tasker -
type Tasker struct {
	DB        *mongo.Client
	Scheduler *cron.Cron
}

type TaskerConfig struct {
	MongoConnectionURL string `required:"true"`
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

// New - Creates a new instance of tasker
func New(tc *TaskerConfig) *Tasker {
	m, err := services.NewMongoConnection(tc.MongoConnectionURL)

	if err != nil {
		log.Panic(err)
	}

	return &Tasker{
		Scheduler: cron.New(),
		DB:        m,
	}
}

// Start - returns a mux router instance
func (t *Tasker) Start() *mux.Router {
	log.Info("Starting Tasker application")
	t.Scheduler.Start()

	taskService := &services.TaskService{
		DB:        t.DB,
		Scheduler: t.Scheduler,
	}

	githubService := services.NewGithubService()

	r := handlers.NewRouter(taskService, githubService)
	return r
}
