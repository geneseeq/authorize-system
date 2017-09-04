package main

import (
	"fmt"
	"os"
	"time"

	"github.com/geneseeq/authorize-system/task/baseJob/control"
	"github.com/geneseeq/authorize-system/task/baseJob/model"
	"github.com/ivpusic/grpool"
)

func Producer(queue chan<- string, baseInfos []model.BaseInfoModel) {
	for _, base := range baseInfos {
		queue <- base.OrderID
	}
}
func Consumer(queue <-chan string, results chan<- int) {
	id := <-queue
	// orderInfo := control.FindSampleOrderInfo(id)
	// for _, order := range orderInfo {
	// 	fmt.Println("===============")
	// 	fmt.Println(order.SaleID)
	// 	fmt.Println("===============")
	// }
	fmt.Println("receive:", id, "this is step:")
	// results <- 0
}
func worker(jobs <-chan string, results chan<- int) {
	i := 1
	for id := range jobs {
		i = i + 1
		fmt.Println("worker", id, "started  job")
		time.Sleep(time.Second)
		orderInfo := control.FindSampleOrderInfo(id)
		for _, order := range orderInfo {
			fmt.Println("===============")
			fmt.Println(order.SaleID)
			fmt.Println("===============")
		}
		fmt.Println("worker", id, "this is step:", i)
		results <- 0
	}
}
func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

const (
	defaultPort = "8080"
)

func main() {
	// jobs := make(chan string, 100)

	err, labCode := control.FindAllLabCode()
	if err == nil {
		// for _, v := range labCode {
		// 	println(v)
		// }
	}
	var baseInfos []model.BaseInfoModel
	baseInfos, err = control.FindSampleBaseInfo(labCode)
	// number of workers, and size of job queue
	pool := grpool.NewPool(100000, 50000)
	defer pool.Release()

	// how many jobs we should wait
	pool.WaitCount(1000000)

	// submit one or more jobs to pool
	for i := 0; i < len(baseInfos); i++ {
		count := i

		pool.JobQueue <- func() {
			// say that job is done, so we can know how many jobs are finished
			defer pool.JobDone()

			fmt.Printf("hello %d\n", count)
		}
	}

	// wait until we call JobDone for all jobs
	pool.WaitAll()

}
