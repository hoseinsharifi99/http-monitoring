package main

import (
	"ec/db"
	"ec/db_manager"
	"ec/handler"
)

func main() {
	db := db.ConnectDatabase("prj.db")
	dm := db_manager.NewDb(db)
	h := handler.Newhandler(dm)
	h.Start()
}
