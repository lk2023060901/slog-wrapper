package slog_wrapper

import "log/slog"

type OutputType int

const (
	OutputConsole OutputType = iota // 仅输出到控制台
	OutputFile                      // 仅输出到文件
	OutputBoth                      // 同时输出到控制台和文件
)

// RotationConfig 定义日志轮转配置
type RotationConfig struct {
	MaxSize    int  // 单个日志文件到最大大小（MB）
	MaxBackups int  // 保留旧日志文件的最大数量
	MaxAge     int  // 保留旧日志文件到最大天数
	Compress   bool // 是否压缩旧日志文件
}

// Config 定义日志库的配置
type Config struct {
	Level      slog.Level     // 日志级别
	OutputType OutputType     // 输出方式
	File       string         // 日志文件路径
	Rotation   RotationConfig // 日志轮转配置
}

// DefaultConfig 返回默认的日志配置
func DefaultConfig() *Config {
	return &Config{
		Level:      slog.LevelInfo,
		OutputType: OutputConsole,
		File:       "app.log",
		Rotation: RotationConfig{
			MaxSize:    100,  // 默认单个日志文件最大 100MB
			MaxBackups: 3,    // 默认保留 3 个旧日志文件
			MaxAge:     30,   // 默认保留 30 天
			Compress:   true, // 默认压缩旧日志文件
		},
	}
}
