package userRepository

type repository struct{}

func New() *repository {
	return &repository{}
}
