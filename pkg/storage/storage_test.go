package storage

import "testing"

func StorageHealthCheckTester(t *testing.T, storage Storage) {
	t.Helper()
	t.Run("Verify HealthCheck method", func(t *testing.T) {
		if err := storage.HealthCheck(); err != nil {
			t.Errorf("HealthCheck() error = %v, wantErr %v", err, nil)
		}
	})
}
