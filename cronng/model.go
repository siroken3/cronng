package cronng

import "code.google.com/p/go-uuid/uuid"
import "net/url"
import "time"

type Job struct {
	uuid               uuid.UUID
	name               string
	description        string
	group              string
	multipleExecutions bool
	timeout            time.Time
	script             Script
	schedule           Schedule
	notification       Notification
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

type Execution struct {
	uuid        uuid.UUID
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

type Statistics struct {
	EnvVar string
	VmPeak TimeSequence // peak value of virtual memory
	VmHWM  TimeSequence // peak value of VmRSS
	VmSwap TimeSequence
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
