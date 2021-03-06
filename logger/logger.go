package logger

import (
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.org/pong3ds/play-with-echo/uuid"
)

// ILogger is for logging
type ILogger interface {
	Info(message string)
	Warn(message string)
	Debug(message string)
	Error(message string)
}

// Logger is the logger utility with information of request context
type Logger struct {
	Type       string
	ProcessID  string
	TrackingID string
	SourceIP   string
	AppID      string
	HTTPMethod string
	EndPoint   string
}

var instance *Logger
var once sync.Once

// CreateLogger will create the logger with context from echo context
func CreateLogger(c echo.Context, uuid uuid.IUUID, level log.Level) {
	once.Do(func() {

		formatter := new(log.JSONFormatter)
		formatter.TimestampFormat = "2018-12-30 23:05:05"
		formatter.DisableTimestamp = false
		log.SetFormatter(formatter)

		file, _ := os.OpenFile("mylog.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
		log.SetLevel(level)

		instance = &Logger{
			Type:       "REQUEST",
			ProcessID:  uuid.GetUUID(),
			SourceIP:   c.Request().RemoteAddr,
			HTTPMethod: c.Request().Method,
			EndPoint:   c.Request().URL.Path,
			TrackingID: "", // User ID
			AppID:      "", // App ID
		}
	})
}

// GetLogger return logger
func GetLogger() ILogger {
	return instance
}

func (logger *Logger) getLogFields(fn string, line int) log.Fields {
	return log.Fields{
		"type":        logger.Type,
		"process_id":  logger.ProcessID,
		"tracking_id": logger.TrackingID,
		"source_ip":   logger.SourceIP,
		"app_id":      logger.AppID,
		"http_method": logger.HTTPMethod,
		"endpoint":    logger.EndPoint,
		"function":    fn,
		"line":        line,
	}
}

// Info log information level
func (logger *Logger) Info(message string) {
	_, fn, line, _ := runtime.Caller(1)
	log.WithFields(logger.getLogFields(fn, line)).Info(message)
}

// Warn log warnning level
func (logger *Logger) Warn(message string) {
	_, fn, line, _ := runtime.Caller(1)
	log.WithFields(logger.getLogFields(fn, line)).Warn(message)
}

// Debug log debug level
func (logger *Logger) Debug(message string) {
	_, fn, line, _ := runtime.Caller(1)
	log.WithFields(logger.getLogFields(fn, line)).Debug(message)
}

// Error log error level
func (logger *Logger) Error(message string) {
	_, fn, line, _ := runtime.Caller(1)
	log.WithFields(logger.getLogFields(fn, line)).Error(message)
}
