package cronng

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func bindStream(readCloser io.ReadCloser, output chan string) {
	reader := bufio.NewReader(readCloser)
	line, err := reader.ReadString('\n')
	for err == nil {
		output <- line
		line, err = reader.ReadString('\n')
	}
}

func saveScript(prefix, content string) (scriptName, error) {
	fw, error := ioutil.TempFile("", prefix)
	if error != nil {
		return nil, error
	}
	defer fw.Close()
	fw.Write([]byte(content))
	scriptName := fw.Name()
	os.Chmod(scriptName, 0755)
	return scriptName, nil
}

func (job *Job) Start(user User, args []Arg, envvar EnvVar, description string) (*Execution, error) {
	// make execution object
	execution = NewExecution(user, job, args, description)
	// save script
	scriptName, error := saveScript(execution.ID, job.Script)
	if error != nil {
		log.Fatal(error)
		return error
	}
	go func() {
		// exec script
		cmd := exec.Command(scriptName, args)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
		execution.StartedAt = time.Now()
		bindStream(stdout, execution.stdout)
		bindStream(stderr, execution.stderr)
		cmd.Wait()
		defer func() {
			execution.EndedAt = time.Now()
			execution.quit <- true
		}()
	}()
	return execution, nil
}

func (execution *Execution) Signal(signal uint, user User) error {
	return nil
}
