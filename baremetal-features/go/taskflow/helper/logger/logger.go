package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	infoLog *log.Logger
	errorLog *log.Logger
}

func New() (*Logger, error) {
	//& The absolute path file logging
		exePath, err := os.Executable()
		if err != nil {
			return nil, fmt.Errorf("failed to get executable path: %w", err)
		}
		baseDir := filepath.Dir(exePath)

		logDir := filepath.Join(baseDir, "logs")
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}

		logFilePath := filepath.Join(logDir, "log.txt")

	//& For muti destnation logging
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //! os.O_APPEND to do not delete anything|| (0644) is The Permissions, others read only 
		if err != nil { log.Fatal(err) }
		//! We don't defer file.Close() here because we need the file to stay open for the life of the Logger!
		multiInfo := io.MultiWriter(file, os.Stdout)
		multiErr := io.MultiWriter(file, os.Stderr)
	
	return &Logger{
		infoLog:  log.New(multiInfo, "[INFO]", log.Ldate|log.Ltime), //! log.New() over log.Println() for more organization
		errorLog: log.New(multiErr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}, nil
}

func (l *Logger) Info(msg string) {
	l.infoLog.Println(msg)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.infoLog.Printf(format, args...)
}

func (l *Logger) Error(msg string) {
	l.errorLog.Println(msg)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.errorLog.Printf(format, args...)
}

func (l *Logger) Fatal(msg string) {
	l.errorLog.Fatal(msg)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.errorLog.Fatalf(format, args...)
}