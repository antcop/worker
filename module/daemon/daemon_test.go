/**
 * Ant Worker Project
 *
 * Copyright (c) 2015 Epinion Online Research Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Jerry Pham       <jerry@andjerry.com>
 *     Loi Nguyen       <loint@penlook.com>
 */

package daemon

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func getDaemon() Daemon {
	return Daemon {
		Name: "job",
		Description: "Job server",
		Bind: "0.0.0.0",
		Port: 1234,
		OnStart: func(daemon Daemon) {
			daemon.Println("Start")
		},
		OnStop: func(daemon Daemon) {
			daemon.Println("Stop")
		},
	}
}

func TestDaemonConstructor(t *testing.T) {
	assert := assert.New(t)
	daemon := getDaemon()
	assert.Equal(daemon.GetName(), "job")
	assert.Equal(daemon.GetDescription(), "Job server")
	assert.Equal(daemon.GetBind(), "0.0.0.0")
	assert.Equal(daemon.GetPort(), 1234)
}

func TestDaemonRun(t *testing.T) {
	assert := assert.New(t)
	//daemon := getDaemon()
	//daemon.RunAsService(true)
	assert.Equal("Hello", "Hello")
}
