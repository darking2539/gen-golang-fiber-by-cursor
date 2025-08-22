package hello

import "context"

type Greeter interface {
	GetHello(ctx context.Context) (string, error)
}

type Usecase struct{}

func New() *Usecase { return &Usecase{} }

func (u *Usecase) GetHello(ctx context.Context) (string, error) {
	return "hello world", nil
}
