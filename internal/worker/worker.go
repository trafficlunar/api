package worker

func StartWorkers() {
	go StartDataStoreWorker()
	go StartLastFMWorker()
	go StartGitHubWorker()
	go StartComputerWorker()
}
