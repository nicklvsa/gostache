package shared

//constants for the log types
const (
	Error LogLevel = iota + 1
	Warn
	Info
)

//constants for store types
const (
	GSCache GSType = iota + 1
	GSStore
)

//LogLevel - severity of log with an integer
type LogLevel int

//GSType - different types of stores
type GSType int

//GSData - standard struct for different gostache store methods (cache, store)
type GSData struct {
	Name    string
	Version int
}

//GSGeneric - generic interface for all gsdata stores
type GSGeneric interface {
	GetName() string
	GetVersion() int
	Init() *GSData
	Store(key interface{}, value interface{}, form GSType) error
	Get(key string, form GSType) (interface{}, error)
	GetAll(keys map[string]GSType) ([]interface{}, error)
	Remove(key string, form GSType) error
}
