package cronng

func GetJob(id JobId) (job *Job) {
	job = &Job{}
	return
}

func UpdateJob(def JobDefinition) (job *Job) {
	job = &Job{}
	return
}

func DeleteJob(def JobId) (job *Job) {
	job := &Job{}
	return
}

func StartProc(job Job) (proc *Proc) {
	proc := &Proc{}
	ProcQueue <- proc
	return
}

func GetProc(id ProcId) (proc *Proc) {
	proc := &Proc{}
	return
}

func StopProc() (proc *Proc) {
	proc := &Proc{}
	return
}

func GetProcsByJob() (procs []Proc) {
	procs := &[]Proc{}
	return
}

func GetProcLogByJob() (procs []Proc) {
	procs := &[]Proc{}
	return
}
