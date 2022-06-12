package api

import (
	"L0/pkg/cache"
	"github.com/go-playground/assert/v2"
	"testing"
)
var api Api
func TestInitRouter(t *testing.T)  {
	cache := cache.NewCache()
	api.InitRouter(cache)
	assert.Equal(t,api.cache, cache)

}

