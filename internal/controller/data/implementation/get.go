package implementation

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
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
func (c *Controller) GetData(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "controller/get"

		if c.errorHandler == nil {
			fmt.Println("XD")
		}

		key := chi.URLParam(r, "key")
		if key == "" {
			c.errorHandler.HandleIncorrectRequestParamError(fmt.Errorf(op), w, r)
			return
		}

		data, err := c.dataService.Get(ctx, key)
		if err != nil {
			c.errorHandler.HandleBusinessError(
				fmt.Errorf(op+":"+err.Error()),
				"",
				w,
				r,
			)
			return
		}

		render.Status(r, 200)
		render.JSON(w, r, data)
		return
	}
}
