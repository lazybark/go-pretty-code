package logs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lazybark/go-pretty-code/console"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	sugaredFile    *zap.SugaredLogger
	sugaredConsole *zap.SugaredLogger
	file           *os.File
	Level          zapcore.Level
}

func Double(filepath string, truncate bool, level zapcore.Level) (*Logger, error) {
	var l Logger

	logFile, err := l.MakeLogPath(filepath, truncate)
	if err != nil {
		return nil, err
	}

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewConsoleEncoder(pe)
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	// Console-only logger
	coreCon := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)
	lCon := zap.New(coreCon)

	// File-only logger
	coreFile := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), level),
	)
	lFile := zap.New(coreFile)

	l.file = logFile
	l.sugaredFile = lFile.Sugar()
	l.sugaredConsole = lCon.Sugar()
	l.Level = level

	return &l, nil
}

func FileOnly(filepath string, truncate bool, level zapcore.Level) (*Logger, error) {
	var l Logger

	logFile, err := l.MakeLogPath(filepath, truncate)
	if err != nil {
		return nil, err
	}

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewConsoleEncoder(pe)

	// File-only logger
	coreFile := zap.New(zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), level))

	l.file = logFile
	l.sugaredFile = coreFile.Sugar()
	l.Level = level

	return &l, nil
}

func ConsoleOnly(level zapcore.Level) *Logger {
	var l Logger

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	// Console-only logger
	coreCon := zap.New(zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level))

	l.sugaredConsole = coreCon.Sugar()
	l.Level = level

	return &l
}

func (l Logger) MakeLogPath(path string, truncate bool) (*os.File, error) {
	dir, _ := filepath.Split(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}

	flags := os.O_CREATE | os.O_APPEND
	if truncate {
		flags += os.O_TRUNC
	}

	logFile, err := os.OpenFile(path, flags, 0666)
	if err != nil {
		return nil, err
	}

	return logFile, nil
}

func (l Logger) Info(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(fmt.Sprint(args...))
	}
}

func (l Logger) Error(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Error(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Error(console.ForeYellow(fmt.Sprint(args...)))
	}
}

func (l Logger) ErrorBackYellow(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Error(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Error(console.BackYellow(fmt.Sprint(args...)))
	}
}

func (l Logger) Warn(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Warn(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Warn(console.ForeYellow(fmt.Sprint(args...)))
	}
}

func (l Logger) WarnBackYellow(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Warn(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Warn(console.BackYellow(fmt.Sprint(args...)))
	}
}

func (l Logger) Fatal(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Fatal(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Fatal(console.ForeRed(fmt.Sprint(args...)))
	}
}

func (l Logger) FatalBackRed(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Fatal(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Fatal(console.BackRed(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoRed(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.ForeRed(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBackRed(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.BackRed(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoGreen(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.ForeGreen(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBackGreen(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.BackGreen(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoYellow(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.ForeYellow(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBackYellow(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.BackYellow(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBlue(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.ForeBlue(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBackBlue(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.BackBlue(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoMagenta(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.ForeMagenta(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBackMagenta(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.BackMagenta(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoCyan(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.ForeCyan(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBackCyan(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.BackCyan(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBlink(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.TextBlink(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoUnderline(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.TextUnderline(fmt.Sprint(args...)))
	}
}

func (l Logger) InfoBold(args ...interface{}) {
	if l.sugaredFile != nil {
		l.sugaredFile.Info(fmt.Sprint(args...))
	}
	if l.sugaredConsole != nil {
		l.sugaredConsole.Info(console.TextBold(fmt.Sprint(args...)))
	}
}
