package logger

import "github.com/sirupsen/logrus"

type LoggerConfig struct {
	level                 logrus.Level
	includeRequestHeaders bool
}

type LoggerConfigBuilder struct {
	loggerConfig *LoggerConfig
}

func NewLoggerConfigBuilder() *LoggerConfigBuilder {
	loggerConfig := &LoggerConfig{
		level:                 logrus.InfoLevel,
		includeRequestHeaders: false,
	}

	return &LoggerConfigBuilder{
		loggerConfig: loggerConfig,
	}
}

func (lcb *LoggerConfigBuilder) WithLoggingLevel(level logrus.Level) *LoggerConfigBuilder {
	lcb.loggerConfig.level = level
	return lcb
}

func (lcb *LoggerConfigBuilder) WithRequestHeadersIncluded() *LoggerConfigBuilder {
	lcb.loggerConfig.includeRequestHeaders = true
	return lcb
}

func (lcb *LoggerConfigBuilder) Build() LoggerConfig {
	return *lcb.loggerConfig
}
