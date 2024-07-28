package mock

import (
	"context"
	"debugger-api/internal/model"
	"fmt"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

const errorDeleteMock = "error while trying to delete mock"

// DeleteMock godoc
// @Summary      Deletes mock
// @Description  Delete json by given key
// @Tags         mock
// @Param        key   path      int  true  "Data key"
// @Success      204
// @Failure      400  {object}  httperr.ResponseDto
// @Failure      404  {object}  httperr.ResponseDto
// @Router       /mock/{id} [delete]
func (c *controller) DeleteMock(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/delete_mock"

		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			c.errorHandler.HandleIncorrectRequestParamError(fmt.Errorf(op), w, r)
			return
		}

		err = c.dataService.Delete(ctx, id)
		if err != nil {
			c.errorHandler.HandleServiceError(
				model.FromError(err, model.BuildSubErrorWithOperation(op, errorDeleteMock)), w, r,
			)
			return
		}

		render.Status(r, 204)
		return
	}
}
