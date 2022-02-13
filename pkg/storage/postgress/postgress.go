package postgress

import (
	"github.com/dsmello/popcornStorage/pkg/storage/storage"
)

type PostgressStorage struct{}

func (s PostgressStorage) GET(key string) ([]byte, error)
func (s PostgressStorage) PUT(key string, value []byte) error
func (s PostgressStorage) DELETE(key string) error
func (s PostgressStorage) KEYS() ([]string, error)
func (s PostgressStorage) Search(key string, filter storage.Filter, size int, page int, crumb string) ([]string, error)
func (s PostgressStorage) Start() error
func (s PostgressStorage) Close() error
func (s PostgressStorage) HealthCheck() error
