package api

import (
	"boltadmin/internal/hook"
	"net/http"
	"github.com/pocketbase/pocketbase/core"
)

type BoltApp interface {
  GetHooks() *hook.HookManager
}


func RegisterCustomRoutes(e *core.ServeEvent, app BoltApp) error {
  e.Router.GET("/api/bolt-test", func (se *core.RequestEvent) error {
    return se.JSON(http.StatusOK, map[string]string{
      "status": "success",
			"message": "BoltAdmin is running from internal/api!",
		})
	})

	return e.Next()
}
