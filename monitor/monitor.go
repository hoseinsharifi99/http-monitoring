package monitor

import (
	"context"
	"ec/db_manager"
	"ec/model"
	"errors"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"net/http"
	"sync"
	"time"
)

type Monitor struct {
	dm         *db_manager.DbInstance
	urls       []*model.Url
	maxSemxize int
	sem        *semaphore.Weighted
}

func NewMonitor(dm *db_manager.DbInstance) *Monitor {
	mnt := new(Monitor)
	mnt.dm = dm
	mnt.maxSemxize = 10
	mnt.sem = semaphore.NewWeighted(int64(mnt.maxSemxize))
	mnt.getUrlFromDB()
	return mnt
}

func (mnt *Monitor) getUrlFromDB() error {
	urls, err := mnt.dm.GetAllUrls()
	if err != nil {
		return err
	}
	mnt.urls = urls
	return nil
}

func (mnt *Monitor) Work() {
	var wg sync.WaitGroup

	for urlIndex := range mnt.urls {
		url := mnt.urls[urlIndex]
		wg.Add(1)
		go func(urlIndex int) {
			if err := mnt.sem.Acquire(context.Background(), 1); err != nil {
				log.Fatal(err)
			}
			defer wg.Done()
			mnt.monitorURL(url)
			defer mnt.sem.Release(1)
		}(urlIndex)
	}
	wg.Wait()
}

func (mnt *Monitor) AddURL(urls []*model.Url) {
	mnt.urls = append(mnt.urls, urls...)
}

func (mnt *Monitor) monitorURL(url *model.Url) {
	req, err := url.SendRequest()
	if err != nil {
		fmt.Println(err, "could not make request")
		req = new(model.Request)
		req.UrlID = url.ID
		req.Result = http.StatusBadRequest
	}
	if err = mnt.dm.AddRequest(req); err != nil {
		fmt.Println(err, "could not save request to database")
	}

	if req.Result/100 == 2 {
		if err = mnt.dm.IncrementSuccess(url); err != nil {
			fmt.Println(err, "could not increment success times for url")
		}
	} else {
		if err = mnt.dm.IncrementFailed(url); err != nil {
			fmt.Println(err, "could not increment failed times for url")
		}
	}

}

type Scheduler struct {
	Mnt  *Monitor
	Stop chan struct{}
}

func NewScheduler(mnt *Monitor) (*Scheduler, error) {
	sch := &Scheduler{Stop: make(chan struct{})}
	if mnt != nil {
		sch.Mnt = mnt
		return sch, nil
	}
	return nil, errors.New("Monitor cannot be null")
}

func (sch *Scheduler) WorkInIntervals(d time.Duration) {
	ticker := time.NewTicker(d)
	go func() {
		for {
			select {
			case <-ticker.C:
				sch.Mnt.Work()
			case <-sch.Stop:
				ticker.Stop()
				return
			}
		}
	}()
}

func (sch *Scheduler) StopSchedule() {
	close(sch.Stop)
}
