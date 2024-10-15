package temporal

import (
	"github.com/audryus/8mix/http/pkg/logger"
	"go.temporal.io/sdk/client"
)

func New(logger *logger.Log) (client.Client, error) {
	return client.Dial(client.Options{
		Logger:   logger,
		Identity: "8mix.http",
	})
}
