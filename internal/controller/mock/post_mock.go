package mock

import (
	"context"
	"debugger-api/internal/model"
	"debugger-api/internal/model/dto"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

const errorPostMock = "error while trying to create mock content"

// PostMock godoc
// @Summary      Creates mock source
// @Description  Stores json by given id
// @Tags         mock
// @Accept       json
// @Param        key   path      int  true  "Mock id"
// @Success      200
// @Failure      400  {object}  httperr.ResponseDto
// @Router       /mock [post]
func (c *controller) PostMock(cxt context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/post_mock"

		var request dto.CreateMockDto
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			c.errorHandler.HandleUnprocessableEntityError(err, w, r)
			return
		}

		if err = c.validator.Struct(request); err != nil {
			c.errorHandler.HandleUnprocessableEntityError(err, w, r)
			return
		}

		response, err := c.dataService.Create(cxt, request)
		if err != nil {
			c.errorHandler.HandleServiceError(
				model.FromError(err, model.BuildSubErrorWithOperation(op, errorPostMock)), w, r,
			)
			return
		}

		render.JSON(w, r, response)
		render.Status(r, 200)
		return
	}
}
