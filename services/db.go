package services

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/roger-king/tasker/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"os"
)

// NewDBConnection -
func NewDBConnection(tc *config.TaskerConfig) *sqlx.DB {
	db := sqlx.MustConnect("postgres", tc.DBConnectionURL)

	if tc.Migrate {
		log.Info("Starting migration")
		log.Info("Pulling latest migration...")
		fs := memfs.New()
		_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{URL: "https://github.com/roger-king/tasker"})

		if err != nil {
			log.Fatal("Failed to pull migrations repo")
		} else {
			files, err := fs.ReadDir("db")

			if err != nil {
				log.Error(err)
			}

			for _, file := range files {
				if file.IsDir() {
					if file.Name() == "migrations" {
						// Using a differnt DB connection here because golang-migrate/migrate
						// does not support sqlx
						m, err := migrate.New(
							fmt.Sprintf("file://%s", "db/migrations"),
							tc.DBConnectionURL)

						if err != nil {
							log.Fatal("Failed to initialize connection:", err)
							os.Exit(1)
						}

						err = m.Up()

						if err != nil {
							log.Errorf("Failed to run migration:", err.Error())
							os.Exit(1)
						}
					}
				}
			}
		}
	}
	return db
}
