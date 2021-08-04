package module

import "github.com/meidomx/msb/api"

var hs []api.HttpApiSimpleHandler
var jobs []api.SchedulingJob

func RegisterHttpApiHandler(h api.HttpApiSimpleHandler) {
	hs = append(hs, h)
}

func GetHttpApiHandlers() []api.HttpApiSimpleHandler {
	return hs
}

func RegisterSchedulingJob(job api.SchedulingJob) {
	jobs = append(jobs, job)
}

func GetSchedulingJob() []api.SchedulingJob {
	return jobs
}
