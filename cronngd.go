package main

import "./cronng"

func main() {
	Strage := cronng.NewLocalStrage()
	defer Strage.DB().Close()
}
