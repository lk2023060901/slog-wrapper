package slog_wrapper

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
	"sync"
)

// Logger 是 slog-wrapper 的核心结构体，封装了 slog.Logger 和 lumberjack.Logger
type Logger struct {
	*slog.Logger
	fileLogger *lumberjack.Logger
	mu         sync.Mutex
}

// NewLogger 创建一个新的 Logger 实例
func NewLogger(opts ...Option) *Logger {
	// 默认配置
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	// 创建 lumberjack.Logger 用于日志轮转
	fileLogger := &lumberjack.Logger{
		Filename:   cfg.File,
		MaxSize:    cfg.Rotation.MaxSize,
		MaxBackups: cfg.Rotation.MaxBackups,
		MaxAge:     cfg.Rotation.MaxAge,
		Compress:   cfg.Rotation.Compress,
	}

	var logger *slog.Logger

	switch cfg.OutputType {
	case OutputConsole:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: cfg.Level,
		}))
	case OutputFile:
		logger = slog.New(slog.NewTextHandler(fileLogger, &slog.HandlerOptions{
			Level: cfg.Level,
		}))
	case OutputBoth:
		multiWriter := NewMultiWriter(os.Stdout, fileLogger)
		logger = slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
			Level: cfg.Level,
		}))
	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: cfg.Level,
		}))
	}
	return &Logger{
		Logger:     logger,
		fileLogger: fileLogger,
	}
}

// Close 关闭文件日志记录器
func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.fileLogger.Close()
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		Logger:     l.Logger.With(args...),
		fileLogger: l.fileLogger,
	}
}

// MultiWriter 用于同时写入多个 io.Writer
type MultiWriter struct {
	writers []io.Writer
}

func (mw *MultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range mw.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
	}
	return len(p), err
}

func NewMultiWriter(writers ...io.Writer) *MultiWriter {
	return &MultiWriter{writers: writers}
}
