package storage

type Filter func([]string) ([]string, error)

type Storage interface {
	GET(key string) ([]byte, error)
	PUT(key string, value []byte) error
	DELETE(key string) error
	KEYS() ([]string, error)
	Search(key string, filter Filter, size int, page int, crumb string) ([]string, error)
	Start() error
	Close() error
	HealthCheck() error
}
