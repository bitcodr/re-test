package shipment

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockPacketRepo MockRepo will satisfy the PacketRepo interface for testing purpose
// we don't want to have an actual insert in db
// to use test cases in CI it is best to use mocks
type MockPacketRepo struct {
	mock.Mock
}

func (m *MockPacketRepo) Get(ctx context.Context) ([]int, error) {
	args := m.Called(ctx)
	return args.Get(0).([]int), args.Error(1)
}

func (m *MockPacketRepo) Update(ctx context.Context, packets []int) ([]int, error) {
	args := m.Called(ctx, packets)
	return args.Get(0).([]int), args.Error(1)
}
