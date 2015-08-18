package cronng

type Job struct {
	Id         string
	Name       string
	Definition string
	CreatedAt  time
}

type Proc struct {
	Id          string
	JobId       string
	TriggerType string
	Crontab     string
	StartAt     time
}

func CreateJob() (job *Job) {
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
