package logging

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"sync"
)

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var once sync.Once
var zeroSingleLogger *zerolog.Logger

var logLevel = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
}

func newZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}

	logger.Init()

	return logger
}

func (l *zeroLogger) getLevel() zerolog.Level {
	level, exists := logLevel[l.cfg.Logger.Level]

	if !exists {
		return zerolog.DebugLevel
	}

	return level
}

func (l *zeroLogger) Init() {
	once.Do(func() {

		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

		if err != nil {
			panic("could not open log file")
		}

		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("AppName", "MyApp").
			Str("LoggerName", "Zerolog").
			Logger()

		zerolog.SetGlobalLevel(l.getLevel())
		zeroSingleLogger = &logger
	})
	l.logger = zeroSingleLogger
}

func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareZeroParams(extra)
	l.logger.Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.
		Debug().
		Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareZeroParams(extra)
	l.logger.Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.
		Info().
		Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareZeroParams(extra)
	l.logger.Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.
		Warn().
		Msgf(template, args...)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareZeroParams(extra)
	l.logger.Error().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.
		Error().
		Msgf(template, args...)
}

func (l *zeroLogger) Panic(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareZeroParams(extra)
	l.logger.Panic().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Panicf(template string, args ...interface{}) {
	l.logger.
		Panic().
		Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareZeroParams(extra)
	l.logger.Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.
		Fatal().
		Msgf(template, args...)
}
