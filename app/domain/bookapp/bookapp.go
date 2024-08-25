package bookapp

import (
	"context"

	"github.com/bentenison/microservice/business/domain/bookbus"
)

type App struct {
	bookbus *bookbus.Business
}

func NewApp(bookbus *bookbus.Business) *App {
	return &App{
		bookbus: bookbus,
	}
}
func (a *App) Query(ctx context.Context) ([]bookbus.Book, error) {
	return a.bookbus.Query(ctx)
}
