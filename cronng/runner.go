package cronng

import "strings"

func (self *Job) Start(user User, args []string, envvar []string, description string) (*Execution, error) {
	// process start
	execution := &Execution{
		StdOutURL:   stdOutUrl,
		StdErrURL:   stdErrUrl,
		Status:      RUNNING,
		UserID:      User.ID,
		JobID:       self.ID,
		Description: description,
		Args:        strings.Join(args[:], " "),
		Monitoring: Monitoring{
			EnvVar: strings.Join(envvar[:], " "),
		},
		Started:   started,
		Ended:     nil,
		AbortedBy: nil,
	}
	Strage.DB().Create(&execution)
	Strage.DB().Create(&monitoring)
	return execution, nil
}

func (execution *Execution) Signal(signal uint, user User) error {
	return nil
}
