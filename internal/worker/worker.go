package worker

func StartWorkers() {
	go StartDataStoreWorker()
	go StartLastFMWorker()
	go StartComputerWorker()
}
