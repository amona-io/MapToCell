package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConn(t *testing.T) {
	asserts := assert.New(t)
	_, err := Conn()
	asserts.NoError(err, "database should connected")
	return
}
