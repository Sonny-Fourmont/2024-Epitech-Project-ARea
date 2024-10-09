package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDatabase(t *testing.T) {
	assert.NotPanics(t, func() {
		ConnectDatabase()
	}, "La fonction ConnectDatabase() ne doit pas provoquer de panic")
}
