package cronng

import "time"

// fundamental types
type UUIDModelID string

type UUIDModel struct {
	ID UUIDModelID `gorm:"primary_key" json:"id"`
}

type Status int

const (
	RUNNING Status = iota
	SUCCEEDED
	FAILED
	ABORTED
)

// User
type User struct {
	UUIDModel
}

// Job
type Job struct {
	UUIDModel
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Timeout     time.Time `json:"timeout"`
	Script      string    `json:"script"`
	Schedule    string    `json:"schdule"`
	Created     time.Time `json:"created"`
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
type Vm struct {
	MonitoringID UUIDModelID
	Time         time.Time `json: time`
	Value        float64   `json:value`
}
type VmRss struct {
	MonitoringID UUIDModelID
	Time         time.Time `json: time`
	Value        float64   `json:value`
}
type VmSwap struct {
	MonitoringID UUIDModelID
	Time         time.Time `json: time`
	Value        float64   `json:value`
}

type Monitoring struct {
	UUIDModel
	ExecutionID UUIDModelID
	EnvVar      string   `json:"env_var"`
	Vms         []Vm     `json:"vm"`     // peak value of virtual memory
	VmRss       []VmRss  `json:"vm_rss"` // peak value of VmRSS
	VmSwap      []VmSwap `json:"vm_swap"`
}

type Execution struct {
	UUIDModel
	StdOutURL   string      `json:"stdout_url"` // websocket streaming output url
	StdErrURL   string      `json:"stderr_url"` // websocket streaming output url
	Status      Status      `json:"status"`
	UserID      UUIDModelID `json:"user_id"`
	JobID       UUIDModelID `json:"job_id"`
	Description string      `json:"description"`
	Args        string      `json:"args"`
	Monitoring  Monitoring  `json:"monitoring"`
	Started     time.Time   `json:"started_at"`
	Ended       time.Time   `json:"ended_at"`
	AbortedBy   UUIDModelID `json:"aborted_by"`
}

func (execution *Execution) Signal(signal int32, user User) error {
	return nil
}
