package main

import (
	"fmt"

	"./cronng"
)

var Strage = cronng.NewLocalStrage()

func main() {
	defer Strage.Db.Close()
	Strage.Insert(&cronng.Job{Name: "myjob"})
	v := Strage.Insert(&cronng.Execution{Description: "my first job", Status: cronng.FAILED})
	fmt.Print(v)
}
