package shipment

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/bitcodr/re-test/internal/domain/model"
	"github.com/bitcodr/re-test/internal/infrastructure/repository/impl"
)

// IShipment ICrawler interface - implement packet entity methods
// in here we can implement our domain logic without any dependency to specific databases and frameworks
type IShipment interface {
	Calculate(ctx context.Context, request int) (*model.Order, error)
	UpdatePacket(ctx context.Context, request []int) ([]int, error)
}

type shipment struct {
	repo impl.PacketRepo

	mu sync.RWMutex
}

// InitService - to initialize packet service and
// pass the repository to it without knowing what kind of DB we are using
func InitService(_ context.Context, repository impl.PacketRepo) IShipment {
	return &shipment{
		repo: repository,
	}
}

// Calculate Show service - store packet logic
func (t *shipment) Calculate(ctx context.Context, request int) (*model.Order, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	packets, err := t.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	if packets == nil {
		return nil, errors.New("there are no packets")
	}

	order := &model.Order{
		Packet: make(map[int]int),
	}

	order.Item = request

	order.Packet, _ = findPacks(request, packets)

	return order, nil
}

func findPacks(remainingItems int, packets []int) (map[int]int, int) {
	// Base case: if there are no remaining items, return an empty pack count.
	if remainingItems == 0 {
		return make(map[int]int), 0
	}

	bestPacks := make(map[int]int)
	bestCount := remainingItems // Initialize with a large count.

	for _, packSize := range packets {
		if remainingItems >= packSize {
			// Try using this pack size.
			newCounts, newCount := findPacks(remainingItems-packSize, packets)
			newCounts[packSize]++
			newCount++ // Include the current pack.

			if newCount < bestCount {
				bestPacks = newCounts
				bestCount = newCount
			}
		}
	}
	fmt.Println(bestPacks, bestCount)
	return bestPacks, bestCount
}

// UpdatePacket Calculate Show service - store packet logic
func (t *shipment) UpdatePacket(ctx context.Context, request []int) ([]int, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.repo.Update(ctx, request)
}
