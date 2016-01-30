package logging


import (
    "os"

    "github.com/op/go-logging"
)

var Logger = logging.MustGetLogger("root")

// Everything except the message has a custom color
// which is dependent on the log level.
var format = logging.MustStringFormatter(
    `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)


func init() {
    logBackend := logging.NewLogBackend(os.Stderr, "", 0)
    backendFormatter := logging.NewBackendFormatter(logBackend, format)
    logging.SetBackend(backendFormatter)
}

