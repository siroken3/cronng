package cronng

type JobDefinition struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

type JobId string

type Job struct {
	Id         JobId
	Name       string
	Definition string
	CreatedAt  time
}

type ProcId string

type Proc struct {
	Id          ProcId
	JobId       JobId
	TriggerType string
	Crontab     string
	StartAt     time
}

func NewJob(def JobDefinition) (job *Job) {
	// Create Job instance
	job := &Job{}
	// Persist Job instance
	return
}

func GetJobs() (jobs []Job) {
	jobs = []Job{}
	return
}
