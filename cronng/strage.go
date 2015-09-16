package cronng

type JobStrage interface {
	Save(j Job) error
	Load(id JobId) (Job, error)
	Delete(id JobId) error
}

type ExecutionStrage interface {
	Save(j *Execution) error
	Load(id ExecutionId) (Execution, error)
	Delete(id ExecutionId) error
}
