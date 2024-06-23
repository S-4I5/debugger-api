package data

import (
	"context"
	"debugger-api/internal/model"
	"fmt"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// PostData godoc
// @Summary      Creates data source
// @Description  Stores json by given key
// @Tags         data
// @Accept       json
// @Param        key   path      int  true  "Data key"
// @Success      200
// @Failure      400  {object}  error.ResponseDto
// @Router       /data [post]
func (c *controller) PostData(cxt context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "controller/post"

		requestToString, err := io.ReadAll(r.Body)
		if err != nil {
			c.errorHandler.HandleIncorrectRequestBodyError(fmt.Errorf(op+":"+err.Error()), w, r)
			return
		}

		id, err := c.dataService.Create(cxt, string(requestToString))
		if err != nil {
			c.errorHandler.HandleBusinessError(
				fmt.Errorf(op+":"+err.Error()),
				"api.data.create.error",
				w, r)
			return
		}

		render.JSON(w, r, model.PostResponse{Id: id})
		render.Status(r, 200)
	}
}
