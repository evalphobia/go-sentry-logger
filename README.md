# go-sentry-logger
Go logger with sentry

this supports logging error data and send it to sentry.

# Configure

set sentry's dsn

```sh
vim log.json

{
    "sentry": "https://<your key>@app.getsentry.com/<your id>",
}

```

# Quick Usage

```go
import(
    log "github.com/evalphobia/go-sentry-logger"
)

func main(){
    _, err := someFunction() // dummy code

    if err != nil {
        e := log.NewData(err, 0) // log.NewData(any data, stack trace depth)
        e.Label = "[ERROR] unknown error occured"
        e.Request = req // set *http.Request

        log.Fatal(err) // logging & send data to senty
        log.Error(err) // logging & send data to senty
        log.Warn(err)  // logging & send data to senty
        log.Info(err)  // do NOT send data to senty
        log.Debug(err) // do NOT send data to senty
        
        log.SentryLevel = 4 // default 3
        log.Fatal(err) // logging & send data to senty, level=5
        log.Error(err) // logging & send data to senty, level=4
        log.Warn(err)  // do NOT send data to senty, level=3
    }
    
    // use for print debugging, 
    log.PrintHeader("!!debug data!!")
    log.Print(err, 0)
}
```

# License

MIT

