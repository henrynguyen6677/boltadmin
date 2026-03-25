package core

import (
	"boltadmin/internal/api"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func Launch() error {
   app := pocketbase.New()
	 app.OnServe().BindFunc(func (e *core.ServeEvent) error {
     return api.RegisterCustomRoutes(e)
	 })
	 return app.Start()
}
