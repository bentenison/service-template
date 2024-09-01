package bookapp

import (
	"context"
	"fmt"

	"github.com/bentenison/microservice/business/domain/bookbus"
	tp "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	bookbus *bookbus.Business
	tracer  trace.Tracer
}

func NewApp(bookbus *bookbus.Business, tp *tp.TracerProvider) *App {
	return &App{
		bookbus: bookbus,
		tracer:  tp.Tracer("BOOKS"),
	}
}
func (a *App) Query(ctx context.Context) {
	ctx, ts := a.tracer.Start(ctx, "BOOK_APP")
	var i int
	for i < 100 {
		fmt.Println("simulating work", i)
		i++
	}

	defer ts.End()
	// return a.bookbus.Query(ctx)
}
