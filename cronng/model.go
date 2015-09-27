package cronng

import "time"

// fundamental types
type UUIDModel struct {
	ID string `gorm:"primary_key" sql:"type:varchar(40)" json:"id"`
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

func (self *Job) GetExecutions() (*[]Execution, error) {
	executions := &[]Execution{}
	return executions, nil
}

// Execution
type Vm struct {
	MonitoringID string    `sql:"type:varchar(40)"`
	Time         time.Time `json: time`
	Value        float64   `json:value`
}
type VmRss struct {
	MonitoringID string    `sql:"type:varchar(40)"`
	Time         time.Time `json: time`
	Value        float64   `json:value`
}
type VmSwap struct {
	MonitoringID string    `sql:"type:varchar(40)"`
	Time         time.Time `json: time`
	Value        float64   `json:value`
}

type Monitoring struct {
	UUIDModel
	ExecutionID string   `sql:"type:varchar(40)"`
	EnvVar      string   `json:"env_var"`
	Vms         []Vm     `json:"vm"`     // peak value of virtual memory
	VmRss       []VmRss  `json:"vm_rss"` // peak value of VmRSS
	VmSwap      []VmSwap `json:"vm_swap"`
}

type Execution struct {
	UUIDModel
	StdOutURL   string     `json:"stdout_url"` // websocket streaming output url
	StdErrURL   string     `json:"stderr_url"` // websocket streaming output url
	Status      Status     `json:"status"`
	UserID      string     `sql:"type:varchar(40)" json:"user_id"`
	JobID       string     `sql:"type:varchar(40)" json:"job_id"`
	Description string     `sql:"type:text" json:"description"`
	Args        string     `sql:"type:text" json:"args"`
	Monitoring  Monitoring `json:"monitoring"`
	Started     time.Time  `json:"started_at"`
	Ended       time.Time  `json:"ended_at"`
	AbortedBy   string     `sql:"type:varchar(40)" json:"aborted_by"`
}
