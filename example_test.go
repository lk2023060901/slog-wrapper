package slog_wrapper

import (
	"log/slog"
	"os"
	"testing"
)

// TestDefaultLogger 测试默认配置的 Logger
func TestDefaultLogger(t *testing.T) {
	logger := NewLogger()
	defer func(logger *Logger) {
		err := logger.Close()
		if err != nil {
			return
		}
	}(logger)

	logger.Info("This is an info message with default configuration")
	logger.Error("This is an error message with default configuration", slog.String("error", "something went wrong"))
}

// TestCustomLevelLogger 测试自定义日志级别的 Logger
func TestCustomLevelLogger(t *testing.T) {
	logger := NewLogger(
		WithLevel(slog.LevelDebug),
	)
	defer func(logger *Logger) {
		err := logger.Close()
		if err != nil {
			return
		}
	}(logger)

	logger.Debug("This is a debug message with custom level")
	logger.Info("This is an info message with custom level")
	logger.Warn("This is a warn message with custom level")
	logger.Error("This is an error message with custom level", slog.String("error", "custom level error"))
}

// TestFileLogger 测试仅输出到文件的 Logger
func TestFileLogger(t *testing.T) {
	logger := NewLogger(
		WithOutputType(OutputFile),
		WithFile("file_only.log"),
	)
	defer func(logger *Logger) {
		err := logger.Close()
		if err != nil {
			return
		}
	}(logger)

	logger.Info("This message is logged only to the file")
	logger.Error("This error is logged only to the file", slog.String("error", "file only error"))
}

// TestBothOutputLogger 测试同时输出到控制台和文件的 Logger
func TestBothOutputLogger(t *testing.T) {
	logger := NewLogger(
		WithOutputType(OutputBoth),
		WithFile("both_output.log"),
	)
	defer func(logger *Logger) {
		err := logger.Close()
		if err != nil {
			return
		}
	}(logger)

	logger.Info("This message is logged to both console and file")
	logger.Error("This error is logged to both console and file", slog.String("error", "both output error"))
}

// TestRotationLogger 测试日志轮转功能的 Logger
func TestRotationLogger(t *testing.T) {
	logger := NewLogger(
		WithOutputType(OutputFile),
		WithFile("rotation.log"),
		WithRotation(1, 3, 7, true), // 1MB 文件大小，保留 3 个备份，保留 7 天，压缩旧日志
	)
	defer func(logger *Logger) {
		err := logger.Close()
		if err != nil {
			return
		}
	}(logger)

	// 模拟写入大量日志以触发轮转
	for i := 0; i < 10000; i++ {
		logger.Info("This is a log message to test rotation", slog.Int("index", i))
	}
}

// TestCustomFormatLogger 测试自定义日志格式的 Logger
func TestCustomFormatLogger(t *testing.T) {
	// 自定义 slog.Handler 以实现自定义格式
	customHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{} // 移除时间戳
			}
			return a
		},
	})

	logger := NewLogger(
		WithLevel(slog.LevelDebug),
		WithOutputType(OutputConsole),
	)
	logger.Logger = slog.New(customHandler) // 替换默认的 Handler

	logger.Debug("This is a debug message with custom format")
	logger.Info("This is an info message with custom format")
}

// TestCloseLogger 测试关闭 Logger 的功能
func TestCloseLogger(t *testing.T) {
	logger := NewLogger(
		WithOutputType(OutputFile),
		WithFile("close_test.log"),
	)

	logger.Info("This message is logged before closing the logger")
	err := logger.Close()
	if err != nil {
		return
	}

	// 尝试在关闭后记录日志（应该不会生效）
	logger.Info("This message should not be logged after closing the logger")
}

// TestLoggerWith 测试 slog 的结构化日志输出功能
func TestLoggerWith(t *testing.T) {
	logger := NewLogger(
		WithOutputType(OutputFile),
		WithLevel(slog.LevelDebug),
		WithFile("item.log"),
	)
	logger = logger.With(slog.String("item-name", "大宝剑"), slog.String("item-count", "1"))
	defer func(logger *Logger) {
		err := logger.Close()
		if err != nil {
			return
		}
	}(logger)
	logger.Debug("This is a debug message with custom format")
}
