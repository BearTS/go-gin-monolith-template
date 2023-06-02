package redis

import (
	"testing"
	"time"

	"github.com/BearTS/go-gin-monolith/config"
	"github.com/stretchr/testify/assert"
)

var testConn Connection

func init() {
	config.LoadConfigs()
}

func TestNewConnection(t *testing.T) {
	err := testConn.NewConnection()
	assert.Empty(t, err)
}

func TestGetClient(t *testing.T) {
	TestNewConnection(t)
	client := testConn.GetClient()
	assert.NotEmpty(t, client)
}

func TestSet(t *testing.T) {
	TestNewConnection(t)
	key := "_test_tez_key"
	value := "_test_tez_val"
	err := testConn.Set(key, value)
	assert.Empty(t, err)
}

func TestSetWithTimeout(t *testing.T) {
	TestNewConnection(t)
	key := "_test_tez_key_t"
	value := "_test_tez_val_t"
	err := testConn.SetWithTimeout(key, value, 2*time.Second)
	assert.Empty(t, err)
}

func TestGet(t *testing.T) {
	TestNewConnection(t)
	TestSet(t)
	key := "_test_tez_key"
	value := "_test_tez_val"
	str, err := testConn.Get(key)
	assert.Empty(t, err)
	assert.Equal(t, value, str)
}

func TestDel(t *testing.T) {
	TestNewConnection(t)
	TestSet(t)
	key := "_test_tez_key"
	err := testConn.Del(key)
	assert.Empty(t, err)
}

func TestDelMilti(t *testing.T) {
	TestNewConnection(t)
	TestSet(t)
	TestSetWithTimeout(t)
	keys := []string{"_test_tez_key", "_test_tez_key_t"}
	err := testConn.DelMulti(keys)
	assert.Empty(t, err)
}

func TestClose(t *testing.T) {
	TestNewConnection(t)
	err := testConn.Close()
	assert.Empty(t, err)

}
