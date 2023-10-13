package packet

import (
	"context"
	"errors"
	"sync"

	"github.com/bitcodr/re-test/internal/infrastructure/config"
	"github.com/bitcodr/re-test/internal/infrastructure/helper"
	"github.com/bitcodr/re-test/internal/infrastructure/repository/impl"
)

type packet struct {
	collection []uint

	mu sync.RWMutex
}

// InitRepo instantiate packet entity memory repository
// with the interface we have in the impl directory we can implement another source of data and pass it to service
// without changing anything in out domain service
func InitRepo(ctx context.Context, cfg *config.Connection) (impl.PacketRepo, error) {
	if cfg == nil {
		return nil, errors.New("config is empty")
	}

	collection, err := helper.NewMemory(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &packet{
		collection: collection,
	}, nil
}

// Get Store - store a packet in memory as a persistent layer
func (p *packet) Get(ctx context.Context) ([]uint, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.collection, nil
}

func (p *packet) Update(_ context.Context, packets []uint) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.collection = packets

	return nil
}
