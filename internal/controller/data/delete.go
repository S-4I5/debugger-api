package data

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

// DeleteData godoc
// @Summary      Deletes data
// @Description  Delete json by given key
// @Tags         data
// @Param        key   path      int  true  "Data key"
// @Success      204
// @Failure      400  {object}  error.ResponseDto
// @Failure      404  {object}  error.ResponseDto
// @Router       /data/{id} [delete]
func (c *controller) DeleteData(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/delete"

		key := chi.URLParam(r, "key")
		id, err := uuid.Parse(key)
		if err != nil {
			c.errorHandler.HandleIncorrectRequestParamError(fmt.Errorf(op), w, r)
			return
		}

		err = c.dataService.Delete(ctx, id)
		if err != nil {
			c.errorHandler.HandleBusinessError(
				fmt.Errorf(op+":"+err.Error()),
				"api.data.delete.error",
				w, r)
			return
		}

		render.Status(r, 204)
	}
}
