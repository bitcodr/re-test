package shipment

import (
	"github.com/stretchr/testify/mock"
)

// MockService will satisfy the ICrawler interface for testing purpose
// we don't want to have an actual insert in db
// to use test cases in CI it is best to use mocks
type MockService struct {
	mock.Mock
}

//func (m *MockService) Show(ctx context.Context, order model.Order) (*model.Order, error) {
//	args := m.Called(ctx, order)
//	return args.Get(0), args.Error(1)
//}
