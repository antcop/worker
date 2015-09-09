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

package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//. "github.com/epinion-online-research/ant-worker/util"
)

func TestJobSendMailEntity( t *testing.T ) {
	assert := assert.New(t)
	/*
	job1 := Job {
		Name: "sendemail",
		Description: "Send Email By Using MailChimp",
		Type: "api_call",
		Endpoint: Json {
			"url" : "http://mailchimp.com/api/v1/",
			"method": "GET", // "POST"
			"data": Json {
				"sender": "root@localhost",
				"receiver": "",
				"user_activation": "Dear, {{ name }}, \n",
			},
		},
		Workload: []map[string] interface{} {
			Json {
				"user_activation": Json {
					"id" : 123456,
					"email": "nguyen.trung.loi@epinion.dk",
					"name": "Loi",
				},
				"user_redemption": Json {
					"id" : 123456,
					"email": "nguyen.trung.loi@epinion.dk",
					"name": "Loi",
				},
			},
			Json {
				"user_activation": Json {
					"id" : 123457,
					"email": "nguyen.trung.loi@epinion.dk",
					"name": "Loi",
				},
				"user_redemption": Json {
					"id" : 123456,
					"email": "nguyen.trung.loi@epinion.dk",
					"name": "Loi",
				},
			},
		},
	}*/
	assert.Equal("test", "test")
}