package logger

import (
	"log"
	"os"
	"path/filepath"
)

// Logger 日志记录器
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// New 创建新的日志记录器
func New() *Logger {
	homeDir, _ := os.UserHomeDir()
	logDir := filepath.Join(homeDir, ".hitryremote", "logs")
	os.MkdirAll(logDir, 0755)

	infoFile, _ := os.OpenFile(filepath.Join(logDir, "info.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errorFile, _ := os.OpenFile(filepath.Join(logDir, "error.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	debugFile, _ := os.OpenFile(filepath.Join(logDir, "debug.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	return &Logger{
		infoLogger:  log.New(infoFile, "[INFO] ", log.LstdFlags),
		errorLogger: log.New(errorFile, "[ERROR] ", log.LstdFlags),
		debugLogger: log.New(debugFile, "[DEBUG] ", log.LstdFlags),
	}
}

// Info 记录信息日志
func (l *Logger) Info(msg string) {
	l.infoLogger.Println(msg)
}

// Error 记录错误日志
func (l *Logger) Error(msg string) {
	l.errorLogger.Println(msg)
}

// Debug 记录调试日志
func (l *Logger) Debug(msg string) {
	l.debugLogger.Println(msg)
}
