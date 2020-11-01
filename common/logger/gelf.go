package logger

import (
	"encoding/json"
	"log/syslog"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// GelfVersion app version for logging
	GelfVersion = "1.0"
)

var (
	levelMap        map[logrus.Level]syslog.Priority
	syslogNameMap   map[syslog.Priority]string
	protectedFields map[string]bool
	// DefaultLevel default level of logging
	DefaultLevel = syslog.LOG_INFO
)

func init() {
	levelMap = map[logrus.Level]syslog.Priority{
		logrus.PanicLevel: syslog.LOG_EMERG,
		logrus.FatalLevel: syslog.LOG_CRIT,
		logrus.ErrorLevel: syslog.LOG_ERR,
		logrus.WarnLevel:  syslog.LOG_WARNING,
		logrus.InfoLevel:  syslog.LOG_INFO,
		logrus.DebugLevel: syslog.LOG_DEBUG,
		logrus.TraceLevel: syslog.LOG_DEBUG,
	}

	syslogNameMap = map[syslog.Priority]string{
		syslog.LOG_EMERG:   "emergency",
		syslog.LOG_ALERT:   "alert",
		syslog.LOG_CRIT:    "critical",
		syslog.LOG_ERR:     "error",
		syslog.LOG_WARNING: "warning",
		syslog.LOG_NOTICE:  "notice",
		syslog.LOG_INFO:    "info",
		syslog.LOG_DEBUG:   "debug",
	}
}

// GelFormatter log prefix
type GelFormatter struct {
	prefix string
}

// NewGelf new logs formatter
func NewGelf(prefix string) GelFormatter {
	return GelFormatter{prefix: prefix}
}

// Format format log entry
func (f GelFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	level := toSyslogLevel(entry.Level)
	levelName := syslogNameMap[level]

	gelfEntry := map[string]interface{}{
		"version":       GelfVersion,
		"short_message": entry.Message,
		"level":         levelName,
		"timestamp":     time.Now().Format(time.RFC3339Nano),
		"level_name":    levelName,
	}

	if _, file, line, ok := runtime.Caller(5); ok {
		gelfEntry["file"] = file
		gelfEntry["line"] = line
	}
	for key, value := range entry.Data {
		if !protectedFields[key] {
			key = f.prefix + key
		}
		gelfEntry[key] = value
	}
	message, err := json.Marshal(gelfEntry)
	return append(message, '\n'), err
}

func toSyslogLevel(level logrus.Level) syslog.Priority {
	syslog, ok := levelMap[level]

	if ok {
		return syslog
	}
	return DefaultLevel
}
