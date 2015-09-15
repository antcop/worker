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

package module

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
)

func TestModuleConfig(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("test", "test")
	module := Module {}
	pwd := os.Getenv("PWD")
	configFile, _ := filepath.Abs(pwd + "/../test.conf")
	_, err := os.Stat(configFile)
	module.Load(configFile)
	assert.Nil(err)
	assert.Equal("admin", module.Config.User)
	assert.Equal("admin", module.Config.Password)
	assert.Equal("ant-worker", module.Config.Name)
	assert.Equal("Simple job service", module.Config.Description)
	assert.Equal("0.0.0.0", module.Config.Bind)
	assert.Equal(2468, module.Config.Port)
	assert.Equal(5, module.Config.JobConcurrency)
	assert.Equal(5, module.Config.MaxWorker)
	assert.Equal(1024, module.Config.MemoryLimit)
	assert.Equal("127.0.0.1", module.Config.RedisHost)
	assert.Equal(6379, module.Config.RedisPort)
	assert.Equal("../ant-worker.db", module.Config.DatabasePath)
}