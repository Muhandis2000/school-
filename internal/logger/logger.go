// internal/logger/logger.go
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.DebugLevel)
}
