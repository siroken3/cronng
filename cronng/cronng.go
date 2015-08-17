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

type ProcLog struct {
	ProcId  string
	Result  string
	Code    int
	StartAt time
	EndAt   time
}

func CreateJob(definition Job) *Job {
	return &Job{}
}

func GetJobs() {
}

func GetJob() {
}

func UpdateJob() {
}

func DeleteJob() {
}

func StartProc() {
}

func GetProc() {
}

func StopProc() {
}

func GetProcsByJob() {
}

func GetProcLogByJob() {
}
