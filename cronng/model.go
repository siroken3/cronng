package cronng

type Job struct {
	uuid               UUID
	name               string
	description        string
	group              string
	multipleExecutions bool
	timeout            Timeout
	sequence           []Command
	dispatch           Dispatch
	schedule           Schedule
	notification       Notification
}

type Timtout string

type Command struct {
	exec   string
	jobref Job
	script Script
}

type Dispatch struct {
	threadcount int
	keepgoing   bool
}

type Crontab string

type Schedule struct {
	crontab Crontab
}

type Notification struct {
	onsuccess string
	onfailure string
	onstart   string
}

// Execution

type Execution struct {
	uuid            UUID
	output          URL
	status          Status
	user            User
	started         Time
	description     string
	job             Job
	args            string
	statistics      Statistics
	ended           Time
	abortedBy       User
	successfulNodes []Node
	failedNodes     []Node
}

const (
	RUNNING = iota
	SUCCEEDED
	FAILED
	ABORTED
)

type URL struct {
}

type User struct {
}

type Statistics struct {
}

type Time stuct {
}

type Node struct {
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
