package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadAppConfig(t *testing.T) {
	LoadConfigs()
	assert.NotEmpty(t, App.Env)
	assert.NotEmpty(t, App.RootPath)
	assert.Equal(t, 0, App.Migrate)
}

func TestLoadDbConfig(t *testing.T) {
	LoadConfigs()
	assert.NotEmpty(t, DB.Database)
	assert.NotEmpty(t, DB.Host)
	assert.NotEmpty(t, DB.Password)
	assert.NotEmpty(t, DB.Port)
	assert.NotEmpty(t, DB.Username)
}

func TestLoadRedisConfig(t *testing.T) {
	LoadConfigs()
	assert.NotEmpty(t, Redis.Host)
	assert.NotEmpty(t, Redis.Port)
	assert.NotEmpty(t, Redis.DB)
	assert.NotEmpty(t, Redis.MaxRetries)
	assert.NotEmpty(t, Redis.MinRetryBackoffMs)
	assert.NotEmpty(t, Redis.MaxRetryBackoffMs)
	assert.NotEmpty(t, Redis.WriteTimeoutMs)
}
