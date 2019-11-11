package services

import (
	"fmt"
	"plugin"
	"time"

	cron "github.com/robfig/cron/v3"
	"github.com/roger-king/tasker/models"
	log "github.com/sirupsen/logrus"
)

type RunnerFunc func()

func (t *TaskService) Runner(m *MongoService, task *models.Task) RunnerFunc {
	return func() {
		this, err := m.FindOne(task.TaskID)

		if err != nil {
			return
		}

		if this.Enabled && !this.Complete {
			plug, err := plugin.Open(fmt.Sprintf("./plugins/%s.so", this.Executor))

			if err != nil {
				fmt.Println(err)
				return
			}

			run, err := plug.Lookup("Run")

			if err != nil {
				fmt.Println(err)
				return
			}

			err = run.(func(map[string]interface{}) error)(this.Args)

			if err != nil {
				fmt.Print(err)
				return
			}

			if !this.IsRepeatable {
				this.Complete = true
				this.Enabled = false
				this.UpdatedAt = time.Now()
				this.DeletedAt = time.Now()

				err = m.Update(this)

				if err != nil {
					log.Error("Failed to mark as complete")
					return
				}

				t.Scheduler.Remove(cron.EntryID(this.EntryID))
			}
		}
	}
}
