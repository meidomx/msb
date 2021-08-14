package main

import (
	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/module"
)

type ModuleJob struct {
}

func (m ModuleJob) CronConfig() string {
	return "*/1 * * * *"
}

func (m ModuleJob) Name() string {
	return "management.job"
}

func (m ModuleJob) Handler(msbHandler api.MsbHandler) (api.JobResultIndicator, error) {
	LOGGER_MODULE.Info("job keep alive")

	return api.JobResultSuccess, nil
}

var mgrjob api.SchedulingJob = ModuleJob{}

func init() {
	module.RegisterSchedulingJob(mgrjob)
}
