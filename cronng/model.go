package cronng

import "code.google.com/p/go-uuid/uuid"
import "net/url"
import "time"

// JobService
type JobService struct {
	Jobstrage Jobstrage
}

func (service *JobService) Get(id JobId) (job *Job) {
	job = service.Jobstrage.Load(id)
	return
}

func (service *JobService) Delete(id JobId) (job *Job) {
	job := &Job{}
	return
}

// Job
type JobId uuid.UUID

type Job struct {
	id           JobId
	name         string
	description  string
	timeout      time.Time
	script       Script
	schedule     Schedule
	notification Notification
}

func (job *Job) Start() (execution *Execution, error Error) {
	execution := &Execution{}
	ProcQueue <- proc
	return
}

func (job *Job) GetExecutions() (executions []Execition, error Error) {
	executions := &[]Execution{}
	return
}

type Schedule string

type Notification struct {
	onstart   [](func(Execution) int)
	onsuccess [](func(Execution) int)
	onfailure [](func(Execution) int)
}

// Execution
const (
	RUNNING = iota
	SUCCEEDED
	FAILED
	ABORTED
)

type ExecutionId uuid.UUID

type Execution struct {
	id          ExecutionId
	output      url.URL
	status      Status
	user        User
	started     time.Time
	description string
	job         Job
	args        []string
	statistics  Statistics
	ended       time.Time
	abortedBy   User
}

func (execution *Execution) Abort() error {
}

type Statistics struct {
	EnvVar string
	VmPeak TimeSequence // peak value of virtual memory
	VmHWM  TimeSequence // peak value of VmRSS
	VmSwap TimeSequence
}

type TimeSequenceEntry struct {
	time  time.Time
	value float64
}

type TimeSequence struct {
	entries []TimeSequenceEntry
}

func (ts *TimeSequence) Max() (result float64) {
	result = math.MinFloat64
	for _, e := range ts.entries {
		result = Fmax(result, e.value)
	}
}

func (ts *TimeSequence) Min() (result float64) {
	result = math.MaxFloat64
	for _, e := range ts.entries {
		result = Fmin(result, e.value)
	}
}

func (ts *TimeSequence) Avg() (result float64) {
	sum := 0.0
	count := len(ts.entries)
	for _, e := range ts.entries {
		sum += e.value
	}
	result = sum / count
}

type User struct {
}

//
// coding snippets are below
//
func NewJob(script Script) (job *Job) {
	// Create Job instance
	job := &Job{}
	// Persist Job instance
	return
}

func GetJobs() (jobs []Job) {
	jobs = []Job{}
	return
}
