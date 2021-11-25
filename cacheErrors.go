package mycache

import (
	"fmt"
)

func ErrKeyNotFound(key interface{}) error {
	return fmt.Errorf("Key %v not found in cache\n", key)
}
