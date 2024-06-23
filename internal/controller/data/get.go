package data

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	//"github.com/andybalholm/brotli"
	_ "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

// GetData godoc
// @Summary      Get data
// @Description  Returns json by given key
// @Tags         data
// @Produce      json
// @Param        key   path      int  true  "Data key"
// @Success      204
// @Failure      400  {object}  error.ResponseDto
// @Failure      404  {object}  error.ResponseDto
// @Router       /data/{id} [get]
func (c *controller) GetData(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/get"

		key := chi.URLParam(r, "key")
		id, err := uuid.Parse(key)
		if err != nil {
			c.errorHandler.HandleIncorrectRequestParamError(fmt.Errorf(op), w, r)
			return
		}

		data, err := c.dataService.Get(ctx, id)
		if err != nil {
			c.errorHandler.HandleBusinessError(
				fmt.Errorf(op+":"+err.Error()),
				"api.data.get.error",
				w, r)
			return
		}

		render.Status(r, 200)
		render.JSON(w, r, data)
		return
	}
}
