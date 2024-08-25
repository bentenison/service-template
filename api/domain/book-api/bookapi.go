package bookapi

import (
	"log"
	"net/http"

	"github.com/bentenison/microservice/app/domain/bookapp"
	"github.com/bentenison/microservice/app/sdk/mid"
)

type api struct {
	bookApp *bookapp.App
}

func newAPI(bookApp *bookapp.App) *api {
	return &api{
		bookApp: bookApp,
	}
}

func (a *api) query(w http.ResponseWriter, r *http.Request) {
	// qp, err := parseQueryParams(r)
	// if err != nil {
	// 	apperrors.NewError(err
	// }
	log.Println("TraceID is ", mid.GetTraceId(r.Context()))
	log.Println("query params:", "Hello World!!")

	// web.Respond(200, r, w, map[string]interface{}{
	// 	"Hello": "I am running",
	// })
	// a.bookApp.Query(r.Context())
}
