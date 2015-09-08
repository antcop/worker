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
	"github.com/epinion-online-research/ant-worker/entity"
)

type Module struct {
	Config entity.Config
	Redis Redis
	Sql Sqlite
}

func (mod *Module) Load() {
	mod.Redis = Redis {
		Server: "localhost:6379",
	}.Connect()
	mod.Sql = Sqlite {
		File: "/tmp/ant-worker.db",
	}.Connect()
}