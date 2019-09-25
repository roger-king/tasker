package pkg

import "github.com/robfig/cron"

// Scheduler -
type Scheduler struct {
	ScriptDirectory string
	Cron            *cron.Cron
}
