package shipment

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/bitcodr/re-test/internal/domain/model"
)

// MockRepo will satisfy the CrawlerRepo interface for testing purpose
// we don't want to have an actual insert in db
// to use test cases in CI it is best to use mocks
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Store(ctx context.Context, order chan model.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}
