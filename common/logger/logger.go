package logger

import (
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger standart logger
type Logger struct {
	*logrus.Logger
}

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// Log standart logger
var Log *Logger

func init() {
	var baseLogger = logrus.New()

	Log = &Logger{baseLogger}

	Log.Formatter = &logrus.JSONFormatter{}
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

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
)

// InvalidArg is a standard error message
func (l *Logger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *Logger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *Logger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}
