package cronng

import (
	"bufio"
	"fmt"
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

func saveScript(prefix, content string) (string, error) {
	fw, error := ioutil.TempFile("", prefix)
	if error != nil {
		return "", error
	}
	defer fw.Close()
	fw.Write([]byte(content))
	scriptName := fw.Name()
	os.Chmod(scriptName, 0755)
	return scriptName, nil
}

func execScript(scriptName string, args []Arg, execution *Execution) {
	// exec script
	fmt.Println("in goroutine: " + scriptName)
	cmd := exec.Command(scriptName, string(args[0]))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	bindStream(stdout, execution.stdout)
	bindStream(stderr, execution.stderr)
	cmd.Wait()
	defer func() {
		execution.Ended = time.Now()
		execution.quit <- true
	}()
}

func (job *Job) Start(user User, args []Arg, envvar EnvVar, description string) (execition *Execution, err error) {
	// make execution object
	execution := NewExecution(user, *job, args, description)
	// save script
	scriptName, error := saveScript(string(execution.ID), job.Script)
	if error != nil {
		log.Fatal(error)
		return nil, error
	}
	//	fmt.Println("outside goroutine: " + scriptName)
	execution.Started = time.Now()
	go execScript(scriptName, args, execution)
	return execution, nil
}

func (execution *Execution) Signal(signal uint, user User) error {
	return nil
}
