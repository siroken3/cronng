package main

import "./cronng"

var Strage = cronng.NewLocalStrage()

func main() {
	job := &cronng.Job{}
	//	Strage.NewRecord(job)
	Strage.Create(job)

	defer Strage.DB().Close()
}
