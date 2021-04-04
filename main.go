package main

import (
	"log"
	"time"

	"github.com/solthoth/like-my-wife/worker"
	"github.com/solthoth/like-my-wife/configuration"
)

func main() {
	logger := log.Default()
	logger.Println("Starting program")
	context := getWorkerContext()
	for ;; {
		worker.Work(context)
		time.Sleep(context.WorkerContext.WaitTime * time.Second)
	}
}

func getWorkerContext() *configuration.Configurations {
	return configuration.Init("like-my-wife")
}