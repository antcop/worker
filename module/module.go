/*
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
	. "github.com/epinion-online-research/ant-worker/module/redis"
	. "github.com/epinion-online-research/ant-worker/module/sqlite"
	. "github.com/epinion-online-research/ant-worker/entity"
	"github.com/jimlawless/cfg"
	"log"
	"os"
	"strconv"
)

type Module struct {
	Config Config
	Redis Redis
	Sql Sqlite
}

func getStrValue(data map[string] string, key string, defaultValue string) string {
	if value, ok := data[key]; ok {
		return value
	}
	return defaultValue
}

func getIntValue(data map[string] string, key string, defaultValue int) int {
	if value, ok := data[key]; ok {
		result, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		return result
	}
	return defaultValue
}

// Load configuration from conf file
func LoadConfig(configPath string) Config {
	data := make(map[string] string)
	if _, err := os.Stat(configPath); err == nil {
		cfg.Load(configPath, data)
	}
	return Config {
		Name:           getStrValue(data, "name", "ant-worker"),
		Description:    getStrValue(data, "description", "Simple job service"),
		Bind :          getStrValue(data, "bind", "0.0.0.0"),
		Port :          getIntValue(data, "port", 2468),
		JobConcurrency: getIntValue(data, "parallel", 5),
		MaxWorker :     getIntValue(data, "worker", 5),
		MemoryLimit:    getIntValue(data, "memory_limit", 5),
		RedisHost:      getStrValue(data, "redis_host", "127.0.0.1"),
		RedisPort:      getIntValue(data, "redis_port", 6379),
		DatabasePath:   getStrValue(data, "db_path", "/tmp/" + getStrValue(data, "name", "ant-worker")),
	}
}

// Load modules for service
func (mod *Module) Load() {
	// Configuration
	// TODO - Fix config path
	mod.Config = LoadConfig("/etc/init/ant-worker.conf")
	// Redis server
	mod.Redis = Redis {
		Server: mod.Config.RedisHost + ":" + strconv.Itoa(mod.Config.RedisPort),
	}.Connect()
	// Sqlite database
	mod.Sql = Sqlite {
		File: mod.Config.DatabasePath,
	}.Connect()
}