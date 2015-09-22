package cronng

import "code.google.com/p/go-uuid/uuid"
import "net/url"
import "time"

type Notification struct {
	OnStart   [](func(Execution) int)
	OnSuccess [](func(Execution) int)
	OnFailure [](func(Execution) int)
}

type User struct {
}

// Job
type JobId uuid.UUID
type Schedule string
type Script string
type Job struct {
	Id          JobId
	Name        string
	Description string
	Timeout     time.Time
	Script      Script
	Schedule    Schedule
	Notfication Notification
	Created     time.Time
}

func NewJob(job Job) (job *Job) {
	// Persist Job instance
	dbmap.Insert(job)
	return
}

func (self *Job) PreInsert(s gorp.SqlExecutor) error {
	self.Id = JobId.NewUUID()
	self.Created = time.Now().UnixNano()
	return nil
}

func (self *Job) Start(user User, args []string) (execution *Execution, error Error) {
	execution := &Execution{}
	ProcQueue <- proc
	return
}

func (self *Job) GetExecutions() (executions []Execition, error Error) {
	executions := &[]Execution{}
	return
}

// Execution
const (
	RUNNING = iota
	SUCCEEDED
	FAILED
	ABORTED
)

type Statistics struct {
	EnvVar string
	VmPeak TimeSequence // peak value of virtual memory
	VmHWM  TimeSequence // peak value of VmRSS
	VmSwap TimeSequence
}

type ExecutionId uuid.UUID
type Status int32
type Execution struct {
	Id          ExecutionId
	Output      url.URL
	Status      Status
	User        User
	Started     time.Time
	Description string
	Job         Job
	Args        []string
	Statistics  Statistics
	Ended       time.Time
	AbortedBy   User
}

func (execution *Execution) Abort(user User) error {
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
