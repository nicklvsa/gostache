package lib

import "gostache/shared"

//GSStore - store type
type GSStore shared.GSData

//Init - initialize store
func (store *GSStore) Init() *GSStore {
	store.Name = "gostache"
	store.Version = 1.0
	return store
}

//GetName - returns the name of the store
func (store *GSStore) GetName() string {
	return store.Name
}

//GetVersion - returns the version of the store
func (store *GSStore) GetVersion() int {
	return store.Version
}

//Store - stores key, value with type of store
func (store *GSStore) Store(key interface{}, value interface{}, form GSType) error {
	switch form {
	case shared.GSCache:
		return CacheMe(key, value)
	case shared.GSStore:
		return StoreMe(key, value)
	default:
		return CacheMe(key, value)
	}
}

//Get - gets data by a stored key
func (store *GSStore) Get(key string, form GSType) (interface{}, error) {

}

//GetAll - returns map of data from a map of keys with their types
func (store *GSStore) GetAll(keys map[string]GSType) ([]interface{}, error) {

}

//Remove - deletes stored data by its key and type
func (store *GSStore) Remove(key string, form GSType) error {

}
