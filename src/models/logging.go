package models

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("update:game")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

// Secret is an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Secret string

// Redacted will hide sensitive information in logs
func (p Secret) Redacted() interface{} {
	return logging.Redact(string(p))
}

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)

	// log.Debugf("debug %s", Secret("secret"))
	// log.Info("info")
	// log.Notice("notice")
	// log.Warning("warning")
	// log.Error("err")
	// log.Critical("crit")
}
