package shipment

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/bitcodr/re-test/internal/domain/model"
)

// MockIShipment will satisfy the IShipment interface for testing purpose
// we don't want to have an actual insert in db
// to use test cases in CI it is best to use mocks
type MockIShipment struct {
	mock.Mock
}

func (m *MockIShipment) Calculate(ctx context.Context, request int) (*model.Order, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *MockIShipment) UpdatePacket(ctx context.Context, request []int) ([]int, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]int), args.Error(1)
}
