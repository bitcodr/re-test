package packet

import (
	"context"
	"errors"
	"slices"

	"github.com/bitcodr/re-test/internal/infrastructure/config"
	"github.com/bitcodr/re-test/internal/infrastructure/helper"
	"github.com/bitcodr/re-test/internal/infrastructure/repository/impl"
)

type packet struct {
	collection []int
}

// InitRepo instantiate packet entity memory repository
// with the interface we have in the impl directory we can implement
// another source of data and pass it to service
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

// Get Store - get packs from memory as a persistent layer
func (p *packet) Get(_ context.Context) ([]int, error) {
	return p.collection, nil
}

// Update packs
func (p *packet) Update(_ context.Context, packets []int) ([]int, error) {
	if !slices.IsSorted(packets) {
		slices.Sort(packets)
	}

	p.collection = packets

	return p.collection, nil
}
