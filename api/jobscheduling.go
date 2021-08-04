package api

type JobResultIndicator int

const (
	JobResultSuccess JobResultIndicator = 1
	JobResultRetry   JobResultIndicator = 2
	JobResultFail    JobResultIndicator = 3
)

type SchedulingJob interface {
	Name() string
	Handler(msbHandler MsbHandler) (JobResultIndicator, error)
	CronConfig() string
}
