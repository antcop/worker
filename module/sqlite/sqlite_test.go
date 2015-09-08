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
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
	"testing"
)

type User struct {
    gorm.Model
    Name string
	Age int
}

func TestSqlite(t *testing.T) {
	assert := assert.New(t)
	sql := Sqlite {
		File : "./ant-worker-test.db",
	}

	db := sql.Connect().Db

	// Reset database :D
	db.DropTable(&User{})
	db.CreateTable(&User{})

	var user User
	db.Create(&User {
		Name: "Loi Nguyen", 
		Age: 22,
	})
	db.First(&user)
	assert.Equal("Loi Nguyen", user.Name)
	assert.Equal(22, user.Age)
}

