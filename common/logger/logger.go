package logger

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
)

// LogConfig config with log level
type LogConfig struct {
	LogLevel string `default:"info" envconfig:"LOG_LEVEL"`
}

// Config logger with level
var Config LogConfig

// NewLogger create new logger
func NewLogger() *logrus.Logger {
	var logger = logrus.New()
	readEnvConfig()
	logger.SetFormatter(NewGelf("x_"))
	logger.SetLevel(parseLogLevel(Config.LogLevel))

	return logger
}

// WithTraceInfo add to log tracing info
func WithTraceInfo(span opentracing.Span, log *logrus.Entry) *logrus.Entry {
	if sc, ok := span.Context().(jaeger.SpanContext); ok {
		traceID := fmt.Sprintf("%x", sc.TraceID().Low)
		log = log.WithFields(logrus.Fields{"uber-trace_id": traceID, "span_id": sc.SpanID().String()})
	}
	return log
}

// RequestLogger common class to log a request
func RequestLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// AppLogger set app name and version to the log
func AppLogger(appName string, appVersion string) *logrus.Entry {
	entry := NewLogger().
		WithFields(logrus.Fields{
			"service":     appName,
			"app_version": appVersion,
		})
	return entry
}

func readEnvConfig() {
	logrus.Debug("Reading envconfig for logger...")

	err := envconfig.Process("", &Config)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func parseLogLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "info":
		return logrus.InfoLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		logrus.Debugf("LogLevel %s not parsed. Using InfoLevel", logLevel)
		return logrus.InfoLevel
	}
}
