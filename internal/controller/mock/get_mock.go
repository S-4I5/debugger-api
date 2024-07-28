package mock

import (
	"context"
	"debugger-api/internal/model"
	"github.com/google/uuid"

	//"github.com/andybalholm/brotli"
	_ "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

const errorGetMock = "error while trying to get mock"

// GetMock godoc
// @Summary      Get mock
// @Description  Returns json by given key
// @Tags         mock
// @Produce      json
// @Param        key   path      int  true  "Data key"
// @Success      204
// @Failure      400  {object}  httperr.ResponseDto
// @Failure      404  {object}  httperr.ResponseDto
// @Router       /mock/{id} [get]
func (c *controller) GetMock(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/get_mock"

		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			c.errorHandler.HandleIncorrectRequestParamError(err, w, r)
			return
		}

		data, err := c.dataService.Get(ctx, id)
		if err != nil {
			c.errorHandler.HandleServiceError(
				model.FromError(err, model.BuildSubErrorWithOperation(op, errorGetMock)), w, r,
			)
			return
		}

		render.Status(r, 200)
		render.JSON(w, r, data)
		return
	}
}
