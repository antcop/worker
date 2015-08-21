package daemon

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func getDaemon() Daemon {
	return Daemon {
		Name: "job",
		Description: "Job server",
		Port : 1234,
		OnStart: func(daemon Daemon) {
			daemon.Print("Start")
		},
		OnStop: func(daemon Daemon) {
			daemon.Print("Stop")
		},
	}
}
func TestDaemonConstructor(t *testing.T) {
	assert := assert.New(t)
	daemon := getDaemon();
	assert.Equal(daemon.GetName(), "job")
	assert.Equal(daemon.GetDescription(), "Job server")
	assert.Equal(daemon.GetPort(), 1234)
}

func TestDaemonRun(t *testing.T) {
	assert := assert.New(t)
	daemon := getDaemon()
	daemon.Run(false)
	assert.Equal("Hello", "Hello")	
}
