package main

import "./cronng"

func main() {
	job := &cronng.Job{
		Name:        "Echo Hello",
		Description: "My First description",
		Script:      "#!/bin/bash\n touch /tmp/hoge",
		Schedule:    "",
	}
	user := cronng.User{ID: ""}
	args := []cronng.Arg{}
	envvar := map[string]string{}
	description := "test"
	cronng.NewRecord(job)
	job.Start(user, args, envvar, description)
}
