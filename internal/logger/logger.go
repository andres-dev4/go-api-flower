// internal/logger/logger.go
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableColors:   false, // true si quieres log limpio sin colores (útil en archivos)
		DisableQuote:    true,  // evita que strings simples vengan entre comillas
	})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}
