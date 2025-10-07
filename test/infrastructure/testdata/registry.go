package testdata

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type TestDataRegistry struct {
	managers map[string]TestDataManager
}

func NewTestDataRegistry() *TestDataRegistry {
	return &TestDataRegistry{managers: make(map[string]TestDataManager)}
}

func (r *TestDataRegistry) Register(manager TestDataManager) {
	r.managers[manager.GetName()] = manager
	log.Info(fmt.Sprintf("Registered test data manager %s", manager.GetName()))
}

func (r *TestDataRegistry) InitializeAll(ctx context.Context, dbPool *pgxpool.Pool) error {
	for name, manager := range r.managers {
		log.Info(fmt.Sprintf("🧪 Setting up: %s", name))
		if err := manager.Initialize(ctx, dbPool); err != nil {
			return fmt.Errorf("failed to initialize %s : %w", name, err)
		}
	}
	return nil
}

func (r *TestDataRegistry) InitializeSpecific(ctx context.Context, dbPool *pgxpool.Pool, names ...string) error {
	for _, name := range names {
		if manager, exists := r.managers[name]; exists {
			log.Info(fmt.Sprintf("🧪 Setting up: %s", name))
			if err := manager.Initialize(ctx, dbPool); err != nil {
				return fmt.Errorf("failed to initialize %s: %w", name, err)
			}
		} else {
			return fmt.Errorf("test data manager not found: %s", name)
		}
	}
	return nil
}

func (r *TestDataRegistry) CleanupAll(ctx context.Context, dbPool *pgxpool.Pool) error {
	for name, manager := range r.managers {
		log.Info(fmt.Sprintf("🧹 Cleaning up: %s", name))
		if err := manager.Cleanup(ctx, dbPool); err != nil {
			return fmt.Errorf("failed to cleanup %s: %w", name, err)
		}
	}
	return nil
}

func (r *TestDataRegistry) CleanupSpecific(ctx context.Context, dbPool *pgxpool.Pool, names ...string) error {
	for i := len(names) - 1; i >= 0; i-- {
		name := names[i]
		if manager, exists := r.managers[name]; exists {
			log.Info(fmt.Sprintf("🧹 Cleaning up: %s", name))
			if err := manager.Cleanup(ctx, dbPool); err != nil {
				return fmt.Errorf("failed to cleanup %s: %w", name, err)
			}
		} else {
			return fmt.Errorf("test data manager not found: %s", name)
		}
	}
	return nil
}
