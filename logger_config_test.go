package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := NewLoggerConfigBuilder().Build()

	assert.NotNil(t, config)
	assert.False(t, config.includeRequestHeaders)
	assert.Equal(t, logrus.InfoLevel, config.level)
}

func TestConfigWithDebugLevel(t *testing.T) {
	config := NewLoggerConfigBuilder().
		WithLoggingLevel(logrus.DebugLevel).
		Build()

	assert.NotNil(t, config)
	assert.False(t, config.includeRequestHeaders)
	assert.Equal(t, logrus.DebugLevel, config.level)
}

func TestConfigWithIncludedHeaders(t *testing.T) {
	config := NewLoggerConfigBuilder().
		WithRequestHeadersIncluded().
		Build()

	assert.NotNil(t, config)
	assert.True(t, config.includeRequestHeaders)
	assert.Equal(t, logrus.InfoLevel, config.level)
}

func TestConfigWithDebugLevelAndIncludedHeaders(t *testing.T) {
	config := NewLoggerConfigBuilder().
		WithLoggingLevel(logrus.DebugLevel).
		WithRequestHeadersIncluded().
		Build()

	assert.NotNil(t, config)
	assert.True(t, config.includeRequestHeaders)
	assert.Equal(t, logrus.DebugLevel, config.level)
}
