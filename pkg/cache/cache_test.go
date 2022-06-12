package cache

import (
	"L0/pkg/model"
	"github.com/go-playground/assert/v2"
	"sync"
	"testing"
)

var cache =Cache{map[string]model.Order{}, sync.RWMutex{}}

func TestPut(t *testing.T) {
	TestTable:=[]model.Order{
		model.Order{OrderUID: "1"},
		model.Order{OrderUID: "2"},
		model.Order{OrderUID: "3"},
	}

	for _, testCase := range TestTable{
		cache.Put(testCase)
		_, ok:=cache.data[testCase.OrderUID]
		assert.Equal(t, ok, true)
	}
}

func TestGet(t *testing.T) {
	TestTable:=[]string{"1","2","3"}
	for _, testCase := range TestTable{
		_,err:=cache.Get(testCase)
		assert.Equal(t, err, nil)
	}
}

func TestIsExist(t *testing.T) {
	TestTable:=[]string{"1","2","3"}
	for _, testCase := range TestTable{
		ok:=cache.IsExist(testCase)
		assert.Equal(t, ok, true)
	}
}
