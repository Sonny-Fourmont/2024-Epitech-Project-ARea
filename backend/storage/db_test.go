package storage

import (
	"area/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDatabase(t *testing.T) {
	config.LoadConfig()
	assert.NotPanics(t, func() {
		ConnectDatabase()
	}, "La fonction ConnectDatabase() ne doit pas provoquer de panic")
}
