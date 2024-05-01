package implementation

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
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
// @Router       /data/{id} [post]
func (c *Controller) PostData(cxt context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "controller/post"

		key := chi.URLParam(r, "key")
		if key == "" {
			c.errorHandler.HandleIncorrectRequestParamError(fmt.Errorf(op), w, r)
			return
		}

		requestToString, err := io.ReadAll(r.Body)
		if err != nil {
			c.errorHandler.HandleIncorrectRequestBodyError(fmt.Errorf(op+":"+err.Error()), w, r)
			return
		}

		err = c.dataService.Create(cxt, string(requestToString), key)
		if err != nil {
			c.errorHandler.HandleBusinessError(
				fmt.Errorf(op+":"+err.Error()),
				"api.data.create.error",
				w, r)
			return
		}

		render.Status(r, 200)
	}
}
