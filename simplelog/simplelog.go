package simplelog

import (
	"io"
	"log"
	"os"

	"github.com/bugfan/logair"
)

func init() {
	logair.RegisterLogger(&SimpleLogger{})
}

type SimpleLogger struct {
	logLevel logair.LogLevel
}

func (l *SimpleLogger) SetLogLevel(level logair.LogLevel) {
	l.logLevel = level
}

func (l *SimpleLogger) Fatal(v ...interface{}) {
	if l.logLevel <= logair.FatalLevel {
		log.Fatal(v...)
	}
	os.Exit(1)
}

func (l *SimpleLogger) Fatalf(format string, v ...interface{}) {
	if l.logLevel <= logair.FatalLevel {
		log.Fatalf(format, v...)
	}
	os.Exit(1)
}

func (l *SimpleLogger) Error(v ...interface{}) {
	if l.logLevel <= logair.ErrorLevel {
		log.Println(v...)
	}
}

func (l *SimpleLogger) Errorf(format string, v ...interface{}) {
	if l.logLevel <= logair.ErrorLevel {
		log.Printf(format, v...)
	}
}

func (l *SimpleLogger) Warn(v ...interface{}) {
	if l.logLevel <= logair.WarnLevel {
		log.Println(v...)
	}
}

func (l *SimpleLogger) Warnf(format string, v ...interface{}) {
	if l.logLevel <= logair.WarnLevel {
		log.Printf(format, v...)
	}
}

func (l *SimpleLogger) Info(v ...interface{}) {
	if l.logLevel <= logair.InfoLevel {
		log.Println(v...)
	}
}

func (l *SimpleLogger) Infof(format string, v ...interface{}) {
	if l.logLevel <= logair.InfoLevel {
		log.Printf(format, v...)
	}
}

func (l *SimpleLogger) Debug(v ...interface{}) {
	if l.logLevel <= logair.AllLevel {
		log.Println(v...)
	}
}

func (l *SimpleLogger) Debugf(format string, v ...interface{}) {
	if l.logLevel <= logair.AllLevel {
		log.Printf(format, v...)
	}
}

func (l *SimpleLogger) Trace(v ...interface{}) {
	if l.logLevel <= logair.AllLevel {
		log.Println(v...)
	}
}

func (l *SimpleLogger) Tracef(format string, v ...interface{}) {
	if l.logLevel <= logair.AllLevel {
		log.Printf(format, v...)
	}
}

func (l *SimpleLogger) SetOutput(io.Writer) {
	// do nothing
}
