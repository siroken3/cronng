package cronng

type JobPool interface {
	Save(j Job) error
	Load(id JobId) (Job, error)
	Delete(id JobId) error
}

type ProcPool interface {
	Save(j Proc) error
	Load(id ProcId) (Job, error)
	Delete(id ProcId) error
}
