package main

import (
	"ec/db"
	"ec/handler"
)

func main() {
	dm := db.ConnectDatabase("prj.db")
	h := handler.Newhandler(dm)
	h.Start()
}
