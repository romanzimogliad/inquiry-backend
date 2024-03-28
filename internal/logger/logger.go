package logger

import (
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/romanzimoglyad/inquiry-backend/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

//nolint:gochecknoglobals //ok
var log zerolog.Logger

func Init() zerolog.Logger {
	var once sync.Once
	once.Do(func() {
		//nolint: reassign //ok
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel := config.Config.LogLevel

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if config.Config.Env != config.EnvTypeDev.String() {
			fileLogger := &lumberjack.Logger{
				Filename:   "inquiry.log",
				MaxSize:    100, //
				MaxBackups: 10,
				MaxAge:     14,
				Compress:   true,
			}

			output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
		}

		var gitRevision string

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().Caller().
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
	})

	return log
}

func Info() *zerolog.Event {
	return log.Info()
}

func Fatal() *zerolog.Event {
	return log.Fatal()
}

func Debug() *zerolog.Event {
	return log.Debug()
}

func Error() *zerolog.Event {
	return log.Error()
}
