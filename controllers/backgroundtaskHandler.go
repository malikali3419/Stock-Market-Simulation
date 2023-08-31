package controllers

import (
	"github.com/gin-gonic/gin"
	"stock_market_simulation/m/models"
	"sync"
)

var wg sync.WaitGroup

type TransactionTask struct {
	Context     *gin.Context
	Transaction models.TransactionData
}

func TransactionWorker(taskCh <-chan TransactionTask) {
	wg.Add(1)
	for {
		select {
		case task := <-taskCh:
			// Call your existing DoTransaction function here but modify it to accept TransactionTask
			DoTransaction(task)

		}
		wg.Wait()

	}

}
