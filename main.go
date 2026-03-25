package main


import (
	"boltadmin/internal/core"
	"log"
)


func main(){
	if err := core.Launch(); err != nil {
    log.Fatal(err)
	}
}
