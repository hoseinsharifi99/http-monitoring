package main

import (
	"ec/db"
	"ec/db_manager"
	"ec/handler"
	"ec/monitor"
	"time"
)

func main() {
	db := db.SetUpDb("prj.db")
	dm := db_manager.NewDb(db)

	mnt := monitor.NewMonitor(dm)
	sch, _ := monitor.NewScheduler(mnt)
	sch.WorkInIntervals(time.Minute)
	h := handler.Newhandler(dm, sch)
	h.Start()
}
