[![Actions Status](https://github.com/roger-king/tasker/workflows/Backend/badge.svg)](https://github.com/roger-king/tasker/actions)
[![Actions Status](https://github.com/roger-king/tasker/workflows/Frontend/badge.svg)](https://github.com/roger-king/tasker/actions)

<p align="center">
    <img src="./www/public/images/flash.png" width="200px">
</p>

# Tasker - [ DEVELOPMENT IN PROGRESS ]

Tasker is a job scheduling service that utilizes `go-cron`. Tasker is a service that will allow for dynamic configration of tasks.

Tasker is inspired by the Android application `tasker` the ability to set scripts and schedule them at whim without needing to touch the server code.

<img src="./static/wip_ui_v1.0.3.png" width="800px" align="center">

# Installation

Import tasker:

```bash
    import "github.com/roger-king/tasker"
```

Install via Go Modules:

```bash
    go mod vendor
```

# Getting Started:

Tasker is simple to get started. Tasker has 2 goals at mind:

- Create a schedule for task to run external scripts
- Create your own handler functions to programmatically run code

```golang
	t, err := tasker.New(&config.TaskerConfig{
		Migrate: false,
		Auth:    true,
		DBConnectionURL: "postgres://appuser:appuser@localhost:5432/tasker",
	})

	if err != nil {
		log.Panic(err)
	}

	router := t.Start()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
```

Go to your server and find the tasker web client at `/tasker/admin`.

## Configuration of your server

Tasker provides a `TaskerConfig` that helps configure your server. I recommend you set those values with a secret manager or using [`12 Factor App configs`](https://12factor.net/config).

| ENV                      |                 Description                  |
| ------------------------ | :------------------------------------------: |
| Migrate                  | true/false flag for auto migrate on start up |
| TASKER_DB_CONNECTION_URL |           Postgres connection uri            |
| GITHUB_CLIENT_ID         |     Github Client ID for authentication      |
| GITHUB_CLIENT_SECRET     |       Github Secret for authentication       |

## TODOS:

View the Github Projects board: https://github.com/roger-king/tasker/projects/1
