package main


import (
	"log"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)


func main(){
  app := pocketbase.New()

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
    e.Router.GET("/api/bolt-test", func(se *core.RequestEvent) error {
      return se.JSON(http.StatusOK, map[string]string {
        "status": "success",
				"message": "BoltAdmin Core is running!",
			})
		})
		return e.Next()
	})
	if err := app.Start(); err != nil {
    log.Fatal(err)
	}
}
