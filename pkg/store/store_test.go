package store

import (
	"github.com/go-playground/assert/v2"
	"testing"
)



func TestNewStore(t *testing.T) {
	store ,err:=NewStore("invalid url")
	assert.Equal(t, err.Error(), "cannot parse `invalid url`: failed to parse as DSN (invalid dsn)")
	assert.Equal(t, store, store )
}

