package cronng

import (
	"math"
	"time"
)

// User
type UserId string
type User struct {
	Id UserId
}

// Job
type Job struct {
	Id          string    `gorm:"id, primarykey" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Timeout     time.Time `db:"timeout" json:"timeout"`
	Script      string    `db:"script" json:"script"`
	Schedule    string    `db:"schedule" json:"schdule"`
	Created     time.Time `db:"created" json:"created"`
}

func (self *Job) Start(user User, args []string) (*Execution, error) {
	execution := &Execution{}
	return execution, nil
}

func (self *Job) GetExecutions() (*[]Execution, error) {
	executions := &[]Execution{}
	return executions, nil
}

// Execution
type Statistics struct {
	EnvVar string
	VmPeak TimeSequence // peak value of virtual memory
	VmHWM  TimeSequence // peak value of VmRSS
	VmSwap TimeSequence
}

type Status int32

const (
	RUNNING Status = iota
	SUCCEEDED
	FAILED
	ABORTED
)

type Execution struct {
	Id          string
	Output      string
	Status      Status
	User        string
	Started     time.Time
	Description string
	JobId       string
	Args        string
	Statistics  Statistics
	Ended       time.Time
	AbortedBy   string
}

func (execution *Execution) Abort(user User) error {
	return nil
}

type TimeSequenceEntry struct {
	time  time.Time
	value float64
}

type TimeSequence struct {
	entries []TimeSequenceEntry
}

func (ts *TimeSequence) Max() (result float64) {
	result = math.SmallestNonzeroFloat64
	for _, e := range ts.entries {
		result = math.Max(result, e.value)
	}
	return
}

func (ts *TimeSequence) Min() (result float64) {
	result = math.MaxFloat64
	for _, e := range ts.entries {
		result = math.Min(result, e.value)
	}
	return
}

func (ts *TimeSequence) Avg() (result float64) {
	sum := 0.0
	count := len(ts.entries)
	for _, e := range ts.entries {
		sum += e.value
	}
	result = sum / float64(count)
	return
}
