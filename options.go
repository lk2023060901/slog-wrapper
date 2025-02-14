package slog_wrapper

import "log/slog"

// Option 用于自定义配置的函数
type Option func(config *Config)

// WithLevel 设置日志级别
func WithLevel(level slog.Level) Option {
	return func(cfg *Config) {
		cfg.Level = level
	}
}

// WithOutputType 设置日志输出方式
func WithOutputType(outputType OutputType) Option {
	return func(cfg *Config) {
		cfg.OutputType = outputType
	}
}

// WithFile 设置日志文件路径
func WithFile(file string) Option {
	return func(cfg *Config) {
		cfg.File = file
	}
}

func WithRotation(maxSize, maxBackups, maxAge int, compress bool) Option {
	return func(cfg *Config) {
		cfg.Rotation.MaxSize = maxSize
		cfg.Rotation.MaxBackups = maxBackups
		cfg.Rotation.MaxAge = maxAge
		cfg.Rotation.Compress = compress
	}
}
