package main


import (
	"boltadmin/internal/core"
	"log"
)


func main(){
	app := core.NewApp()
	if err := app.Launch(); err != nil {
    log.Fatal(err)
	}
}
