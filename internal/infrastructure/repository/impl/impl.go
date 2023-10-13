package impl

import "context"

type PacketRepo interface {
	Get(ctx context.Context) ([]uint, error)
	Update(ctx context.Context, packets []uint) error
}
