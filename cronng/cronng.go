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

func NewJob() (job *JobDefinition) {
	// Create Job instance
	job := &Job{}
	// Persist Job instance
	return
}

func GetJobs() (jobs []Job) {
	jobs = []Job{}
	return
}

func GetJob() (job *Job) {
	job = &Job{}
}

func UpdateJob() (job *Job) {
	job = &Job{}
}

func DeleteJob() (job *Job) {
	job := &Job{}
}

func StartProc() (proc *Proc) {
	proc := &Proc{}
}

func GetProc() *Proc {
	proc := &Proc{}
}

func StopProc() *Proc {
	proc := &Proc{}
}

func GetProcsByJob() []Proc {
	proc := &Proc{}
}

func GetProcLogByJob() []Proc {
	proc := &Proc{}
}
