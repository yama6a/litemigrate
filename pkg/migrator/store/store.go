package store

import (
	"github.com/ymakhloufi/light-migrate/internal/pkg/migration/store"
	"go.uber.org/zap"
)

func NewPostgresStore(logger *zap.Logger, migrationTableName, connectionString string) (pgStore *store.PostgresStore, err error) {
	return store.NewPostgresStore(logger, migrationTableName, connectionString)
}
