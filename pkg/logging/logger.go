package logger

import (
	"io"
	"log"
	"os"
)

// Logger представляет собой структурированный логгер
type Logger struct {
	infoLog    *log.Logger
	errorLog   *log.Logger
	warningLog *log.Logger
	debugLog   *log.Logger
}

// New создает новый экземпляр логгера с настраиваемыми выходами
func New(infoWriter, errorWriter io.Writer) *Logger {
	if infoWriter == nil {
		infoWriter = os.Stdout
	}
	if errorWriter == nil {
		errorWriter = os.Stderr
	}

	return &Logger{
		infoLog:    log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog:   log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLog: log.New(infoWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLog:   log.New(infoWriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info логирует информационное сообщение
func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLog.Printf(format, v...)
}

// Error логирует сообщение об ошибке
func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLog.Printf(format, v...)
}

// Warning логирует предупреждение
func (l *Logger) Warning(format string, v ...interface{}) {
	l.warningLog.Printf(format, v...)
}

// Debug логирует отладочное сообщение
func (l *Logger) Debug(format string, v ...interface{}) {
	l.debugLog.Printf(format, v...)
}

// Default возвращает логгер с настройками по умолчанию
func Default() *Logger {
	return New(os.Stdout, os.Stderr)
}
