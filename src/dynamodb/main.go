package dynamodb

import (
	"models"
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("http-handler")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x} ▶ %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

// Client interface
type Client interface {
	// GetGameByID() <-chan struct{}
	GetGameByID(gameID string) (models.Game, error)
}
