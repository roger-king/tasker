# Tasker - [ DEVELOPEMENT IN PROGRESS ]

Tasker is a job scheduling service that utilizes `go-cron`. Tasker is a service that will allow for dynamic configration of tasks.

Tasker is inspired by the Android application `tasker` the ability to set scripts and schedule them at whim without needing to touch the server code.

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
    t := tasker.New()
	router := t.Start()

    // Add to your HTTP Server
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
```

Go to your server and find the tasker web client at `/tasker`.

## Configuration of your server

Tasker utilizes [`12 Factor App configs`](https://12factor.net/config). Tasker relies environment variables to handle configruation of the application

Below are the availble of environment variables that Tasker looks for (note: prefixed with `TASKER_`):

| ENV                |                                        Description                                         |
| ------------------ | :----------------------------------------------------------------------------------------: |
| TASKER_DB_TYPE     |        type of presistent data store for your scheduled tasks (e.g. redis or mongo)        |
| TASKER_DB_HOST     |                            host address for your database type                             |
| TASKER_DB_USER     |                       username used to authenticate against your db                        |
| TASKER_DB_PASSWORD |                       password used to authenticate against your db                        |
| TASKER_DB_PORT     | port your database is listening on (if not specified will use the default your desired db) |
| TASKER_SCRIPT_ROOT |                        root directory where your scripts are stored                        |
