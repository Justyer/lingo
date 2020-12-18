package log

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rs/zerolog"
)

func init() {
	ConfigHook = func() {}
}

var ConfigHook func()

var logger zerolog.Logger

func PrepareAll(name, path string) {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	ConfigHook()
	writer, err := rotatelogs.New(
		path,
		rotatelogs.WithLinkName(name),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}

	mulW := io.MultiWriter(os.Stderr, writer)
	logger = zerolog.New(mulW).With().Timestamp().Logger()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error()
}
