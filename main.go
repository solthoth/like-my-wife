package main

import (
	"time"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/solthoth/like-my-wife/worker"
	"github.com/solthoth/like-my-wife/configuration"
)

func main() {
	ConfigureLogger()
	context := getWorkerContext()
	StartUp(context)
}

func ConfigureLogger()  {
	jww.SetStdoutThreshold(jww.LevelTrace)
}

func StartUp(context *configuration.Configurations)  {
	jww.INFO.Println("Starting program")
	
	for ;; {
		worker.Work(context)
		time.Sleep(context.WorkerContext.WaitTime * time.Second)
	}
}

func getWorkerContext() *configuration.Configurations {
	return configuration.Init("like-my-wife")
}