package db

import (
	"context"
	"sync"
)

var (
	// onceInit guarantees initialization of properties only once
	onceInit      = new(sync.Once)
	concreteImpls = make(map[string]interface{})
)

const (
	dbImplKey = "dbImpl"
)

// Load is an exported method that loads props depending on environment
func Load(ctx context.Context) error {
	var dbErr error
	onceInit.Do(func() {
		dbErr = loadImpls(ctx)
	})
	return dbErr
}

func loadImpls(ctx context.Context) error {
	if concreteImpls[dbImplKey] == nil {
		dbErr := InitDB(ctx)
		if dbErr != nil {
			return dbErr
		}
		concreteImpls[dbImplKey] = Postgres
	}
	return nil
}
