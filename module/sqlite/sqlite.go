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

package sqlite

import (
    "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	File string
	Db gorm.DB
}

func (sql Sqlite) Connect() Sqlite {
	db, err := gorm.Open("sqlite3", sql.File)
	if err != nil {
		panic(err)
	}
	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)
	db.SingularTable(true)
	sql.Db = db
	return sql
}

