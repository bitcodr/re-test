package shipment

import (
	"context"
	"sync"

	"github.com/bitcodr/re-test/internal/domain/model"
	"github.com/bitcodr/re-test/internal/infrastructure/repository/impl"
)

// IShipment ICrawler interface - implement packet entity methods
// in here we can implement our domain logic without any dependency to specific databases and frameworks
type IShipment interface {
	Calculate(ctx context.Context, request uint) (*model.Order, error)
	UpdatePacket(ctx context.Context, request []uint) ([]uint, error)
}

type shipment struct {
	repo impl.PacketRepo

	mu sync.Mutex
}

// InitService - to initialize packet service and
// pass the repository to it without knowing what kind of DB we are using
func InitService(_ context.Context, repository impl.PacketRepo) IShipment {
	return &shipment{
		repo: repository,
	}
}

// Calculate Show service - store packet logic
func (t *shipment) Calculate(ctx context.Context, request uint) (*model.Order, error) {

	return nil, nil
}

// UpdatePacket Calculate Show service - store packet logic
func (t *shipment) UpdatePacket(ctx context.Context, request []uint) ([]uint, error) {

	return nil, nil
}
