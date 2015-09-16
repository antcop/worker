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

package util

import (
	"io"
	"encoding/json"
	"strconv"
)

type Json map[string] interface {}

// Convert string to JSON Object type
func ToJsonObject(data io.Reader) Json {
	decoder := json.NewDecoder(data)
	var json Json
	decoder.Decode(&json)
	return json
}

// Convert string to JSON Array type
func ToJsonArray(data io.Reader) []Json {
	decoder := json.NewDecoder(data)
	var json []Json
	decoder.Decode(&json)
	return json
}

func Uint2Str(number uint) string {
	return strconv.FormatUint(uint64(number), 10)
}