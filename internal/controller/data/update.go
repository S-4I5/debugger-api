package data

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// UpdateData godoc
// @Summary      Update data
// @Description  Updates json by given key
// @Tags         data
// @Accept 		 json
// @Param        key   path      int  true  "Data key"
// @Success      200
// @Failure      400  {object}  error.ResponseDto
// @Router       /data/{id} [put]
func (c *Controller) UpdateData(cxt context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/update"

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

		err = c.dataService.Update(cxt, string(requestToString), key)
		if err != nil {
			c.errorHandler.HandleBusinessError(
				fmt.Errorf(op+":"+err.Error()),
				"api.data.update.error",
				w, r)
			return
		}

		render.Status(r, 200)
	}
}
