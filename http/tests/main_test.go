package tests_test

import (
	"os"
	"testing"

	"github.com/audryus/8mix/http/pkg/logger"
	"github.com/audryus/8mix/http/pkg/mongo"
	"github.com/audryus/8mix/http/tools/container"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	exitcode := m.Run()
	shutDown()
	os.Exit(exitcode)
}

func setup() {
	container.InitMongo()
}

func shutDown() {
	logger := logger.New()
	container.TerminateMongo(logger)
}

func Test(t *testing.T) {
	t.Run("context load", func(t *testing.T) {
		_, err := mongo.New()
		if err != nil {
			assert.Fail(t, "Mongo can't connect", err)
			return
		}
	})
}
