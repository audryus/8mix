package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type iLog interface {
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Warn(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
}

type Log struct {
	logger zerolog.Logger
}

var _ iLog = (*Log)(nil)

func New() *Log {
	l := zerolog.InfoLevel
	zerolog.SetGlobalLevel(l)
	skipFrameCount := 3

	return &Log{
		logger: zerolog.New(os.Stdout).
			With().
			Timestamp().
			CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).
			Logger(),
	}
}

func (l *Log) Core() *zerolog.Logger {
	return &l.logger
}

func (l *Log) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

func (l *Log) Debug(message string, args ...interface{}) {
	l.log(message, args...)
}

func (l *Log) Warn(message string, args ...interface{}) {
	l.log(message, args...)
}

func (l *Log) Error(message string, args ...interface{}) {
	l.log(message, args...)
}

func (l *Log) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}
