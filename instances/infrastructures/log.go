package infrastructures

import (
	"fmt"
	"sync"

	structures "github.com/burhanmubarok/microservice/structures/infrastructures"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	conf "github.com/spf13/viper"
)

// Logger doc
type Logger struct {
	log     *log.Logger
	content interface{}
	mutex   sync.Mutex
}

func (l *Logger) set(data interface{}) {
	l.content = data
}

// Debug doc
func (l *Logger) Debug(data interface{}) {
	l.set(data)
	l.write("debug")
}

// Info doc
func (l *Logger) Info(data structures.Log) {
	// l.write("info", data)
}

// Warning doc
func (l *Logger) Warning(data structures.Log) {
	// l.write("warning", data)
}

// Error doc
func (l *Logger) Error(data structures.Log) {
	// l.write("error", data)
}

func (l *Logger) write(level string) {
	l.init()

	switch level {
	case "info", "warning":
		l.log.WithFields(log.Fields{
			// "data": data,
		}).Info(l.content)
	case "error":
		l.log.WithFields(log.Fields{
			// "stacks": data,
		}).Error(l.content)
	default:
		l.log.WithFields(log.Fields{
			"data": l.content,
		}).Debug(l.content)
	}
}

func (l *Logger) init() {
	l.log = log.New()
	l.log.SetLevel(log.DebugLevel)
	l.log.Formatter = &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: log.FieldMap{
			"time": "timestamp",
			"msg":  "message",
		},
	}

	logName := "go.log"
	logPathConf := conf.GetString("microservice.logPath")
	logPath := fmt.Sprintf("%s/%s", logPathConf, logName)
	l.log.Out = &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1000, // megabytes after which new file is created
		MaxBackups: 3,    // number of backups
		MaxAge:     30,   //days
	}
}
