package mock

import (
	"context"
	"debugger-api/internal/model"
	"debugger-api/internal/model/dto"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

const errorUpdateMock = "error while trying to update mock content"

// UpdateMock godoc
// @Summary      Update mock
// @Description  Updates json by given key
// @Tags         mock
// @Accept 		 json
// @Param        key   path      int  true  "Data key"
// @Success      200
// @Failure      400  {object}  httperr.ResponseDto
// @Router       /mock/{id} [put]
func (c *controller) UpdateMock(cxt context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/update_mock"

		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			c.errorHandler.HandleIncorrectRequestParamError(err, w, r)
			return
		}

		var request dto.UpdateMockDto
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			c.errorHandler.HandleUnprocessableEntityError(err, w, r)
			return
		}

		if err = c.validator.Struct(request); err != nil {
			c.errorHandler.HandleUnprocessableEntityError(err, w, r)
			return
		}

		err = c.dataService.Update(cxt, request, id)
		if err != nil {
			c.errorHandler.HandleServiceError(
				model.FromError(err, model.BuildSubErrorWithOperation(op, errorUpdateMock)), w, r,
			)
			return
		}

		render.Status(r, 200)
		return
	}
}
