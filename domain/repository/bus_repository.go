package repository

type BusRepository interface {
	FindAll(ctx context.Context) (*model.Bus, error)
}
