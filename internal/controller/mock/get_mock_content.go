package mock

import (
	"context"
	"debugger-api/internal/model"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

const errorGetMockContent = "error while trying to get mock content"

// GetMockContent godoc
// @Summary      Get mock content
// @Description  Returns mock content as json
// @Tags         mock
// @Produce      json
// @Param        key   path      int  true  "Data key"
// @Success      204
// @Failure      400  {object}  httperr.ResponseDto
// @Router       /mock/{id} [get]
func (c *controller) GetMockContent(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/get_mock_content"

		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			c.errorHandler.HandleIncorrectRequestParamError(err, w, r)
			return
		}

		content, err := c.dataService.GetContent(ctx, id)
		if err != nil {
			c.errorHandler.HandleServiceError(
				model.FromError(err, model.BuildSubErrorWithOperation(op, errorGetMockContent)), w, r,
			)
			return
		}

		render.Status(r, 200)
		render.JSON(w, r, content)
		return
	}
}
