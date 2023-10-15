package shipment

import (
	"context"
	"errors"
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

	var sum int
	findPacks(request, packets, order, &sum)
	checkPacksEquivalent(packets, order, &sum)

	return order, nil
}

func findPacks(remainingItems int, packets []int, order *model.Order, sum *int) {

	// Base case: if there are no remaining items, return an empty pack count.
	if remainingItems <= 0 {
		return
	}

	// Find the largest pack size that can be used.
	largestPackSize := 0
	for _, packSize := range packets {
		if packSize <= remainingItems {
			largestPackSize = packSize
		}
	}

	if largestPackSize == 0 {
		// If still no suitable pack size is found, choose the smallest available pack size.
		largestPackSize = packets[0]
	}

	order.Packet[largestPackSize]++
	*sum += largestPackSize
	findPacks(remainingItems-largestPackSize, packets, order, sum)
}

// checkPacksEquivalent
// checks whether the sum of the packs can be changed to another one to reduce the number of packs
func checkPacksEquivalent(packets []int, order *model.Order, sum *int) {
	for _, p := range packets {
		if *sum == p {
			order.Packet = map[int]int{p: 1}
			break
		}
	}
}

// UpdatePacket Calculate Show service - store packet logic
func (t *shipment) UpdatePacket(ctx context.Context, request []int) ([]int, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.repo.Update(ctx, request)
}
