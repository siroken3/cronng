package cronng

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// fundamental types
type Arg string

type EnvVar map[string]string

type UUIDHex string

type Status int

const (
	BEFORE_RUNNING Status = iota
	RUNNING
	EXIT_SUCCEEDED
	EXIT_FAILED
	EXIT_ABORTED
)

// User
type User struct {
	ID UUIDHex `json:"id"`
}

// Job
type Job struct {
	ID          UUIDHex   `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Timeout     time.Time `json:"timeout"`
	Script      string    `json:"script"`
	Schedule    string    `json:"schdule"`
	Created     time.Time `json:"created"`
}

func (self *Job) GetExecutions() (*[]Execution, error) {
	executions := &[]Execution{}
	return executions, nil
}

// Execution
type MonitoringItem struct {
	Time  time.Time `json:time`
	Value float64   `json:value`
}

type Monitoring struct {
	ID     UUIDHex          `json:"id"`
	EnvVar EnvVar           `json:"env_var"`
	Vms    []MonitoringItem `json:"vm"`     // peak value of virtual memory
	VmRss  []MonitoringItem `json:"vm_rss"` // peak value of VmRSS
	VmSwap []MonitoringItem `json:"vm_swap"`
}

type Execution struct {
	ID          UUIDHex    `json:"id"`
	StdOutURL   string     `json:"stdout_url"` // websocket streaming output url
	StdErrURL   string     `json:"stderr_url"` // websocket streaming output url
	Status      Status     `json:"status"`
	User        User       `json:"user"`
	Job         Job        `json:"job"`
	Description string     `json:"description"`
	Args        []Arg      `json:"args"`
	Monitoring  Monitoring `json:"monitoring"`
	Started     time.Time  `json:"started_at"`
	Ended       time.Time  `json:"ended_at"`
	AbortedBy   User       `json:"aborted_by"`

	stdout chan string
	stderr chan string
	quit   chan bool
}

func NewExecution(user User, job Job, args []Arg, description string) (execution *Execution) {
	execution = &Execution{
		ID:          uuid.NewV1().String(),
		Status:      BEFORE_RUNNING,
		User:        user,
		Job:         job,
		Description: description,
		Args:        args,
		Monitoring:  nil,
		StartedAt:   nil,
		EndedAt:     nil,
		AbortedBy:   nil,
		stdout:      make(chan string),
		stderr:      make(chan string),
		quit:        make(chan bool),
	}
}
