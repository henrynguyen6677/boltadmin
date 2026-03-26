package core

import (
	"boltadmin/internal/api"
	"boltadmin/internal/hook"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type App struct {
	PB *pocketbase.PocketBase
	Hooks *hook.HookManager
}

func NewApp() *App {
  return &App {
		PB: pocketbase.New(),
		Hooks: hook.NewManager(),
	}
}

func (a *App) GetHooks() *hook.HookManager {
  return a.Hooks
}

func (a *App) Launch() error {
	 a.PB.OnRecordCreate().BindFunc(func(e *core.RecordEvent) error {
		 collectionName := e.Record.Collection().Name
     data := e.Record.PublicExport()
		 if err := a.Hooks.TriggerBeforeSave(collectionName, data); err != nil {
			 return err
		 }
		 return e.Next()
   })

	 a.PB.OnServe().BindFunc(func (e *core.ServeEvent) error {
     return api.RegisterCustomRoutes(e, a)
	 })
	 return a.PB.Start()
}
