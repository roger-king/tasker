# Tasker Plugins

Tasker Plugins are at its core. Go Plugins. The mantra of tasker is scheduling tasks without reloading your server. Utilizing Go Plugins allows us to build and test our tasks in isolation. 

## Skeleton of a task plugin

```go
    import "github.com/roger-king/tasker/pkg"

    type payload struct {
        Name string `json:"name"`
    }

    func Run(args map[string]interface{}) error {
        var p payload

        err := pkg.ValidateArgs(args, &p)

        if err != nil {
            return err
        }

        fmt.Println("Hello Tasker: ", p.Name)
        return nil
    }
```

Here we are first setting up a payload struct. This will allow us to verify the args passed in are indeed valid args (key/value) map.