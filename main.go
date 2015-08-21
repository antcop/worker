package main

import (
	. "github.com/epinion-online-research/ant-worker/daemon"
)

func main() {
	daemon := Daemon {
		Name: "job",
		Description: "Job server",
		Port : 1234,
		OnStart: func(daemon Daemon) {
			daemon.Print("Service start")
		},
		OnStop: func(daemon Daemon) {
			daemon.Print("Service stop")
		},
	}
	daemon.Run(true)
}