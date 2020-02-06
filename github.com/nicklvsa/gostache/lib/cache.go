package lib

import (
	"errors"
	"fmt"
)

var cacheData = make(map[string]interface{})

//CacheMe - caches a dataset locally
func CacheMe(key interface{}, val interface{}) error {
	if key != nil && val != nil {
		cleanKey := fmt.Sprintf("%s", key)
		if _, ok := cacheData[cleanKey]; !ok {
			cacheData[cleanKey] = val
		}
		return errors.New("key already exists")
	}
}
