# Tasker Plugins

Tasker Plugins are at its core. Go Plugins. The mantra of tasker is scheduling tasks without reloading your server. Utilizing Go Plugins allows us to build and test our tasks in isolation. 

## Skeleton of a task plugin

```go
    import "github.com/roger-king/tasker/utils"

    type payload struct {
        Name string `json:"name" validate:"required"`
    }

    func Run(args map[string]interface{}) error {
        var p payload

        err := utils.Validate(args, p)

        if err != nil {
            return err
        }

        fmt.Println("Hello Tasker: ", p.Name)
        return nil
    }
```

Here we are first setting up a payload struct. This will allow us to verify the args passed in are indeed valid args (key/value) map. We achieve this with a nice library called [validator](https://github.com/go-playground/validator) by go-playground