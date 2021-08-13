package jobscheduler

import (
	"time"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/config"
	"github.com/meidomx/msb/module"

	"github.com/go-co-op/gocron"
)

type JobScheduler struct {
	sch *gocron.Scheduler

	jobs     []api.SchedulingJob
	tempJobs []api.SchedulingJob
}

type JobMsbHandler struct {
}

func (h JobMsbHandler) CallProcess(process string, param interface{}) (interface{}, error) {
	return module.GetProcess(process).Call(param)
}

var shared api.MsbHandler = JobMsbHandler{}

func NewJobScheduler(config config.MsbConfig) *JobScheduler {
	loc, err := time.LoadLocation(config.Shared.JobScheduling.Timezone)
	if err != nil {
		LOGGER.Error("parse timezone for job scheduler failed.", err)
		panic(err)
	}
	sch := gocron.NewScheduler(loc)

	js := new(JobScheduler)
	js.sch = sch

	return js
}

func (this *JobScheduler) BuildJobs(jobs []api.SchedulingJob) {
	this.tempJobs = jobs
}

func (this *JobScheduler) ReloadJobs(jobs []api.SchedulingJob) {
	//TODO
}

func (this *JobScheduler) Start() error {
	this.jobs = this.tempJobs
	this.tempJobs = nil

	for _, v := range this.jobs {
		_, err := this.sch.Cron(v.CronConfig()).Do(func() {
			// MsbHandler required
			_, err := v.Handler(shared)
			if err != nil {
				LOGGER.Error("run job handler failed.", err)
			}
		})
		if err != nil {
			LOGGER.Error("add cron job failed.", err)
			return err
		}
	}
	this.sch.WaitForSchedule()
	this.sch.StartAsync()

	return nil
}
