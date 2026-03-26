package api

import (
	"net/http"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterCustomRoutes(e *core.ServeEvent, _ any) error {
  e.Router.GET("/api/bolt-test", func (se *core.RequestEvent) error {
    return se.JSON(http.StatusOK, map[string]string{
      "status": "success",
			"message": "BoltAdmin is running from internal/api!",
		})
	})

	return e.Next()
}
