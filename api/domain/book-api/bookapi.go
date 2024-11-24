package bookapi

import (
	"log"
	"net/http"

	"github.com/bentenison/microservice/app/domain/bookapp"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/gin-gonic/gin"
)

type api struct {
	bookApp *bookapp.App
}

func newAPI(bookApp *bookapp.App) *api {
	return &api{
		bookApp: bookApp,
	}
}

func (a *api) query(c *gin.Context) {
	// qp, err := parseQueryParams(r)
	// if err != nil {
	// 	apperrors.NewError(err
	// }
	id := mid.GetTraceId(c.Request.Context())
	log.Println("TraceID is ", id)
	log.Println("query params:", "Hello World!!")
	// web.Respond(200, r, w, map[string]interface{}{
	// 	"Hello": "I am running",
	// })
	a.bookApp.Query(c)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
