package shipment_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bitcodr/re-test/internal/domain/model"
	"github.com/bitcodr/re-test/internal/domain/service/shipment"
	mockrepo "github.com/bitcodr/re-test/internal/infrastructure/mock/shipment"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		name        string
		items       int
		expectedRes *model.Order
	}{
		{
			name:  "Test case 0",
			items: 501,
			expectedRes: &model.Order{
				Item: 501,
				Packet: map[int]int{
					250: 1,
					500: 1,
				},
			},
		},
		{
			name:  "Test case 1",
			items: 1,
			expectedRes: &model.Order{
				Item: 1,
				Packet: map[int]int{
					250: 1,
				},
			},
		},
		{
			name:  "Test case 2",
			items: 250,
			expectedRes: &model.Order{
				Item: 250,
				Packet: map[int]int{
					250: 1,
				},
			},
		},
		{
			name:  "Test case 3",
			items: 251,
			expectedRes: &model.Order{
				Item: 251,
				Packet: map[int]int{
					500: 1,
				},
			},
		},
		{
			name:  "Test case 4",
			items: 12001,
			expectedRes: &model.Order{
				Item: 12001,
				Packet: map[int]int{
					250:  1,
					2000: 1,
					5000: 2,
				},
			},
		},
		{
			name:  "Test case 5",
			items: 750,
			expectedRes: &model.Order{
				Item: 750,
				Packet: map[int]int{
					500: 1,
					250: 1,
				},
			},
		},
		{
			name:  "Test case 6",
			items: 751,
			expectedRes: &model.Order{
				Item: 751,
				Packet: map[int]int{
					1000: 1,
				},
			},
		},
		{
			name:  "Test case 7",
			items: 111,
			expectedRes: &model.Order{
				Item: 111,
				Packet: map[int]int{
					250: 1,
				},
			},
		},
		{
			name:  "Test case 8",
			items: 3455,
			expectedRes: &model.Order{
				Item: 3455,
				Packet: map[int]int{
					250:  2,
					1000: 1,
					2000: 1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()

			mockRepo := new(mockrepo.MockPacketRepo)
			mockRepo.On("Update", ctx, []int{250, 500, 1000, 2000, 5000}).Return(
				[]int{250, 500, 1000, 2000, 5000}, nil,
			)
			mockRepo.On("Get", ctx).Return([]int{250, 500, 1000, 2000, 5000}, nil)

			sh := shipment.InitService(ctx, mockRepo)

			order, err := sh.Calculate(ctx, tc.items)
			if err != nil {
				assert.Error(t, err)

				return
			}

			assert.Equal(t, tc.expectedRes, order)
		})
	}
}

func TestUpdatePacket(t *testing.T) {
	testCases := []struct {
		name        string
		items       []int
		expectedRes []int
	}{
		{
			name:        "Test case 0",
			items:       []int{250, 500, 1000, 2000, 5000},
			expectedRes: []int{250, 500, 1000, 2000, 5000},
		},
		{
			name:        "Test case 1",
			items:       []int{1, 3, 4, 5, 6},
			expectedRes: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()

			mockRepo := new(mockrepo.MockPacketRepo)
			mockRepo.On("Update", ctx, tc.items).Return(tc.expectedRes, nil)

			sh := shipment.InitService(ctx, mockRepo)

			res, err := sh.UpdatePacket(ctx, tc.items)
			if err != nil {
				assert.Error(t, err)

				return
			}

			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
