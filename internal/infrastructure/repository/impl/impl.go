package impl

import "context"

type PacketRepo interface {
	Get(ctx context.Context) ([]int, error)
	Update(ctx context.Context, packets []int) ([]int, error)
}
